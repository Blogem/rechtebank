import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';
import fs from 'fs';

export default defineConfig({
	plugins: [sveltekit()],
	test: {
		// Vitest configuration for unit and component testing
		include: ['src/**/*.{test,spec}.{js,ts}'],
		environment: 'jsdom',
		globals: true,
		setupFiles: ['./src/test-setup.ts'],
		// Svelte 5 support
		alias: {
			$lib: '/src/lib'
		}
	},
	resolve: {
		conditions: ['browser']
	},
	server: {
		https: {
			key: fs.readFileSync('./localhost+1-key.pem'),
			cert: fs.readFileSync('./localhost+1.pem')
		}
	},
	// Environment variable prefix for client-side access
	envPrefix: 'PUBLIC_'
});
