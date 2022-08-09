let ws
export const wsConnect = (url: string) => {
  if (ws === undefined) {
    ws = new WebSocket(url)
  }

  ws.onmessage = (res) => {
    const resData = JSON.parse(res.data)

    if (resData.task) {
      // todo 如果当前页在扫描任务列表的页面，需要刷新列表
      // todo 如果不在当前页需要一个跳转按钮
    }
    window.$notification?.create({
      title: resData.title,
      content: resData.context,
      type: resData.state,
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
