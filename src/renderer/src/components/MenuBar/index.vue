<template>
  <div class="menu-bar">
    <div class="os-button">
      <tiny-button type="primary" :icon="IconMinus" circle @click="minimizeWindow"></tiny-button>
      <tiny-button v-if="isMaximized" type="primary" :icon="IconExpand" circle
        @click="windowedOrFullscreen"></tiny-button>
      <tiny-button v-if="!isMaximized" type="primary" :icon="IconPutAway" circle
        @click="windowedOrFullscreen"></tiny-button>
      <tiny-button type="primary" :icon="IconClose" circle @click="closeWindow"></tiny-button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
//导入三个icon图标的执行函数
import { iconPutAway, iconExpand, iconMinus, iconClose } from '@opentiny/vue-icon'
//三(四)个图标
const IconPutAway = iconPutAway()
const IconExpand = iconExpand()
const IconMinus = iconMinus()
const IconClose = iconClose()
//是否最大化
let isMaximized = ref(false)
//三个图标的回调函数
const closeWindow = () => {
  window.api.closeWindow()
}
const minimizeWindow = () => {
  window.api.minimizeWindow()
}
const windowedOrFullscreen = () => {
  if (!isMaximized.value) {
    window.api.maximizeWindow()
    isMaximized.value = true
  } else {
    window.api.windowedWindow()
    isMaximized.value = false
  }
}
</script>

<style scoped lang="scss">
@use '../../styles/variables.scss' as variables;

.menu-bar {
  height: variables.$menu-bar-height;
  background-color: var(--base-menu-background-color);
  -webkit-app-region: drag;

  .os-button {
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: flex-end;
    padding-right: 5px;

    .tiny-button {
      -webkit-app-region: no-drag;
      margin: 0 2px;
    }
  }
}
</style>