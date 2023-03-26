import { sveltekit } from '@sveltejs/kit/vite';
import type { UserConfig } from 'vite';
import { sentryVitePlugin } from '@sentry/vite-plugin';
const config: UserConfig = {
	build: {
		sourcemap: true // Source map generation must be turned on
	},
	plugins: [
		sveltekit(),
		process.env.SENTRY_ORG && process.env.SENTRY_PROJECT && process.env.SENTRY_AUTH_TOKEN
			? [
					sentryVitePlugin({
						org: process.env.SENTRY_ORG,
						project: process.env.SENTRY_PROJECT,
						include: './dist',
						// Auth tokens can be obtained from https://sentry.io/settings/account/api/auth-tokens/
						// and need `project:releases` and `org:read` scopes
						authToken: process.env.SENTRY_AUTH_TOKEN,
						ext: ['mjs', 'map']
					})
			  ]
			: undefined
	],
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
};

export default config;
