package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/text/encoding/charmap"
)

type LinearRing struct {
	PosList string `xml:"posList"`
}

type PolygonPatch struct {
	Exterior LinearRing `xml:"exterior>LinearRing"`
}

type Surface struct {
	Patches PolygonPatch `xml:"patches>PolygonPatch"`
}

type MultiSurface struct {
	SurfaceMember Surface `xml:"surfaceMember>Surface"`
}

type Geometry struct {
	MultiSurface MultiSurface `xml:"MultiSurface"`
}

type CadastralParcel struct {
	Geometry Geometry `xml:"geometry"`
}

type FeatureCollection struct {
	XMLName xml.Name        `xml:"FeatureCollection"`
	Member  CadastralParcel `xml:"member>CadastralParcel"`
}

func getCoordinatesFromGML(gmlData string) ([]Point, error) {
	var gml FeatureCollection

	// Define the CharsetReader function to handle "ISO-8859-1" encoding
	charsetReader := func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "ISO-8859-1" {
			return charmap.ISO8859_1.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("unsupported charset: %s", charset)
	}

	// Create a new XML decoder and set the CharsetReader function

	decoder := xml.NewDecoder(bytes.NewReader([]byte(gmlData)))
	decoder.CharsetReader = charsetReader

	// Decode the XML data into a FeatureCollection struct
	if err := decoder.Decode(&gml); err != nil {
		return nil, fmt.Errorf("failed to parse GML: %v", err)
	}

	coordinatesStr := gml.Member.Geometry.MultiSurface.SurfaceMember.Patches.Exterior.PosList
	coordinatePairs := strings.Split(coordinatesStr, " ")
	var coordinates []Point
	for i := 0; i < len(coordinatePairs)-1; i += 2 {
		point := Point{
			X: parseCoordinate(coordinatePairs[i]),
			Y: parseCoordinate(coordinatePairs[i+1]),
		}
		coordinates = append(coordinates, point)
	}

	return coordinates, nil
}

func parseCoordinate(coordStr string) float64 {
	var coord float64
	fmt.Sscanf(coordStr, "%f", &coord)
	return coord
}

func getGMLFromCatastro(codigoCatastral string, redisClient *redis.Client) (string, error) {
	// check if the response is already cached in Redis
	cacheKey := fmt.Sprintf("gml:%s", codigoCatastral)
	context := context.Background()
	cachedGML, err := redisClient.Get(context, cacheKey).Result()
	if err == nil {
		// response found in cache, return it
		return cachedGML, nil
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	const baseURL = "https://ovc.catastro.meh.es/INSPIRE/wfsCP.aspx?service=wfs&version=2&request=getfeature&STOREDQUERIE_ID=GetParcel"

	reqURL := fmt.Sprintf("%s&refcat=%s&srsname=EPSG::25830", baseURL, codigoCatastral)
	req, err := http.NewRequest("GET", reqURL, nil)

	if err != nil {
		fmt.Println("Error creating request:", err)
		// return
	}
	// Set the header
	req.Header.Set("User-Agent", "curl/7.85.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	// cache the response in Redis
	err = redisClient.Set(context, cacheKey, string(body), 24*time.Hour).Err()
	if err != nil {
		fmt.Println("Failed to cache response:", err)
	}

	return string(body), nil
}
