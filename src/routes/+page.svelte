<script>
	import CatastroInput from './CatastroInput.svelte';
	import BillOfMaterials from '$lib/BillOfMaterials.svelte';

</script>

<svelte:head>
	<title>Solarstro</title>
	<meta name="description" content="Calculadora para proyecto de huerto solar" />
	<link
		rel="stylesheet"
		href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css"
	/>
</svelte:head>


<section class="hero">

	<div class="container">

		<div class="p-6">
		<BillOfMaterials />
	</div>
	</div>
</section>

