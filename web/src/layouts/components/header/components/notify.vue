<script lang="ts">
export default {
  name: 'YNotify',
}
</script>
<script setup lang="ts">
import { YIcon } from '@/components'
</script>
<template>
  <n-popover trigger="click" placement="bottom-end" :width="300">
    <template #trigger>
      <!--          todo 可能要用tsx去加载-->
      <n-badge dot processing>
        <y-icon :size="22" icon-type="NotificationsCircle" :depth="2" />
      </n-badge>
    </template>
    <n-tabs type="line" justify-content="space-evenly" style="--pane-padding: 0">
      <!--            todo 没有消息时显示空数据 有数据时固定高度-->
      <n-tab-pane name="notifications" tab="notifications（5）">
        <n-list style="margin: 0">
          <n-list-item v-for="i in 8" :key="i">No {{ i }}</n-list-item>
        </n-list>
      </n-tab-pane>
      <n-tab-pane name="message" tab="messages（5）">
        <n-list style="margin: 0">
          <n-list-item v-for="i in 8" :key="i">msg {{ i }}</n-list-item>
        </n-list>
      </n-tab-pane>
    </n-tabs>
  </n-popover>
</template>
<style scoped></style>
