<script lang="ts">
export default {
  name: 'YLayout',
}
</script>
<script lang="ts" setup>
import { YHeader } from './components'
import { YContent } from './components'
import { YFooter } from './components'
import { YSideBar } from './components'
</script>
<template>
  <n-layout has-sider position="absolute">
    <y-side-bar />
    <n-layout>
      <y-header />
      <y-content />
      <y-footer />
    </n-layout>
  </n-layout>
</template>
