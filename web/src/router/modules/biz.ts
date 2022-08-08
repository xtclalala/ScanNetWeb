import { AppRouteRecordRaw } from '@/router/types'
import { rPath } from '@/enums/rPath'
import { rName } from '@/enums/rName'

const routes: AppRouteRecordRaw = {
  path: rPath.BIZ_LINUX_SCAN,
  name: rName.BIZ_LINUX_SCAN,
  component: () => import('@/pages/biz/linuxScan/index.vue'),
  meta: {
    // title: t('routes.dashboard.about'),
    title: '扫描',
    icon: 'LogoGithub',
    hideMenu: false,
  },
}

export default routes
