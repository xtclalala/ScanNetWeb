<script lang="ts">
export default {
  name: 'YRefresh',
}
</script>
<script setup lang="ts">
import { unref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { YIcon } from '@/components'
import { rPath } from '@/enums/rPath'

const router = useRouter()
const route = useRoute()

const refresh = () => {
  router.push({
    path: rPath.REDIRECT + unref(route).fullPath,
  })
}
</script>
<template>
  <n-button text @click="refresh">
    <y-icon size="22" icon-type="RefreshOutline" />
  </n-button>
</template>

<style scoped></style>
