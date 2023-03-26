<script lang="ts">

	import type { PageData } from './$types';

	import { List, Heading, DescriptionList } from 'flowbite-svelte';

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

	export let data: PageData;
	import { page } from '$app/stores';
	// let spanClass = 'flex-1 ml-3 whitespace-nowrap';
	$: areaUsed = $page.url.searchParams.get('percentage');
	function getAreaUsed(area: number, percentage: string | null) {
		if (percentage) {
			return (area * parseFloat(percentage)) / 100;
		}
		return area;
	}
</script>

<div class="p-6 rounded-xl shadow-lg">
	<Heading id="parcel" color="text-purple-400 dark:text-white" tag="h2" class="mb-4"
		>Parcel Information</Heading
	>
	<List tag="dl" color="text-gray-900">
		<div class="flex flex-col pb-3">
			<DescriptionList tag="dt" class="mb-1">Cadastral code:</DescriptionList>
			<DescriptionList tag="dd">{data.dashboard.cadastral_code}</DescriptionList>
		</div>
		<div class="flex flex-col pb-3">
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

<style>
</style>
