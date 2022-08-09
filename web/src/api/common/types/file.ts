import { Authorization } from '@/service'
import { BaseUUID } from '@/api/common/types/login'

export type UploadParam = {
  url: string
  header: { [Authorization]: string }
}

export type DownloadParam = {
  id: string
}

export type File = Partial<BaseUUID> & {
  name: string
  type: string
  path: string
}
