<script lang="ts">
	import Chart from 'chart.js/auto';
	import { onMount, onDestroy } from 'svelte';
	export let data: any;
    export let title: string="chart";
	let chart: Chart<"bar", unknown[], string>;

	onMount(() => {
		const canvas = document.getElementById(title) as HTMLCanvasElement;
        if (!canvas) {
            return;
        }
		const ctx = canvas.getContext('2d');
        if (!ctx) {
            return;
        }
		const labels = Object.keys(data);
        const chartData = Object.values(data);
		const chart = new Chart(ctx, {
			type: 'bar',
			data: {
				labels: labels,
				datasets: [
					{
						label: title,
						data: chartData,
						backgroundColor: 'rgba(255, 178, 51, 0.2)', // orange background color
						borderColor: 'rgba(255, 165, 0, 1)', // orange border color
						borderWidth: 1,
                        borderRadius: 5 
					}
				]
			},
			options: {
				scales: {
					y: {
						beginAtZero: true,
				
					},
					x: {
						beginAtZero: true,
						grid: {
							display: false
						}
					}
				}
			}
		});
	});

	onDestroy(() => {
        if (chart == null) {
            return;
        }
		chart.destroy();
	});
</script>

<canvas id={title} />
