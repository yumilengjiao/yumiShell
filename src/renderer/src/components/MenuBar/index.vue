<template>
  <div class="menu-bar"></div>
</template>

<script lang="ts" setup>

</script>

<style scoped lang="scss">
@use '../../styles/variables.scss' as variables;

.menu-bar {
  height: variables.$menu-bar-height;
  background-color: #3e3e42;
  -webkit-app-region: drag;
}
</style>