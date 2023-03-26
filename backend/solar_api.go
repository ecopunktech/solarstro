package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Inputs struct {
	Location       Location       `json:"location"`
	MeteoData      MeteoData      `json:"meteo_data"`
	MountingSystem MountingSystem `json:"mounting_system"`
	PVModule       PVModule       `json:"pv_module"`
	EconomicData   EconomicData   `json:"economic_data"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Elevation float64 `json:"elevation"`
}

type MeteoData struct {
	RadiationDB string `json:"radiation_db"`
	MeteoDB     string `json:"meteo_db"`
	YearMin     int    `json:"year_min"`
	YearMax     int    `json:"year_max"`
	UseHorizon  bool   `json:"use_horizon"`
	HorizonDB   string `json:"horizon_db"`
}

type MountingSystem struct {
	Fixed struct {
		Slope struct {
			Value   int  `json:"value"`
			Optimal bool `json:"optimal"`
		} `json:"slope"`
		Azimuth struct {
			Value   int  `json:"value"`
			Optimal bool `json:"optimal"`
		} `json:"azimuth"`
		Type string `json:"type"`
	} `json:"fixed"`
}

type PVModule struct {
	Technology string  `json:"technology"`
	PeakPower  float64 `json:"peak_power"`
	SystemLoss float64 `json:"system_loss"`
}

type EconomicData struct {
	SystemCost float64 `json:"system_cost"`
	Interest   float64 `json:"interest"`
	Lifetime   int     `json:"lifetime"`
}

type Outputs struct {
	Monthly Monthly `json:"monthly"`
	Totals  Totals  `json:"totals"`
}

type Monthly struct {
	Fixed []struct {
		Month int     `json:"month"`
		ED    float64 `json:"E_d"`
		EM    float64 `json:"E_m"`
		HID   float64 `json:"H(i)_d"`
		HIM   float64 `json:"H(i)_m"`
		SDM   float64 `json:"SD_m"`
	} `json:"fixed"`
}

type Totals struct {
	Fixed struct {
		ED     float64 `json:"E_d"`
		EM     float64 `json:"E_m"`
		EY     float64 `json:"E_y"`
		HID    float64 `json:"H(i)_d"`
		HIM    float64 `json:"H(i)_m"`
		HIY    float64 `json:"H(i)_y"`
		SDM    float64 `json:"SD_m"`
		SDY    float64 `json:"SD_y"`
		LAOI   float64 `json:"l_aoi"`
		Lspec  string  `json:"l_spec"`
		Ltg    float64 `json:"l_tg"`
		Ltotal float64 `json:"l_total"`
		LCOEPv float64 `json:"LCOE_pv"`
	} `json:"fixed"`
}

type PVResponse struct {
	Inputs  Inputs  `json:"inputs"`
	Outputs Outputs `json:"outputs"`
}

type PVGISResponse struct {
	Outputs struct {
		EPVTimeSeries []float64 `json:"EPV_time_series"`
	} `json:"outputs"`
}

// create function to calculate total cost of solar panels installation
func GetCost(panels int, panelInfo SolarPanel) float64 {
	// cost of solar panels
	costPanels := panelInfo.GetPrice(panels)
	// cost of installation
	installation := 0.0
	// cost of inverter
	inverter := 0.0
	// cost of racking
	racking := 0.0
	// cost of wiring
	wiring := 0.0
	// cost of permits
	permits := 0.0
	// cost of interconnection
	interconnection := 0.0
	// cost of financing
	financing := 0.0
	// cost of maintenance
	maintenance := 0.0
	return costPanels + installation + inverter + racking + wiring + permits + interconnection + financing + maintenance
}

func MapPVResponseToSolarFarmResponse(pvResponse PVResponse, panels int, panel SolarPanel, compensation float64) SolarFarmResponse {
	powerGeneration := make(map[string]float64)
	monthlyIncome := make(map[string]float64)
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	for _, monthData := range pvResponse.Outputs.Monthly.Fixed {
		month := months[monthData.Month-1]
		powerGeneration[month] = monthData.EM
		monthlyIncome[month] = monthData.EM * compensation
	}

	annualEnergy := pvResponse.Outputs.Totals.Fixed.EY
	annualIncome := annualEnergy * compensation
	totalInstallationCost := pvResponse.Inputs.EconomicData.SystemCost
	paybackPeriod := totalInstallationCost / annualIncome

	solarFarmResponse := SolarFarmResponse{
		Area:                       0, // Add the area value here
		NumSolarPanels:             panels,
		TotalPower:                 panel.GetTotalPower(panels),
		PowerGeneration:            powerGeneration,
		MonthlyIncome:              monthlyIncome,
		TotalInstallationCost:      totalInstallationCost,
		PaybackPeriod:              paybackPeriod,
		AnnualCO2EmissionReduction: 0, // Add the annual CO2 emission reduction value here
		// Add other fields like IncentivesAndSubsidies, Costs, MaintenanceCost, etc.
		Lifetime:        float64(panel.Warranty),
		DegradationRate: 0,                // Add the degradation rate value here
		Efficiency:      panel.Efficiency, // Add the efficiency value here
	}

	return solarFarmResponse
}

func GetSolarInfo(lat, lon string, panels int, panelInfo SolarPanel) (SolarFarmResponse, error) {

	u := &url.URL{
		Scheme: "https",
		Host:   "re.jrc.ec.europa.eu",
		Path:   "/api/PVcalc",
	}

	systemCost := GetCost(panels, panelInfo) // €/kWp
	interestRate := 0.05                     // 5%
	lost := 4.5                              // 14%
	powerInstalled := panelInfo.GetTotalPower(panels)
	// La compensación por excedentes es una tarifa que se aplica a los propietarios de granjas solares que generan más energía de la que consumen. El precio de la compensación por excedentes depende de cada comercializadora y de las tarifas que ofrezca, pero la media del mercado nacional está en torno a los 0,100 €/kWh1.
	compen := 0.1 // €/kWh
	q := u.Query()
	q.Set("lat", lat)
	q.Set("lon", lon)
	q.Set("peakpower", fmt.Sprint(powerInstalled))
	q.Set("loss", fmt.Sprint(lost))
	q.Set("outputformat", "json")
	q.Set("pvprice", "1")
	q.Set("systemcost", fmt.Sprint(systemCost))
	q.Set("interest", fmt.Sprint(interestRate))
	q.Set("optimalinclination", "1")
	u.RawQuery = q.Encode()
	pvgisURL := u.String()

	resp, err := http.Get(pvgisURL)
	if err != nil {
		fmt.Println("Error fetching PVGIS data:", err)
		return SolarFarmResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading PVGIS response:", err)
		return SolarFarmResponse{}, err
	}

	var pvgisResponse PVResponse
	err = json.Unmarshal(body, &pvgisResponse)
	if err != nil {
		fmt.Println("Error unmarshaling PVGIS response:", err)
		return SolarFarmResponse{}, err
	}

	return MapPVResponseToSolarFarmResponse(pvgisResponse, panels, panelInfo, compen), nil
}
