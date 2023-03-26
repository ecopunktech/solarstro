<script lang="ts">
	import type { PageData } from './$types';

	import { Heading, DescriptionList } from 'flowbite-svelte';

	function formatNumber(num: number): string {
		const options = {
			useGrouping: true,
			maximumFractionDigits: 2,
			minimumFractionDigits: 2,
			currencyDisplay: 'symbol',
			style: 'decimal'
		};

		return num.toLocaleString('es-ES', options);
	}

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

	export let data: PageData;
</script>

<div color="purple" class="my-8 p-6 rounded-xl shadow-lg">
	<Heading id="installation" color="text-purple-400 dark:text-white" tag="h2" class="mb-4"
		>Installation Info</Heading
	>
	<div class="flex flex-col pb-3">
		<DescriptionList tag="dt" class="mb-1">Total Power:</DescriptionList>
		<DescriptionList tag="dd"
			>{formatNumber(data.dashboard.solar_info.totalPower)} kW</DescriptionList
		>
	</div>
	<div class="flex flex-col pb-3">
		<DescriptionList tag="dt" class="mb-1">Total Panels:</DescriptionList>
		<DescriptionList tag="dd">{data.dashboard.solar_info.numSolarPanels}</DescriptionList>
	</div>
	<div class="flex flex-col pb-3">
		<DescriptionList tag="dt" class="mb-1">Total Installation Cost:</DescriptionList>
		<DescriptionList tag="dd"
			>{formatNumber(data.dashboard.solar_info.totalInstallationCost)} €</DescriptionList
		>
	</div>
	<div class="flex flex-col pb-3">
		<DescriptionList tag="dt" class="mb-1">Payback Period:</DescriptionList>
		<DescriptionList tag="dd"
			>{formatYears(data.dashboard.solar_info.paybackPeriod)}</DescriptionList
		>
	</div>

	<div class="flex flex-col pb-3">
		<DescriptionList tag="dt" class="mb-1">Lifetime:</DescriptionList>
		<DescriptionList tag="dd">{data.dashboard.solar_info.lifetime} years</DescriptionList>
	</div>

	<div class="flex flex-col pb-3">
		<DescriptionList tag="dt" class="mb-1">Efficiency:</DescriptionList>
		<DescriptionList tag="dd"
			>{formatNumber(data.dashboard.solar_info.efficiency)} %</DescriptionList
		>
	</div>
</div>

<style>
</style>
