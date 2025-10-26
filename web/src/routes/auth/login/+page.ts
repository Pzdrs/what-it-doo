import { AppRoute } from '$lib/constants.js';
import { requireNoAuth } from '$lib/utils/auth.js';
import { getFormatter } from '$lib/utils/i18n.js';

export const load = async ({ url }) => {
	await requireNoAuth();

	const $t = await getFormatter();

	return {
		continueUrl: url.searchParams.get('continue') || AppRoute.HOME,
		meta: {
			title: $t('sign_in'),
		}
	};
};
