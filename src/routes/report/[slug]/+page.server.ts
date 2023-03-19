import type { PageServerLoad } from './$types';

export const load = (async ({ params }) => {
	return {
		dashboard: {
			cadastral_code: '33034A02000032',
			latitude: '43.5452340306152',
			longitude: '-6.54927168566082',
			address: 'Polígono 20 Parcela 32 S MARTIN. VALDES (ASTURIAS)',
			area: 1908.4088113311705,
			num_panels: 1061,
			solar_info: {
				area: 0,
				numSolarPanels: 1061,
				totalPower: 344825,
				powerGeneration: {
					April: 41442.77,
					August: 47426.34,
					December: 26142.47,
					February: 28343.95,
					January: 23991.61,
					July: 46007.08,
					June: 42943.29,
					March: 39130.58,
					May: 42882.9,
					November: 24631.86,
					October: 36047.75,
					September: 43355.57
				},
				monthlyIncome: {
					April: 4144.277,
					August: 4742.634,
					December: 2614.2470000000003,
					February: 2834.3950000000004,
					January: 2399.161,
					July: 4600.7080000000005,
					June: 4294.329000000001,
					March: 3913.0580000000004,
					May: 4288.29,
					November: 2463.186,
					October: 3604.775,
					September: 4335.557
				},
				totalInstallationCost: 265250,
				paybackPeriod: 5.996435144739023,
				annualCO2EmissionReduction: 0,
				incentivesAndSubsidies: {
					distrubutorCompensation: 0,
					stateTaxCredit: 0,
					localIncentives: 0
				},
				costs: {
					permitting: 0,
					sitePreparation: 0,
					equipment: 0,
					installationLabor: 0,
					interconnection: 0,
					operationsAndMaintenance: 0,
					monitoringAndControl: 0,
					insurance: 0
				},
				maintenanceCost: { annual: 0, perSolarPanel: 0 },
				lifetime: 25,
				degradationRate: 0,
				efficiency: 19,
				averageSunHoursPerDay: 0
			}
		}
	};
}) satisfies PageServerLoad;
