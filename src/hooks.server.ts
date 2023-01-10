import * as Sentry from '@sentry/node';
import type { HandleServerError } from '@sveltejs/kit';
import crypto from 'crypto';

Sentry.init({ dsn: "https://d31fd8b4648c4f3493040e7328bb33a0@o4504481621213184.ingest.sentry.io/4504481625800705" });
 
export const handleError = (({ error, event }) => {
  // example integration with https://sentry.io/
  const errorId = crypto.randomUUID();
  Sentry.captureException(error);
 
  return {
    message: 'Whoops!',
    errorId
  };
}) satisfies HandleServerError;