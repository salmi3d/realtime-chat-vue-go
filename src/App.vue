<template>
  <div id="app">
    <Header />
    <Chat :messages="messages" />
  </div>
</template>

<script>
import Header from './components/Header.vue'
import Chat from './components/Chat.vue'

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
  },
}
</script>
