<template>
  <div id="app">
    <Header />
    <Chat :messages="messages" />
    <InputBox :send="send" />
  </div>
</template>

<script>
import Header from './components/Header.vue'
import Chat from './components/Chat.vue'
import InputBox from './components/InputBox.vue'

import { connect, sendMsg } from './api.js'

export default {
  name: 'app',
  data() {
    return {
      messages: []
    }
  },
  components: {
    Header,
    Chat,
    InputBox,
  },
  mounted() {
    let name = prompt("What's your name fella?")
    connect(name, msg => {
      this.messages.push(JSON.parse(msg.data))
    })
  },
  methods: {
    send(event) {
      let msg = event.target.value.trim()
      if (msg === "") {
        return
      }
      this.messages.push({
        author: 'You',
        isMe: true,
        body: msg
      })
      sendMsg(event.target.value)
      event.target.value = ""
    }
  }
}
</script>
