<script lang="ts">
export default {
  name: 'YHeader',
}
</script>
<script lang="ts" setup>
import { YAbout } from './components'
import { YTheme } from './components'
import { YNotify } from './components'
import { YRefresh } from './components'
import { YGithub } from './components'
import { YCenter } from './components'
import { YTabBar } from '@/layouts/components/tabbar'
</script>
<template>
  <n-layout-header bordered style="padding: 0">
    <!--    <current-view />-->
    <n-layout>
      <n-layout-header bordered>
        <y-refresh />
        <n-space :size="20" align="center" style="line-height: 1">
          <y-theme />
          <y-about />
          <y-github />
          <y-notify />
          <y-center />
        </n-space>
      </n-layout-header>
      <n-layout-content>
        <y-tab-bar />
      </n-layout-content>
    </n-layout>
  </n-layout-header>
</template>

<style scoped>
.n-layout-header {
  position: sticky;
  top: 0;
  z-index: 1;
  display: flex;
  align-items: center;
  padding: 9px 18px;
}

.n-button {
  margin-right: 15px;
}

.n-space {
  margin-left: auto;
}
</style>
