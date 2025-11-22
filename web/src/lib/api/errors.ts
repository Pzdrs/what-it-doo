import { HttpError } from '@oazapfts/runtime';
import type { ProblemDetails } from './client';

export interface ApiHttpError extends HttpError {
  data: ProblemDetails;
}

export function isHttpError(error: unknown): error is ApiHttpError {
  return error instanceof HttpError;
}
