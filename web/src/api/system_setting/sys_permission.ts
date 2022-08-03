import { defHttp } from '@/service'
import { RequestOptions, Result } from '#axios'
import {
  BasePer,
  PerId,
  registerPer,
  SearchPerWithPage,
} from '@/api/system_setting/types/sys_permission'
import { MenuId } from '@/api/system_setting/types/sys_menu'

enum Api {
  per = '/permission/per',
  allPer = '/permission/perAll',
}

export const register = <T = Result>(params: registerPer, options?: RequestOptions) =>
  defHttp.post<T>({ url: Api.per, params }, options)

export const updatePer = <T = Result>(params: BasePer, options?: RequestOptions) =>
  defHttp.put<T>({ url: Api.per, params }, options)

export const allPerByMenuId = <T = Result>(params: MenuId, options?: RequestOptions) =>
  defHttp.get<T>({ url: Api.allPer, params }, options)

export const deletePer = <T = Result>(params: PerId, options?: RequestOptions) =>
  defHttp.delete<T>({ url: Api.per, params }, options)

export const searchPer = <T = Result>(params: SearchPerWithPage, options?: RequestOptions) =>
  defHttp.get<T>({ url: Api.per, params }, options)
