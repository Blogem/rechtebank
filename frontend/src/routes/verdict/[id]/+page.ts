import { error } from '@sveltejs/kit';
import { ApiAdapter } from '$lib/adapters/api/ApiAdapter';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params }) => {
    // Use default API adapter - in browser, will proxy through nginx to backend
    const api = new ApiAdapter();
    const { id } = params;

    if (!id) {
        throw error(400, 'Ongeldige vonnis-ID');
    }

    try {
        const data = await api.getVerdictById(id);

        return {
            verdict: data.verdict,
            imageData: data.image
        };
    } catch (err) {
        console.error('Failed to load verdict:', err);

        // Check error type and return appropriate error
        if (err instanceof Error) {
            if (err.message.includes('404')) {
                throw error(404, 'Vonnis niet gevonden');
            }
            if (err.message.includes('400')) {
                throw error(400, 'Ongeldige vonnis-ID');
            }
        }

        throw error(500, 'Fout bij ophalen van vonnis');
    }
};
