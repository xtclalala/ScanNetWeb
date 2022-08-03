import type { AppRouteRecordRaw } from '@/router/types'

import { LAYOUT } from '@/router/constant'
import { rPath } from '@/enums/rPath'
import { rName } from '@/enums/rName'
// import { t } from '@/utils/i18n/useI18n'

// 404
export const PAGE_NOT_FOUND_ROUTE: AppRouteRecordRaw = {
  path: rPath.NOT_FOUND,
  name: rName.NOT_FOUND,
  component: LAYOUT,
  redirect: rPath.NOT_FOUND_404,
  meta: {
    title: 'ErrorPage',
    hideBreadcrumb: true,
    hideMenu: true,
    white: true,
  },
  children: [
    {
      path: rPath.NOT_FOUND,
      name: rName.NOT_FOUND,
      component: () => import('@/pages/exception/404.vue'),
      meta: {
        title: 'ErrorPage',
        hideBreadcrumb: true,
        hideMenu: true,
        white: true,
      },
    },
  ],
}

// 重定向
export const REDIRECT_ROUTE: AppRouteRecordRaw = {
  path: rPath.REDIRECT,
  name: rName.REDIRECT,
  component: LAYOUT,
  meta: {
    title: 'redirect',
    hideBreadcrumb: true,
    hideMenu: true,
    white: true,
  },
  children: [
    {
      path: rPath.REDIRECT_ROUTE,
      name: rName.REDIRECT_ROUTE,
      component: () => import('@/pages/redirect/index.vue'),
      meta: {
        title: 'redirect',
        hideBreadcrumb: true,
        white: true,
      },
    },
  ],
}

// 登录
export const LOGIN_ROUTE: AppRouteRecordRaw = {
  path: rPath.LOGIN,
  name: rName.LOGIN,
  component: () => import('@/pages/login/index.vue'),
  meta: {
    title: '登录',
    white: true,
  },
}
