package main

type SolarFarmResponse struct {
	Area                       float64                `json:"area"`
	NumSolarPanels             int                    `json:"numSolarPanels"`
	TotalPower                 float64                `json:"totalPower"`
	PowerGeneration            map[string]float64     `json:"powerGeneration"`
	MonthlyIncome              map[string]float64     `json:"monthlyIncome"`
	TotalInstallationCost      float64                `json:"totalInstallationCost"`
	PaybackPeriod              float64                `json:"paybackPeriod"`
	AnnualCO2EmissionReduction float64                `json:"annualCO2EmissionReduction"`
	IncentivesAndSubsidies     IncentivesAndSubsidies `json:"incentivesAndSubsidies"`
	Costs                      Costs                  `json:"costs"`
	MaintenanceCost            MaintenanceCost        `json:"maintenanceCost"`
	Lifetime                   float64                `json:"lifetime"`
	DegradationRate            float64                `json:"degradationRate"`
	Efficiency                 float64                `json:"efficiency"`
	AverageSunHoursPerDay      float64                `json:"averageSunHoursPerDay"`
}

type IncentivesAndSubsidies struct {
	DistributorCompensation float64 `json:"distrubutorCompensation"`
	StateTaxCredit          float64 `json:"stateTaxCredit"`
	LocalIncentives         float64 `json:"localIncentives"`
}

type Costs struct {
	Permitting               float64 `json:"permitting"`
	SitePreparation          float64 `json:"sitePreparation"`
	Equipment                float64 `json:"equipment"`
	InstallationLabor        float64 `json:"installationLabor"`
	Interconnection          float64 `json:"interconnection"`
	OperationsAndMaintenance float64 `json:"operationsAndMaintenance"`
	MonitoringAndControl     float64 `json:"monitoringAndControl"`
	Insurance                float64 `json:"insurance"`
}

type MaintenanceCost struct {
	Annual        float64 `json:"annual"`
	PerSolarPanel float64 `json:"perSolarPanel"`
}
