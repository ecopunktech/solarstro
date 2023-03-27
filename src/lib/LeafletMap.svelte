<script>
    import { onMount, onDestroy } from 'svelte';
    import { browser } from '$app/environment';
	import { stringify } from 'postcss';

    let mapElement;
    let map;
    export  let lanString="";
    export  let lonString="";
    onMount(async () => {
        if(browser) {
            const leaflet = await import('leaflet');
            // set lan a to number
            const lan=Number(lanString);
            const lon = Number(lonString);


            map = leaflet.map(mapElement).setView([lan, lon], 15);

            leaflet.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
                attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
            }).addTo(map);

            leaflet.marker([lan, lon]).addTo(map)
                .bindPopup('Aqui esta tu parcela!!')
                .openPopup();
        }
    });

    onDestroy(async () => {
        if(map) {
            console.log('Unloading Leaflet map.');
            map.remove();
        }
    });
</script>


<main>
    <div bind:this={mapElement}></div>
</main>

<style>
    @import 'leaflet/dist/leaflet.css';
    main div {
        height: 400px;
        width: 400px;
    }
</style>