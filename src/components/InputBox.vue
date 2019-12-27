<template>
  <div class="inputbox">
    <input class="inputbox__input" @keyup.enter="send" placeholder="Type a message... Hit Enter to Send"/>
  </div>
</template>

<script>
export default {
  props: {
    send: { type: Function },
  }
}
</script>
