<template>
  <div class="setting-layout">
    <SettingView />
  </div>
</template>

<script lang="ts" setup>
import SettingView from '@renderer/components/view/SettingView/index.vue'
</script>

<style scoped lang="scss">
@use '@renderer/styles/variables.scss' as variables;

.setting-layout {
  position: fixed;
  top: variables.$menu-bar-height;
  right: 0;
  bottom: 0;
  left: calc(1/24 * 100vw);
  background-color: var(--base-background-color);
  z-index: 5;
}
</style>