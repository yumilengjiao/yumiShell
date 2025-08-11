<template>
  <div class="setting-view">
    <Hr content="basicSetting" percent="10%" />
    <div class="form">
      <FormItem label="username">
        <tiny-input v-model="configStore.config.basicConfig.username" placeholder="please input username" />
      </FormItem>
      <FormItem label="shellOpacity">
        <input type="range" min="0" max="1" step="0.01" v-model="configStore.config.basicConfig.shellOpacity">
        {{ configStore.config.basicConfig.shellOpacity }}
      </FormItem>
      <FormItem label="shellBgColor">
        <input type="color" placeholder="please input shell bg color"
          v-model="configStore.config.basicConfig.shellBgColor">
      </FormItem>
      <FormItem label="shellTextColor">
        <input type="color" placeholder="please input shell text color"
          v-model="configStore.config.basicConfig.shellTextColor">
      </FormItem>
      <FormItem label="shellCursorColor">
        <input type="color" placeholder="please input shell cursor color"
          v-model="configStore.config.basicConfig.cursorColor">
      </FormItem>
      <FormItem label="shellCursorBlink">
        <tiny-radio v-model="configStore.config.basicConfig.isBlink" label="true">是</tiny-radio>
        <tiny-radio v-model="configStore.config.basicConfig.isBlink" label="false">否</tiny-radio>
      </FormItem>
    </div>
    <Hr content="themeSetting" percent="10%" />
    <div class="form">
      <FormItem label="backgroundColor">
        <input type="color" placeholder="please input background color"
          v-model="configStore.config.theme.backgroundColor">
      </FormItem>
      <FormItem label="menuBgColor">
        <input type="color" placeholder="please input menu bg color" v-model="configStore.config.theme.menuBgColor">
      </FormItem>
      <FormItem label="textColor">
        <input type="color" placeholder="please input text color" v-model="configStore.config.theme.textColor">
      </FormItem>
      <FormItem label="borderColor">
        <input type="color" placeholder="please input border color" v-model="configStore.config.theme.borderColor">
      </FormItem>
      <FormItem label="toolbarBgColor">
        <input type="color" placeholder="please input toolbar bg color"
          v-model="configStore.config.theme.toolbarBgColor">
      </FormItem>
      <FormItem label="directoryBgColor">
        <input type="color" placeholder="please input directory bg color"
          v-model="configStore.config.theme.directoryBgColor">
      </FormItem>
      <FormItem label="homeDisplayBoxBgColor">
        <input type="color" placeholder="please input home display box bg color"
          v-model="configStore.config.theme.showBoxBgColor">
      </FormItem>
    </div>
    <Hr content="resetAndSave" percent="10%" />
    <div class="btn">
      <tiny-button type="primary" @click="resetTheme">reset</tiny-button>
      <tiny-button type="primary" @click="saveTheme">save</tiny-button>
    </div>
    <div class="btn" style="color: var(--base-text-color);line-height: 30px;">---Take effect after saving---</div>
  </div>
</template>

<script lang="ts" setup>
import FormItem from './FormItem.vue'
import { useConfigStore } from '@renderer/stores/useConfigStore'
import Hr from './Hr.vue'
import { toRaw } from 'vue';
// 配置store
const configStore = useConfigStore()

// 重置主题
const resetTheme = () => {
  configStore.config.basicConfig = {
    username: 'user',
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