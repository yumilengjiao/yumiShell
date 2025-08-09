<template>
  <div class="file-item" @dragover="handleDragOver" @dragleave="handleDragLeave" @dragenter="handleDragEnter"
    @click="enterDirectory" @drop="handleDrop" ref="fileItemRef">
    <svg v-if="props.item.isDir" t="1754619954211" class="icon" viewBox="0 0 1024 1024" version="1.1"
      xmlns="http://www.w3.org/2000/svg" p-id="1673" width="50" height="50">
      <path
        d="M981.333333 266.666667v418.133333c-2.133333 32-27.733333 57.6-59.733333 55.466667H294.4c-19.2 2.133333-36.266667-6.4-46.933333-21.333334-10.666667-14.933333-14.933333-34.133333-10.666667-51.2l100.266667-290.133333H98.133333c-29.866667 2.133333-53.333333-21.333333-55.466666-49.066667v-149.333333C44.8 149.333333 68.266667 125.866667 98.133333 128h296.533334c21.333333 0 40.533333 10.666667 49.066666 27.733333l29.866667 55.466667h448c32-2.133333 57.6 23.466667 59.733333 55.466667z m0 0"
        fill="#185ABD" p-id="1674" data-spm-anchor-id="a313x.collections_detail.0.i8.16293a81UvVaXb" class="selected">
      </path>
      <path
        d="M87.466667 298.666667h849.066666c29.866667 0 44.8 14.933333 44.8 44.8V853.333333c0 29.866667-14.933333 44.8-44.8 44.8H87.466667C57.6 896 42.666667 881.066667 42.666667 851.2V343.466667c0-29.866667 14.933333-44.8 44.8-44.8z"
        fill="#34b2f3" p-id="1675" data-spm-anchor-id="a313x.collections_detail.0.i3.16293a81UvVaXb" class=""></path>
      <path
        d="M533.333333 454.4l34.133334 68.266667 76.8 10.666666c8.533333 2.133333 17.066667 8.533333 19.2 17.066667 2.133333 8.533333 0 19.2-6.4 25.6l-55.466667 53.333333 12.8 74.666667c2.133333 8.533333-2.133333 19.2-8.533333 23.466667-8.533333 6.4-17.066667 6.4-25.6 2.133333L512 693.333333l-66.133333 36.266667c-8.533333 4.266667-17.066667 4.266667-25.6-2.133333-6.4-6.4-10.666667-14.933333-8.533334-23.466667l12.8-74.666667-55.466666-53.333333c-6.4-6.4-8.533333-17.066667-6.4-25.6 2.133333-8.533333 10.666667-14.933333 19.2-17.066667l74.666666-10.666666 34.133334-68.266667c4.266667-8.533333 12.8-14.933333 21.333333-14.933333 10.666667 0 19.2 6.4 21.333333 14.933333z m0 0"
        fill="#34b2f3" opacity=".75" p-id="1676" data-spm-anchor-id="a313x.collections_detail.0.i1.16293a81UvVaXb"
        class=""></path>
    </svg>
    <svg v-if="!props.item.isDir" t="1754620070396" class="icon" viewBox="0 0 1024 1024" version="1.1"
      xmlns="http://www.w3.org/2000/svg" p-id="1929" width="50" height="50">
      <path
        d="M663.466667 755.2c78.933333 19.2 183.466667-10.666667 232.533333-46.933333-57.6 78.933333-200.533333 224-283.733333 268.8 42.666667-44.8 59.733333-164.266667 51.2-221.866667z"
        fill="#185ABD" p-id="1930"></path>
      <path
        d="M853.333333 42.666667H170.666667c-23.466667 0-42.666667 19.2-42.666667 42.666666v853.333334c0 23.466667 19.2 42.666667 42.666667 42.666666h281.6c185.6 0 128-313.6 128-313.6 115.2 29.866667 315.733333 17.066667 315.733333-117.333333V85.333333c0-23.466667-19.2-42.666667-42.666667-42.666666z"
        fill="#41A5EE" p-id="1931"></path>
      <path
        d="M746.666667 469.333333c0 12.8-8.533333 21.333333-21.333334 21.333334H298.666667c-12.8 0-21.333333-8.533333-21.333334-21.333334v-21.333333c0-12.8 8.533333-21.333333 21.333334-21.333333h426.666666c12.8 0 21.333333 8.533333 21.333334 21.333333v21.333333zM746.666667 298.666667c0 12.8-8.533333 21.333333-21.333334 21.333333H298.666667c-12.8 0-21.333333-8.533333-21.333334-21.333333v-21.333334c0-12.8 8.533333-21.333333 21.333334-21.333333h426.666666c12.8 0 21.333333 8.533333 21.333334 21.333333v21.333334z"
        fill="#FFFFFF" p-id="1932"></path>
    </svg>
    <div class="name">{{ props.item.name }}</div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { useDirectoryStore } from '@renderer/stores/useDirectoryStore'
import { SftpRequest } from '@renderer/types/sftp'

