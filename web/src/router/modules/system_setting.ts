import { AppRouteRecordRaw } from '@/router/types'
import { rPath } from '@/enums/rPath'
import { rName } from '@/enums/rName'
import { PageEnum } from '@/enums/pageEnum'
import { PAGE } from '@/router/constant'

const routes: AppRouteRecordRaw = {
  path: rPath.SYSTEM,
  name: rName.SYSTEM,
  redirect: PageEnum.SYSTEM_USER,
  component: PAGE,
  meta: {
    title: '系统设置',
    icon: 'AtCircleOutline',
  },
  children: [
    {
      path: rPath.SYSTEM_MENU,
      name: rName.SYSTEM_MENU,
      component: () => import('@/pages/system_setting/sys_menu/index.vue'),
      meta: {
        // title: t('routes.system.sysMenu'),
        title: '菜单设置',
        hideBreadcrumb: true,
      },
    },
    {
      path: rPath.SYSTEM_ORGANIZE,
      name: rName.SYSTEM_ORGAMIZE,
      component: () => import('@/pages/system_setting/sys_organization/index.vue'),
      meta: {
        // title: t('routes.system.sysOrganization'),
        title: '部门设置',
        hideBreadcrumb: true,
      },
    },
    {
      path: rPath.SYSTEM_ROLE,
      name: rName.SYSTEM_ROLE,
      component: () => import('@/pages/system_setting/sys_role/index.vue'),
      meta: {
        // title: t('routes.system.sysRole'),
        title: '角色设置',
        hideBreadcrumb: true,
      },
    },
    {
      path: rPath.SYSTEM_USER,
      name: rName.SYSTEM_USER,
      component: () => import('@/pages/system_setting/sys_user/index.vue'),
      meta: {
        // title: t('routes.system.sysUser'),
        title: '用户设置',
        hideBreadcrumb: true,
      },
    },
  ],
}

export default routes
