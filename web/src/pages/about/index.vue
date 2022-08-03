<script lang="ts">
export default {
  name: 'Y1tAbout',
}
</script>
<script setup lang="ts">
import { ref } from 'vue'
import HelloWorld from '@/components/HelloWorld.vue'

const msg = ref<string>('about')
</script>
<template>
  <h1>{{ msg }}</h1>
  <hello-world msg="test" />
</template>

<style scoped></style>
