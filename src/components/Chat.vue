<template>
  <div class="chat">
    <div class="chat__entry" v-for="(message, index) in messages" :key="index">
      <div class="chat__person" :title="message.author">{{ getShortName(message.author) }}</div>
      <div class="chat__message" :class="{'chat__message_me': message.isMe}">
        <div class="chat__time">{{ message.time }}</div>
        {{ message.text }}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    messages: Array
  },
  methods: {
    getShortName(name) {
      if(name === 'You'){
        return name
      }
      return name.substring(0, 1).toUpperCase()
    }
  }
}
</script>
