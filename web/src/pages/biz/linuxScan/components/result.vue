<script setup lang="ts">
import { ref } from 'vue'
import { NSpace } from 'naive-ui'
import {
  BizSSHResult,
  SearchResult,
  SearchSSHResult,
  SearchSSHResultWithPage,
} from '@/api/biz/linuxScan'
import { PageResult } from '#axios'
import { Page } from '@/api/common/types/login'
import { completeMerger } from '@/utils/helper/objectHelper'
import { useTable } from '@/hooks/comHooks/useTable'
import { formatToDate } from '@/utils/dateUtil'
import { isNullOrUnDef } from '@/utils/is'
import { renderTaskResult } from '@/utils/render'

const show = ref<boolean>(false)

const columns = [
  {
    type: 'expand',
    expandable: () => true,
    renderExpand: (rowData) => {
      return renderTaskResult(rowData.items)
    },
  },
  {
    title: 'IP',
    key: 'addr',
    width: 150,
  },
  {
    title: '用户',
    key: 'user',
    width: 100,
  },
  {
    title: '密码',
    key: 'password',
    width: 150,
  },
  {
    title: '创建时间',
    key: 'createTime',
    width: 100,
    render: (row) => formatToDate(new Date(row.createTime)),
  },
  {
    title: '更新时间',
    key: 'updateTime',
    width: 100,
    render: (row) => (isNullOrUnDef(row.updateTime) ? '' : formatToDate(new Date(row.updateTime))),
  },
]
const sTmpData = {
  taskId: null,
}
const tableApi = async (page: Page, searchData: any) => {
  return SearchResult<PageResult<Array<BizSSHResult>>>(
    completeMerger<SearchSSHResultWithPage>(page, searchData.value),
    {
      isMessage: false,
    }
  )
}
const [pagination, loading, data, searchData, getData, , , key2id, tableHeight] =
  useTable<BizSSHResult, SearchSSHResult>(
    tableApi,
    { page: 1, pageSize: 10, desc: false },
    sTmpData,
    'SSHResult',
    0.6
  )

const open = (id: number) => {
  searchData.value.taskId = id
  getData({ page: pagination.page, pageSize: pagination.pageSize, desc: false })
  show.value = true
}
const close = () => (show.value = false)

defineExpose({
  open,
  close,
})
</script>

<template>
  <n-drawer v-model:show="show" :width="1000">
    <n-drawer-content title="页面按钮">
      <n-space vertical class="qa-table-box">
        <n-data-table
          :columns="columns"
          :data="data"
          :scroll-x="500"
          :max-height="tableHeight"
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
    </n-drawer-content>
  </n-drawer>
</template>

<style scoped></style>
