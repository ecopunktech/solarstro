import * as Sentry from '@sentry/node';
import type { HandleServerError } from '@sveltejs/kit';
import crypto from 'crypto';
import { SENTRY_DSN } from '$env/static/private';
Sentry.init({ dsn: SENTRY_DSN });
 
export const handleError = (({ error, event }) => {
  // example integration with https://sentry.io/
  const errorId = crypto.randomUUID();
  Sentry.captureException(error);
 
  return {
    message: 'Whoops!',
    errorId
  };
}) satisfies HandleServerError;