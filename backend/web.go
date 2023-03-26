package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func svgHandler(redisClient *redis.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		rc := query.Get("rc")
		if rc == "" {
			http.Error(w, "Missing 'rc' parameter", http.StatusBadRequest)
			return
		}

		gmlFile, err := getGMLFromCatastro(rc, redisClient)
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
}
func reportHandler(redisClient *redis.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
		resultsChan := make(chan [2]interface{}, 2)
		go func() {
			coord, err := getCoordinatesFromCadastralCode(rc, redisClient)
			resultsChan <- [2]interface{}{coord, err}
		}()

		go func() {
			gml, err := getGMLFromCatastro(rc, redisClient)
			resultsChan <- [2]interface{}{gml, err}
		}()
		// wait for the two functions to complete and receive their results
		var coord CatastroResponse
		var gmlFile string
		for i := 0; i < 2; i++ {
			result := <-resultsChan
			if result[1] != nil {
				fmt.Println("Error:", result[1])
				http.Error(w, "Error getting data", http.StatusBadRequest)
				return
			}
			switch v := result[0].(type) {
			case CatastroResponse:
				coord = v
			case string:
				gmlFile = v
			default:
				fmt.Println("Unknown result type:", v)
				http.Error(w, "Error getting data", http.StatusBadRequest)
				return
			}
		}

		lon := coord.ConsultaCPMRCResult.Coordenadas.Coord[0].Geo.Xcen
		lat := coord.ConsultaCPMRCResult.Coordenadas.Coord[0].Geo.Ycen
		address := coord.ConsultaCPMRCResult.Coordenadas.Coord[0].Ldt

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
}
