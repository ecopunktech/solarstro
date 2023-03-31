<script lang="ts">
	import { getMonthlyIncomeAndPowerGeneration } from './catastro/solar';
	import TableChart from './TableChart.svelte';
	import type { SolarPanelData } from './catastro/solar';
	import { Heading, DescriptionList, List, P } from 'flowbite-svelte';
	export let data;
	const redTag = 'text-lg font-semibold text-red-500';
	const greenTag = 'text-lg font-semibold text-green-500';
	let monthlyIncome = getMonthlyIncomeAndPowerGeneration(data.dashboard).monthlyIncome;
	function formatYears(years: number) {
		const wholeYears = Math.floor(years);
		const months = Math.floor((years - wholeYears) * 12);
		let result = `${wholeYears} year`;
		if (wholeYears !== 1) {
			result += 's';
		}
		if (months > 0) {
			result += ` ${months} month`;
			if (months !== 1) {
				result += 's';
			}
		}
		return result;
	}

	function formatNumber(number: number): string {
		const options = {
			useGrouping: true,
			maximumFractionDigits: 2,
			minimumFractionDigits: 2,
			currencyDisplay: 'symbol',
			style: 'decimal'
		};

		return number.toLocaleString('es-ES', options);
	}

	function calYearlyIncome() {
		const dash = data.dashboard as SolarPanelData;

		return monthlyIncome * 12;
	}


    function calAnualIncomeNumber() {
        const dash = data.dashboard as SolarPanelData;
		const years = dash.solar_info.lifetime;
		// iterrate over the years and add the monthly income
		let d = 0;
		Object.keys(dash.solar_info.monthlyIncome).forEach((key: string) => {
			if (dash.solar_info) {
				d += dash.solar_info.monthlyIncome[key];
			}
		});
        return d;
    }
	function calAnualIncome() {
		const d = calAnualIncomeNumber();
		return formatNumber(d);
	}

    function calTotalIncoming() {
        const dash = data.dashboard as SolarPanelData;
        const years = dash.solar_info.lifetime;
        const totalCost = years * calAnualIncomeNumber();
        return formatNumber(totalCost);
    }

    function calcCost() {
        const dash = data.dashboard as SolarPanelData;
        // const totalCost = years * calAnualIncomeNumber();
        // const totalCost = years * calAnualIncomeNumber();
        return formatNumber(dash.solar_info.maintenanceCost.annual);
    }

    function  totalCostNumber() {
        const dash = data.dashboard as SolarPanelData;
        const years = dash.solar_info.lifetime;
        const totalCost = years * dash.solar_info.maintenanceCost.annual + dash.solar_info.totalInstallationCost;
        return totalCost;
    }
    

    function totalCost() {
        return formatNumber(totalCostNumber());
    }

    function calcProfit() {
        const dash = data.dashboard as SolarPanelData;
        const years = dash.solar_info.lifetime;
        const profit = years * calAnualIncomeNumber() - totalCostNumber();
        // const totalCost = years * calAnualIncomeNumber();
        // const totalCost = years * calAnualIncomeNumber();
        return formatNumber(profit);
    }



</script>

<div>
	<Heading color="ml-6 text-purple-400 dark:text-white" class="flex items-center" tag="h2">
		Redimiento Económico
	</Heading>
</div>
<div color="purple" class="p-2">
	<div class="md:flex">
		<List class="p-6" tag="dl" color="text-gray-900">
			<div class="flex flex-col pb-3">
				<DescriptionList tag="dt" class="mb-1">Ganancia Anual:</DescriptionList>
				<DescriptionList class={greenTag} tag="dd">{calAnualIncome()} €</DescriptionList>
			</div>
			<div class="flex flex-col pb-3">
				<DescriptionList tag="dt" class="mb-1">Ciclo de vida:</DescriptionList>
				<DescriptionList class={greenTag} tag="dd">{data.dashboard.solar_info.lifetime} years</DescriptionList>
			</div>
			<div class="flex flex-col pb-3">
				<DescriptionList tag="dt" class="mb-1">Ganancia Total:</DescriptionList>
				<DescriptionList class={greenTag} tag="dd">{calTotalIncoming()} €</DescriptionList>
			</div>
			<div class="flex flex-col pb-3">
				<DescriptionList tag="dt" class="mb-1">Beneficio:</DescriptionList>
				<DescriptionList class={greenTag} tag="dd">{calcProfit()} €</DescriptionList>
			</div>
		</List>

		<List class="p-6" tag="dl" color="text-gray-900">
			<div class="flex flex-col pb-3">
				<DescriptionList tag="dt" class="mb-1">Coste Inicial:</DescriptionList>
				<DescriptionList class={redTag} tag="dd">
					{formatNumber(data.dashboard.solar_info.totalInstallationCost)} €
				</DescriptionList>
			</div>
			<div class="flex flex-col pb-3">
				<DescriptionList tag="dt" class="mb-1 flex">Coste Anual:</DescriptionList>
				<DescriptionList class={redTag} tag="dd">{calcCost()} €</DescriptionList>
			</div>
			<div class="flex flex-col pb-3">
				<DescriptionList tag="dt" class="mb-1 flex">Coste Total:</DescriptionList>
				<DescriptionList class={redTag} tag="dd">{totalCost()} €</DescriptionList>
			</div>

			<div class="flex flex-col pb-3">
				<DescriptionList tag="dt" class="mb-1">Retorno de inversion:</DescriptionList>
				<DescriptionList tag="dd"
					>{formatYears(data.dashboard.solar_info.paybackPeriod)}</DescriptionList
				>
			</div>
		</List>
	</div>
</div>

<TableChart data={monthlyIncome} title="Distribución Mesual" titleChart="Euros" />