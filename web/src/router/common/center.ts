import { AppRouteRecordRaw } from '@/router/types'
import { rPath } from '@/enums/rPath'
import { rName } from '@/enums/rName'

const routes: AppRouteRecordRaw = {
  path: rPath.CENTER,
  name: rName.CENTER,
  component: () => import('@/pages/user_center/center/index.vue'),
  meta: {
    // title: t('routes.dashboard.about'),
    title: '个人中心',
    hideMenu: true,
  },
}

export default routes
