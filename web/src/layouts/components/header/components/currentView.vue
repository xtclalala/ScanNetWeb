<script lang="ts">
export default {
  name: 'YCurrentView',
}
</script>
<script setup lang="ts">
import { useViewStore } from '@/store/module/views'
const viewStore = useViewStore()
</script>
<template>
  <n-breadcrumb>
    <n-breadcrumb-item>{{ viewStore.currentView }}</n-breadcrumb-item>
    <!--          <n-breadcrumb-item>Home</n-breadcrumb-item>-->
  </n-breadcrumb>
</template>
<style scoped></style>
