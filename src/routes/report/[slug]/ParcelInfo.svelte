<script lang="ts">
	import type { PageData } from './$types';
	import LeafletMap from '$lib/LeafletMap.svelte';
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

<div color="purple" class="p-2">
	<Heading id="parcel" color="text-purple-400 dark:text-white" tag="h2" class="ml-2 mb-1"
		>Infomacion General</Heading
	>
	<div class="md:flex p-2">
		<div class="flex ml-2 rounded-xl shadow-lg">
			<LeafletMap lonString={data.dashboard.longitude} lanString={data.dashboard.latitude} />
		</div>
		<div class="flex p-2 ml-6 rounded-xl shadow-lg">
			{@html data.svg}
		</div>
	</div>
</div>

<style>
</style>
