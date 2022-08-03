<script setup lang="ts" name="ChooseActiveRole">
import { ref } from 'vue'
import { useUserStore } from '@/store/module/user'
import { storeToRefs } from 'pinia'
import { useRouter, RouteRecordRaw } from 'vue-router'
import { useRouteStore } from '@/store/module/router'
import { PageEnum } from '@/enums/pageEnum'
import { buildDynamicRoute } from '@/utils/yRouter'
import { rName } from '@/enums/rName'
import { isNull } from '@/utils/is'

const show = ref<boolean>(false)
const router = useRouter()
const currentRole = ref<number | null>(null)
const userStore = useUserStore()
const { roles } = storeToRefs(userStore)
const routeStore = useRouteStore()
// 选择活跃身份
const submitCallback = async () => {
  if (isNull(currentRole.value)) {
    window.$message?.warning('请选择身份')
    return
  }
  await userStore.setCurrentRole(currentRole.value)
  await generate()
  show.value = false
  await router.push(PageEnum.BASE_HOME)
}

const open = () => {
  show.value = true
}
const generate = async () => {
  await routeStore.generateMenus()
  const routeList = await routeStore.generateRoute()
  await buildDynamicRoute(routeList)
  routeList.forEach((route) => {
    router.addRoute(rName.TAB_VIEW, route as RouteRecordRaw)
  })
}
defineExpose({
  generate,
  open,
})
</script>
<template>
  <n-modal :show="show">
    <n-card style="width: 400px" title="请选择活跃身份" footer-style="login-button-fixe">
      <n-select v-model:value="currentRole" :options="roles" value-field="id" label-field="name" />
      <template #action>
        <n-button style="float: right" @click="submitCallback">确定</n-button>
      </template>
    </n-card>
  </n-modal>
</template>
<style lang="sass" scoped></style>
