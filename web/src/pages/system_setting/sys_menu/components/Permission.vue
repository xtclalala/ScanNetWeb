<script setup lang="ts" name="YPermission">
import { h, toRef, Ref } from 'vue'
import { YIcon } from '@/components'
import { BasePer, registerPer, SearchPer } from '@/api/system_setting/types/sys_permission'
import { deletePer, register, searchPer, updatePer } from '@/api/system_setting/sys_permission'
import { FormRules, NButton, NDivider, NPopconfirm, NSpace } from 'naive-ui'
import { Page } from '@/api/system_setting/types/sys_role'
import { useTable } from '@/hooks/comHooks/useTable'
import { useModal } from '@/hooks/comHooks/useModal'

type PerProps = {
  show: boolean
  menuId: number
}
const props = withDefaults(defineProps<PerProps>(), {
  show: false,
  menuId: 0,
})
const show = toRef(props, 'show')
const emit = defineEmits(['update:show'])

const columns = [
  {
    title: '按钮名称',
    key: 'title',
    width: 150,
  },
  {
    title: '按钮编码',
    key: 'code',
    width: 150,
  },
  {
    title: '操作',
    key: 'actions',
    fixed: 'right',
    width: 200,
    render(row) {
      return h(
        NSpace,
        {
          size: 1,
        },
        {
          default: () => {
            const options = [
              h(
                NButton,
                {
                  onClick: () => {
                    isAdd.value = false
                    perModel.value = row
                    openModal()
                  },
                  text: true,
                },
                { default: () => '编辑' }
              ),
              h(NDivider, { vertical: true }),
              h(
                NPopconfirm,
                {
                  onPositiveClick: async () => {
                    await deletePer({ id: row.id }, { isMessage: true })
                    await getData()
                  },
                },
                {
                  trigger: () => h(NButton, { text: true }, { default: () => '删除' }),
                  default: () => '请确认是否删除!',
                }
              ),
            ]
            return options
          },
        }
      )
    },
  },
]

// @ts-ignore
const tableApi = async (page: Page, searchData: Ref<any>) => {
  return searchPer<Array<BasePer>>(searchData.value, { isMessage: false })
}
const [, loading, data, searchData, getData, , , key2id] = useTable<BasePer, SearchPer>(
  tableApi,
  { page: 1, pageSize: 10, desc: false },
  { menuId: 0 },
  'Permission'
)

const rules: FormRules = {
  title: {
    required: true,
    message: '请填写按钮名称！',
    trigger: ['input', 'blur'],
  },
  code: {
    required: true,
    message: '请填写按钮编码！',
    trigger: ['input', 'blur'],
  },
}
const registerApi = async (params: registerPer) => {
  return register<string>(params, { isMessage: true })
}
const updateApi = async (params: BasePer) => {
  return updatePer<string>(params, { isMessage: true })
}
const afterApi = async () => {
  loading.value = false
  await getData()
}
const menuId = toRef(props, 'menuId')
const [
  isAdd,
  showModal,
  form,
  perModel,
  modalStyle,
  handleRegister,
  submitCallback,
  ,
  openModal,
  cancelCallback,
  modalTitle,
] = useModal<registerPer>(
  registerApi,
  updateApi,
  afterApi,
  {
    menuId: menuId.value,
    title: '',
    code: '',
    sort: 100,
  },
  {},
  'Permission'
)

const beforeRegister = () => {
  perModel.value.menuId = menuId.value
  handleRegister()
}
// 更新 show
const closeDrawer = () => {
  data.value = []
  emit('update:show', false)
}
const openAfter = () => {
  loading.value = true
  searchData.value.menuId = menuId.value
  getData()
}
</script>

<template>
  <n-drawer v-model:show="show" :width="700" @after-enter="openAfter" @mask-click="closeDrawer">
    <n-drawer-content title="页面按钮">
      <n-space vertical>
        <n-button type="primary" @click="beforeRegister">
          <template #icon>
            <y-icon icon-type="Add" :depth="2" :size="17" color="white" />
          </template>
          新增
        </n-button>
        <n-data-table
          :columns="columns"
          :data="data"
          :scroll-x="500"
          :max-height="750"
          :loading="loading"
          :row-key="key2id"
        />
      </n-space>
    </n-drawer-content>
  </n-drawer>
  <n-modal v-model:show="showModal" :title="modalTitle()" :style="modalStyle" preset="card">
    <n-form
      ref="form"
      :model="perModel"
      :rules="rules"
      label-placement="left"
      require-mark-placement="right-hanging"
      label-width="auto"
    >
      <n-form-item label="按钮名称" path="title">
        <n-input v-model:value="perModel.title" placeholder="标题" />
      </n-form-item>
      <n-form-item label="按钮编码" path="code">
        <n-input v-model:value="perModel.code" placeholder="编码" />
      </n-form-item>
      <n-form-item label="排序" path="sort">
        <n-input-number v-model:value="perModel.sort" placeholder="排序" />
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

<style scoped></style>
