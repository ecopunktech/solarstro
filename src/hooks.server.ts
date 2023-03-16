import { Toucan } from 'toucan-js';
import type { HandleServerError } from '@sveltejs/kit';
import { SENTRY_DSN } from '$env/static/private'; 
export const handleError = (({ error, event }) => {
  const sentry = new Toucan({
    dsn: SENTRY_DSN,
    release: '1.0.0',
  });
  // console.log('error', error);
  const errorId = sentry.captureException(error);
 
  return {
    message: 'Whoops!',
    errorId: errorId
  };
}) satisfies HandleServerError;