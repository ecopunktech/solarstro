<script>
	import { page } from '$app/stores';
	$: activeUrl = $page.url.pathname;
  </script>
  
  <slot></slot>
  
  <style>
  </style>
  