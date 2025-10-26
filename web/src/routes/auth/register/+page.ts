import { requireNoAuth } from '$lib/utils/auth.js';
import { getFormatter } from '$lib/utils/i18n.js';

export const load = async () => {
	await requireNoAuth();

	const $t = await getFormatter();

	return {
		meta: {
			title: $t('sign_up')
		}
	};
};
