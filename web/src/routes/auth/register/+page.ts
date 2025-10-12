import { requireNoAuth } from '$lib/utils/auth.js';

export const load = async ({ url }) => {
	await requireNoAuth(url);
};
