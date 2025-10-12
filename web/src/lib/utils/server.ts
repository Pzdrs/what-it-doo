import { defaults } from "$lib/api/client";
import { initLanguage } from "./i18n";

type Fetch = typeof fetch;

export const init = async (fetch: Fetch) =>{
    defaults.fetch = fetch;
    defaults.baseUrl = '/api/v1';
    await initLanguage();
};