<script>
	// import {validateRC} from '$lib/catastro/catastro.ts';
	export let rc = '';
	export let m2 = '';
	export let img_source = '';
	export let direccion = '';

	import { Card, Button, P } from 'flowbite-svelte';
	import { Img } from 'flowbite-svelte';
	import { Select, Label, Heading } from 'flowbite-svelte';
	import { Range } from 'flowbite-svelte';

	let selected = '';
	let budget = [
		{ value: 'min', name: 'Economico' },
		{ value: 'med', name: 'Medio' },
		{ value: 'max', name: 'Go Big' }
	];


	// create method to tranform m2 string value to number
	function m2ToNumber() {
		return parseFloat(m2);
	}
	// create a method return the m2 value depeding on the step value
	/**
	 * @param {number} [stepValue]
	 */
	function getM2(stepValue) {
		if (!stepValue) {
			stepValue = 75;
		}
		return (m2ToNumber() * stepValue) / 100;
	}
	$: href = `report/${rc}?percentage=${stepValue}&budget=${selected}`;

	function handleRangeChange() {
		href = `report/${rc}?percentage=${stepValue}&budget=${selected}`;
	}

	let stepValue = 75;
</script>

<div class="p-2 items-center justify-center">
	<div class="max-w-2xl w-full mx-auto">
		<Card color="primary" size="xl" class="grid md:grid-cols-2">
			<div class="p-2">
				<Img src={img_source} alt="image finca" class="rounded-lg" alignment="mx-auto" />
			</div>
			<div class="p-3 rounded-t-lg">
				<h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
					Datos de la parcela
				</h5>
				<p class="mb-3 font-normal text-gray-700 dark:text-gray-400 leading-tight">
					El numero de catastro es <strong>{rc}</strong>
				</p>
				<p class="mb-3 font-normal text-gray-700 dark:text-gray-400 leading-tight">
					La direccion es <strong>{direccion}</strong>
				</p>
				<p class="mb-3 font-normal text-gray-700 dark:text-gray-400 leading-tight">
					La superficie disponible es de <strong>{m2}</strong> m2
				</p>
			</div>
		</Card>
	</div>
	<div class="p-2 basis-1/3 md:basis-1/3 sm:basis-1/3">
		<div class="grid md:grid-cols-2 gap-4">
			<div>
				<Heading tag="h6">Cuantos m2 quieres usar?</Heading>
			</div>
			<div>
				<P class="md:text-right">{getM2(stepValue)} m2</P>
			</div>
		</div>

		<Range id="range-steps" min="0" max="100" bind:value={stepValue} step="25" on:change={handleRangeChange}/>
		<Heading class="mb-3" tag="h6">Selecciona el presupuesto que tienes disponible</Heading>
		<Select
			class="mb-3 mt-2 grid md:grid-cols-2"
			items={budget}
			bind:value={selected}
			on:change={handleRangeChange}
		/>

		<Button color="primary" {href}>
			Elegir <svg
				xmlns="http://www.w3.org/2000/svg"
				fill="none"
				viewBox="0 0 24 24"
				stroke-width="1.5"
				stroke="currentColor"
				class="w-5 h-5 ml-2"
				><path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="M13.5 4.5L21 12m0 0l-7.5 7.5M21 12H3"
				/></svg
			>
		</Button>
	</div>
</div>
