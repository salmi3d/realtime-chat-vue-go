<template functional>
  <div class="chat">
    <div class="chat__message" v-for="(message, index) in messages" :key="index">{{ message.body }}</div>
  </div>
</template>

<script>
export default {
  props: {
    messages: Array
  }
}
</script>
