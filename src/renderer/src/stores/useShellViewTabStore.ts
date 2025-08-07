import { ref } from 'vue'
import { defineStore } from "pinia";
import type { tabItem } from '@renderer/types/tab'


export const useShellViewTabStore = defineStore('shellViewTabStore', () => {
  const tabs = ref<tabItem[]>([])
  let index = ref<number>(0)
  const closeTab = (name: string) => {
    tabs.value = tabs.value.filter(tab => tab.name !== name)
  }
  return {
    tabs,
    index,
    closeTab
  }
});
