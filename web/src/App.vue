<template>
  <n-config-provider :theme="theme" :locale="zhCN">
    <!--    <n-theme-editor>-->
    <n-global-style />
    <n-loading-bar-provider>
      <n-message-provider>
        <n-dialog-provider>
          <n-notification-provider>
            <naive-provider-content />
            <RouterView />
          </n-notification-provider>
        </n-dialog-provider>
      </n-message-provider>
    </n-loading-bar-provider>
    <!--    </n-theme-editor>-->
  </n-config-provider>
</template>

<script lang="ts" setup>
import { computed, defineComponent, h } from 'vue'
import { darkTheme, zhCN, lightTheme } from 'naive-ui'
import { useLoadingBar, useDialog, useMessage, useNotification } from 'naive-ui'
import { useAppStore } from '@/store/module/app'
import { ThemeEnum } from '@/enums/appEnum'
import { wsConnect } from '@/api/common/ws'
import { useWsUrl } from '@/utils/env'

function registerNaiveTools() {
  window.$loadingBar = useLoadingBar()
  window.$dialog = useDialog()
  window.$message = useMessage()
  window.$notification = useNotification()
}

const NaiveProviderContent = defineComponent({
  setup() {
    registerNaiveTools()
  },
  render() {
    return h('div')
  },
})

const app = useAppStore()
const theme = computed(() => (app.getDarkMode === ThemeEnum.LIGHT ? lightTheme : darkTheme))
wsConnect(useWsUrl())
</script>
