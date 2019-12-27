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
  methods: {
    send(event) {
      alert(event.target.value)
    }
  }
}
</script>
