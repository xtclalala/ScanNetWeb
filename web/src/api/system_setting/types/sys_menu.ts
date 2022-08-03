import { Page } from '@/api/system_setting/types/sys_role'
import { BasePer } from '@/api/system_setting/types/sys_permission'

export type SearchMenu = Pick<BaseMenu, 'name' | 'path' | 'title' | 'component'>

export type SearchMenuWithPage = Page & SearchMenu

export type MenuId = {
  id: number
}

export type BaseMenu = MenuId & {
  name: string
  title: string
  path: string
  hidden: boolean
  component: string
  icon: string
  sort: number
  pid?: number
  permissions?: BasePer[]
  disabled?: boolean
}

export type PerMenu = MenuId & {
  sysRoleId: Array<number>
}

export type registerMenu = Omit<BaseMenu, 'id'> &
  Partial<PerMenu> & {
    hiddenNumber: string
  }

export type SelectMenu = Pick<BaseMenu, 'id' | 'name'>
