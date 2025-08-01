<template>
  <div class="session-view">
    <tiny-tree-menu :show-filter="false" :data="sessionsStore.sessions" draggable></tiny-tree-menu>
  </div>
</template>

<script lang="ts" setup>
import { useSessionsStore } from '@renderer/stores/useSessionsStore'
const sessionsStore = useSessionsStore()
</script>

<style scoped lang="scss">
@use '@renderer/styles/variables.scss' as variables;

.session-view {
  width: 100%;
  height: calc(100vh - variables.$menu-bar-height);
  background-color: var(--base-background-color);
  border-left: 1px solid var(--base-border-color);
}
</style>