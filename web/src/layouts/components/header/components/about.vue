<script lang="ts">
export default {
  name: 'YAbout',
}
</script>
<script setup lang="ts">
import { YIcon } from '@/components'
import { rPath } from '@/enums/rPath'
</script>
<template>
  <n-tooltip>
    <template #trigger>
      <router-link :to="rPath">
        <y-icon :size="22" icon-type="HelpCircleOutline" :depth="2" />
      </router-link>
    </template>
    about me
  </n-tooltip>
</template>
<style scoped></style>
