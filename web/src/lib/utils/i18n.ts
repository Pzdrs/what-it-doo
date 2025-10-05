import { defaultLang, langs } from '$lib/constants';
import { init, locale, register, t, waitLocale } from 'svelte-i18n';
import { get, type Unsubscriber } from 'svelte/store';

export async function getFormatter() {
	let unsubscribe: Unsubscriber | undefined;
	await new Promise((resolve) => {
		unsubscribe = locale.subscribe((value) => value && resolve(value));
	});
	unsubscribe?.();

	await waitLocale();
	return get(t);
}

export const initLanguage = async () => {
	for (const { code, loader } of langs) register(code, loader);

	await init({
		fallbackLocale: 'dev',
		initialLocale: defaultLang.code
	});
};
