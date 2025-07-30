<template>
  <div class="file-view">
  </div>
</template>

<script lang="ts" setup>

</script>

<style scoped lang="scss">
@use '../../styles/variables.scss' as variables;

.file-view {
  width: 100%;
  height: calc(100vh - variables.$menu-bar-height);
  background-color: pink;
}
</style>