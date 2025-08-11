import { defineStore } from 'pinia'
import { onMounted, ref } from 'vue'
export const useOtherConnStore = defineStore('otherConnStore', () => {
  let loading = ref(true)
  //尝试后是否连接失败
  let isFailed = ref(false)


  const ws = ref<WebSocket | null>(null)
  const isConnected = ref(false)
  const isBackendReady = ref(false)
  const retryCount = ref(0)
  const maxRetries = 20  // 最大重试次数
  const retryDelay = 500  // 重试间隔(毫秒)

  // 尝试连接后端
  const connectToBackend = () => {

    // 如果已连接或超过最大重试次数，则不再尝试
    if (isConnected.value || retryCount.value >= maxRetries) {
      isFailed.value = true
      loading.value = false
      return
    }

    retryCount.value++
    console.log(`第 ${retryCount.value} 次尝试连接后端...`)

    // 先尝试发送HTTP请求检查后端是否就绪
    fetch('http://localhost:8977/health')
      .then(response => {
        if (response.ok) {
          console.log('后端已就绪，开始建立WebSocket连接')
          loading.value = false
          isBackendReady.value = true
          initWebSocket()
        } else {
          throw new Error('后端未就绪')
        }
      })
      .catch(error => {
        console.log(`连接检查失败: ${error.message}`)
        // 延迟后重试
        setTimeout(connectToBackend, retryDelay)
      })
  }

  // 初始化WebSocket连接
  const initWebSocket = () => {
    if (ws.value || !isBackendReady.value) {
      return
    }

    try {
      ws.value = new WebSocket("ws://localhost:8977/ws/others")

      ws.value.onopen = () => {
        console.log('WebSocket连接成功')
        isConnected.value = true
      }

      ws.value.onmessage = (e) => {
        if (e.data === 'ok') {
          console.log('后端程序初始化完成')
          // 可以在这里触发全局事件或更新状态
        }
      }

      ws.value.onerror = (error) => {
        console.error('WebSocket错误:', error)
        isConnected.value = false
        // 错误处理，可选择重试
        setTimeout(() => {
          ws.value = null
          initWebSocket()
        }, 1000)
      }

      ws.value.onclose = () => {
        console.log('WebSocket连接关闭')
        isConnected.value = false
        // 连接关闭后尝试重连
        setTimeout(() => {
          ws.value = null
          initWebSocket()
        }, 1000)
      }
    } catch (error) {
      console.error('创建WebSocket失败:', error)
      setTimeout(connectToBackend, retryDelay)
    }
  }

  onMounted(() => {
    // 应用启动后立即开始尝试连接
    connectToBackend()
  })
  return {
    ws,
    loading,
    isFailed
  }
})
