/**
 * @Description: src\utils\yMenu\index.ts
 * @author: y1t
 * @date 2022/7/26
 **/
import { AppRouteRecordRaw, Menu } from '@r/types'
import { renderIcon } from '@/utils/render'
import { list2Tree } from '@/utils/helper/treeHelper'
import { listSort } from '@/utils/helper/listHelper'

/**
 * 把路由构建成菜单
 * @param routes
 */
export const router2menuDeep = (routes: AppRouteRecordRaw[]): Menu[] => {
  const menu: Menu[] = []
  for (const item of routes) {
    const { meta, name, children, path } = item
    if (meta.hideMenu) {
      continue
    }
    const menuItem: Menu = {
      id: meta.id,
      pid: meta.pid ? meta.pid : 0,
      name: name,
      path: path,
      title: meta.title,
      hidden: meta.hideMenu ? meta.hideMenu : false,
      sort: 1,
      keepAlive: true,
      icon: renderIcon(!meta.icon ? 'LogoGithub' : meta.icon),
      children: children !== undefined ? router2menuDeep(children) : null,
    }
    menu.push(menuItem)
  }
  return menu
}

/**
 * 路由转化成菜单
 * @param routes
 */
export const router2menu = (routes: AppRouteRecordRaw): Menu => {
  const { name, meta } = routes
  return { name, title: meta.title } as Menu
}

/**
 * 菜单渲染图标
 * @param menu
 */
export const renderMenuIcon = (menu: Menu[]): Menu[] => {
  menu.flatMap((item) => {
    if (typeof item.icon === 'string') {
      item.icon = renderIcon(item.icon)
    }
    if (item.children !== null) {
      item.children = renderMenuIcon(item.children)
    }
    return item
  })
  return menu
}

/**
 * 构建侧边菜单兰内容
 * @param routes 用户可以使用的菜单
 */
export const buildMenusTree = (routes: Menu[]): Menu[] => {
  sortMenu(routes)
  return list2Tree(routes)
}

/**
 * 路由排序
 * @param routes 需要排序的路由
 */
export const sortMenu = (routes: Menu[]) => {
  listSort(routes, 'pid', 0)
  listSort(routes, 'sort', 100)
}
