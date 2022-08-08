export enum rPath {
  ROOT = '/',
  TAB_VIEW = '/tabview',
  LOGIN = '/login',

  NOT_FOUND = '/:path(.*)*',
  NOT_FOUND_404 = '/404',
  REDIRECT = '/redirect',
  REDIRECT_ROUTE = '/redirect/:path(.*)',

  BIZ_LINUX_SCAN = '/linuxScan',
  SYSTEM_MENU = 'menu',
  SYSTEM_ROLE = 'role',
  SYSTEM_USER = 'user',
  SYSTEM_ORGANIZE = 'organize',
  ABOUT = '/about',
  CENTER = '/center',
}
