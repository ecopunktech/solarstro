import type { RequestHandler } from './$types';
import { json } from '@sveltejs/kit';
import { Catastro, validateRC } from '$lib/catastro/catastro';
/** @type {import('./$types').RequestHandler} */
export const GET: RequestHandler = async ({ url }) => {
	const rc = url.searchParams.get('rc');
	if (!rc) {
		return json({ error: 'RC is required' });
	}
	if (!validateRC(rc)) {
		return json({ error: 'RC is not valid' });
	}
	const catatro = new Catastro(rc);
	try {
		catatro.direccion = "mi direccion";
		await catatro.getRemoteData();
	} catch (error) {
		return json({ error: error });
	}
	
	return new Response(JSON.stringify(catatro));
};
