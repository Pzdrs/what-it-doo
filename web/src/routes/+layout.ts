import { init } from '$lib/utils/server.js';
import type { LayoutLoad } from './$types';

export const ssr = false;

export const load = (async ({ fetch }) => {
	await init(fetch);
}) satisfies LayoutLoad;
