import { defaults } from "./client";

export interface InitOptions {
	baseUrl: string;
	apiKey: string;
	headers?: Record<string, string>;
}

export const init = ({ baseUrl, apiKey, headers }: InitOptions) => {
	setBaseUrl(baseUrl);
};
export const setBaseUrl = (baseUrl: string) => {
	defaults.baseUrl = baseUrl;
};
