import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useSessionsStore = defineStore('sessions', () => {
  const sessions = ref([
    {
      id: 1,
      label: 'Session 1',
      children: [
        {
          id: 1,
          label: 'Tab 1',
        },
      ],
    }
  ])
  return {
    sessions,
  }
})
