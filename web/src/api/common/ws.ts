import { useViewStoreWithOut } from '@/store/module/views'
import { h } from 'vue'
import { useEmit } from '@/hooks/comHooks/useEmit'
import { NButton } from 'naive-ui'
import { useRouter } from 'vue-router'

let ws
export const wsConnect = (url: string) => {
  if (ws === undefined) {
    ws = new WebSocket(url)
  }

  ws.onmessage = (res) => {
    const resData = JSON.parse(res.data)

    const view = useViewStoreWithOut()
    const router = useRouter()
    let str
    if (view.currentView === resData.task) {
      const { GetEmit } = useEmit()
      const fn = GetEmit(resData.task)
      if (fn !== null) {
        fn()
      }
    } else {
      str = '前往'
    }
    let n = window.$notification?.create({
      title: resData.title,
      content: resData.context,
      type: resData.state,
      action: () =>
        h(
          NButton,
          {
            text: true,
            type: 'primary',
            onClick: () => {
              router.push({ name: resData.task })
              n?.destroy()
            },
          },
          {
            default: () => str,
          }
        ),
    })
  }

  ws.onopen = (event) => {
    console.log('open', event)
    window.$notification?.success({
      title: 'WebSocket',
      content: 'ws连接成功！',
    })
  }

  ws.onerror = (a) => {
    console.log(a)
    window.$notification?.error({
      title: 'WebSocket',
      content: 'ws连接失败！',
    })
  }
}

export const send = (msg: string) => {
  ws.send(msg)
}
