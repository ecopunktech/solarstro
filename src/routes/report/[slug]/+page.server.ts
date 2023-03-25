import type { PageServerLoad } from './$types';
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



export const load = (async ({ params, url }) => {
	console.log(url.searchParams.get('percentage'));
	const dashboard = await fetch(`https://solar-backend.fly.dev/rectangles?rc=${params.slug}`).then(
		(r) => r.json()
	);
	const dasObj: SolarPanelData = dashboard as SolarPanelData;
	return { dashboard: dasObj };
}) satisfies PageServerLoad;
