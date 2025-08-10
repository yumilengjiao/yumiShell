<template>
  <div class="tab-bar">
    <tiny-tabs v-model="activeName" size="small" :active-name="tabs[0]?.name" show-more-tabs @close="closeTab"
      :with-close="true">
      <tiny-tab-item :key="item.name" v-for="item in tabs" :title="item.title" :name="item.name">
      </tiny-tab-item>
      <template #moreIcon>
        <span style="color: var(--base-text-color);">...</span>
      </template>
    </tiny-tabs>
    <template v-for="item in tabs" :key="item.name">
      <ShellView :activeName="activeName" :sessionInfo="item" class="templateShell" v-show="activeName === item.name"
        ref="shellView" />
    </template>
    <div class="default-view" v-if="tabs.length === 0">
      <div class="default-view-content">
        <span class="title">Welcome to YumiShell</span>
        <span class="description">😁please click the new tab button to create a new tab😁</span>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, watch } from 'vue'
import ShellView from '@renderer/components/ShellView/index.vue'
import { useShellViewTabStore } from '@renderer/stores/useShellViewTabStore'
import { computed } from 'vue'
import { useDirectoryStore } from '@renderer/stores/useDirectoryStore'


const activeName = ref('0')
const shellViewTabStore = useShellViewTabStore()
const directoryStore = useDirectoryStore()
const tabs = computed(() => shellViewTabStore.tabs)
//在avtive改变的时候改变当前仓库里的currentUniqId
watch(activeName, (newVal) => {
  directoryStore.currentSessionUniqId = newVal
})

const closeTab = (name) => {
  if (name == activeName.value) {
    activeName.value = tabs.value[0]?.name
  }
  shellViewTabStore.closeTab(name)
}
</script>

<style scoped lang="scss">
@use '../../styles/variables.scss' as variables;

.tab-bar {
  height: calc(100vh - variables.$menu-bar-height);

  .default-view {
    height: 100%;
    width: 100%;
    background-color: var(--base-background-color);
    display: flex;
    justify-content: center;
    align-items: center;

    .default-view-content {
      height: 85%;
      width: 90%;
      background-color: var(--base-show-box-background-color);
      border-radius: 10px;
      box-shadow: 0 0 10px var(--base-shadow-color);

      .title {
        display: flex;
        justify-self: center;
        margin-top: 10%;
        font-size: 48px;
        font-weight: 500;
        color: var(--base-text-color);
      }

      .description {
        display: flex;
        justify-self: center;
        margin-top: 10%;
        font-size: 24px;
        font-weight: 500;
        color: var(--base-text-color);
      }


    }


  }

  .tiny-tabs {

    :deep(.tiny-tabs__header) {
      background: var(--base-background-color);
    }

    :deep(.tiny-tabs__content) {
      margin: 0;
      padding: 0;
    }
  }

  .templateShell {
    height: calc(100% - 40px);
    width: 100%;
  }
}
</style>