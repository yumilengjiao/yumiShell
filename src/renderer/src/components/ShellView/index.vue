<template>
  <div class="shell">
    <div class="shell-container" ref="container">
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Terminal } from '@xterm/xterm';
import { onMounted, ref } from 'vue';
//引入调整大小的插件
import { FitAddon } from '@xterm/addon-fit';
//导入xterm样式
import '@xterm/xterm/css/xterm.css'
//xterm的容器
const container = ref<HTMLDivElement | null>(null);
//调整大小插件
const fitAddon = new FitAddon();
//xterm实例
let xterm: Terminal;
onMounted(() => {
  xterm = new Terminal({
    cursorBlink: true,
    allowTransparency: true
  });
  xterm.open(container.value!);
  //添加调整大小插件
  xterm.loadAddon(fitAddon);
  //调整大小
  fitAddon.fit();
  xterm.write('Welcome to Yumi Shell\r\n$');
})

</script>

<style scoped lang="scss">
@use '../../styles/variables.scss' as variables;

.shell {
  width: 100%;
  height: 100%;

  .shell-container {
    width: 100%;
    height: 100%;
    background: variables.$shell-bg-color;
    opacity: variables.$shell-opacity;
  }
}
</style>
