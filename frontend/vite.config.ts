import adapter from '@sveltejs/adapter-static';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';
import fs from 'fs';

// Check if SSL cert files exist (for local development only)
const keyPath = '../192.168.1.37+1-key.pem';
const certPath = '../192.168.1.37+1.pem';
const hasSSLCerts = fs.existsSync(keyPath) && fs.existsSync(certPath);

export default defineConfig({
	// Kit 3 removed svelte.config.js and the `kit` namespace — the config now
	// goes straight to the plugin.
	plugins: [
		sveltekit({
			// Kit 3 removed `$lib` in favour of `#lib`. There are 23 `$lib`
			// imports here; this alias is the escape hatch the upgrade guide
			// documents, so the migration stays reviewable. Renaming them to
			// `#lib` and dropping this line is a follow-up.
			alias: { $lib: 'src/lib' },
			// adapter-static for SPA mode - compiles to static files served by Nginx
			adapter: adapter({
				// default options - outputs to build/
				fallback: 'index.html', // SPA mode - all routes serve index.html
				precompress: false
			})
		})
	],
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
		...(hasSSLCerts && {
			https: {
				key: fs.readFileSync(keyPath),
				cert: fs.readFileSync(certPath)
			}
		}),
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
