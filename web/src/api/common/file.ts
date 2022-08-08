import { Authorization } from '@/service'
import { useGlobalSetting } from '@/utils/env'
import { getToken } from '@/utils/auth'
import { DownloadParam, UploadParam } from '@/api/common/types/file'

enum Api {
  FileUrl = '/file/file',
  DownloadUrl = '/file',
}

export const Upload = (): UploadParam => {
  const { urlPrefix } = useGlobalSetting()
  const url = `${urlPrefix}${Api.FileUrl}`

  return {
    url,
    header: {
      [Authorization]: getToken(),
    },
  }
}

export const Download = (data: DownloadParam): string => {
  const { urlPrefix, domain } = useGlobalSetting()
  return `${domain}${urlPrefix}${Api.DownloadUrl}/${data.id}`
}
