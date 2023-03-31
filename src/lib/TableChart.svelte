<script lang="ts">
	import { Heading } from 'flowbite-svelte';
	import BarChart from '$lib/Chart.svelte';
	import {
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';
	import { Toggle } from 'flowbite-svelte';
	function formatNumber(num: number | null): string {
		const options = {
			useGrouping: true,
			maximumFractionDigits: 2,
			minimumFractionDigits: 2,
			currencyDisplay: 'symbol',
			style: 'decimal'
		};
		if (num === null) {
			return '0';
		}

		return num.toLocaleString('es-ES', options);
	}

	export let data: any;
	export let show = false;
	function changeShow() {
		showLabel = changeShowText();
		show = !show;
	}
	function changeShowText() {
		if (show) {
			return 'Chart';
		}
		return 'Table';
	}
	let showLabel = 'Chart';
	export let title = 'Monthly Income';
	export let titleChart = 'Power Generation';
	export let unit = '€';
</script>

<div color="purple" class="p-6">
	<div class="flex">
		<Heading id="income" color="text-purple-400 dark:text-white" tag="h3" class="mb-4"
			>{title}</Heading
		>
		<Toggle color="purple" checked on:change={changeShow}>{showLabel}</Toggle>
	</div>
	<div>
		{#if show}
			<Table hoverable={true} shadow>
				<TableHead class="bg-purple-300">
					<TableHeadCell>Month</TableHeadCell>
					<TableHeadCell>Euros</TableHeadCell>
				</TableHead>
				<TableBody>
					{#each Object.entries(data) as [month, value]}
						<TableBodyRow class="hover:bg-purple-100">
							<TableBodyCell>{month}</TableBodyCell>
							<TableBodyCell >{formatNumber(value)} {unit}</TableBodyCell>
						</TableBodyRow>
					{/each}
				</TableBody>
			</Table>
		{/if}
		{#if !show}
		<div color="purple" class="p-6 rounded-xl shadow-lg">

			<BarChart {data} title={titleChart} />
			</div>
		{/if}
	</div>
</div>

<style>
</style>
