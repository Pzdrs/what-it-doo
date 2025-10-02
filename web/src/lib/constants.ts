export const SESSION_COOKIE_NAME = 'wid_session';

export enum AppRoute {
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
	{ name: 'English', code: 'en', loader: () => import('$i18n/en.json') }
];
