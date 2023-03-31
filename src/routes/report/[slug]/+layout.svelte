<script>
	import { ChartPie, Document, CurrencyEuro, Bolt, HandRaised,
		WrenchScrewdriver,PresentationChartBar } from 'svelte-heros-v2';
	import {
		SidebarDropdownWrapper,
		Sidebar,
		SidebarGroup,
		SidebarItem,
		SidebarWrapper
	} from 'flowbite-svelte';
	// this function changes the value of index
	// it takes in a parameter of index
	// it sets the context of myVariable to the value of the parameter
	import message from '$lib/message';
	function changeIndex(index) {
		$message = index;
	}
	const activeClass =
		'flex items-center p-2 text-base font-normal text-purple-900 bg-purple-400  rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700';
	const nonActiveClass =
		'flex items-center p-2 text-base font-normal text-gray-900 bg-purple-200 dark:bg-gray-700 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700';
</script>

<div class="flex">
	<div class="fixed hidden sm:block overflow-y-auto">
		<Sidebar color="white">
			<SidebarWrapper class="'overflow-y-auto py-4 px-3 bg-white rounded dark:bg-gray-800">
				<SidebarGroup>
					<SidebarItem
						label="Informaction General"
						active={$message === 'general'}
						on:click={() => changeIndex('general')}
						{activeClass}
					>
						<svelte:fragment slot="icon">
							<Document size="25" />
						</svelte:fragment>
					</SidebarItem>
					<SidebarDropdownWrapper label="Redimiento">
						<svelte:fragment slot="icon">
							<ChartPie size="25" />
						</svelte:fragment>
						<SidebarItem
							label="Económico"
							active={$message === 'economic'}
							{activeClass}
							on:click={() => changeIndex('economic')}
						>
							<svelte:fragment slot="icon">
								<CurrencyEuro size="25" />
							</svelte:fragment>
						</SidebarItem>
						<SidebarItem
							label="Energético"
							active={$message === 'energy'}
							{activeClass}
							on:click={() => changeIndex('energy')}
						>
							<svelte:fragment slot="icon">
								<Bolt size="25" />
							</svelte:fragment></SidebarItem
						>
					</SidebarDropdownWrapper>
					<SidebarDropdownWrapper label="Presupuesto">
						<svelte:fragment slot="icon">
							<PresentationChartBar size="25" />
						</svelte:fragment>
						<SidebarItem
							label="Materiales"
							active={$message === 'materials'}
							{activeClass}
							on:click={() => changeIndex('materials')}
						>
							<svelte:fragment slot="icon">
								<WrenchScrewdriver size="25" />
							</svelte:fragment></SidebarItem
						>
						<SidebarItem
							label="Mano de Obra"
							active={$message === 'handwork'}
							{activeClass}
							on:click={() => changeIndex('handwork')}
						>
							<svelte:fragment slot="icon">
								<HandRaised size="25" />
							</svelte:fragment></SidebarItem
						>
					</SidebarDropdownWrapper>
				</SidebarGroup>
			</SidebarWrapper>
		</Sidebar>
	</div>
	<div class="md:ml-64 w-full">
		<slot />
	</div>
</div>

<style>
</style>
