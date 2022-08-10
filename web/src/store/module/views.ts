/**
 * @Description: src\store\module\views.ts
 * @author: y1t
 * @date 2022/7/26
 **/
import { defineStore } from 'pinia'
import { Menu } from '@/router/types'
import { rName } from '@/enums/rName'
import { store } from '@/store'

interface ITabsViewStore {
  viewList: Menu[]
  currentView: string
  tBlacklist: string[]
}

export const useViewStore = defineStore('tabsView', {
  state: (): ITabsViewStore => {
    return {
      viewList: [],
      currentView: '',
      tBlacklist: [rName.REDIRECT, rName.LOGIN, rName.REDIRECT_ROUTE],
    }
  },
  getters: {
    /**
     * 返回视图页面最后一个
     */
    viewListLast(): Menu {
      return this.viewList.slice(-1)[0]
    },
  },
  actions: {
    /**
     * 删除指定视图
     * @param index
     */
    removeTab(index: number) {
      this.viewList.splice(index, 1)
    },
    /**
     * 找到指定视图
     * @param menu
     */
    findTab(menu: Menu): boolean {
      return this.findTabIndex(menu) >= 0
    },
    findTabByName(name: string): boolean {
      return this.findTab({ name } as Menu)
    },
    findTabIndex(menu: Menu): number {
      return this.viewList.findIndex((item) => item.name === menu.name)
    },
    findTabIndexByName(name: string): number {
      return this.findTabIndex({ name } as Menu)
    },
    /**
     * 通过点击侧边菜单栏方式，跳转指定视图
     * @param to
     */
    routerPush(to: Menu) {
      this.currentView = to.name
      if (this.findTab(to)) {
        return
      }
      if (this.tBlacklist.findIndex((self: string) => self === to.name) !== -1) {
        return
      }
      this.viewList.push(to)
    },
  },
})

// Need to be used outside the setup
export function useViewStoreWithOut() {
  return useViewStore(store)
}
