import type { PageServerLoad } from './$types';

export const load = (async ({ params }) => {
    const dashboard = await fetch(`https://solar-backend.fly.dev/rectangles?rc=${params.slug}`).then(r => r.json());
	return {dashboard:dashboard};
}) satisfies PageServerLoad;
