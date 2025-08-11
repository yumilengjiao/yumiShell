<template>
  <div class="loading-spinner-container" v-if="visible">
    <div class="loading-spinner">
      <div class="spinner-circle"></div>
      <div class="spinner-text">{{ label }}</div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, watch } from 'vue'
import { useOtherConnStore } from '@renderer/stores/useOtherConnStore'
const otherConnStore = useOtherConnStore()
const label = ref('等待初始化连接...')
const visible = ref(true)
watch(() => otherConnStore.loading, (newVal) => {
  setTimeout(() => {
    visible.value = newVal
  }, 400)
})
watch(() => otherConnStore.isFailed, (newVal) => {
  if (newVal) {
    label.value = '连接失败,请重查看是否启动远程主机并重启程序...'
  }
})
</script>

<style scoped lang="scss">
.loading-spinner-container {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: rgba(0, 0, 0, 0.3);
  z-index: 9999;
}

.loading-spinner {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background-color: #fff;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.spinner-circle {
  width: 50px;
  height: 50px;
  border: 4px solid #f3f3f3;
  border-top: 4px solid #3498db;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 1rem;
}

.spinner-text {
  color: #333;
  font-size: 14px;
  font-weight: 500;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(360deg);
  }
}

/* 尺寸变体 */
.loading-spinner-container.small .spinner-circle {
  width: 30px;
  height: 30px;
  border-width: 3px;
}

.loading-spinner-container.medium .spinner-circle {
  width: 50px;
  height: 50px;
}

.loading-spinner-container.large .spinner-circle {
  width: 70px;
  height: 70px;
  border-width: 5px;
}
</style>