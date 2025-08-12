<template>
  <div class="shell">
    <div class="shell-container" ref="container">
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Terminal } from '@xterm/xterm'
import { onMounted, ref, nextTick, watch } from 'vue'
//导入directory仓库
import { useDirectoryStore } from '@renderer/stores/useDirectoryStore'
//引入调整大小的插件
import { FitAddon } from '@xterm/addon-fit'
//引入webgl插件
import { WebglAddon } from '@xterm/addon-webgl'
//导入终端信息存储
import { useTerminalStore } from '@renderer/stores/useTerminalStore'
//导入配置store
import { useConfigStore } from '@renderer/stores/useConfigStore'
//导入其他连接store
import { useOtherConnStore } from '@renderer/stores/useOtherConnStore'
//导入xterm样式
import '@xterm/xterm/css/xterm.css'
//xterm的容器
const container = ref<HTMLDivElement | null>(null)
//调整大小插件
const fitAddon = new FitAddon()
//webgl插件
const webglAddon = new WebglAddon()
const terminalStore = useTerminalStore()
const directoryStore = useDirectoryStore()
const configStore = useConfigStore()
const otherConnStore = useOtherConnStore()

//用于处理交互的websocket
let terminalWebsocket: WebSocket
let directoryWebsocket: WebSocket
//xterm实例
let xterm: Terminal;
const props = defineProps(['activeName', 'sessionInfo'])
onMounted(() => {
  xterm = new Terminal({
    cursorBlink: configStore.config.basicConfig.isBlink,
    allowTransparency: true,
    rows: 33,
    cols: 90,
    theme: {
      foreground: configStore.config.basicConfig.shellTextColor,
      background: configStore.config.basicConfig.shellBgColor,
      cursor: configStore.config.basicConfig.cursorColor,
    }
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
  // 2) 复制：监听鼠标松开事件
  xterm.onSelectionChange(() => {
    const selected = xterm.getSelection();
    if (selected) {
      navigator.clipboard.writeText(selected);
    }
  });

  // 3) 粘贴：监听右键菜单或 Ctrl+Shift+V
  xterm.element?.addEventListener('contextmenu', (e) => {
    e.preventDefault();
    navigator.clipboard.readText().then(text => xterm.paste(text));
  });

  // 4) 可选：键盘粘贴（Ctrl+Shift+V）
  xterm.attachCustomKeyEventHandler((e) => {
    if (e.ctrlKey && e.shiftKey && e.code === 'KeyV') {
      navigator.clipboard.readText().then(text => xterm.paste(text));
      return false; // 阻止默认
    }
    return true;
  });
  window.addEventListener('resize', handleResize)
})
//监听活动会话变化
watch(() => props.activeName, (newVal, oldVal) => {
  if (newVal == oldVal) {
    return
  }
  nextTick(() => {
    xterm.resize(terminalStore.currentcols, terminalStore.currentrows)
  })
})
//处理变化的配置
watch(() => configStore.config, (newVal) => {
  console.log("config改变", newVal);
  xterm.options.cursorBlink = newVal.basicConfig.isBlink
  xterm.options.theme!.foreground = newVal.basicConfig.shellTextColor
  xterm.options.theme!.background = newVal.basicConfig.shellBgColor
  xterm.options.theme!.cursor = newVal.basicConfig.cursorColor
}, {
  deep: true
})


//处理窗口大小变化
const handleResize = (_) => {
  if (props.activeName == props.sessionInfo.name) {
    fitAddon.fit();
    terminalStore.currentcols = xterm.cols
    terminalStore.currentrows = xterm.rows
    //调整后端pty大小
    otherConnStore.ws!.send(JSON.stringify({
      reqType: 'fitness',
      content: {
        row: xterm.rows,
        col: xterm.cols,
      }
    }))

  }
}
const initWebsocket = () => {
  terminalWebsocket = new WebSocket(`ws://localhost:8977/ws/term/${props.sessionInfo.sessionId}/${props.sessionInfo.name}`);
  directoryWebsocket = new WebSocket(`ws://localhost:8977/ws/file/${props.sessionInfo.name}`);
  //设置二进制类型
  terminalWebsocket.binaryType = 'arraybuffer';
  directoryWebsocket.binaryType = 'arraybuffer';

  directoryStore.fileOperationObjList.push({
    uniqId: props.sessionInfo.name,
    fileList: [],
    currentPath: '/',
    websocket: directoryWebsocket
  })
}
const handleWebsocket = () => {
  terminalWebsocket.onopen = () => {
    console.log('terminalWebsocket连接成功');
  }
  terminalWebsocket.onmessage = (e) => {
    try {
      let data = e.data;
      // 处理二进制数据
      if (data instanceof ArrayBuffer) {
        // 尝试将ArrayBuffer解码为UTF-8字符串
        const decoder = new TextDecoder('utf-8');
        data = decoder.decode(data);
      }
      // 确保是字符串类型再写入终端
      if (typeof data === 'string') {
        xterm.write(data);
      } else {
        console.error('接收到非字符串数据:', data);
      }
    } catch (error) {
      console.error('WebSocket消息处理错误:', error);
    }
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
  background-color: var(--base-shell-bg-color);
  opacity: var(--base-shell-opacity);
  padding: 0 5px 0;

  .shell-container {
    width: 100%;
    height: 100%;
  }
}
</style>
