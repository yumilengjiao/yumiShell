<template>
  <div class="file-view">
    <div class="directory-container" v-for="(item, index) in currentPathSplit" :key="index">
      <a href="#" @click="handleDirectoryClick(item)">{{ item }}</a>
    </div>
    <div class="file-list">
      <FileItem v-for="(item, index) in currentOperationObj?.fileList" :key="index" :item="item" />
      <div class="default-view" v-if="currentOperationObj?.fileList.length === 0 || currentOperationObj == undefined">
        <svg t="1754743743438" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
          p-id="11570" width="200" height="200">
          <path d="M609.52381 97.52381H560.761905v243.809523h243.809524V292.571429h-195.047619z" fill="#A6A6A6"
            p-id="11571"></path>
          <path
            d="M659.260952 0H121.904762v1024h780.190476V228.547048L659.260952 0z m190.805334 975.238095H173.933714V48.761905h463.872l212.260572 199.923809V975.238095z"
            fill="#A6A6A6" p-id="11572"></path>
        </svg>
      </div>
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
const handleDirectoryClick = (item: string) => {
  //如果/切完长度大于1
  if (item.split('/').length > 2) {
    item = item.split('/')[1]
  }
  const index = directoryStore.currentOperationObj!.currentPath.indexOf(item)
  directoryStore.currentOperationObj!.currentPath = directoryStore.currentOperationObj!.currentPath.slice(0, index) + `${item}`
  directoryStore.currentOperationObj?.websocket?.send(JSON.stringify({
    requestType: 'get',
    path: directoryStore.currentOperationObj!.currentPath,
  }))


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
  color: var(--base-text-color);

  .directory-container {
    max-height: 100px;
    border: 1px dashed var(--base-border-color);
    background-color: var(--base-directory-background-color);
    line-height: 1.6;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .file-list {
    height: 100%;
    border: 1px dashed var(--base-border-color);
    display: flex;
    justify-content: flex-start;
    align-items: flex-start;
    align-content: flex-start;
    gap: 10px;
    flex-wrap: wrap;
    overflow: scroll;
    -ms-overflow-style: none;

    &::-webkit-scrollbar {
      display: none;
    }

    .default-view {
      height: 100%;
      width: 100%;
      background-color: var(--base-background-color);
      display: flex;
      justify-content: center;
      align-items: center;
    }

  }
}
</style>