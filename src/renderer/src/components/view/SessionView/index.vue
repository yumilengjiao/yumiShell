<template>
  <div class="session-view">
    nihao
  </div>
</template>

<script lang="ts" setup>

</script>

<style scoped lang="scss">
@use '@renderer/styles/variables.scss' as variables;

.session-view {
  width: 100%;
  height: calc(100vh - variables.$menu-bar-height);
  background-color: variables.$controller-bg-color;
}
</style>