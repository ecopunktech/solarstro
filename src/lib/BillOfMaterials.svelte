<script>
	import {
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell,
        Heading,
	} from 'flowbite-svelte';
	let data = {
		items: [
			{
				item: 'Placas fotovoltaicas',
				amount: 10,
				unit_price: 100,
				total_price: 1000,
				link: 'https://example.com/buy/Placas%20fotovoltaicas',
                info: 'https://example.com/info/Placas%20fotovoltaicas'
			},
			{
				item: 'Protecciones de corriente continua',
				amount: 5,
				unit_price: 50,
				total_price: 250,
				link: 'https://example.com/buy/Protecciones%20de%20corriente%20continua',
                info: 'https://example.com/info/Protecciones%20de%20corriente%20continua'
			},
			{
				item: 'Inversores',
				amount: 2,
				unit_price: 500,
				total_price: 1000,
				link: 'https://example.com/buy/Inversores',
                info: 'https://example.com/info/Inversores'
			},
			{
				item: 'Protecciones de corriente alterna',
				amount: 5,
				unit_price: 40,
				total_price: 200,
				link: 'https://example.com/buy/Protecciones%20de%20corriente%20alterna',
                info: 'https://example.com/info/Protecciones%20de%20corriente%20alterna'

			},
			{
				item: 'Contadores',
				amount: 2,
				unit_price: 20,
				total_price: 40,
				link: 'https://example.com/buy/Contadores',
                info: 'https://example.com/info/Contadores'

			},
			{
				item: 'Sistemas de seguimiento y monitoreo',
				amount: 1,
				unit_price: 1000,
				total_price: 1000,
				link: 'https://example.com/buy/Sistemas%20de%20seguimiento%20y%20monitoreo',
                info: 'https://example.com/info/Sistemas%20de%20seguimiento%20y%20monitoreo'
			},
			{
				item: 'Cableado',
				amount: 500,
				unit_price: 1,
				total_price: 500,
				link: 'https://example.com/buy/Cableado',
                info: 'https://example.com/info/Cableado'
			},
			{
				item: 'Instalación de placas, obras necesarias de soporte',
				amount: 1,
				unit_price: 2000,
				total_price: 2000,
				link: 'https://example.com/buy/Instalación%20de%20placas%2C%20obras%20necesarias%20de%20soporte',
                info: 'https://example.com/info/Instalación%20de%20placas%2C%20obras%20necesarias%20de%20soporte'
			},
			{
				item: 'Instalación del sistema, conexionados, y conexión a red (casetas, transformadores, etc.)',
				amount: 1,
				unit_price: 5000,
				total_price: 5000,
				link: 'https://example.com/buy/Instalación%20del%20sistema%2C%20conexionados%2C%20y%20conexión%20a%20red%20(casetas%2C%20transformadores%2C%20etc.)',
                info: 'https://example.com/info/Instalación%20del%20sistema%2C%20conexionados%2C%20y%20conexión%20a%20red%20(casetas%2C%20transformadores%2C%20etc.)'
			},
			{
				item: 'Elementos de seguridad (vallado, alarmas, conexionado por GPS)',
				amount: 1,
				unit_price: 3000,
				total_price: 3000,
				link: 'https://example.com/buy/Elementos%20de%20seguridad%20(vallado%2C%20alarmas%2C%20conexionado%20por%20GPS)',
                info: 'https://example.com/info/Elementos%20de%20seguridad%20(vallado%2C%20alarmas%2C%20conexionado%20por%20GPS)'
			},
			{ item: 'Mantenimiento del sistema',
                amount: 1,
                unit_price: 1000,
                total_price: 1000,
                link: 'https://example.com/buy/Mantenimiento%20del%20sistema',
                info: 'https://example.com/info/Mantenimiento%20del%20sistema'
         }
		]
	};
</script>

<Heading id="parcel" color="text-purple-400 dark:text-white" tag="h2" class="mb-4"
		>Presupuesto</Heading
	>
<Table hoverable={true} shadow >
	<TableHead class="bg-purple-300">
		<TableHeadCell>Componente</TableHeadCell>
		<TableHeadCell>Cantidad</TableHeadCell>
        <TableHeadCell>Precio Unidad</TableHeadCell>
        <TableHeadCell>Precio Total</TableHeadCell>
        <TableHeadCell>Comprar</TableHeadCell>
	</TableHead>
	<TableBody>
        {#each data.items as item}
		<TableBodyRow class="hover:bg-purple-100">
			<TableBodyCell >{item.item}</TableBodyCell>
			<TableBodyCell >{item.amount}</TableBodyCell>
            <TableBodyCell >{item.unit_price}</TableBodyCell>
            <TableBodyCell >{item.total_price}</TableBodyCell>
            <TableBodyCell ><a href="{item.link}" class="font-medium text-purple-600 hover:underline dark:text-purple-500">Comprar</a></TableBodyCell>
		</TableBodyRow>
        
        {/each}
       
	</TableBody>
    <tfoot>
        <tr class="font-semibold bg-purple-200 text-gray-900 dark:text-white">
          <th scope="row" class="py-3 px-6 text-base">Total</th>
          <td class="py-3 px-6"></td>
          <td class="py-3 px-6"></td>
          <td class="py-3 px-6">21,000</td>
            <td class="py-3 px-6"></td>
        </tr>
      </tfoot>
</Table>
