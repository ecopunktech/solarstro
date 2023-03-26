package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/go-redis/redis/v8"
)

type Coord struct {
	PC struct {
		PC1 string `json:"pc1"`
		PC2 string `json:"pc2"`
	} `json:"pc"`
	Geo struct {
		Xcen string `json:"xcen"`
		Ycen string `json:"ycen"`
		Srs  string `json:"srs"`
	} `json:"geo"`
	Ldt string `json:"ldt"`
}

type Coordenadas struct {
	Coord []Coord `json:"coord"`
}

type Control struct {
	Cucoor int `json:"cucoor"`
}

type ConsultaCPMRCResult struct {
	Control     Control     `json:"control"`
	Coordenadas Coordenadas `json:"coordenadas"`
}

type CatastroResponse struct {
	ConsultaCPMRCResult ConsultaCPMRCResult `json:"Consulta_CPMRCResult"`
}

func getCoordinatesFromCadastralCode(cadastralCode string, redisClient *redis.Client) (CatastroResponse, error) {
	// check if the response is already cached in Redis
	cacheKey := fmt.Sprintf("catastro:%s", cadastralCode)
	ctx := context.Background()
	cachedResponse, err := redisClient.Get(ctx, cacheKey).Bytes()
	if err == nil {
		// response found in cache, unmarshal and return it
		var coord CatastroResponse
		err = json.Unmarshal(cachedResponse, &coord)
		if err != nil {
			return coord, fmt.Errorf("failed to unmarshal cached response: %v", err)
		}
		return coord, nil
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	var coord CatastroResponse
	baseURL := "https://ovc.catastro.meh.es/OVCServWeb/OVCWcfCallejero/COVCCoordenadas.svc/json/Consulta_CPMRC?SRS=EPSG:4326"
	reqURL, err := url.Parse(baseURL)
	if err != nil {
		return coord, err
	}

	q := reqURL.Query()
	q.Set("refcat", cadastralCode)
	reqURL.RawQuery = q.Encode()
	// resp, err := http.Get(reqURL.String())
	req, err := http.NewRequest("GET", reqURL.String(), nil)

	if err != nil {
		fmt.Println("Error creating request:", err)
		// return
	}
	// Set the header
	req.Header.Set("User-Agent", "curl/7.85.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
	resp, err := client.Do(req)
	if err != nil {
		return coord, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return coord, err
	}
	err = json.Unmarshal(body, &coord)
	if err != nil {
		return coord, err
	}

	// cache the response in Redis
	responseBytes, err := json.Marshal(coord)
	if err != nil {
		fmt.Println("Failed to marshal response:", err)
	} else {
		err = redisClient.Set(ctx, cacheKey, responseBytes, 24*time.Hour).Err()
		if err != nil {
			fmt.Println("Failed to cache response:", err)
		}
	}

	return coord, nil
}
