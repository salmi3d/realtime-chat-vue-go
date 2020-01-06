const websocketServerLocation = 'ws://localhost:4444'

let ws

let connect = (name, cb) => {

  if (!window["WebSocket"]) {
    alert('Your browser does not support WebSockets')
    return
  }

  ws = new WebSocket(`${websocketServerLocation}?name=${name}`)

  ws.onmessage = msg => {
    cb(msg)
  }

  ws.onclose = () => {
    setTimeout(() => {
      connect(name, cb)
    }, 5000)
  }

}

let sendMsg = msg => ws.send(msg)

export { connect, sendMsg }
