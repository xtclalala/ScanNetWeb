import {RequestOptions, Result} from "#axios";
import {LoginParams} from "@/api/common/types/login";
import {defHttp} from "@/service";

enum Api {
    Login = '/login',
    Informational = '/user/routerAndRole',
    demo = '/demo/demo',
}

export const doLogin = <T = Result>(params: LoginParams, options?: RequestOptions) =>
    defHttp.post<T>({ url: Api.Login, params }, options)
