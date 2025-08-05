<template>
  <div class="shell">
    <div class="shell-container" ref="container">
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Terminal } from '@xterm/xterm';
import { onMounted, ref, nextTick, watch } from 'vue';
//引入调整大小的插件
import { FitAddon } from '@xterm/addon-fit';
//引入webgl插件
import { WebglAddon } from '@xterm/addon-webgl';
//导入xterm样式
import '@xterm/xterm/css/xterm.css'
//xterm的容器
const container = ref<HTMLDivElement | null>(null);
//调整大小插件
const fitAddon = new FitAddon();
//webgl插件
const webglAddon = new WebglAddon();
//xterm实例
let xterm: Terminal;
const props = defineProps(['activeName'])
onMounted(() => {
  xterm = new Terminal({
    cursorBlink: true,
    allowTransparency: true,
  });
  xterm.open(container.value!);
  //添加调整大小插件
  xterm.loadAddon(fitAddon);
  //添加webgl插件
  xterm.loadAddon(webglAddon);
  //调整大小
  fitAddon.fit();
  xterm.write('Welcome to Yumi Shell\r\n$');

  window.addEventListener('resize', handleResize)
})
//监视activeName的变化
watch(() => props.activeName, (newVal, oldVal) => {
  if (newVal === oldVal) {
    return
  }
  //调整大小
  nextTick(() => {
    fitAddon.fit();
  })
})
//处理窗口大小变化
const handleResize = (e) => {
  fitAddon.fit();
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
