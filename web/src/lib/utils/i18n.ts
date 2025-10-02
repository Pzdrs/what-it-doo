import { defaultLang, langs } from '$lib/constants';
import { init, register } from 'svelte-i18n';


export const initLanguage = async () => {
	for (const { code, loader } of langs) register(code, loader);

	await init({
		fallbackLocale: 'dev',
		initialLocale: defaultLang.code
	});
};
