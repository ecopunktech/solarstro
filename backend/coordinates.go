package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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

func getCoordinatesFromCadastralCode(cadastralCode string) (CatastroResponse, error) {
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

	return coord, nil
}
