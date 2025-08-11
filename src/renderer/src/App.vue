<template>
  <div class="app">
    <Layout />
    <SettingLayout v-if="toolBarStore.nowToolBarModel === 'setting'" />
    <MessageDialog />
  </div>
</template>

<script lang="ts" setup>
import Layout from "@renderer/layout/index.vue"
import SettingLayout from "@renderer/layout/SettingLayout.vue"
import MessageDialog from '@renderer/components/MessageDialog/index.vue'
import { useToolBarStore } from '@renderer/stores/useToolBarStore'
import { useConfigStore } from '@renderer/stores/useConfigStore'

const toolBarStore = useToolBarStore()
//@ts-ignore
const configStore = useConfigStore()

</script>

<style scoped lang="scss">
.app {
  border-radius: 10px;
  height: 100vh;
  width: 100vw;
}
</style>
