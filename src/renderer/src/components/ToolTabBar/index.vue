<template>
  <div class="tool-tab-bar">
    <func-button icon="session"></func-button>
    <func-button icon="file"></func-button>
    <func-button icon="setting"></func-button>
  </div>
</template>

<script lang="ts" setup>
//引入按钮组件
import FuncButton from './FuncButton.vue'
</script>

<style scoped lang="scss">
@use '../../styles/variables.scss' as variables;

.tool-tab-bar {
  height: calc(100vh - variables.$menu-bar-height);
  width: 100%;
  background-color: var(--base-background-color);
  display: flex;
  flex-direction: column;
}
</style>