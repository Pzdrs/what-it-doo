import { AppRoute } from '$lib/constants.js';
import { requireNoAuth } from '$lib/utils/auth.js';

export const load = async ({ url }) => {
	await requireNoAuth(url);
	
	return {
		continueUrl: url.searchParams.get('continue') || AppRoute.HOME
	};
};
