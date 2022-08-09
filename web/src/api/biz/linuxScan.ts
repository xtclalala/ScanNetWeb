import { PageResult, RequestOptions, Result } from '#axios'
import { BaseId, Page } from '@/api/common/types/login'
import { defHttp } from '@/service'
import { File } from '@/api/common/types/file'
import { TaskState } from '@/enums/bizEnum'

enum Api {
  SSH = '/ssh/ssh',
  RunSSH = '/ssh/run',
}

export type BizSSH = BaseId & {
  name: string
  desc: string
  state: TaskState | null
  fileId: string | null
  thread: number
  sheet: string
  ip: number | null
  port: number | null
  user: number | null
  password: number | null
  timeout: number

  file?: File
}

export type SearchSSH = Partial<Pick<BizSSH, 'name' | 'state'>>

export type SearchWithPage = Page & SearchSSH

export type CreateSSH = Omit<BizSSH, 'state' | 'id' | 'createTime' | 'updateTime' | 'deleted'>

export type UpdateSSH = CreateSSH & Pick<Omit<BizSSH, 'state'>, 'id'>

export type DeleteSSH = Pick<BizSSH, 'id'>

export type RunSSH = Pick<BizSSH, 'id'>

export const Search = <T = PageResult>(params: SearchWithPage, options?: RequestOptions) =>
  defHttp.get<T>({ url: Api.SSH, params }, options)

export const Create = <T = Result>(params: CreateSSH, options?: RequestOptions) =>
  defHttp.post<T>({ url: Api.SSH, params }, options)

export const Update = <T = Result>(params: UpdateSSH, options?: RequestOptions) =>
  defHttp.put<T>({ url: Api.SSH, params }, options)

export const Delete = <T = Result>(params: DeleteSSH, options?: RequestOptions) =>
  defHttp.delete<T>({ url: Api.SSH, params }, options)

export const Run = <T = Result>(params: RunSSH, options?: RequestOptions) =>
  defHttp.post<T>({ url: Api.RunSSH, params }, options)

// export const RunResult = <T = Result>(params: CreateSSH, options?: RequestOptions) =>
//     defHttp.post<T>({ url: Api.Login, params }, options)
