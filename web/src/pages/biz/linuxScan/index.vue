<script setup lang="ts">
import { ref, h } from 'vue'
import { useTable } from '@/hooks/comHooks/useTable'
import { YIcon } from '@/components'
import {
  FormRules,
  NButton,
  NDivider,
  NPopconfirm,
  NSpace,
  SelectGroupOption,
  SelectOption,
  UploadFileInfo,
  useMessage,
} from 'naive-ui'
import { Download, Upload } from '@/api/common/file'
import { Page } from '@/api/common/types/login'
import { PageResult, Result } from '#axios'
import {
  Search,
  SearchWithPage,
  BizSSH,
  CreateSSH,
  SearchSSH,
  Create,
  UpdateSSH,
  Update,
  Delete,
  Run,
} from '@/api/biz/linuxScan'
import { completeMerger } from '@/utils/helper/objectHelper'
import { useModal } from '@/hooks/comHooks/useModal'
import { renderTaskState } from '@/utils/render'
import { TaskState, TaskStateMap } from '@/enums/bizEnum'
import { rName } from '@/enums/rName'
import { useEmit } from '@/hooks/comHooks/useEmit'
import BizSshResult from './components/result.vue'
import { formatToDate } from '@/utils/dateUtil'
import { isNullOrUnDef } from '@/utils/is'

const message = useMessage()
const bizSshResultRef = ref<InstanceType<typeof BizSshResult> | null>(null)

const fileList = ref<UploadFileInfo[]>([])
const { url, header } = Upload()
const handleDownload = async (file: UploadFileInfo) => {
  const url = Download({ id: file.id })
  const a = document.createElement('a')
  a.style.display = 'none'
  a.href = url
  a.download = file.name
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(a)
  URL.revokeObjectURL(url)
}
// @ts-ignore
const CallBackUpload = ({ file, event }: { file: UploadFileInfo; event?: ProgressEvent }) => {
  if (fileList.value.length > 0) {
    fileList.value.pop()
  }
  const res = JSON.parse((event?.target as XMLHttpRequest).response)
  Object.entries(res.result).forEach(([key, value]) => {
    sshModal.value.file = {
      id: value as string,
      name: key,
      type: '',
      path: '',
    }
    sshModal.value.fileId = value as string
  })
}
const CallBackRemove = ({ file, fileList }): boolean => {
  if (sshModal.value.fileId !== file.id || fileList.length > 1) {
    message.warning('程序出错，请重启！')
    return false
  }
  sshModal.value.fileId = null
  return true
}

const selectOptions: Array<SelectOption | SelectGroupOption> = [
  {
    label: TaskStateMap[TaskState.Build].str,
    value: TaskState.Build,
  },
  {
    label: TaskStateMap[TaskState.Ready].str,
    value: TaskState.Ready,
  },
  {
    label: TaskStateMap[TaskState.Doing].str,
    value: TaskState.Doing,
  },
  {
    label: TaskStateMap[TaskState.Finish].str,
    value: TaskState.Finish,
  },
]

const columns = [
  {
    title: '任务名称',
    key: 'name',
    width: 150,
  },
  {
    title: '任务简介',
    key: 'desc',
    width: 150,
  },
  {
    title: '任务状态',
    key: 'state',
    width: 150,
    render: (row) => {
      return renderTaskState(row.state)
    },
  },
  {
    title: '文件',
    key: 'file.name',
    width: 150,
  },
  {
    title: '并发量',
    key: 'thread',
    width: 100,
  },
  {
    title: '工作表',
    key: 'sheet',
    width: 100,
  },
  {
    title: 'IP所在列',
    key: 'ip',
    width: 100,
  },
  {
    title: '端口所在列',
    key: 'port',
    width: 100,
  },
  {
    title: '账号所在列',
    key: 'user',
    width: 100,
  },
  {
    title: '密码所在列',
    key: 'password',
    width: 100,
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
                    bizSshResultRef.value?.open(row.id)
                  },
                  text: true,
                },
                { default: () => '扫描结果' }
              ),
            ]
            if (row.state !== TaskState.Doing) {
              options.push(h(NDivider, { vertical: true }))
              options.push(
                h(
                  NButton,
                  {
                    onClick: () => {
                      isAdd.value = false
                      sshModal.value = row
                      fileList.value.pop()
                      if (row.file.name !== '') {
                        sshModal.value.fileId = row.file.id
                        fileList.value.push({
                          id: row.file.id,
                          name: row.file.name,
                          status: 'finished',
                        })
                      }
                      openModal()
                    },
                    text: true,
                  },
                  { default: () => '编辑' }
                )
              )
              options.push(h(NDivider, { vertical: true }))
              options.push(
                h(
                  NPopconfirm,
                  {
                    onPositiveClick: async () => {
                      await Delete<Result<string>>({ id: row.id }, { isMessage: true })
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
                )
              )
            }
            if (row.state === TaskState.Ready) {
              options.push(h(NDivider, { vertical: true }))
              options.push(
                h(
                  NPopconfirm,
                  {
                    onPositiveClick: async () => {
                      await Run<Result<string>>({ id: row.id }, { isMessage: true })
                      await getData({
                        page: pagination.page,
                        pageSize: pagination.pageSize,
                        desc: false,
                      })
                    },
                  },
                  {
                    trigger: () => h(NButton, { text: true }, { default: () => '运行' }),
                    default: () => '开始任务后将无法修改，请确认是否运行任务!',
                  }
                )
              )
            }
            return options
          },
        }
      )
    },
  },
]
const sTmpData = {
  name: '',
  state: null,
}
const tableApi = async (page: Page, searchData: any) => {
  return Search<PageResult<Array<BizSSH>>>(completeMerger<SearchWithPage>(page, searchData.value), {
    isMessage: false,
  })
}
const [pagination, loading, data, searchData, getData, doSearch, doReset, key2id, tableHeight] =
  useTable<BizSSH, SearchSSH>(
    tableApi,
    { page: 1, pageSize: 10, desc: false },
    sTmpData,
    rName.BIZ_LINUX_SCAN
  )

