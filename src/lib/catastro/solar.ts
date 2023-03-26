export function getMonthlyIncomeAndPowerGeneration(solarInfo: SolarPanelData) {
	const powerGeneration = Object.entries(solarInfo.solar_info.powerGeneration)
		.sort(
			(a, b) =>
				new Date(Date.parse(`01 ${a[0]}`)).getTime() -
				new Date(Date.parse(`01 ${b[0]}`)).getTime()
		)
		.reduce((obj, [key, value]) => ({ ...obj, [key]: value }), {});

	const monthlyIncome = Object.entries(solarInfo.solar_info.monthlyIncome)
		.sort(
			(a, b) =>
				new Date(Date.parse(`01 ${a[0]}`)).getTime() -
				new Date(Date.parse(`01 ${b[0]}`)).getTime()
		)
		.reduce((obj, [key, value]) => ({ ...obj, [key]: value }), {});

	return { powerGeneration, monthlyIncome };
}
export interface SolarPanelData {
	cadastral_code: string;
	latitude: string;
	longitude: string;
	address: string;
	area: number;
	num_panels: number;
	solar_info: {
		area: number;
		numSolarPanels: number;
		totalPower: number;
		powerGeneration: {
			April: number;
			August: number;
			December: number;
			February: number;
			January: number;
			July: number;
			June: number;
			March: number;
			May: number;
			November: number;
			October: number;
			September: number;
		};
		monthlyIncome: {
			April: number;
			August: number;
			December: number;
			February: number;
			January: number;
			July: number;
			June: number;
			March: number;
			May: number;
			November: number;
			October: number;
			September: number;
		};
		totalInstallationCost: number;
		paybackPeriod: number;
		annualCO2EmissionReduction: number;
		incentivesAndSubsidies: {
			distrubutorCompensation: number;
			stateTaxCredit: number;
			localIncentives: number;
		};
		costs: {
			permitting: number;
			sitePreparation: number;
			equipment: number;
			installationLabor: number;
			interconnection: number;
			operationsAndMaintenance: number;
			monitoringAndControl: number;
			insurance: number;
		};
		maintenanceCost: {
			annual: number;
			perSolarPanel: number;
		};
		lifetime: number;
		degradationRate: number;
		efficiency: number;
		averageSunHoursPerDay: number;
	};
}
