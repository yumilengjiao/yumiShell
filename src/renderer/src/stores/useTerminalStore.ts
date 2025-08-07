import { defineStore } from "pinia"
import { ref } from "vue"

export const useTerminalStore = defineStore('terminalStore', () => {
  const currentSessionUniqId = ref<string>('')
  let currentrows = ref<number>(34)
  let currentcols = ref<number>(90)
  return {
    currentSessionUniqId,
    currentrows,
    currentcols
  }
})