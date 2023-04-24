<script lang="ts">
	import type { PageData } from './$types';
	import LeafletMap from '$lib/LeafletMap.svelte';
	import { Heading } from 'flowbite-svelte';
	export let data: PageData;
</script>

<div color="orange" class="p-2">
	<Heading id="parcel" color="text-orange-400 dark:text-white" tag="h2" class="ml-2 mb-1">
		Infomación General
	</Heading>
	<div class="md:flex p-2">
		<div class="flex flex-col items-center justify-center ml-2 rounded-xl shadow-lg">
			<LeafletMap lonString={data.dashboard.longitude} lanString={data.dashboard.latitude} />
		</div>
		<div class="flex flex-col items-center justify-center p-2 ml-6 rounded-xl shadow-lg">
			{@html data.svg}
		</div>
	</div>
</div>

<style>
</style>
