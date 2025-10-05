/**
 * What-it-doo API
 * 1.0
 * DO NOT MODIFY - This file has been generated using oazapfts.
 * See https://www.npmjs.com/package/oazapfts
 */
import * as Oazapfts from "@oazapfts/runtime";
import * as QS from "@oazapfts/runtime/query";
export const defaults: Oazapfts.Defaults<Oazapfts.CustomHeaders> = {
    headers: {},
    baseUrl: "/",
};
const oazapfts = Oazapfts.runtime(defaults);
export const servers = {};
export type LoginRequest = {
    email?: string;
    password?: string;
    remember_me?: boolean;
};
export type ProblemDetails = {
    detail?: string;
    instance?: string;
    status?: number;
    title?: string;
    "type"?: string;
};
export type DtoLogoutResponse = {
    redirect_url?: string;
    success?: boolean;
};
export type RegisterRequest = {
    email?: string;
    password?: string;
};
export type ModelUser = {
    avatar_url?: string;
    bio?: string;
    created_at?: string;
    email?: string;
    first_name?: string;
    id?: string;
    last_name?: string;
    updated_at?: string;
};
/**
 * Authenticate user
 */
export function login(loginRequest: LoginRequest, opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
    } | {
        status: 400;
        data: ProblemDetails;
    } | {
        status: 401;
        data: ProblemDetails;
    }>("/auth/login", oazapfts.json({
        ...opts,
        method: "POST",
        body: loginRequest
    })));
}
/**
 * Logout user
 */
export function logout(opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: DtoLogoutResponse;
    } | {
        status: 401;
        data: ProblemDetails;
    }>("/auth/logout", {
        ...opts,
        method: "POST"
    }));
}
/**
 * Register user
 */
export function register(registerRequest: RegisterRequest, opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchText("/auth/register", oazapfts.json({
        ...opts,
        method: "POST",
        body: registerRequest
    })));
}
/**
 * Get current user
 */
export function getMyself(opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: ModelUser;
    }>("/users/me", {
        ...opts
    }));
}
