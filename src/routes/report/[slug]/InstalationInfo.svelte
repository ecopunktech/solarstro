<script lang="ts">
	import type { PageData } from './$types';
	import { page } from '$app/stores';

	import { Heading, DescriptionList, Img, List } from 'flowbite-svelte';

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

	function getAreaUsed(area: number, percentage: string | null) {
		if (percentage) {
			return (area * parseFloat(percentage)) / 100;
		}
		return area;
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
	$: areaUsed = $page.url.searchParams.get('percentage');
	export let data: PageData;
</script>

<div color="purple" class="p-2">
	<div class="md:flex">
		<List class="p-6" tag="dl" color="text-gray-900">
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
		</List>
		
		<List class="p-6" tag="dl" color="text-gray-900">
			<div class="flex flex-col pb-2">
				<DescriptionList tag="dt" class="mb-1">Cadastral code:</DescriptionList>
				<DescriptionList tag="dd">{data.dashboard.cadastral_code}</DescriptionList>
			</div>
			<div class="flex flex-col pb-2">
				<DescriptionList tag="dt" class="mb-1">Address:</DescriptionList>
				<DescriptionList tag="dd">{data.dashboard.address}</DescriptionList>
			</div>
			<div class="flex flex-col pb-3">
				<DescriptionList tag="dt" class="mb-1">Total Area:</DescriptionList>
				<DescriptionList tag="dd">{formatNumber(data.dashboard.area)} m²</DescriptionList>
			</div>
			<div class="flex flex-col pb-3">
				<DescriptionList tag="dt" class="mb-1">Area Used:</DescriptionList>
				<DescriptionList tag="dd"
					>{formatNumber(getAreaUsed(data.dashboard.area, areaUsed))} m²</DescriptionList
				>
			</div>
		</List>

		
	
	</div>
	
</div>

<style>
</style>
