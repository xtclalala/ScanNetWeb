<script setup lang="ts" name="Y1tLogin">
import { computed, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useMessage } from 'naive-ui'
import { useUserStore } from '@/store/module/user'
import { LoginParams } from '@/api/common/types/login'
import { PageEnum } from '@/enums/pageEnum'
import { useRouteStore } from '@/store/module/router'
import { usePagesAsyncComponent } from '@/hooks/comHooks/useAsyncComponent'
import { demo111 } from '@/api/common/login'

const ChooseActiveRole = usePagesAsyncComponent('/login', 'ChooseActiveRole')

const rules = {
  loginName: {
    required: true,
    message: 'Username is required.',
    trigger: 'blur',
  },
  password: {
    required: true,
    message: 'Password is required.',
    trigger: 'blur',
  },
}
const model = ref<LoginParams>({
  loginName: 'admin11',
  password: '123456',
})
const disabled = computed<boolean>(
  () => model.value.loginName === '' || model.value.password === ''
)
const loading = ref<boolean>(false)
const router = useRouter()
const routeStore = useRouteStore()
const userStore = useUserStore()
const message = useMessage()
const res = await demo111()
console.log(res.message)
const chooseRoleRef = ref<InstanceType<typeof ChooseActiveRole> | null>(null)

const handleLogin = async (e: Event): Promise<void> => {
  e.preventDefault()
  loading.value = true
  try {
    const res = await userStore.login(model.value)
    if (res) {
      // 动态路由 需要选择身份
      if (routeStore.getIsDynamicAddedRoute) {
        // @ts-ignore
        chooseRoleRef.value?.open()
      } else {
        // @ts-ignore
        await chooseRoleRef.value?.generate()
        await router.push(PageEnum.BASE_HOME)
      }
    }
  } catch (e) {
    message.error(e instanceof Error ? e.message : 'unknown error')
  } finally {
    loading.value = false
  }
}
</script>
<template>
  <n-h1 style="--font-size: 60px; --font-weight: 100"> PenetrationHub </n-h1>
  <n-card size="large" style="--padding-bottom: 30px">
    <n-h2 style="--font-weight: 400">登录</n-h2>
    <n-form size="large" :rules="rules" :model="model">
      <n-form-item-row label="Username" path="loginName">
        <n-input v-model:value="model.loginName" placeholder="请输入用户名" />
      </n-form-item-row>
      <n-form-item-row label="Password" path="password">
        <n-input v-model:value="model.password" type="password" placeholder="请输入密码" />
      </n-form-item-row>
    </n-form>
    <n-button
      type="primary"
      size="large"
      block
      :loading="loading"
      :disabled="disabled"
      @click="handleLogin"
      >登录</n-button
    >
    <choose-active-role ref="chooseRoleRef" />
  </n-card>
</template>

<style scoped>
.n-h1 {
  margin: 20vh auto 20px;
  text-align: center;
  letter-spacing: 5px;
  opacity: 0.8;
}

.n-card {
  margin: 0 auto;
  max-width: 380px;
  box-shadow: var(--box-shadow);
}
</style>
