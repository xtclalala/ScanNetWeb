import { Authorization } from '@/service'

export type UploadParam = {
  url: string
  header: { [Authorization]: string }
}

export type DownloadParam = {
  id: string
}
