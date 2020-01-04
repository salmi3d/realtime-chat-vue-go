<template>
  <div class="chat">
    <div class="chat__entry" v-for="(message, index) in messages" :key="index">
      <div class="chat__person">{{ message.author }}</div>
      <div class="chat__message" :class="{'chat__message_me': message.isMe}">{{ message.body }}</div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    messages: Array
  }
}
</script>
