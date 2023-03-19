<!-- Header.svelte -->
<script>
	import { Navbar, NavBrand, NavLi, NavUl, NavHamburger } from 'flowbite-svelte';
    import "../app.postcss"
</script>

<Navbar let:hidden let:toggle>
	<NavBrand href="/">
		<img
			src="/solarstro-logo/cover.png"
			class="mr-3 h-9 sm:h-16"
			alt="Flowbite Logo"
		/>
	</NavBrand>
	<NavHamburger color="purple"  on:click={toggle} />
	<NavUl activeClass="text-white bg-purple-700 md:bg-transparent md:text-purple-700 md:dark:text-white dark:bg-purple-600 md:dark:bg-transparent" nonActiveClass="text-gray-700 hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-purple-700 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent" {hidden}>
		<NavLi color="purple" href="/" active={true}>Home</NavLi>
		<NavLi color="purple" href="/about">About</NavLi>
		<NavLi color="purple" href="/docs">Docs</NavLi>
	</NavUl>
</Navbar>
<slot></slot>
