import type { PageServerLoad } from './$types';
import type { SolarPanelData } from '$lib/catastro/solar';


export const load = (async ({ params, url }) => {
    const percentage = url.searchParams.get('percentage');
    const budget = url.searchParams.get('budget');
    const localURI = `http://localhost:8080/report?rc=${params.slug}&percentage=${percentage}&budget=${budget}`;
    const backendURI = `https://solar-backend.fly.dev/report?rc=${params.slug}&percentage=${percentage}&budget=${budget}`;
	const dashboard = await fetch(backendURI).then(
		(r) => r.json()
	);
	const dasObj: SolarPanelData = dashboard as SolarPanelData;
	return { dashboard: dasObj };
}) satisfies PageServerLoad;
