let ws
export const wsConnect = (url: string) => {
  ws = new WebSocket(url)

  ws.onmessage = (res) => {
    console.log(JSON.parse(res.data))
  }

  ws.onopen = (event) => {
    console.log('open', event)
  }
}

export const send = (msg: string) => {
  ws.send(msg)
}
