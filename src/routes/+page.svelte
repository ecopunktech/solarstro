<script>
	import CatastroInput from './CatastroInput.svelte';
</script>
<svelte:head>
	<title>Solarstro</title>
	<meta name="description" content="Calculadora para proyecto de huerto solar" />
</svelte:head>
<section>

<h1>
    Bienvenido a <strong>Solarstro</strong>
</h1>
<CatastroInput />
<button
  type="button"
  on:click={() => {
    throw new Error("Sentry Frontend Error");
  }}
>
  Throw error
</button>
</section>

