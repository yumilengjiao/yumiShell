import { ref, watchEffect } from "vue";
import { defineStore } from "pinia";
import { FileOperationObjList, FileOperationObj } from "@renderer/types/directory";
import { watch } from "fs";
import { SftpRequest } from "@renderer/types/sftp";

export const useDirectoryStore = defineStore('directoryStore', () => {
  const fileOperationObjList = ref<FileOperationObjList>([]);
  const currentSessionUniqId = ref<string>('');
  const currentOperationObj = ref<FileOperationObj>()
  //每个websocket的处理
  watchEffect(() => {
    fileOperationObjList.value.forEach((item) => {
      let isFirstMessageProcessed = true;
      //处理第一次传输的path
      item.websocket.onmessage = (event) => {
        console.log('收到消息:', event.data);
        if (isFirstMessageProcessed) {
          item.currentPath = event.data
          isFirstMessageProcessed = false;
          return
        }
        //之后发送来的都是当前目录的数据文件
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
  //用于设置当前的文件操作对象
  watchEffect(() => {
    if (currentSessionUniqId.value) {
      currentOperationObj.value = fileOperationObjList.value.find((item) => item.uniqId === currentSessionUniqId.value)
    }
  })
  return {
    fileOperationObjList,
    currentSessionUniqId,
    currentOperationObj
  };
}) 