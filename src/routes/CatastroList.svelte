<script>
	// import {validateRC} from '$lib/catastro/catastro.ts';
	export let rc = '';
	export let m2 = '';
	export let img_source = '';
	export let direccion = '';

	import { Card, Button } from 'flowbite-svelte';
    import { Img } from 'flowbite-svelte';

</script>

<div class="p-2 flex items-center justify-center">
		<Card color="primary" horizontal size="xl" class="grid md:grid-cols-2">
            <div class="p-2">
            <Img src={img_source} alt="image finca" class="rounded-t-lg" />
        </div>
            <div class="p-2 rounded-t-lg">
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
			<Button color="primary" href={`report/${rc}`}>
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
		</Card>
	</div>

