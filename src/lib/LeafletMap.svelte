<script>
    import { onMount, onDestroy } from 'svelte';
    import { browser } from '$app/environment';

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

            var marker = leaflet.marker([lan, lon]).addTo(map)
                .bindPopup('Aqui esta tu parcela!!')
                .setIcon(leaflet.icon({
                    iconUrl: 'https://cdn.rawgit.com/pointhi/leaflet-color-markers/master/img/marker-icon-2x-violet.png',
                    shadowUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/0.7.7/images/marker-shadow.png',
                    iconSize: [25, 41],
                    iconAnchor: [12, 41],
                    popupAnchor: [1, -34],
                    shadowSize: [41, 41]
                }))
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


<main class="flex items-center justify-center">
    <div bind:this={mapElement}></div>
</main>

<style>
    @import 'leaflet/dist/leaflet.css';
    main div {
        align-content: center;
        height: 420px;
        width: 320px;
    }
</style>