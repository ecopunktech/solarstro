<script>
	import { page } from '$app/stores';
	import {
		SidebarDropdownWrapper,
		Sidebar,
		SidebarGroup,
		SidebarItem,
		SidebarWrapper
	} from 'flowbite-svelte';
	$: activeUrl = $page.url.pathname;
	let spanClass = 'flex-1 ml-3 whitespace-nowrap';
</script>

<div class="flex">
	<div class="mt-4 fixed  bg-white overflow-y-auto">
		<Sidebar>
			<SidebarWrapper>
				<SidebarGroup>
					<SidebarItem
						label="Informaction General"
						href="/dashboard"
						active={activeUrl === '/dashboard'}
					/>
					<SidebarDropdownWrapper label="Redimiento" active={activeUrl === '/components/sidebar'}>
						<SidebarItem
							label="Económico"
							href="/component/table"
							active={activeUrl === '/components/table'}
						/>
						<SidebarItem
							label="Energético"
							href="/component/table"
							active={activeUrl === '/components/table'}
						/>
					</SidebarDropdownWrapper>
					<SidebarDropdownWrapper
						label="Presupuesto"
						href="/component/table"
						active={activeUrl === '/components/table'}
					>
						<SidebarItem
							label="Materiales"
							href="/component/table"
							active={activeUrl === '/components/table'}
						/>
						<SidebarItem
							label="Mano de Obra"
							href="/component/table"
							active={activeUrl === '/components/table'}
						/>
						<SidebarItem
							label="Subtotales"
							href="/component/table"
							active={activeUrl === '/components/table'}
						/>
					</SidebarDropdownWrapper>
				</SidebarGroup>
			</SidebarWrapper>
		</Sidebar>
	</div>
	<div class="ml-60">
		<slot />
	</div>
</div>

<style>
</style>
