<script>
  import CatastroList from './CatastroList.svelte';

	// import {validateRC} from '$lib/catastro/catastro.ts';
	let rc = '';
	let m2 = '';
	let lat = 0;
	let lon = 0;
	let img_source = '';

	// let catastroData;
	let direccion = '';
	/**
	 * @param {any[]} subparcelas
	 */
	function getTotalSuperficie(subparcelas) {
		return subparcelas.reduce((acc, item) => {
			return +acc + +item.superficie;
		}, '');
	}
	async function getCatrastroInfo() {
		const url = `/catastro?rc=${rc}`;
		const response = await fetch(url, {
			method: 'GET',
			headers: {
				'content-type': 'application/json'
			}
		});
		const catastroData = await response.json();
		if (catastroData.errorId) {
			alert(catastroData.error);
			unique = {};
			return;
		} else {
			direccion = catastroData.direccion;
			m2 = getTotalSuperficie(catastroData.subparcelas);
			lat = catastroData.coor.y;
			lon = catastroData.coor.x;
			console.log(catastroData);
			console.log(catastroData.direccion);
			console.log(direccion);
			const pID = rc.slice(0, 2);
			const mID = rc.slice(2, 5);
			img_source = `http://www1.sedecatastro.gob.es/Cartografia/GeneraGraficoParcela.aspx?del=${pID}&mun=${mID}&refcat=${rc}&AnchoPixels=300&AltoPixels=300`;
		}
		unique = {};
	}
	let unique = {};

	function restart() {
		unique = {}; // every {} is unique, {} === {} evaluates to false
	}
	const regex = /^(\d{2})(\d{3})([A-Z])(\d{3})(\d{5})$/;
	/**
	 * @param {string} rc1
	 */
	function isValidRC(rc1) {
		return !regex.test(rc1);
	}
	import { Button, Search, ImagePlaceholder } from 'flowbite-svelte';
</script>

<div>
	{#if !direccion}
	<div color="purple" class="my-8 p-6 rounded-xl shadow-lg">
		<form color="purple" id="search" on:submit={getCatrastroInfo}>
		<Search placeholder="Escribe el numero de catastro" bind:value={rc}>
			<Button color="primary" type="submit" disabled={isValidRC(rc) }>Search</Button>
		  </Search>
		</form>
	</div>
	<div class="p-6 rounded-xl shadow-lg">
		<ImagePlaceholder>
		</ImagePlaceholder>
	
	</div>
	{/if}
	{#if direccion}
		<CatastroList rc={rc} img_source={img_source} m2={m2} direccion={direccion}></CatastroList>
	{/if}
	
</div>
