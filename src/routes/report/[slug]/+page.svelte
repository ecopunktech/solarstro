<script lang="ts">
	import InstalationInfo from './InstalationInfo.svelte';

	import ParcelInfo from './ParcelInfo.svelte';

	import TableChart from './TableChart.svelte';

	import type { PageData } from './$types';

	import { getMonthlyIncomeAndPowerGeneration } from '$lib/catastro/solar';
	export let data: PageData;
	let powerGeneration = getMonthlyIncomeAndPowerGeneration(data.dashboard).powerGeneration;
	let monthlyIncome = getMonthlyIncomeAndPowerGeneration(data.dashboard).monthlyIncome;

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
