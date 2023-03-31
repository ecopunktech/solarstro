<script lang="ts">
	import InstalationInfo from './InstalationInfo.svelte';

	import ParcelInfo from './ParcelInfo.svelte';

	import TableChart from '$lib/TableChart.svelte';

	import type { PageData } from './$types';
	import BillOfMaterials from '$lib/BillOfMaterials.svelte';

	import { getMonthlyIncomeAndPowerGeneration } from '$lib/catastro/solar';
	export let data: PageData;
	let powerGeneration = getMonthlyIncomeAndPowerGeneration(data.dashboard).powerGeneration;
	import message from '$lib/message';
	import Economic from '$lib/Economic.svelte';
</script>

<div class="">
	<div class="ml-6">
		{#if $message === 'general'}
			<div class="p-2">
				<ParcelInfo {data} />
			</div>
			<div class="p-2">
				<InstalationInfo {data} />
			</div>
		{/if}
		{#if $message === 'economic'}
			<div class="ml-6 mr-2">
				<Economic {data} />
			</div>
		{/if}
		{#if $message === 'energy'}
			<div class="ml-6 mr-2">
				<TableChart data={powerGeneration} title="Energia Generada" titleChart="kWh" unit="kWh" />
			</div>
		{/if}
		{#if $message === 'materials'}
			<div class="mr-2">
				<BillOfMaterials />
			</div>
		{/if}
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
