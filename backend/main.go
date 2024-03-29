package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
)

func main() {
	// get env redis
	redisURI := os.Getenv("REDIS_URL")
	opt, _ := redis.ParseURL(redisURI)
	redisClient := redis.NewClient(opt)

	// GetSolarInfo()
	http.HandleFunc("/svg", svgHandler(redisClient))
	http.HandleFunc("/report", reportHandler(redisClient))

	fmt.Println("Server listening on :8080")

	// GetSolarInfo()
	// TestMain()
	http.ListenAndServe(":8080", nil)
}

// func TestMain() {

// 	rc := "33034A02000029"
// 	cadastralCode := "33034A02000041"
// 	coord, err := getCoordinatesFromCadastralCode(cadastralCode)

// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	lon := coord.ConsultaCPMRCResult.Coordenadas.Coord[0].Geo.Xcen
// 	lat := coord.ConsultaCPMRCResult.Coordenadas.Coord[0].Geo.Ycen
// 	address := coord.ConsultaCPMRCResult.Coordenadas.Coord[0].Ldt

// 	fmt.Printf("Coordinates for cadastral code %s: X: %s, Y: %s\nAdrress: %s\n", rc, lon, lat, address)
// 	redisClient := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379", // Redis server address
// 		Password: "",               // Redis password
// 		DB:       0,                // Redis database number
// 	})

// 	gmlFile, err := getGMLFromCatastro(rc, redisClient)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	polygon, err := getCoordinatesFromGML(gmlFile)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	area := polygonArea(polygon)
// 	panasonic := PanasonicSolarPanel()

// 	panels := rectanglesInsidePolygon(polygon, Rectangle{Width: panasonic.Width / 1000, Height: panasonic.Height / 1000})

// 	fmt.Println("Panels:", len(panels), "Area:", area, "m2")
// 	fmt.Println("Panasonic:")
// 	GetSolarInfo(lat, lon, len(panels), panasonic)

// }
