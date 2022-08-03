/**
 * @Description: src\store\module\app.ts
 * @author: y1t
 * @date 2022/7/26
 **/
import type { BeforeMiniState } from '#/store'

import { defineStore } from 'pinia'
import { store } from '@/store'

import { ThemeEnum } from '@/enums/appEnum'
import { APP_DARK_MODE_KEY_ } from '@/enums/cacheEnum'
import { Persistent } from '@/utils/cache/persistent'
import { resetRouter } from '@/router/index'
import { useUserStore } from '@/store/module/user'
import { useRouteStore } from '@/store/module/router'
import { useViewStore } from '@/store/module/views'

interface AppState {
  // 模式
  darkMode: ThemeEnum | undefined
  // Page loading status
  pageLoading: boolean
  // When the window shrinks, remember some states, and restore these states when the window is restored
  beforeMiniInfo: BeforeMiniState
}
let timeId: TimeoutHandle
export const useAppStore = defineStore({
  id: 'app',
  state: (): AppState => ({
    darkMode: undefined,
    pageLoading: false,
    beforeMiniInfo: {},
  }),
  getters: {
    getPageLoading(): boolean {
      return this.pageLoading
    },
    /**
     * 获取模式，默认明亮模式
     */
    getDarkMode(): 'light' | 'dark' | string {
      return this.darkMode || localStorage.getItem(APP_DARK_MODE_KEY_) || ThemeEnum.LIGHT
    },
    getBeforeMiniInfo(): BeforeMiniState {
      return this.beforeMiniInfo
    },
  },
  actions: {
    setPageLoading(loading: boolean): void {
      this.pageLoading = loading
    },
    /**
     * 设置模式
     * @param mode
     */
    setDarkMode(mode: ThemeEnum): void {
      this.darkMode = mode
      localStorage.setItem(APP_DARK_MODE_KEY_, mode)
    },
    setBeforeMiniInfo(state: BeforeMiniState): void {
      this.beforeMiniInfo = state
    },
    async resetAllState() {
      resetRouter()
      Persistent.clearAll()
    },
    async setPageLoadingAction(loading: boolean): Promise<void> {
      if (loading) {
        clearTimeout(timeId)
        // Prevent flicker
        timeId = setTimeout(() => {
          this.setPageLoading(loading)
        }, 50)
      } else {
        this.setPageLoading(loading)
        clearTimeout(timeId)
      }
    },
    /**
     * 清空其他store
     */
    async clearOtherStore(): Promise<void> {
      useUserStore().$reset()
      useRouteStore().$reset()
      useViewStore().$reset()
    },
  },
})

// Need to be used outside the setup
export function useAppStoreWithOut() {
  return useAppStore(store)
}
