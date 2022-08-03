import { defHttp } from '@/service'
import { LoginParams } from '@/api/common/types/login'
import { RequestOptions, Result } from '#axios'

enum Api {
  Login = '/login',
  Informational = '/user/routerAndRole',
}

/**
 * 登录
 * @param params
 * @param options
 */
export const doLogin = <T = Result>(params: LoginParams, options?: RequestOptions) =>
  defHttp.post<T>({ url: Api.Login, params }, options)

/**
 * 获取登录用户信息
 * @param options
 */
export const getInformation = <T = Result>(options?: RequestOptions) =>
  defHttp.get<T>({ url: Api.Informational }, options)
