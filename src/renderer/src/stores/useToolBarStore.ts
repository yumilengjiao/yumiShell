import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useToolBarStore = defineStore('toolBar', () => {
  const nowToolBarModel = ref('session')
  const toolBarModelList = ref(['session', 'file', 'setting'])
  //判断工具条是否存在
  function isToolBarModelExist(model: string) {
    return toolBarModelList.value.includes(model)
  }
  //切换工具条
  function switchToolBarModel(model: string) {
    if (isToolBarModelExist(model)) {
      nowToolBarModel.value = model
    }
  }
  return {
    nowToolBarModel,
    switchToolBarModel,
  }
})