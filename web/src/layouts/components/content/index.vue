<script lang="ts">
export default {
  name: 'YContent',
}
</script>
<script lang="ts" setup>
import { useRouteStore } from '@/store/module/router'
const routeStore = useRouteStore()
</script>
<template>
  <n-layout-content
    embedded
    content-style="padding: 12px; min-height: 80vh"
    :native-scrollbar="false"
  >
    <router-view v-slot="{ Component, route }">
      <!--      todo 添加缓存配置 使用的组件需要有 name属性-->
      <keep-alive :include="routeStore.getCacheList">
        <n-card>
          <component :is="Component" :key="route.name" />
        </n-card>
      </keep-alive>
    </router-view>
  </n-layout-content>
</template>
