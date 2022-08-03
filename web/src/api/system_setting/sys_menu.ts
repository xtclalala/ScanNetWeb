import { defHttp } from '@/service'
import { RequestOptions, Result } from '#axios'

import {
  BaseMenu,
  MenuId,
  registerMenu,
  SearchMenuWithPage,
} from '@/api/system_setting/types/sys_menu'

enum Api {
  menu = '/menu/menu',
  allMenu = '/menu/allMenu',
}

export const register = <T = Result>(params: registerMenu, options?: RequestOptions) =>
  defHttp.post<T>({ url: Api.menu, params }, options)

export const updateMenu = <T = Result>(params: BaseMenu, options?: RequestOptions) =>
  defHttp.put<T>({ url: Api.menu, params }, options)

export const searchMenu = <T = Result>(params: SearchMenuWithPage, options?: RequestOptions) =>
  defHttp.get<T>({ url: Api.menu, params }, options)

export const allMenu = <T = Result>(options?: RequestOptions) =>
  defHttp.get<T>({ url: Api.allMenu }, options)

export const deleteMenu = <T = Result>(params: MenuId, options?: RequestOptions) =>
  defHttp.delete<T>({ url: Api.menu, params }, options)
