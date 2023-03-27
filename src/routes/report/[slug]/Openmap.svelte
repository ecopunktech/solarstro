<script lang="ts">
	import { onMount } from 'svelte';
	import { afterUpdate } from 'svelte';
	import L from 'leaflet';
	import { onMount as domOnMount } from 'svelte';
	import { get } from 'svelte/store';
	import { tick } from 'svelte';
	import { onDestroy } from 'svelte';
	import { setContext } from 'svelte';

	export let coordinates = [-3.70379, 40.416775];

	// Define the map options
	const mapOptions = {
		center: [coordinates[0], coordinates[1]],
		zoom: 15
	};

	// Create the map instance
	let map;
	let mapContainer;

	onMount(() => {
		// Initialize the map
		map = L.map(mapContainer, mapOptions);
		setContext('leaflet-map', map);

		// Add the tile layer (e.g. OpenStreetMap)
		L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
			attribution: '&copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors',
			maxZoom: 18
		}).addTo(map);

		// Add a marker at the parcel coordinates
		L.marker(coordinates).addTo(map);
	});

	let tileLayer;

	afterUpdate(() => {
		tileLayer = get(mapContainer).querySelector('.leaflet-tile-pane');
	});

	domOnMount(() => {
		// Force a reflow to ensure tiles are rendered
		tileLayer && tick().then(() => tileLayer.offsetWidth);
	});

	onDestroy(() => {
		// Clean up the map
		map.remove();
	});
</script>

<svelte:head>
	<link
		rel="stylesheet"
		href="https://unpkg.com/leaflet@1.6.0/dist/leaflet.css"
		integrity="sha512-xwE/Az9zrjBIphAcBb3F6JVqxf46+CDLwfLMHloNu6KEQCAWi6HcDUbeOfBIptF7tcCzusKFjFw2yuvEpDL9wQ=="
		crossorigin=""
	/>

	<script
		src="https://unpkg.com/leaflet@1.6.0/dist/leaflet.js"
		integrity="sha512-gZwIG9x3wUXg2hdXF6+rVkLF/0Vi9U8D2Ntg4Ga5I5BZpVkVxlJWbSQtXPSiUTtC0TjtGOmxa1AJPuV0CPthew=="
		crossorigin=""
	>
	</script>
</svelte:head>
<div bind:this={mapContainer} style="height: 480px;" />
