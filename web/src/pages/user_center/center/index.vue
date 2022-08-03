<script setup lang="ts" name="UserCenter">
import { FormItemRule, FormRules, useMessage } from 'naive-ui'
import { useModal } from '@/hooks/comHooks/useModal'
import { changePwd } from '@/api/system_setting/sys_user'
import { userPassword } from '@/api/system_setting/types/sys_user'
import { useUserStore } from '@/store/module/user'
const message = useMessage()
const userStore = useUserStore()
const changeAvatar = async () => {
  message.success('哈哈！没有做修改头像功能')
}
const changePassword = async () => {
  isAdd.value = false
  await clearModel()
  changePwdModel.value.id = userStore.getUserId
  await openModal()
}

const updateApi = async (params: userPassword) => {
  return changePwd<string>(params, { isMessage: true })
}

const rules: FormRules = {
  oldPassword: [
    { required: true, message: '请填写老密码！', trigger: ['input', 'blur'] },
    { min: 6, message: '密码不能小于6位！', trigger: ['input', 'blur'] },
    { max: 15, message: '密码不能大于15位！', trigger: ['input', 'blur'] },
  ],
  newPassword: [
    { required: true, message: '请填写新密码！', trigger: ['input', 'blur'] },
    { min: 6, message: '密码不能小于6位！', trigger: ['input', 'blur'] },
    { max: 15, message: '密码不能大于15位！', trigger: ['input', 'blur'] },
  ],
  newPasswordAgain: [
    { required: true, message: '请填写新密码！', trigger: ['input', 'blur'] },
    {
      // @ts-ignore
      validator: (rule: FormItemRule, value: string) => value === changePwdModel.value.newPassword,
      message: '两次密码不一致',
      trigger: ['input', 'blur'],
    },
  ],
}

const [
  isAdd,
  showModal,
  form,
  changePwdModel,
  modalStyle,
  ,
  submitCallback,
  clearModel,
  openModal,
  cancelCallback,
  modalTitle,
] = useModal(
  () => {},
  updateApi,
  () => {},
  {
    id: '',
    oldPassword: '',
    newPassword: '',
    newPasswordAgain: '',
  },
  {},
  'ChangePwd'
)
</script>

<template>
  <div :class="[`y-center-layout`]">
    <div :class="[`y-center-left`]">
      <n-space vertical>
        <n-avatar
          round
          :size="150"
          src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
        />
        <n-button text @click="changeAvatar">修改头像</n-button>
      </n-space>
    </div>
    <div :class="[`y-center-right`]">
      <n-button text @click="changePassword">修改密码</n-button>
    </div>
  </div>
  <n-modal v-model:show="showModal" :title="modalTitle()" :style="modalStyle" preset="card">
    <n-form
      ref="form"
      :model="changePwdModel"
      :rules="rules"
      label-placement="left"
      require-mark-placement="right-hanging"
      label-width="auto"
    >
      <n-form-item label="老密码" path="oldPassword">
        <n-input v-model:value="changePwdModel.oldPassword" placeholder="老密码" />
      </n-form-item>
      <n-form-item label="新密码" path="newPassword">
        <n-input v-model:value="changePwdModel.newPassword" placeholder="新密码" />
      </n-form-item>
      <n-form-item label="二次输入" path="newPasswordAgain">
        <n-input v-model:value="changePwdModel.newPasswordAgain" placeholder="二次密码" />
      </n-form-item>
    </n-form>
    <template #action>
      <n-space :size="5" justify="end">
        <n-button type="tertiary" size="small" @click="cancelCallback">取消</n-button>
        <n-button type="success" size="small" @click="submitCallback">确认</n-button>
      </n-space>
    </template>
  </n-modal>
</template>

<style scoped lang="sass">
.y-center-layout
  display: flex
.y-center-left
  border-right: solid 1px #efeff5
  width: 28%
  text-align: center
.y-center-right
  padding: 8px
  width: 72%
</style>
