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
    password?: string;
    username?: string;
};
export type RegisterRequest = {
    email?: string;
    password?: string;
    username?: string;
};
/**
 * Authenticate user
 */
export function postAuthLogin(loginRequest: LoginRequest, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchText("/auth/login", oazapfts.json({
        ...opts,
        method: "POST",
        body: loginRequest
    }));
}
/**
 * Logout user
 */
export function postAuthLogout(opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchText("/auth/logout", {
        ...opts,
        method: "POST"
    });
}
/**
 * Register user
 */
export function postAuthRegister(registerRequest: RegisterRequest, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchText("/auth/register", oazapfts.json({
        ...opts,
        method: "POST",
        body: registerRequest
    }));
}
