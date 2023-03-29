<script lang="ts">
	import InstalationInfo from './InstalationInfo.svelte';

	import ParcelInfo from './ParcelInfo.svelte';

	import TableChart from './TableChart.svelte';

	import type { PageData } from './$types';
	import BillOfMaterials from '$lib/BillOfMaterials.svelte';

	import { getMonthlyIncomeAndPowerGeneration } from '$lib/catastro/solar';
	export let data: PageData;
	let powerGeneration = getMonthlyIncomeAndPowerGeneration(data.dashboard).powerGeneration;
	let monthlyIncome = getMonthlyIncomeAndPowerGeneration(data.dashboard).monthlyIncome;
</script>

<div>
	<div>
		<div class="p-2">
			<ParcelInfo {data} />
		</div>
		<div class="p-2">
			<InstalationInfo {data} />
		</div>
		
		<!-- <div>
			<div class="p-2 w-full">
				<TableChart data={powerGeneration} title="Energia Generada" titleChart="kWh" unit="kWh" />
			</div>
			<div class="p-2 w-full">
				<TableChart data={monthlyIncome} title="Ganancias Mes" titleChart="Euros" />
			</div>
		</div>
		<div class="xl:flex">
			<div class="p-2 w-full">
				<BillOfMaterials />
			</div>
		</div> -->
		<!-- Add more fields as needed -->
	</div>
</div>

<style>
</style>
