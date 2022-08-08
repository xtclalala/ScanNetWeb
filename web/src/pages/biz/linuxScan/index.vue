<script setup lang="ts">
import { ref } from 'vue'
import { UploadFileInfo } from 'naive-ui'
import { Download, Upload } from '@/api/common/file'
import { send, wsConnect } from '@/api/common/ws'
const fileList = ref<UploadFileInfo[]>([
  {
    id: '4d9760c3-4c23-42ed-a13e-e4fdf64bc3bb',
    name: 'in.xlsx',
    status: 'finished',
  },
])
wsConnect('ws://127.0.0.1:33333/ws')

const handleDownload = async (file: UploadFileInfo) => {
  const url = Download({ id: file.id })
  const a = document.createElement('a')
  a.style.display = 'none'
  a.href = Download({ id: file.id })
  a.download = file.name
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(a)
  URL.revokeObjectURL(url)
}
const st = () => {
  send('111111')
}

const { url, header } = Upload()
</script>
<template>
  <h1>test</h1>
  <n-upload
    :action="url"
    :headers="header"
    :default-file-list="fileList"
    list-type="image"
    show-download-button
    @download="handleDownload"
  >
    <n-button>上传文件</n-button>
  </n-upload>
  <n-button @click="st">send</n-button>
</template>

<style scoped></style>
