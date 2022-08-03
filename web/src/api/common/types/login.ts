import { Menu } from '@r/types'

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
