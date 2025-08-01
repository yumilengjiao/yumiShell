<template>
  <div class="func-button-container">
    <button class="func-button" :class="{ 'clicked': isClicked }" @click="handleClick">
      <component :is="iconMap[icon]"></component>
    </button>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import Session from '@renderer/icon/Session.vue';
import File from '@renderer/icon/File.vue';
import Setting from '@renderer/icon/Setting.vue';
//导入piniaToolBar仓库，用于不同tool栏点击的时候切换视图
import { useToolBarStore } from '@renderer/stores/useToolBarStore'
const toolBarStore = useToolBarStore()

// 图标映射表
const iconMap = {
  session: Session,
  file: File,
  setting: Setting,
}
const icon = ref('')
const props = defineProps(['icon'])
switch (props.icon) {
  case 'session':
    icon.value = 'session'
    break;
  case 'file':
    icon.value = 'file'
    break;
  case 'setting':
    icon.value = 'setting'
    break;
}
//检查是否点击
const isClicked = ref<boolean>(false)
//点击事件
const handleClick = () => {
  isClicked.value = true
  setTimeout(() => {
    isClicked.value = false
  }, 400)
  //切换工具条
  toolBarStore.switchToolBarModel(props.icon)
}
</script>

<style scoped lang="scss">
.func-button-container {
  width: 100%;
  aspect-ratio: 1/1;

  .func-button {
    // 清除默认样式
    border: none;
    background-color: transparent;
    padding: 0;
    margin: 0;
    font: inherit;
    color: inherit;
    outline: none;
    cursor: pointer;
    -webkit-appearance: none;
    -moz-appearance: none;
    appearance: none;
    text-align: center;
    line-height: normal;
    overflow: visible;
    //自定义样式
    width: 100%;
    border-bottom: 1px solid #000000;
    height: 100%;
    background-color: #393939;
    transition: background-color 0.3s;

    &.clicked {
      background-color: #595959; // 点击后的颜色
    }
  }
}
</style>
