<script lang="ts" setup name="YIcon">
/**
 * @Description: 图标 根据 vicons/ionicons5 类型的图标名称
 * @author: y1t
 * @Date: 2022/7/27
 **/
import { NIcon } from 'naive-ui'
import * as vIcons from '@vicons/ionicons5'
const props = defineProps({
  ...NIcon.props,
  iconType: {
    type: String,
    default: 'fallback',
  },
})
</script>
<template>
  <n-icon v-bind="props" :component="vIcons[props.iconType]" />
</template>
