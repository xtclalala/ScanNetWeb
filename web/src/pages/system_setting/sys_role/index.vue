<script lang="ts">
export default {
  name: 'Y1tSysRole',
}
</script>
<script setup lang="ts">
import { h, ref } from 'vue'
import { FormRules, NButton, NDivider, NPopconfirm, NSpace } from 'naive-ui'
import type {
  BaseRole,
  Page,
  registerRole,
  SearchRole,
  SearchRoleWithPage,
} from '@/api/system_setting/types/sys_role'
import { PageResult } from '#axios'
import { completeMerger } from '@/utils/helper/objectHelper'
import { useTable } from '@/hooks/comHooks/useTable'
import {
  copyRole,
  deleteRole,
  register,
  searchRole,
  updateRole,
} from '@/api/system_setting/sys_role'
import { YIcon } from '@/components'
import { useModal } from '@/hooks/comHooks/useModal'
import { BaseOrg } from '@/api/system_setting/types/sys_organization'
import { selectOrg } from '@/api/system_setting/sys_organize'
import { SetPermissions } from './components'

const drawerShow = ref<boolean>(false)
const currentRole = ref<number>(0)
const defaultMenuData = ref<number[]>([])
const defaultPerData = ref<number[]>([])
const options = ref<Array<BaseOrg>>([])

const allOrganizations = async () => {
  options.value = await selectOrg<Array<BaseOrg>>({ name: '' }, { isMessage: false })
}
const setPerAfter = async () => {
  await getData({
    page: pagination.page,
    pageSize: pagination.pageSize,
    desc: false,
  })
}

const columns = [
  {
    type: 'selection',
    options: ['all', 'none'],
  },
  {
    title: '角色名称',
    key: 'name',
    width: 150,
  },
  {
    title: '角色编码',
    key: 'code',
    width: 150,
  },
  {
    title: '组织',
    key: 'organize.name',
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
                    roleModel.value.name = row.name
                    roleModel.value.code = row.code
                    roleModel.value.id = row.id
                    roleModel.value.sort = row.sort
                    roleModel.value.orgId = row.organize.id
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
                    await deleteRole({ id: row.id }, { isMessage: true })
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
                NPopconfirm,
                {
                  onPositiveClick: async () => {
                    await copyRole({ id: row.id }, { isMessage: true })
                    await getData({
                      page: pagination.page,
                      pageSize: pagination.pageSize,
                      desc: false,
                    })
                  },
                },
                {
                  trigger: () => h(NButton, { text: true }, { default: () => '复制角色' }),
                  default: () => '请确认是否复制!',
                }
              ),
              h(NDivider, { vertical: true }),
              h(
                NButton,
                {
                  onClick: () => {
                    currentRole.value = row.id
                    defaultMenuData.value = row.menus.flatMap((item) => item.id)
                    defaultPerData.value = row.permissions.flatMap((item) => item.id)
                    drawerShow.value = true
                  },
                  text: true,
                },
                { default: () => '授权' }
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
  orgId: null,
}
const tableApi = async (page: Page, searchData: any) => {
  return searchRole<PageResult<Array<registerRole>>>(
    completeMerger<SearchRoleWithPage>(page, searchData.value),
    { isMessage: false }
  )
}
const [pagination, loading, data, searchData, getData, doSearch, doReset, key2id] = useTable<
  registerRole,
  SearchRole
>(tableApi, { page: 1, pageSize: 10, desc: false }, sTmpData, 'Role')
const checkedRowKeys = ref([])

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
  orgName: {
    required: true,
    message: '请选择组织！',
    trigger: ['blur', 'change'],
  },
}
const registerApi = async (params: BaseRole) => {
  return register<string>(params, { isMessage: true })
}
const updateApi = async (params: BaseRole) => {
  return updateRole<string>(params, { isMessage: true })
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
  roleModel,
  modalStyle,
  handleRegister,
  submitCallback,
  ,
  openModal,
  cancelCallback,
  modalTitle,
] = useModal<BaseRole>(
  registerApi,
  updateApi,
  afterApi,
  { name: '', code: '', sort: 100, orgId: null },
  {},
  'Organize'
)

allOrganizations()
getData({ page: pagination.page, pageSize: pagination.pageSize, desc: false })
</script>
<template>
  <n-space vertical class="y1t-table-box">
    <n-form :label-width="'auto'" label-placement="left">
      <n-space>
        <n-form-item label="角色名称" path="name">
          <n-input v-model:value="searchData.name" type="text" placeholder="搜索条件" />
        </n-form-item>
        <n-form-item label="组织" path="orgId">
          <n-select
            v-model:value="searchData.orgId"
            value-field="id"
            label-field="name"
            filterable
            :options="options"
            clearable
          />
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
      :scroll-x="800"
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
      :model="roleModel"
      :rules="rules"
      label-placement="left"
      require-mark-placement="right-hanging"
      label-width="auto"
    >
      <n-form-item label="角色名称" path="name">
        <n-input v-model:value="roleModel.name" placeholder="名称" />
      </n-form-item>
      <n-form-item label="角色编码" path="code">
        <n-input v-model:value="roleModel.code" placeholder="编码" />
      </n-form-item>
      <n-form-item label="排序" path="sort">
        <n-input-number v-model:value="roleModel.sort" placeholder="排序" />
      </n-form-item>
      <n-form-item label="组织" path="orgId">
        <n-select
          v-model:value="roleModel.orgId"
          value-field="id"
          label-field="name"
          filterable
          :options="options"
          clearable
        />
      </n-form-item>
    </n-form>
    <template #action>
      <n-space :size="5" justify="end">
        <n-button type="tertiary" size="small" @click="cancelCallback">取消</n-button>
        <n-button type="success" size="small" @click="submitCallback">确认</n-button>
      </n-space>
    </template>
  </n-modal>
  <set-permissions
    v-model:show="drawerShow"
    :current-role="currentRole"
    :default-menu-data="defaultMenuData"
    :default-per-data="defaultPerData"
    @after-update="setPerAfter"
  />
</template>
<style scoped></style>
