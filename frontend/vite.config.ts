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
		host: '0.0.0.0', // Allow access from other devices on network
		https: {
			key: fs.readFileSync('../192.168.1.37+1-key.pem'),
			cert: fs.readFileSync('../192.168.1.37+1.pem')
		},
		proxy: {
			'/v1': {
				target: 'http://localhost:8080',
				changeOrigin: true,
				secure: false
			}
		}
	},
	// Environment variable prefix for client-side access
	envPrefix: 'PUBLIC_'
});
