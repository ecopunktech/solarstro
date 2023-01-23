import * as Sentry from '@sentry/svelte';
import type { HandleClientError } from '@sveltejs/kit';
import { PUBLIC_SENTRY_DSN } from '$env/static/public';
import { BrowserTracing } from "@sentry/tracing";

Sentry.init({
  dsn: PUBLIC_SENTRY_DSN,
  integrations: [new BrowserTracing()],
  tracesSampleRate: 1.0,
}); 
export const handleError = (({ error, event }) => {
  const errorId = crypto.randomUUID();
  // example integration with https://sentry.io/
  Sentry.captureException(error);
 
  return {
    message: 'Whoops!',
    errorId
  };
}) satisfies HandleClientError;