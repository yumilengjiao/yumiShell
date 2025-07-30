<template>
  <div class="tab-bar">
    <tiny-tabs v-model="activeName" size="small">
      <tiny-tab-item :key="item.name" v-for="item in tabs" :title="item.title" :name="item.name">
        <ShellView />
      </tiny-tab-item>
    </tiny-tabs>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive } from 'vue'
import { TinyTabs, TinyTabItem } from '@opentiny/vue'
import ShellView from '@renderer/components/ShellView/index.vue'

const activeName = ref('1')
const tabs = reactive<any>([])

// 创建 tabs
for (let i = 1; i < 5; i++) {
  const title = `Tab ${i}`
  tabs.push({
    title,
    name: i + '',
    content: `${title} content `
  })
}
//切换tab的回调
const toggleActive = (e) => {
  console.log(e);
}
</script>

<style scoped lang="scss">
@use '../../styles/variables.scss' as variables;

.tab-bar {
  height: calc(100vh - variables.$menu-bar-height);
  width: 100%;

  .tiny-tabs {
    height: 100%;
    width: 100%;

    :deep(.tiny-tabs__header) {
      background: variables.$tab-bar-bg-color;
    }

    :deep(.tiny-tabs__content) {
      height: calc(100% - 40px);
      width: 100%;
      margin: 0;
      padding: 0;

      .active-item {
        height: 100%;
        width: 100%;
      }
    }
  }
}
</style>