<script>
	import Header from './Header.svelte';
	import CatastroInput from './CatastroInput.svelte';

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
	<Header />

	<div class="container">

		<div class="p-6">
		<CatastroInput />
	</div>
	</div>
</section>

