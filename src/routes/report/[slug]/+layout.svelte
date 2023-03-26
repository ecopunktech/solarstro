<script>
	import { Sidebar, SidebarGroup, SidebarItem, SidebarWrapper } from 'flowbite-svelte';
	import { page } from '$app/stores';
	// let spanClass = 'flex-1 ml-3 whitespace-nowrap';
	$: activeUrl = $page.url.pathname;
  </script>
  
  <slot></slot>
  
  <style>
	/* add some styles to the sidebar container */
	.sticky {
	  position: sticky;
	  left: 0;
	}
  </style>
  