<script lang="ts">
	import InstalationInfo from './InstalationInfo.svelte';

	import ParcelInfo from './ParcelInfo.svelte';

	import TableChart from './TableChart.svelte';

	import type { PageData } from './$types';

	import { getMonthlyIncomeAndPowerGeneration } from '$lib/catastro/solar';
	export let data: PageData;
	let powerGeneration = getMonthlyIncomeAndPowerGeneration(data.dashboard).powerGeneration;
	let monthlyIncome = getMonthlyIncomeAndPowerGeneration(data.dashboard).monthlyIncome;
	import { page } from '$app/stores';
	// let spanClass = 'flex-1 ml-3 whitespace-nowrap';
	$: areaUsed = $page.url.searchParams.get('percentage');
	function getAreaUsed(area: number, percentage: string) {
		if (percentage) {
			return (area * parseFloat(percentage)) / 100;
		}
		return area;
	}
</script>

<div>
	<div class="p-6">
		<ParcelInfo {data} />
		<InstalationInfo {data} />
		<TableChart
			data={powerGeneration}
			title="Montly Power"
			titleChart="kWh"
			unit="kWh"
		/>
		<TableChart data={monthlyIncome} title="Montly Incoming" titleChart="Euros" />
		<!-- Add more fields as needed -->
	</div>
</div>

<style>
</style>
