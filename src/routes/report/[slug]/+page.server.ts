import type { PageServerLoad } from './$types';
import type { SolarPanelData } from '$lib/catastro/solar';

export const load = (async ({ params, url }) => {
	const percentage = url.searchParams.get('percentage');
	const budget = url.searchParams.get('budget');
	const local = false;
	const host = local ? 'http://localhost:8080' : 'https://solar-backend.fly.dev';
	const backendURI = `${host}/report?rc=${params.slug}&percentage=${percentage}&budget=${budget}`;
	const svgURI = `${host}/svg?rc=${params.slug}&percentage=${percentage}`;
	const dashboard = await fetch(backendURI).then((r) => r.json());
	const svg = await fetch(svgURI).then((svgURI) => svgURI.text());
	const dasObj: SolarPanelData = dashboard as SolarPanelData;
	return { dashboard: dasObj, svg: svg };
}) satisfies PageServerLoad;
