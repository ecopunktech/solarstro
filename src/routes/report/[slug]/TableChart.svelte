<script lang="ts">
	import { Heading } from 'flowbite-svelte';
    import BarChart from './Chart.svelte';
	import {
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';
    import { Toggle } from 'flowbite-svelte'
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

<div color="purple" class="my-8 p-6 rounded-xl shadow-lg">
    <div class="flex">
	<Heading id="income" color="text-purple-400 dark:text-white" tag="h2" class="mb-4"
		>{title}</Heading
	> 
    <Toggle color="purple" checked on:change={changeShow} >{showLabel}</Toggle>
</div>

	{#if show}
		<Table>
			<TableHead>
				<TableHeadCell>Month</TableHeadCell>
				<TableHeadCell>Euros</TableHeadCell>
			</TableHead>
			<TableBody>
				{#each Object.entries(data) as [month, value]}
					<TableBodyRow>
						<TableBodyCell>{month}</TableBodyCell>
						<TableBodyCell>{formatNumber(value)} {unit}</TableBodyCell>
					</TableBodyRow>
				{/each}
			</TableBody>
		</Table>
	{/if}
	{#if !show}
		<div>
			<BarChart data={data} title={titleChart}/>
		</div>
	{/if}
</div>

<style>
</style>
