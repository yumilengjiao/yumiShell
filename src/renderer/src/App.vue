<template>
  <div class="app">
    <Layout />
    <SettingLayout v-if="toolBarStore.nowToolBarModel === 'setting'" />
  </div>
</template>

<script lang="ts" setup>
import Layout from "@renderer/layout/index.vue"
import SettingLayout from "@renderer/layout/SettingLayout.vue"
import { useToolBarStore } from '@renderer/stores/useToolBarStore'
import { useConfigStore } from '@renderer/stores/useConfigStore'

const toolBarStore = useToolBarStore()
const configStore = useConfigStore()

</script>

<style scoped lang="scss">
.app {
  border-radius: 10px;
  height: 100vh;
  width: 100vw;
}
</style>
