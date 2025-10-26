import { authenticate } from '$lib/utils/auth';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ url }) => {
	await authenticate(url);

	return {
		meta: {
			title: "<chat label>"
		}
	}
};
