export const SESSION_COOKIE_NAME = 'wid_is_authenticated';

export enum AppRoute {
	HOME = '/',
	CHAT = '/chat',

	AUTH_LOGIN = '/auth/login',
	AUTH_REGISTER = '/auth/register'
}

interface Lang {
	name: string;
	code: string;
	loader: () => Promise<{ default: object }>;
}

export const defaultLang: Lang = { name: 'English', code: 'en', loader: () => import('$i18n/en.json') };

export const langs: Lang[] = [
	{ name: 'English', code: 'en', loader: () => import('$i18n/en.json') },
	{ name: 'Czech', code: 'cs', loader: () => import('$i18n/cs.json') }
];

export const INIT_LOAD_MESSAGES_COUNT = 20;
export const LOAD_OLDER_MESSAGES_COUNT = 5;