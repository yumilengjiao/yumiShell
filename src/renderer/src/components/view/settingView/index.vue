<template>
  <div class="setting-view">

  </div>
</template>

<script lang="ts" setup>

</script>

<style scoped lang="scss">
@use '@renderer/styles/variables.scss' as variables;

.setting-view {
  width: 100%;
  height: calc(100vh - variables.$menu-bar-height);
  background-color: var(--base-background-color);
  border-left: 1px solid var(--base-border-color);
}
</style>