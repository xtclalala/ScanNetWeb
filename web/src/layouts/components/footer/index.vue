<script lang="ts">
export default {
  name: 'YFooter',
}
</script>
<script setup lang="ts">
import ProjectInfo from '@/settings/projectInfo'
const { projectGithub, authorGithub } = ProjectInfo
</script>
<template>
  <n-layout-footer>
    <n-p>
      <n-a :href="projectGithub">&lt;/&gt;</n-a> with <n-text type="error">♥</n-text> by
      <n-a :href="authorGithub">y1t</n-a>. Licensed under the
      <n-a :href="projectGithub">MIT</n-a> License.
    </n-p>
    <n-p>&copy; {{ new Date().getFullYear() }} Lei’s. All rights reserved.</n-p>
  </n-layout-footer>
</template>
<style scoped>
.n-layout-footer {
  height: 118px;
  opacity: 0.6;
  padding: 30px;
  text-align: center;
}
.n-a {
  text-decoration: none;
}
</style>
