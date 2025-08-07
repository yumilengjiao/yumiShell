<template>
  <div class="shell">
    <div class="shell-container" ref="container">
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Terminal } from '@xterm/xterm'
import { onMounted, ref, nextTick, watch } from 'vue'
//引入调整大小的插件
import { FitAddon } from '@xterm/addon-fit'
//引入webgl插件
import { WebglAddon } from '@xterm/addon-webgl'
//导入终端信息存储
import { useTerminalStore } from '@renderer/stores/useTerminalStore'
//导入xterm样式
import '@xterm/xterm/css/xterm.css'
//xterm的容器
const container = ref<HTMLDivElement | null>(null)
//调整大小插件
const fitAddon = new FitAddon()
//webgl插件
const webglAddon = new WebglAddon()
const terminalStore = useTerminalStore()
//用于处理交互的websocket
let terminalWebsocket: WebSocket
let directoryWebsocket: WebSocket
//xterm实例
let xterm: Terminal;
const props = defineProps(['activeName', 'sessionInfo'])
onMounted(() => {
  xterm = new Terminal({
    cursorBlink: false,
    allowTransparency: true,
    rows: 34,
    cols: 90
  });
  xterm.open(container.value!);
  //添加调整大小插件
  xterm.loadAddon(fitAddon);
  //添加webgl插件
  xterm.loadAddon(webglAddon);
  //初始化并处理websocket
  initWebsocket()
  handleWebsocket()
  //处理输入
  xterm.onData((data) => {
    terminalWebsocket.send(data)
  })
  window.addEventListener('resize', handleResize)
})

watch(() => props.activeName, (newVal, oldVal) => {
  if (newVal == oldVal) {
    return
  }
  nextTick(() => {
    xterm.resize(terminalStore.currentcols, terminalStore.currentrows)
  })
})


//处理窗口大小变化
const handleResize = (e) => {
  if (props.activeName == props.sessionInfo.name) {
    fitAddon.fit();
    terminalStore.currentcols = xterm.cols
    terminalStore.currentrows = xterm.rows
  }
}
const initWebsocket = () => {
  terminalWebsocket = new WebSocket("ws://localhost:8977/ws/term/" + props.sessionInfo.sessionId + '/' + props.sessionInfo.name)
  directoryWebsocket = new WebSocket("ws://localhost:8977/ws/file/" + props.sessionInfo.name)
}
const handleWebsocket = () => {
  terminalWebsocket.onopen = () => {
    console.log('terminalWebsocket连接成功');
  }
  terminalWebsocket.onmessage = (e) => {
    console.log(e.data);
    xterm.write(e.data)
  }
  terminalWebsocket.onerror = (e) => {
    console.log(e);
  }
  terminalWebsocket.onclose = (e) => {
    console.log(e);
  }
}

</script>

<style scoped lang="scss">
@use '../../styles/variables.scss' as variables;

.shell {
  width: 100%;
  height: 100%;

  .shell-container {
    width: 100%;
    height: 100%;
    background: var(--base-shell-bg-color);
    opacity: 0.95;
  }
}
</style>
