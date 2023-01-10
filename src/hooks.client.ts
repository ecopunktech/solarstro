import * as Sentry from '@sentry/svelte';
import type { HandleClientError } from '@sveltejs/kit';
import { BrowserTracing } from "@sentry/tracing";

Sentry.init({
  dsn: "https://f7f09aff98e7454aa9e60e3d5befadc9@o4504481621213184.ingest.sentry.io/4504481625800706",
  integrations: [new BrowserTracing()],

  // Set tracesSampleRate to 1.0 to capture 100%
  // of transactions for performance monitoring.
  // We recommend adjusting this value in production
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