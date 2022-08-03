<script lang="ts">
export default {
  name: 'YTheme',
}
</script>
<script setup lang="ts">
import { ref } from 'vue'
import { useAppStore } from '@/store/module/app'
import { ThemeEnum } from '@/enums/appEnum'

const active = ref<boolean>(false)
const app = useAppStore()
const handleChange = () => {
  app.setDarkMode(active.value ? ThemeEnum.DARK : ThemeEnum.LIGHT)
}
</script>

<template>
  <n-switch v-model:value="active" @update:value="handleChange" />
</template>

<style scoped></style>
