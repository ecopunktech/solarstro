import type { RequestHandler } from './$types';
import { json } from '@sveltejs/kit';
import { parseStringPromise } from 'xml2js';
import {Catastro} from '$lib/catastro/catastro'
/** @type {import('./$types').RequestHandler} */
export const GET: RequestHandler = async ({ url }) => {
	const rc = url.searchParams.get('rc');
	if (!rc) {
		return json({ error: 'RC is required' });
	}
	const catatro = new Catastro(rc);
	await catatro.getRemoteData();

	return new Response(JSON.stringify(catatro));
};

