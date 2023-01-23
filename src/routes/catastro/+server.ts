import type { RequestHandler } from './$types';
import { json } from '@sveltejs/kit';
import {Catastro} from '$lib/catastro/catastro'
import { Toucan } from 'toucan-js';
type Env = {
	SENTRY_DSN: string;
  };
/** @type {import('./$types').RequestHandler} */
export const GET: RequestHandler = async ({ url }) => {
	// const sentry = new Toucan({
    //   dsn: env.SENTRY_DSN,
    //   context,
    //   request,
    // });
	const rc = url.searchParams.get('rc');
	if (!rc) {
		return json({ error: 'RC is required' });
	}
	const catatro = new Catastro(rc);
	await catatro.getRemoteData();

	return new Response(JSON.stringify(catatro));
	// try {
	// 	await catatro.getRemoteData();

	// return new Response(JSON.stringify(catatro));
	// } catch (e) {
	// 	sentry.captureException(e);

	// 	return new Response('Something went wrong! Team has been notified.', {
	// 	  status: 500,
	// 	});
	
};

