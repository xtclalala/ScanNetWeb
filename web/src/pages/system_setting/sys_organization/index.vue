<script lang="ts">
export default {
  name: 'Y1tSysMenu',
}
</script>
<script setup lang="ts">
import { h, ref } from 'vue'
import { YIcon } from '@/components'
import { FormRules, NButton, NDivider, NPopconfirm, NSpace } from 'naive-ui'
import { Page } from '@/api/system_setting/types/sys_role'
import { PageResult } from '#axios'
import { completeMerger } from '@/utils/helper/objectHelper'
import { useTable } from '@/hooks/comHooks/useTable'
import { deleteOrg, searchOrg, updateOrg, register } from '@/api/system_setting/sys_organize'
import {
  BaseOrg,
  registerOrg,
  SearchOrg,
  SearchOrgWithPage,
} from '@/api/system_setting/types/sys_organization'
import { useModal } from '@/hooks/comHooks/useModal'

const checkedRowKeys = ref([])

const columns = [
  {
    type: 'selection',
    options: ['all', 'none'],
  },
  {
    title: '部门名称',
    key: 'name',
    width: 150,
  },
  {
    title: '部门编码',
    key: 'code',
    width: 150,
  },
  {
    title: '排序',
    key: 'sort',
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
        { size: 1 },
        {
          default: () => {
            const options = [
              h(
                NButton,
                {
                  onClick: () => {
                    isAdd.value = false
                    orgModel.value = row
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
                    await deleteOrg({ id: row.id }, { isMessage: true })
                    await getData({
                      page: pagination.page,
                      pageSize: pagination.pageSize,
                      desc: false,
                    })
                  },
                },
                {
                  trigger: () => h(NButton, { text: true }, { default: () => '删除' }),
                  default: () => '请确认是否删除!',
                }
              ),
              h(NDivider, { vertical: true }),
              h(
                NButton,
                {
                  onClick: () => {
                    clearModel()
                    orgModel.value.pid = row.id
                    openModal()
                  },
                  text: true,
                },
                { default: () => '添加下级' }
              ),
            ]
            return options
          },
        }
      )
    },
  },
]
const sTmpData = {
  name: '',
}
const tableApi = async (page: Page, searchData: any) => {
  return searchOrg<PageResult<Array<BaseOrg>>>(
    completeMerger<SearchOrgWithPage>(page, searchData.value),
    {
      isMessage: false,
    }
  )
}
const [pagination, loading, data, searchData, getData, doSearch, doReset, key2id] = useTable<
  registerOrg,
  SearchOrg
>(tableApi, { page: 1, pageSize: 10, desc: false }, sTmpData, 'Organize')

const rules: FormRules = {
  name: {
    required: true,
    message: '请填写部门名称！',
    trigger: ['input', 'blur'],
  },
  code: {
    required: true,
    message: '请填写部门编码！',
    trigger: ['input', 'blur'],
  },
}
const registerApi = async (params: registerOrg) => {
  return register<string>(params, { isMessage: true })
}
const updateApi = async (params: BaseOrg) => {
  return updateOrg<string>(params, { isMessage: true })
}
const afterApi = async () => {
  loading.value = false
  await getData({
    page: pagination.page,
    pageSize: pagination.pageSize,
    desc: false,
  })
}
const [
  isAdd,
  showModal,
  form,
  orgModel,
  modalStyle,
  handleRegister,
  submitCallback,
  clearModel,
  openModal,
  cancelCallback,
  modalTitle,
] = useModal<registerOrg>(
  registerApi,
  updateApi,
  afterApi,
  {
    name: '',
    code: '',
    sort: 100,
    pid: 0,
  },
  {},
  'Organize'
)

getData({ page: pagination.page, pageSize: pagination.pageSize, desc: false })
</script>

<template>
  <n-space vertical class="y1t-table-box">
    <n-form :label-width="'auto'" label-placement="left">
      <n-space>
        <n-form-item label="部门名称" path="name">
          <n-input v-model:value="searchData.name" type="text" placeholder="搜索条件" />
        </n-form-item>
        <n-button type="primary" @click="doSearch">
          <template #icon>
            <y-icon icon-type="SearchOutline" :depth="2" :size="15" color="white" />
          </template>
          搜索
        </n-button>
        <n-button @click="doReset">
          <template #icon>
            <y-icon icon-type="Reload" :depth="2" :size="15" />
          </template>
          重置
        </n-button>
      </n-space>
    </n-form>
    <n-space :size="5">
      <n-button type="primary" @click="handleRegister">
        <template #icon>
          <y-icon icon-type="Add" :depth="2" :size="17" color="white" />
        </template>
        新增
      </n-button>
    </n-space>
    <n-data-table
      v-model:checked-row-keys="checkedRowKeys"
      :columns="columns"
      :data="data"
      :scroll-x="650"
      :max-height="750"
      :loading="loading"
      :row-key="key2id"
    />
    <n-pagination
      v-model:page="pagination.page"
      v-model:pageSize="pagination.pageSize"
      :item-count="pagination.itemCount"
      @update:page="pagination.onChange"
    />
  </n-space>
  <n-modal v-model:show="showModal" :title="modalTitle()" :style="modalStyle" preset="card">
    <n-form
      ref="form"
      :model="orgModel"
      :rules="rules"
      label-placement="left"
      require-mark-placement="right-hanging"
      label-width="auto"
    >
      <n-form-item label="部门名称" path="name">
        <n-input v-model:value="orgModel.name" placeholder="名称" />
      </n-form-item>
      <n-form-item label="部门编码" path="code">
        <n-input v-model:value="orgModel.code" placeholder="编码" />
      </n-form-item>
      <n-form-item label="排序" path="sort">
        <n-input-number v-model:value="orgModel.sort" placeholder="排序" />
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
