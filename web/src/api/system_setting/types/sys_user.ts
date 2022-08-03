import { BaseRole, Page } from '@/api/system_setting/types/sys_role'
import { BaseOrg } from '@/api/system_setting/types/sys_organization'

export type SearchUser = Required<
  Pick<registerUser, 'username' | 'loginName' | 'roleIds' | 'orgIds'>
>

export type SearchUserWithPage = Page & SearchUser

export type UserId = {
  id: number
  CreateTime: string
  UpdateTime: string
}

export type BaseUser = UserId & {
  username: string
  loginName: string
}

export type PerUser = UserId & {
  roles: Array<BaseRole>
  roleIds: Array<number>
  organizes: Array<BaseOrg>
  orgIds: Array<number>
}

export type registerUser = Omit<BaseUser, 'id' | 'CreateTime' | 'UpdateTime'> & Partial<PerUser>

export type userPassword = Pick<UserId, 'id'> & {
  oldPassword: string
  newPassword: string
  newPasswordAgain: string
}
