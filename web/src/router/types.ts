import type { RouteRecordRaw, RouteMeta } from 'vue-router'
import { RoleEnum } from '@/enums/roleEnum'
import { defineComponent, VNode } from 'vue'

export type Component<T = any> = ReturnType<typeof defineComponent> | (() => Promise<T>)

// @ts-ignore
export interface AppRouteRecordRaw extends Omit<RouteRecordRaw, 'meta'> {
  name: string
  meta: RouteMeta
  component?: Component | string
  children?: AppRouteRecordRaw[]
  props?: Recordable
  fullPath?: string
}

export interface Menu {
  id: number

  pid: number

  // menu route name
  name: string

  // menu route path
  path: string

  // menu route title
  title: string

  icon: string | (() => VNode)

  permissions?: [] | null

  sort: number

  children: Menu[] | null

  // component path
  component?: string

  roles?: RoleEnum[]

  // hidden menu
  hidden: boolean

  keepAlive: boolean
}
