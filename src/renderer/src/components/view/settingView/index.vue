<template>
  <div class="setting-view">
    <Hr content="基本设置" percent="10%" />
    <div class="form">
      <FormItem label="用户名">
        <tiny-input v-model="configStore.config.basicConfig.username" placeholder="请输入用户名" />
      </FormItem>
      <FormItem label="shell透明度">
        <input type="range" min="0" max="1" step="0.01" v-model="configStore.config.basicConfig.shellOpacity">
        {{ configStore.config.basicConfig.shellOpacity }}
      </FormItem>
      <FormItem label="shell背景颜色">
        <input type="color" placeholder="请输入" v-model="configStore.config.basicConfig.shellBgColor">
      </FormItem>
      <FormItem label="shell字体颜色">
        <input type="color" placeholder="请输入" v-model="configStore.config.basicConfig.shellTextColor">
      </FormItem>
      <FormItem label="shell光标颜色">
        <input type="color" placeholder="请输入" v-model="configStore.config.basicConfig.cursorColor">
      </FormItem>
      <FormItem label="shell光标是否闪烁">
        <tiny-radio v-model="configStore.config.basicConfig.isBlink" label="true">是</tiny-radio>
        <tiny-radio v-model="configStore.config.basicConfig.isBlink" label="false">否</tiny-radio>
      </FormItem>
    </div>
    <Hr content="主题设置" percent="10%" />
    <div class="form">
      <FormItem label="基本背景颜色">
        <input type="color" placeholder="请输入" v-model="configStore.config.theme.backgroundColor">
      </FormItem>
      <FormItem label="菜单背景颜色">
        <input type="color" placeholder="请输入" v-model="configStore.config.theme.menuBgColor">
      </FormItem>
      <FormItem label="文字颜色">
        <input type="color" placeholder="请输入" v-model="configStore.config.theme.textColor">
      </FormItem>
      <FormItem label="边框颜色">
        <input type="color" placeholder="请输入" v-model="configStore.config.theme.borderColor">
      </FormItem>
      <FormItem label="工具栏背景颜色">
        <input type="color" placeholder="请输入" v-model="configStore.config.theme.toolbarBgColor">
      </FormItem>
      <FormItem label="目录背景颜色">
        <input type="color" placeholder="请输入" v-model="configStore.config.theme.directoryBgColor">
      </FormItem>
      <FormItem label="显示框背景颜色">
        <input type="color" placeholder="请输入" v-model="configStore.config.theme.showBoxBgColor">
      </FormItem>

    </div>
    <Hr content="重置和保存" percent="10%" />
    <div class="btn">
      <tiny-button type="primary" @click="resetTheme">reset</tiny-button>
      <tiny-button type="primary" @click="saveTheme">save</tiny-button>
    </div>
    <div class="btn" style="color: var(--base-text-color);">---保存后生效</div>
  </div>
</template>

<script lang="ts" setup>
import FormItem from './FormItem.vue'
import { useConfigStore } from '@renderer/stores/useConfigStore'
import Hr from './Hr.vue'
import { toRaw, watch } from 'vue';
// 配置store
const configStore = useConfigStore()

// 重置主题
const resetTheme = () => {
  configStore.config.basicConfig = {
    username: '',
    shellOpacity: 0.95,
    shellBgColor: '#0e0e0e',
    shellTextColor: 'white',
    cursorColor: 'white',
    isBlink: true
  }
  configStore.config.theme = {
    backgroundColor: "#efefef",
    menuBgColor: "#b3b3b3",
    textColor: "#000000",
    borderColor: "#000000",
    toolbarBgColor: "#383838",
    directoryBgColor: "#f5e0e0",
    showBoxBgColor: "#f2f2f2",
  }
}
//保存配置文件
const saveTheme = () => {
  window.api.saveConfig({
    basicConfig: toRaw(configStore.config.basicConfig),
    theme: toRaw(configStore.config.theme),
  })
  configStore.changeStyle()
}
</script>

<style scoped lang="scss">
@use '@renderer/styles/variables.scss' as variables;

.setting-view {
  height: 100%;
  width: 100%;
  background-color: var(--base-background-color);
  border-left: 1px solid var(--base-border-color);
  padding: 20px;

  .form {
    display: flex;
    flex-wrap: wrap;
    gap: 20px;
  }

  .btn {
    display: flex;
    justify-content: flex-end;
    margin-right: 20px;
  }

}
</style>