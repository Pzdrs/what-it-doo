import type { ProblemDetails } from '$lib/api/client';
import { isHttpError } from '$lib/api/errors';

export const toProblemDetails = (error: unknown): ProblemDetails | null => {
	if (isHttpError(error)) {
		return error.data;
	}
	return null;
};

export const isProblemType = (problem : ProblemDetails | null, typePath: string): boolean => {
	if (!problem || !problem.type) {
		return false;
	}
	return problem.type.endsWith(typePath);
};

export const problemTypeToKey = (typeUri: string): string => {
	try {
		const url = new URL(typeUri);
		// Split the path, drop empty parts, e.g. ["probs", "auth", "incorrect-credentials"]
		const parts = url.pathname.split('/').filter(Boolean);

		// We only want everything after "probs"
		const probsIndex = parts.indexOf('probs');
		if (probsIndex >= 0) {
			const relevant = parts.slice(probsIndex + 1);
			return relevant.join('.');
		}

		// fallback: use whole path
		return parts.join('.');
	} catch {
		// if it's not a valid URL, just return it as-is
		return typeUri;
	}
};

export const getTranslatedError = ($t, error: unknown, tOpts: {} = {}): string => {
	if (!isHttpError(error)) {
		return $t('errors.unknown_error');
	}

	const data: ProblemDetails = error.data;
	if (data.type) {
		const key = `errors.${problemTypeToKey(data.type)}`;
		const translated = $t(key, { default: '', ...tOpts });

		if (translated) {
			return translated;
		}
	}

	return data.detail || $t('errors.unknown_error');
};
