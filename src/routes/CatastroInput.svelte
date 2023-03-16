<script>
	// import {validateRC} from '$lib/catastro/catastro.ts';
	let provincia = '';
	let municipio = '';
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
</script>

<p>
	Escribe el numero de catastro de tu finca
	<input type="text" placeholder="Codigo catastral" bind:value={rc} />
	<button disabled={isValidRC(rc)} on:click={getCatrastroInfo}> Empezar </button>
	{#if direccion}
		<p>El numero de catastro es <strong>{rc}</strong></p>
		<p>La direccion es <strong>{direccion}</strong></p>
		<p>La superficie disponible es de <strong>{m2}</strong> m2</p>
		<!-- {#key unique}
			<Map bind:lat bind:lon />
		{/key} -->
		<img src={img_source} alt="figura" />
	{/if}
</p>
