import adapter from '@sveltejs/adapter-static';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		// adapter-static for SPA mode - compiles to static files served by Nginx
		adapter: adapter({
			// default options - outputs to build/
			fallback: 'index.html', // SPA mode - all routes serve index.html
			precompress: false
		})
	}
};

export default config;
