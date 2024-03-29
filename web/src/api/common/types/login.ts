import { Menu } from '@/router/types'

/**
 * 登录接口参数
 */
export interface LoginParams {
  loginName: string
  password: string
}

/**
 * 获取登录用户接口的返回
 */
export interface Information {
  id: string
  username: string
  roles: IRole[]
  orgs: IOrg[]
}

export interface IRole {
  id: number
  name: string
  pid?: number
  sort?: number
  organizeId?: number
  menus: Menu[]
}

export interface IOrg {
  ID: number
  name: string
  pid?: number
  sort?: number
  children?: IOrg[]
}

export type Page = {
  pageSize: number
  page: number
  desc: boolean
}

export type BaseId = {
  id: number
  createTime: string
  updateTime: string
  deleted: string
}

export type BaseUUID = {
  id: string
  createTime: string
  updateTime: string
  deleted: string
}
