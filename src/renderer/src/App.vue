<template>
  <div class="app">
    <Layout />
    <button class="switchTheme" @click="switchTheme">切换主题</button>
  </div>
</template>

<script lang="ts" setup>
import Layout from "@renderer/layout/index.vue"
import { useThemeStore } from "@renderer/stores/useThemeStore"

const themeStore = useThemeStore()

const switchTheme = () => {
  if (document.documentElement.classList.contains('dark')) {
    document.documentElement.classList.remove('dark')
    themeStore.theme = 'light'
  } else {
    document.documentElement.classList.add('dark')
    themeStore.theme = 'dark'
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
