<template>
  <div class="app">
    <Layout />
    <button class="switchTheme" @click="switchTheme">切换主题</button>
  </div>
</template>

<script lang="ts" setup>
import Layout from "@renderer/layout/index.vue"

const switchTheme = () => {
  if (document.documentElement.classList.contains('dark')) {
    document.documentElement.classList.remove('dark')
  } else {
    document.documentElement.classList.add('dark')
  }
}
</script>

<style scoped lang="scss">
.app {
  border-radius: 10px;
  height: 100vh;
  width: 100vw;

  .switchTheme {
    position: fixed;
    top: 20px;
    right: 20px;
  }
}
</style>
