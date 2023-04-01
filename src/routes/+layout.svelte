<!-- Header.svelte -->
<script>
	import { Navbar, NavBrand, NavLi, NavUl, NavHamburger } from 'flowbite-svelte';
	import {Sun} from 'svelte-heros-v2'
    import "../app.postcss"

</script>

<Navbar let:hidden let:toggle>
	
	<NavBrand href="/">
		<span class="p-2 text-orange-400 text-xl font-semibold dark:text-white">
			Solar
		  </span>
		<Sun class="ml-2" size="35" />
		
	</NavBrand>
	
	<NavHamburger color="orange"  on:click={toggle} />
	<NavUl activeClass="text-white bg-orange-700 md:bg-transparent md:text-orange-700 md:dark:text-white dark:bg-orange-600 md:dark:bg-transparent" nonActiveClass="text-gray-700 hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-orange-700 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent" {hidden}>
		<NavLi color="orange" href="/" active={true}>Home</NavLi>
		<NavLi color="orange" href="/about">About</NavLi>
		<NavLi color="orange" href="/docs">Docs</NavLi>
	</NavUl>
	
</Navbar>

<slot></slot>
