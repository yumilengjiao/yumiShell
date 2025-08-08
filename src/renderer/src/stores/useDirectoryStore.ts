import { ref, watchEffect } from "vue";
import { defineStore } from "pinia";
import { FileOperationObjList, FileOperationObj } from "@renderer/types/directory";

export const useDirectoryStore = defineStore('directoryStore', () => {
  const fieleOperationObjList = ref<FileOperationObjList>([]);
  const currentSessionUniqId = ref<string>('');
  const currentOperationObj = ref<FileOperationObj>()
  //每个websocket的处理
  watchEffect(() => {
    fieleOperationObjList.value.forEach((item) => {
      let isFirstMessageProcessed = true;
      item.websocket.onmessage = (event) => {
        console.log('收到消息:', event.data);
        if (isFirstMessageProcessed) {
          item.currentPath = event.data
          isFirstMessageProcessed = false;
          return
        }
        item.fileList = JSON.parse(event.data)
      };
      item.websocket.onerror = (event) => {
        console.log('发生错误:', event);
      };
      item.websocket.onclose = (event) => {
        console.log('连接关闭:', event);
      };
    })
  })
  watchEffect(() => {
    if (currentSessionUniqId.value) {
      currentOperationObj.value = fieleOperationObjList.value.find((item) => item.uniqId === currentSessionUniqId.value)
    }
  })

  return {
    fieleOperationObjList,
    currentSessionUniqId,
    currentOperationObj
  };
}) 