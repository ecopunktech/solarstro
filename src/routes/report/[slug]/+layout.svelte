<script>
	import { Sidebar, SidebarGroup, SidebarItem, SidebarWrapper } from 'flowbite-svelte';
	import { page } from '$app/stores';
	// let spanClass = 'flex-1 ml-3 whitespace-nowrap';
	$: activeUrl = $page.url.pathname;
  </script>
  
  <div class="relative flex">
	<div class="sticky top-0 h-screen w-64 overflow-y-auto">
	  <Sidebar>
		<SidebarWrapper>
		  <SidebarGroup>
			<SidebarItem label="Parcel Information" href='#parcel' active={activeUrl === '#parcel'} />
			<SidebarItem label="Installation" href='#installation' active={activeUrl === '#installation'} />
			<SidebarItem label="Power" href='#power' active={activeUrl === '#power'} />
			<SidebarItem label="Income" href='#income' active={activeUrl === '#income'} />
		  </SidebarGroup>
		</SidebarWrapper>
	  </Sidebar>
	</div>
  
	<div class="flex-1">
	  <slot></slot>
	</div>
  </div>
  
  <style>
	/* add some styles to the sidebar container */
	.sticky {
	  position: sticky;
	  left: 0;
	}
  </style>
  