<template>
  <div class="file-view">
    <div class="directory-container" v-for="(item, index) in currentPathSplit" :key="index">
      <div class="directory">
        {{ item }}
      </div>
    </div>
    <div class="file-list">
      <FileItem v-for="(item, index) in currentOperationObj?.fileList" :key="index" :item="item" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useDirectoryStore } from '@renderer/stores/useDirectoryStore';
import { computed, ref, watch } from 'vue';

const directoryStore = useDirectoryStore();
const currentOperationObj = computed(() => directoryStore.currentOperationObj)
const currentPathSplit = ref<string[]>([])
watch(() => currentOperationObj.value, (newVal) => {
  if (newVal) {
    currentPathSplit.value = mergePath(newVal.currentPath.split('/'))
  }
}, { deep: true })

//用于合并索引为3之后的元素
const mergePath = (pathSplit: string[]) => {
  if (!pathSplit) {
    return ['/']
  }
  if (pathSplit.every(item => item === '')) {
    return ['/']
  }

  if (pathSplit.length <= 4) {
    pathSplit.map((item, index) => {
      pathSplit![index] = `/${item}`
    })
    return pathSplit
  } else {
    pathSplit = [
      ...pathSplit.slice(0, 3),
      pathSplit.slice(3).join('/')
    ]
    pathSplit.map((item, index) => {
      pathSplit![index] = `/${item}`
    })
    return pathSplit
  }
}
</script>

<style scoped lang="scss">
@use '@renderer/styles/variables.scss' as variables;

.file-view {
  width: 100%;
  height: calc(100vh - variables.$menu-bar-height);
  background-color: var(--base-background-color);
  display: flex;
  flex-direction: column;
  font-size: 20px;
  line-height: 2;
  color: var(--base-text-color);

  .directory-container {
    border: 1px dashed var(--base-border-color);

    .directory {
      background-color: var(--base-directory-background-color);
    }
  }

  .directory-container:last-child {
    border: 1px dashed var(--base-border-color);

    .directory {
      background-color: var(--base-directory-background-color);
    }
  }

  .file-list {
    height: 100%;
    border: 1px dashed var(--base-border-color);
    display: flex;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;
    overflow: scroll;
    -ms-overflow-style: none;

    &::-webkit-scrollbar {
      display: none;
    }
  }
}
</style>