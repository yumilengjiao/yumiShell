import { defineStore } from 'pinia'
export const useOtherConnStore = defineStore('otherConnStore', () => {
  const ws = new WebSocket("ws://localhost:8977/ws/others")
  return {
    ws
  }
})
