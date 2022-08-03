import { defHttp } from '@/service'
import { RequestOptions, Result } from '#axios'
import {
  BaseOrg,
  OrgId,
  registerOrg,
  SearchOrg,
  SearchOrgWithPage,
} from '@/api/system_setting/types/sys_organization'

enum Api {
  organize = '/organize/organize',
  organizeSelect = '/organize/organize/select',
}

export const register = <T = Result>(params: registerOrg, options?: RequestOptions) =>
  defHttp.post<T>({ url: Api.organize, params }, options)

export const updateOrg = <T = Result>(params: BaseOrg, options?: RequestOptions) =>
  defHttp.put<T>({ url: Api.organize, params }, options)

export const deleteOrg = <T = Result>(params: Pick<OrgId, 'id'>, options?: RequestOptions) =>
  defHttp.delete<T>({ url: Api.organize, params }, options)

export const searchOrg = <T = Result>(params: SearchOrgWithPage, options?: RequestOptions) =>
  defHttp.get<T>({ url: Api.organize, params }, options)

export const selectOrg = <T = Result>(params: Pick<SearchOrg, 'name'>, options?: RequestOptions) =>
  defHttp.get<T>({ url: Api.organizeSelect, params }, options)
