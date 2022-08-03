// import { useNotification } from 'naive-ui'
// import { useI18n } from 'vue-i18n'
// import router from '/@/router';
// import { PageEnum } from '/@/enums/pageEnum';
// import { useUserStoreWithOut } from '@/store/modules/user'
import projectSetting from '@/settings/projectSetting'
import { SessionTimeoutProcessingEnum } from '@/enums/appEnum'

// const { error } = useNotification()
const stp = projectSetting.sessionTimeoutProcessing

export function checkStatus(status: number, msg: string): void {
  // const { t } = useI18n()
  // const userStore = useUserStoreWithOut()
  let errMessage = ''

  switch (status) {
    case 400:
      errMessage = `${msg}`
      break
    // 401: Not logged in
    // Jump to the login page if not logged in, and carry the path of the current page
    // Return to the current page after successful login. This step needs to be operated on the login page.
    case 401:
      // userStore.setToken(undefined)
      // errMessage = msg || t('sys_role.api.errMsg401')
      errMessage = msg || 'sys_role.api.errMsg401'
      if (stp === SessionTimeoutProcessingEnum.PAGE_COVERAGE) {
        // userStore.setSessionTimeout(true)
      } else {
        // userStore.logout(true)
      }
      break
    case 403:
      // errMessage = t('sys_role.api.errMsg403')
      errMessage = 'sys_role.api.errMsg403'
      break
    // 404请求不存在
    case 404:
      // errMessage = t('sys_role.api.errMsg404')
      errMessage = 'sys_role.api.errMsg404'
      break
    case 405:
      // errMessage = t('sys_role.api.errMsg405')
      errMessage = 'sys_role.api.errMsg405'
      break
    case 408:
      // errMessage = t('sys_role.api.errMsg408')
      errMessage = 'sys_role.api.errMsg408'
      break
    case 500:
      // errMessage = t('sys_role.api.errMsg500')
      errMessage = 'sys_role.api.errMsg500'
      break
    case 501:
      // errMessage = t('sys_role.api.errMsg501')
      errMessage = 'sys_role.api.errMsg501'
      break
    case 502:
      // errMessage = t('sys_role.api.errMsg502')
      errMessage = 'sys_role.api.errMsg502'
      break
    case 503:
      // errMessage = t('sys_role.api.errMsg503')
      errMessage = 'sys_role.api.errMsg503'
      break
    case 504:
      // errMessage = t('sys_role.api.errMsg504')
      errMessage = 'sys_role.api.errMsg504'
      break
    case 505:
      // errMessage = t('sys_role.api.errMsg505')
      errMessage = 'sys_role.api.errMsg505'
      break
    default:
  }

  if (errMessage) {
    // error({ content: errMessage, title: status.toString() })
  }
}
