<script lang="ts">
export default {
  name: 'YCenter',
}
</script>
<script setup lang="ts">
import { computed, h } from 'vue'
import { renderIcon } from '@/utils/render'
import { RouterLink, useRouter } from 'vue-router'
import { NButton, useMessage } from 'naive-ui'
import { useUserStore } from '@/store/module/user'
import { rPath } from '@/enums/rPath'
import { rName } from '@/enums/rName'
import { useAppStore } from '@/store/module/app'

const router = useRouter()
const message = useMessage()
const userStores = useUserStore()

// 选中下拉框中的回调
const handleOptionsSelect = async (key: unknown): Promise<void> => {
  if ((key as string) === 'me') {
    message.success(`Welcome back, ${userStores.username as string}!`)
  }
}

// 退出登录
const logout = async () => {
  await useAppStore().clearOtherStore()
  router.push({ name: rName.LOGIN })
}

const options = computed(() => [
  { key: 'me', icon: renderIcon('Barcode'), label: `Hey, ${userStores.username as string}!` },
  { key: 'divider', type: 'divider' },
  // todo 权限模式切换身份
  {
    key: 'profile',
    icon: renderIcon('PeopleCircleOutline'),
    label: () => h(RouterLink, { to: rPath.CENTER }, { default: () => 'Your Settings' }),
  },
  {
    key: 'logout',
    icon: renderIcon('LogInOutline'),
    label: () => h(NButton, { text: true, onClick: () => logout() }, { default: () => 'Sign out' }),
  },
])
</script>
<template>
  <n-dropdown placement="bottom-end" show-arrow :options="options" @select="handleOptionsSelect">
    <n-avatar size="small" round src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg" />
  </n-dropdown>
</template>

<style scoped></style>
