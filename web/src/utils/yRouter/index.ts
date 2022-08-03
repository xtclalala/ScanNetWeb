import { AppRouteRecordRaw, Menu } from '@r/types'
import { PAGE } from '@r/constant'

/**
 * 引入 pages 下的 sfc
 */
const modules = import.meta.glob('../../pages/**/*.vue')

/**
 * 动态导入 sfc
 * @param routeList
 */
export const buildDynamicRoute = async (routeList: AppRouteRecordRaw[]): Promise<void> => {
  routeList.forEach((item) => {
    if (typeof item.component === 'string') {
      if (item.component === 'PAGE') {
        item.component = PAGE
      } else {
        item.component = modules[`../../${item.component}`]
      }
    }
    if (item.children !== undefined) {
      buildDynamicRoute(item.children)
    }
  })
}

/**
 * 菜单构建路由
 * @param routes
 */
export const buildRoute = (routes: Menu[]): AppRouteRecordRaw[] => {
  const m: AppRouteRecordRaw[] = []
  for (const route of routes) {
    const temp: AppRouteRecordRaw = {
      ...route,
      meta: {
        hideMenu: route.hidden,
        title: route.title,
      },
      children: route.children !== null ? buildRoute(route.children) : undefined,
    }
    m.push(temp)
  }
  return m
}
