<script lang="ts">
export default {
  name: 'YSideBar',
}
</script>
<script lang="ts" setup>
import { SideBarLogo } from './components'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useViewStore } from '@/store/module/views'
import { Menu } from '@r/types'
import { useRouteStore } from '@/store/module/router'

const router = useRouter()
const routeStore = useRouteStore()
const collapsed = ref<boolean>(false)
const viewStore = useViewStore()
const options = routeStore.getMenus
const expandedKeys = ref<string[]>([])

// 菜单跳转
const handleAlt = (key: string, item: Menu) => {
  viewStore.routerPush(item)
  router.push({ name: key })
}
</script>
<template>
  <n-layout-sider
    bordered
    :width="220"
    :collapsed="collapsed"
    :native-scrollbar="false"
    collapse-mode="width"
    show-trigger
    @collapse="collapsed = true"
    @expand="collapsed = false"
  >
    <side-bar-logo />
    <n-menu
      :value="viewStore.currentView"
      :collapsed="collapsed"
      key-field="name"
      label-field="title"
      :default-expanded-keys="expandedKeys"
      :options="options"
      :indent="24"
      @update:value="handleAlt"
    />
  </n-layout-sider>
</template>

<style scoped></style>
