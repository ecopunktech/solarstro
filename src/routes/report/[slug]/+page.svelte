<script lang="ts">
	import InstalationInfo from './InstalationInfo.svelte';

	import ParcelInfo from './ParcelInfo.svelte';

	import TableChart from './TableChart.svelte';

	import type { PageData } from './$types';

	import { getMonthlyIncomeAndPowerGeneration } from '$lib/catastro/solar';
	import { List, Heading, DescriptionList } from 'flowbite-svelte';
	import {
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';

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
			titleChart="Power Generation"
			unit="kWh"
		/>
		<TableChart data={monthlyIncome} title="Montly Incoming" titleChart="Euros" />
		<!-- Add more fields as needed -->
	</div>
</div>

<style>
</style>
