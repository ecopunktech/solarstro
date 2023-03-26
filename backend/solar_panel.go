package main

type SolarPanel struct {
	Price        float64
	Warranty     int
	Width        float64
	Height       float64
	High         float64
	Tension      float64
	Power        float64
	Efficiency   float64
	Manufacturer string
	Model        string
}

func (s *SolarPanel) GetPrice(n int) float64 {
	return s.Price * float64(n)
}

func (s *SolarPanel) GetTotalPower(n int) float64 {
	return s.Power / 1000 * float64(n)
}

// Nº	Fabricante	Modelo. Tipo Monocristalino	Eficiencia	Potencia en Watios	Tension de Conexión	Tamaño en mm	Precio en €. Iva incluido	Garantía en años	Observaciones
// 1	Sun Power	Max3 400 Wp	Hasta 22%	400	24 Voltios	1046x1690x40	375	40/40	72 celdas
// 2	Panasonic (Sanyo)	VBHN325SJ47	Hasta 19%	325	24/32 voltios	1053x1590x35	250	25/10	96 celdas
// 3	LG	Neon R LG370Q1C-V5	"Hasta 17%"	370	12 voltios	1016x1700x40	300	25/10	60 celdas

// 4	Sharp	NU-AH370	Hasta 19%	370	24 voltios	1956x992x35	160	25/10	72 celdas
// 5	JinkoSolar	Tiger Pro 460Wp 120cél 1500v 	Hasta 18,7%	460	24 voltios	1903x1134x30	160	25/12	120 celdas
// 6	RisenEnergy	Perc Titan 40	Hasta 21,6%	450	24 voltios	1754x1096x30	178	25/12	120 celdas
// 7	EGingPV	Star Pro EG-460NT60-HLV	Hasta 21,24%	460	24 voltios	1909x1134x30 	180	25/12	120 celdas
// 8	Atersa	A-450M GS (M6x24) PERC	Hasta 20,7%	450	24 voltios	2094 × 1038x40	192	25/10	72 celdas
// 9	Eurener	MEPV144.HALF-CUT PLUS	Hasta 21,3%	450	24 voltios	2094x1038x35	160	25/20	72 celdas
// 10	EastchSolar	Amerisolar 450	Hasta 20,58%	450	24 voltios	2102x1040x35	295	25/12	72 celdas
// 11	Tamesol	TM 430-450M-144HC	Hasta 20,09	450	24 voltios	2094x1038x35	201	35/25	72 celdas
// 12	JASolar	460W 24V Perc	Hasta 20,7%	460	24 voltios	2120x1052x 35	205	25/12	72 celdas
func SharpSolarPanel() SolarPanel {
	return SolarPanel{
		Price:        160,
		Warranty:     25,
		Width:        1956,
		Height:       992,
		High:         35,
		Tension:      24,
		Power:        370,
		Efficiency:   19,
		Manufacturer: "Sharp",
		Model:        "NU-AH370",
	}
}

func JinkoSolarPanel() SolarPanel {
	return SolarPanel{
		Price:        160,
		Warranty:     25,
		Width:        1903,
		Height:       1134,
		High:         30,
		Tension:      24,
		Power:        460,
		Efficiency:   18.7,
		Manufacturer: "JinkoSolar",
		Model:        "Tiger Pro 460Wp 120cél 1500v",
	}
}

func SunPowerSolarPanel() SolarPanel {
	return SolarPanel{
		Price:        375,
		Warranty:     40,
		Width:        1046,
		Height:       1690,
		High:         40,
		Tension:      24,
		Power:        400,
		Efficiency:   22,
		Manufacturer: "Sun Power",
		Model:        "Max3 400 Wp",
	}
}

func PanasonicSolarPanel() SolarPanel {
	return SolarPanel{
		Price:        250,
		Warranty:     25,
		Width:        1053,
		Height:       1590,
		High:         35,
		Tension:      24,
		Power:        325,
		Efficiency:   19,
		Manufacturer: "Panasonic (Sanyo)",
		Model:        "VBHN325SJ47",
	}
}

func LGSolarPanel() SolarPanel {
	return SolarPanel{
		Price:        300,
		Warranty:     25,
		Width:        1016,
		Height:       1700,
		High:         40,
		Tension:      12,
		Power:        370,
		Efficiency:   17,
		Manufacturer: "LG",
		Model:        "Neon R LG370Q1C-V5",
	}
}

func GetSolarPanelByBudget(budgetType string) SolarPanel {
	switch budgetType {
	case "low":
		return SharpSolarPanel()
	case "med":
		return LGSolarPanel()
	case "high":
		return SunPowerSolarPanel()
	}
	return SunPowerSolarPanel()
}
