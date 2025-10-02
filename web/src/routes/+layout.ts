import { initLanguage } from '$lib/utils/i18n.js';

export const load = async ({ }) => {
    await initLanguage();
};