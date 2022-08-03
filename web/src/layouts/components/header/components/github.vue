<script lang="ts">
export default {
  name: 'YGithub',
}
</script>
<script setup lang="ts">
import { YIcon } from '@/components'
import ProjectInfo from '@/settings/projectInfo'
const { projectGithub } = ProjectInfo
</script>
<template>
  <n-tooltip>
    <template #trigger>
      <a :href="projectGithub" target="_blank">
        <y-icon :size="22" icon-type="LogoGithub" :depth="2" />
      </a>
    </template>
    View on github
  </n-tooltip>
</template>
<style scoped></style>
