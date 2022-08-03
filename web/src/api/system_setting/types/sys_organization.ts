import { Page } from '@/api/system_setting/types/sys_role'

export type OrgId = {
  id: number
  ids: number[]
}

export type SearchOrg = Pick<BaseOrg, 'name'>

export type SearchOrgWithPage = Page & SearchOrg

export type BaseOrg = Partial<OrgId> & {
  name: string
  code: string
  sort: number
  pid: number
}

export type registerOrg = Omit<BaseOrg, 'id' | 'ids'>