const rules: FormRules = {
  name: {
    required: true,
    message: '请填写页面标题！',
    trigger: ['input', 'blur'],
  },
  thread: {
    type: 'number',
    max: 2000,
    min: 1,
    message: '并发量必须在1-2000之间',
    trigger: ['input'],
  },
  sheet: {
    type: 'string',
    max: 50,
    min: 1,
    message: '工作表名必须在1-50之间',
    trigger: ['input'],
  },
  timeout: {
    type: 'number',
    max: 2000,
    min: 1,
    message: '超时等待时间必须在1-20之间',
    trigger: ['input'],
  },
}
const registerApi = async (params: CreateSSH): Promise<string> => {
  return Create<string>(params, { isMessage: true })
}
const updateApi = async (params: UpdateSSH): Promise<string> => {
  return Update<string>(params, { isMessage: true })
}
const afterApi = async (): Promise<void> => {
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
  sshModal,
  modalStyle,
  handleRegister,
  submitCallback,
  ,
  openModal,
  cancelCallback,
  modalTitle,
] = useModal<CreateSSH>(
  registerApi,
  updateApi,
  afterApi,
  {
    name: '',
    desc: '',
    fileId: null,
    thread: 5,
    sheet: '',
    ip: null,
    port: null,
    user: null,
    password: null,
    timeout: 5,
  },
  {},
  rName.BIZ_LINUX_SCAN
)
const { SaveEmit } = useEmit()

SaveEmit(
  () => getData({ page: pagination.page, pageSize: pagination.pageSize, desc: false }),
  rName.BIZ_LINUX_SCAN
)

getData({ page: pagination.page, pageSize: pagination.pageSize, desc: false })
</script>
<template>
  <n-space vertical class="qa-table-box">
    <n-form :label-width="'auto'" label-placement="left">
      <n-space>
        <n-form-item label="任务名称" path="name">
          <n-input v-model:value="searchData.name" type="text" placeholder="搜索条件" />
        </n-form-item>
        <n-form-item label="任务状态" path="state">
          <n-select v-model:value="searchData.state" :options="selectOptions" />
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
      :columns="columns"
      :data="data"
      :scroll-x="1600"
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
  <n-modal v-model:show="showModal" :title="modalTitle()" :style="modalStyle" preset="card">
    <n-form
      ref="form"
      :model="sshModal"
      :rules="rules"
      label-placement="left"
      require-mark-placement="right-hanging"
      label-width="auto"
    >
      <n-form-item label="任务名称" path="name">
        <n-input v-model:value="sshModal.name" placeholder="任务名称" />
      </n-form-item>
      <n-form-item label="任务简介" path="desc">
        <n-input v-model:value="sshModal.desc" placeholder="任务简介" />
      </n-form-item>
      <n-form-item label="文件" path="fileList">
        <n-upload
          :action="url"
          :headers="header"
          :default-file-list="fileList"
          list-type="image"
          show-download-button
          @finish="CallBackUpload"
          @remove="CallBackRemove"
          @download="handleDownload"
        >
          <n-button>上传文件</n-button>
        </n-upload>
      </n-form-item>
      <n-form-item label="并发量" path="thread">
        <n-input-number v-model:value="sshModal.thread" placeholder="并发量" />
      </n-form-item>
      <n-form-item label="工作表" path="sheet">
        <n-input v-model:value="sshModal.sheet" placeholder="工作表" />
      </n-form-item>
      <n-form-item label="IP所在列" path="ip">
        <n-input-number v-model:value="sshModal.ip" placeholder="IP所在列" />
      </n-form-item>
      <n-form-item label="端口所在列" path="port">
        <n-input-number v-model:value="sshModal.port" placeholder="端口所在列" />
      </n-form-item>
      <n-form-item label="账号所在列" path="user">
        <n-input-number v-model:value="sshModal.user" placeholder="账号所在列" />
      </n-form-item>
      <n-form-item label="密码所在列" path="password">
        <n-input-number v-model:value="sshModal.password" placeholder="密码所在列" />
      </n-form-item>
      <n-form-item label="连接超时事件" path="timeout">
        <n-input-number v-model:value="sshModal.timeout" placeholder="连接超时事件" />
      </n-form-item>
    </n-form>
    <template #action>
      <n-space :size="5" justify="end">
        <n-button type="tertiary" size="small" @click="cancelCallback">取消</n-button>
        <n-button type="success" size="small" @click="submitCallback">确认</n-button>
      </n-space>
    </template>
  </n-modal>
  <biz-ssh-result ref="bizSshResultRef"></biz-ssh-result>
</template>

<style scoped></style>
