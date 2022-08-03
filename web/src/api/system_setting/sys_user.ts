import { defHttp } from '@/service'
import { RequestOptions, Result } from '#axios'
import {
  BaseUser,
  PerUser,
  registerUser,
  SearchUserWithPage,
  UserId,
  userPassword,
} from '@/api/system_setting/types/sys_user'

enum Api {
  user = '/user/user',
  password = '/user/password',
  per = '/user/per',
  complete = '/menu/completeInfo',
}

export const register = <T = Result>(params: registerUser, options?: RequestOptions) =>
  defHttp.post<T>({ url: Api.user, params }, options)

export const updateUser = <T = Result>(params: BaseUser, options?: RequestOptions) =>
  defHttp.put<T>({ url: Api.user, params }, options)

export const updateUserPer = <T = Result>(params: PerUser, options?: RequestOptions) =>
  defHttp.put<T>({ url: Api.per, params }, options)

export const searchUser = <T = Result>(params: SearchUserWithPage, options?: RequestOptions) =>
  defHttp.get<T>({ url: Api.user, params }, options)

export const completeInfo = <T = Result>(params: Pick<UserId, 'id'>, options?: RequestOptions) =>
  defHttp.get<T>({ url: Api.complete, params }, options)

export const deleteUser = <T = Result>(params: Pick<UserId, 'id'>, options?: RequestOptions) =>
  defHttp.delete<T>({ url: Api.user, params }, options)

export const resetPwd = <T = Result>(params: Pick<UserId, 'id'>, options?: RequestOptions) =>
  defHttp.put<T>({ url: Api.password, params }, options)

export const changePwd = <T = Result>(params: userPassword, options?: RequestOptions) =>
  defHttp.put<T>({ url: Api.password, params }, options)
