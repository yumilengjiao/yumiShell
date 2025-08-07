<template>
  <div class="file-view">
    <div class="directory-container"></div>
  </div>
</template>

<script lang="ts" setup>

</script>

<style scoped lang="scss">
@use '@renderer/styles/variables.scss' as variables;

.file-view {
  width: 100%;
  height: calc(100vh - variables.$menu-bar-height);
  background-color: var(--base-background-color);
  display: flex;
  flex-direction: column;


}

// ... existing code ...</style>