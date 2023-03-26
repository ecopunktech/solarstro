package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func svgHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	rc := query.Get("rc")
	if rc == "" {
		http.Error(w, "Missing 'rc' parameter", http.StatusBadRequest)
		return
	}

	gmlFile, err := getGMLFromCatastro(rc)
	if err != nil {
		fmt.Println(err)
	}
	polygon, err := getCoordinatesFromGML(gmlFile)
	if err != nil {
		fmt.Println(err)
	}

	minX, minY, maxX, maxY := PoligonBB(polygon)

	minPoint := Point{minX, minY}
	maxPoint := Point{maxX, maxY}
	panels := rectanglesInsidePolygon(polygon, Rectangle{Width: 2, Height: 2})

	// scaleFactor := 10.0
	scaledPolygon := scaleCoordinates(polygon, 1000, 1000)
	rectangule := make([]Rectangle, len(panels))
	for i, p := range panels {
		rectangule[i] = scaleRectangle(p, minPoint, maxPoint, 1000, 1000)
	}
	svgData, err := createSVG(scaledPolygon, rectangule)
	if err != nil {
		fmt.Println("Error creating SVG data:", err)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write(svgData)
}

func rectanglesHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	rc := query.Get("rc")
	if rc == "" {
		http.Error(w, "Missing 'rc' parameter", http.StatusBadRequest)
		return
	}
	areaPercentageQuery := query.Get("percentage")
	if areaPercentageQuery == "" {
		areaPercentageQuery = "100"
	}

	areaPercentage, err := strconv.ParseFloat(areaPercentageQuery, 64)
	if err != nil && areaPercentage < 0 && areaPercentage > 100 {
		http.Error(w, "Invalid 'area_percentage' parameter", http.StatusBadRequest)
		return
	}

	budget := query.Get("budget")
	if budget == "" {
		budget = "med"
	}

	coord, err := getCoordinatesFromCadastralCode(rc)

	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, "Error getting coordinates", http.StatusBadRequest)
		return
	}

	lon := coord.ConsultaCPMRCResult.Coordenadas.Coord[0].Geo.Xcen
	lat := coord.ConsultaCPMRCResult.Coordenadas.Coord[0].Geo.Ycen
	address := coord.ConsultaCPMRCResult.Coordenadas.Coord[0].Ldt

	gmlFile, err := getGMLFromCatastro(rc)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting GML", http.StatusBadRequest)
		return
	}
	polygon, err := getCoordinatesFromGML(gmlFile)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting coordinates from GML", http.StatusBadRequest)
		return
	}
	area := polygonArea(polygon)
	solarPanelType := GetSolarPanelByBudget(budget)

	panelsTotal := rectanglesInsidePolygon(polygon, Rectangle{Width: solarPanelType.Width / 1000, Height: solarPanelType.Height / 1000})
	panels := panelsTotal[:int(float64(len(panelsTotal))*areaPercentage/100)]

	solarFarmResponse, err := GetSolarInfo(lat, lon, len(panels), solarPanelType)
	if err != nil {
		fmt.Println("Error getting solar info:", err)
		http.Error(w, "Error getting solar info", http.StatusInternalServerError)
		return
	}

	response := struct {
		CadastralCode string            `json:"cadastral_code"`
		Latitude      string            `json:"latitude"`
		Longitude     string            `json:"longitude"`
		Address       string            `json:"address"`
		Area          float64           `json:"area"`
		NumPanels     int               `json:"num_panels"`
		SolarInfo     SolarFarmResponse `json:"solar_info"`
	}{
		CadastralCode: rc,
		Latitude:      lat,
		Longitude:     lon,
		Address:       address,
		Area:          area,
		NumPanels:     len(panels),
		SolarInfo:     solarFarmResponse,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
