<template>
  <div class="tab-bar">
    <tiny-tabs v-model="activeName" size="small" @click="toggleActive">
      <tiny-tab-item :key="item.name" v-for="item in tabs" :title="item.title" :name="item.name">
      </tiny-tab-item>
    </tiny-tabs>
    <template v-for="(item, index) in tabs" :key="item.name">
      <ShellView :activeName="activeName" class="templateShell" v-show="activeName === index.toString()"
        ref="shellView" />
    </template>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive } from 'vue'
import { TinyTabs, TinyTabItem } from '@opentiny/vue'
import ShellView from '@renderer/components/ShellView/index.vue'

const activeName = ref('0')
const tabs = reactive<any>([])
// 创建 tabs
for (let i = 0; i < 4; i++) {
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

  .tiny-tabs {

    :deep(.tiny-tabs__header) {
      background: variables.$tab-bar-bg-color;
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