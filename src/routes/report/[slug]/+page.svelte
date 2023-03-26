<script lang="ts">
	import type { PageData } from '../$types';
	import type { SolarPanelData } from './+page.server';
	import { List, Heading, DescriptionList } from 'flowbite-svelte';
	import {
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';

	function roundNumber(num: string) {
		return parseFloat(num).toFixed(2);
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
	function getMonthlyIncomeAndPowerGeneration(solarInfo: SolarPanelData) {
		const powerGeneration = Object.entries(solarInfo.solar_info.powerGeneration)
			.sort(
				(a, b) =>
					new Date(Date.parse(`01 ${a[0]}`)).getTime() -
					new Date(Date.parse(`01 ${b[0]}`)).getTime()
			)
			.reduce((obj, [key, value]) => ({ ...obj, [key]: value }), {});

		const monthlyIncome = Object.entries(solarInfo.solar_info.monthlyIncome)
			.sort(
				(a, b) =>
					new Date(Date.parse(`01 ${a[0]}`)).getTime() -
					new Date(Date.parse(`01 ${b[0]}`)).getTime()
			)
			.reduce((obj, [key, value]) => ({ ...obj, [key]: value }), {});

		return { powerGeneration, monthlyIncome };
	}
	export let data: PageData;
	let powerGeneration = getMonthlyIncomeAndPowerGeneration(data.dashboard).powerGeneration
	let monthlyIncome = getMonthlyIncomeAndPowerGeneration(data.dashboard).monthlyIncome
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
	<div class="my-8 p-6">
		<div class="my-8 p-6 rounded-xl shadow-lg">
			<Heading id="parcel" color="text-purple-400 dark:text-white" tag="h2" class="mb-4"
				>Parcel Information</Heading
			>
			<List tag="dl" color="text-gray-900 dark:text-white">
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
					<DescriptionList tag="dd">{data.dashboard.area.toFixed(2)} m²</DescriptionList>
				</div>
				<div class="flex flex-col pb-3">
					<DescriptionList tag="dt" class="mb-1">Area Used:</DescriptionList>
					<DescriptionList tag="dd">{getAreaUsed(data.dashboard.area,areaUsed).toFixed(2)} m²</DescriptionList>
				</div>
			</List>
		</div>
		<div color="purple" class="my-8 p-6 rounded-xl shadow-lg">
			<Heading id="installation" color="text-purple-400 dark:text-white" tag="h2" class="mb-4"
				>Installation Info</Heading
			>
			<div class="flex flex-col pb-3">
				<DescriptionList tag="dt" class="mb-1">Total Power:</DescriptionList>
				<DescriptionList tag="dd"
					>{data.dashboard.solar_info.totalPower.toFixed(2)} kW</DescriptionList
				>
			</div>
			<div class="flex flex-col pb-3">
				<DescriptionList tag="dt" class="mb-1">Total Panels:</DescriptionList>
				<DescriptionList tag="dd">{data.dashboard.solar_info.numSolarPanels}</DescriptionList>
			</div>
			<div class="flex flex-col pb-3">
				<DescriptionList tag="dt" class="mb-1">Total Installation Cost:</DescriptionList>
				<DescriptionList tag="dd"
					>{data.dashboard.solar_info.totalInstallationCost} €</DescriptionList
				>
			</div>
			<div class="flex flex-col pb-3">
				<DescriptionList tag="dt" class="mb-1">Payback Period:</DescriptionList>
				<DescriptionList tag="dd"
					>{formatYears(data.dashboard.solar_info.paybackPeriod)} years</DescriptionList
				>
			</div>

			<div class="flex flex-col pb-3">
				<DescriptionList tag="dt" class="mb-1">Lifetime:</DescriptionList>
				<DescriptionList tag="dd">{data.dashboard.solar_info.lifetime} years</DescriptionList>
			</div>

			<div class="flex flex-col pb-3">
				<DescriptionList tag="dt" class="mb-1">Efficiency:</DescriptionList>
				<DescriptionList tag="dd">{data.dashboard.solar_info.efficiency} %</DescriptionList>
			</div>
		</div>
		<div color="purple" class="my-8 p-6 rounded-xl shadow-lg">
			<Heading id="power" color="text-purple-400 dark:text-white" tag="h2" class="mb-4"
				>Monthly Power</Heading
			>
			<Table>
				<TableHead>
					<TableHeadCell>Month</TableHeadCell>
					<TableHeadCell>Kw</TableHeadCell>
				</TableHead>
				<TableBody>
					{#each Object.entries(powerGeneration) as [month, value]}
						<TableBodyRow>
							<TableBodyCell>{month}</TableBodyCell>
							<TableBodyCell>{value.toFixed(0)} kW</TableBodyCell>
						</TableBodyRow>
					{/each}
				</TableBody>
			</Table>
		</div>

		<div color="purple" class="my-8 p-6 rounded-xl shadow-lg">
			<Heading id="income" color="text-purple-400 dark:text-white" tag="h2" class="mb-4"
				>Monthly Income</Heading
			>
			<Table>
				<TableHead>
					<TableHeadCell>Month</TableHeadCell>
					<TableHeadCell>Euros</TableHeadCell>
				</TableHead>
				<TableBody>
					{#each Object.entries(monthlyIncome) as [month, value]}
						<TableBodyRow>
							<TableBodyCell>{month}</TableBodyCell>
							<TableBodyCell>{value.toFixed(0)} €</TableBodyCell>
						</TableBodyRow>
					{/each}
				</TableBody>
			</Table>
		</div>
		<!-- Add more fields as needed -->
	</div>
</div>

<style>
</style>