const DOUBLE_CLICK_THRESHOLD = 300
const props = defineProps(['item'])
const fileItemRef = ref<HTMLDivElement>()
const directoryStore = useDirectoryStore()
//定义一个起始时间的变量
let lastClickTime = ref<number>(0)

//拖拽进入时的处理方法
const handleDragEnter = (e: DragEvent) => {
  if (!props.item.isDir) return
  console.log("开始推拽")
  e.preventDefault()
  //加上蒙版
  fileItemRef.value?.classList.add('drag-over')
}

//拖拽进入后的处理方法
const handleDragOver = (e: DragEvent) => {
  if (!props.item.isDir) return
  console.log("开始推拽")
  e.preventDefault()
}
//拖拽离开的处理方法
const handleDragLeave = (e: DragEvent) => {
  if (!props.item.isDir) return
  console.log("结束推拽")
  e.preventDefault()
  //移除蒙版
  setTimeout(() => {
    fileItemRef.value?.classList.remove('drag-over')
  }, 200)
}
//处理拖拽放下的事件
const handleDrop = (e: DragEvent) => {
  if (!props.item.isDir) return
  e.preventDefault()
  setTimeout(() => {
    fileItemRef.value?.classList.remove('drag-over')
  }, 200)
  //将文件发送后端处理
  //获取当前数据和拖拽放下的文件夹并拼在一起
  console.log(directoryStore);
  console.log(e.dataTransfer?.files[0].name);
  console.log(e.dataTransfer?.files[0].size);
  console.log(props.item);
  //上传文件的路径
  const uploadPath = directoryStore.currentOperationObj!.currentPath
  //上传文件的对象
  const files = Array.from(e.dataTransfer!.files)
  files.forEach(async (item) => {
    //将每个文件发送给后端进行上传
    const sftpOnj: SftpRequest = {
      requestType: 'upload',
      path: uploadPath + "/" + props.item.name,

      file: {
        name: item.name,
        content: await readFileAsBase64(item),
      }
    }
    console.log("读取完成发送请求");
    // 获取现在的websocket发送请求
    directoryStore.currentOperationObj?.websocket?.send(JSON.stringify(sftpOnj))
  })
}
//双击文件夹进入
const enterDirectory = () => {
  if (!props.item.isDir) return
  //判断是否是双击
  const currentTime = Date.now();
  const isDoubleClick = (currentTime - lastClickTime.value) < DOUBLE_CLICK_THRESHOLD;
  lastClickTime.value = currentTime;
  // 仅在双击时执行操作
  if (isDoubleClick) {
    //改变当前路径
    if (props.item.name === '..') {
      //返回上一级
      const currentPath = directoryStore.currentOperationObj!.currentPath
      currentPath.lastIndexOf('/')
      directoryStore.currentOperationObj!.currentPath = currentPath.slice(0, currentPath.lastIndexOf('/'))
      if (currentPath.lastIndexOf('/') === 0) {
        directoryStore.currentOperationObj!.currentPath = '/'
      }
      //发送请求
      directoryStore.currentOperationObj?.websocket?.send(JSON.stringify({
        requestType: 'get',
        path: directoryStore.currentOperationObj!.currentPath,
      }))
      return
    }
    if (props.item.name === '.') {
      //返回当前目录
      return
    }
    if (directoryStore.currentOperationObj!.currentPath === '/') {
      directoryStore.currentOperationObj!.currentPath = directoryStore.currentOperationObj!.currentPath + props.item.name
    } else {
      directoryStore.currentOperationObj!.currentPath = directoryStore.currentOperationObj!.currentPath + "/" + props.item.name
    }
    directoryStore.currentOperationObj?.websocket?.send(JSON.stringify({
      requestType: 'get',
      path: directoryStore.currentOperationObj!.currentPath,
    }))
  }
}
// 添加文件读取工具函数
const readFileAsBase64 = (file: File): Promise<string> => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader()
    reader.onload = () => {
      // 提取Base64内容（去掉data URL前缀）
      const result = reader.result as string
      const base64Content = result.split(',')[1]
      resolve(base64Content)
    }
    reader.onerror = reject
    // 以DataURL形式读取文件
    reader.readAsDataURL(file)
  })
}
</script>

<style scoped lang="scss">
@use '@renderer/styles/variables.scss' as variables;

.file-item {
  width: 75px;
  height: 75px;
  aspect-ratio: 1/1;
  margin-bottom: 30px;
  position: relative;
  overflow: hidden;

  &.drag-over::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    border-radius: 10%;
    background-color: variables.$drag-over-color;
    opacity: 1;
    /* 默认隐藏 */
    transition: opacity 0.2s;
    /* 平滑过渡 */
    z-index: 1;
    /* 确保覆盖在内容上方 */
  }

  .icon {
    display: block;
    margin: 0;
    padding: 0;
    pointer-events: none;
  }

  .name {
    margin-top: 0;
    font-size: 10px;
  }
}
</style>