/**
 * @Description: src\store\module\router.ts
 * @author: y1t
 * @date 2022/7/26
 **/
import { defineStore } from 'pinia'
import { store } from '@/store'
import { AfterBusinessRoutes } from '@r/index'
import { PermissionModeEnum } from '@/enums/appEnum'
import { buildMenusTree, renderMenuIcon, router2menuDeep } from '@/utils/yMenu'
import { AppRouteRecordRaw, Menu } from '@r/types'
import projectSetting from '@/settings/projectSetting'
import { getAuthCache, setAuthCache } from '@/utils/auth'
import { CACHELIST_CACHE_KEY, MENU_CACHE_KEY } from '@/enums/cacheEnum'
import { useUserStore } from '@/store/module/user'
import { buildRoute } from '@/utils/yRouter'

export interface IAsyncRouteState {
  menus: Menu[] | undefined
  cacheList: string[]
  isDynamicAddedRoute: boolean | undefined
}

export const useRouteStore = defineStore({
  id: 'route',
  state: (): IAsyncRouteState => ({
    menus: undefined,
    cacheList: [],
    // Whether the route has been dynamically added
    isDynamicAddedRoute: undefined,
  }),
  getters: {
    getMenus(): Menu[] {
      return this.menus || renderMenuIcon(getAuthCache<Menu[]>(MENU_CACHE_KEY))
    },
    getCacheList(): string[] {
      return this.cacheList || getAuthCache<string[]>(CACHELIST_CACHE_KEY)
    },
    getIsDynamicAddedRoute(): boolean {
      if (this.isDynamicAddedRoute === undefined) {
        const { permissionMode } = projectSetting
        switch (permissionMode) {
          case PermissionModeEnum.BACK:
          case PermissionModeEnum.ROLE:
            this.isDynamicAddedRoute = true
            break

          case PermissionModeEnum.ROUTE_MAPPING:
            this.isDynamicAddedRoute = false
            break

          default:
            this.isDynamicAddedRoute = undefined
        }
      }
      return this.isDynamicAddedRoute || false
    },
  },
  actions: {
    setDynamicAddedRoute(added: boolean) {
      this.isDynamicAddedRoute = added
    },
    /**
     * 设置菜单
     * @param menus
     */
    setMenus(menus: Menu[]) {
      setAuthCache(MENU_CACHE_KEY, menus)
      this.menus = renderMenuIcon(menus)
    },
    /**
     * 设置缓存组件
     */
    setCacheList(menus: Menu[]) {
      const cacheList = menus
        .filter((self) => self.keepAlive)
        .flatMap((self) => self.name)
        .concat('About')
      this.cacheList = cacheList
      setAuthCache(CACHELIST_CACHE_KEY, cacheList)
    },
    /**
     * 生成菜单
     */
    async generateMenus(): Promise<void> {
      const buildMenus: Menu[] = []
      const { permissionMode } = projectSetting

      switch (permissionMode) {
        case PermissionModeEnum.BACK: {
          this.setDynamicAddedRoute(true)
          const user = useUserStore()
          let menus = user.getCurrentRole.menus
          this.setCacheList(menus)
          menus = buildMenusTree(menus)
          buildMenus.push(...menus)
          break
        }

        case PermissionModeEnum.ROUTE_MAPPING:
        default: {
          this.setDynamicAddedRoute(false)
        }
      }
      this.setMenus(buildMenus)
    },
    /**
     * 生成路由
     */
    async generateRoute(): Promise<AppRouteRecordRaw[]> {
      let buildRoutes: AppRouteRecordRaw[]
      const { permissionMode } = projectSetting

      switch (permissionMode) {
        case PermissionModeEnum.BACK: {
          this.setDynamicAddedRoute(true)
          buildRoutes = buildRoute(this.getMenus)
          break
        }

        case PermissionModeEnum.ROUTE_MAPPING:
        default: {
          this.setDynamicAddedRoute(false)
          buildRoutes = AfterBusinessRoutes
          const menus = router2menuDeep(buildRoutes)
          this.setMenus(menus)
        }
      }

      return buildRoutes
    },
  },
})

// Need to be used outside the setup
export function useRouteStoreWidthOut() {
  return useRouteStore(store)
}
