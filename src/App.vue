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
    connect(msg => {
      this.messages.push(JSON.parse(msg.data))
    })
  },
  methods: {
    send(event) {
      if (event.target.value.trim() === "") {
        return
      }
      sendMsg(event.target.value)
      event.target.value = ""
    }
  }
}
</script>
