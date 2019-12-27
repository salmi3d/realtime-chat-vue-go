const websocketServerLocation = 'ws://localhost:4444/ws'

let ws

let connect = cb => {
  ws = new WebSocket(websocketServerLocation)

  ws.onmessage = msg => {
    cb(msg)
  }

  ws.onclose = () => {
    setTimeout(() => {
      connect(cb)
    }, 5000)
  }

}

let sendMsg = msg => ws.send(msg)

export { connect, sendMsg }
