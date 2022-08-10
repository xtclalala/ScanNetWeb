<script setup lang="ts">
import { ref } from 'vue'
import { YIcon } from '@/components'
import { NButton, NSpace } from 'naive-ui'
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
import { renderTaskResult } from '@/utils/render'

const show = ref<boolean>(false)

// todo 创建时间 根据创建时间去搜索
const columns = [
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
    title: '操作系统',
    key: 'os',
    width: 100,
  },
  {
    title: '结果',
    key: 'result',
    width: 300,
    render: (row) => {
      const result = JSON.parse(row.result)
      return renderTaskResult(result)
    },
  },
]
const sTmpData = {
  os: '',
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
const [pagination, loading, data, searchData, getData, doSearch, doReset, key2id, tableHeight] =
  useTable<BizSSHResult, SearchSSHResult>(
    tableApi,
    { page: 1, pageSize: 10, desc: false },
    sTmpData,
    'SSHResult'
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
        <n-form :label-width="'auto'" label-placement="left">
          <n-space>
            <n-form-item label="操作系统" path="os">
              <n-input v-model:value="searchData.os" type="text" placeholder="搜索条件" />
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
        <n-data-table
          :columns="columns"
          :data="data"
          :scroll-x="800"
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
