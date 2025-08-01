<template>
  <div class="content">
    <tiny-layout :cols="24">
      <!-- 菜单行栏-->
      <tiny-row nospace gutter="0">
        <tiny-col :span="24">
          <MenuBar />
        </tiny-col>
      </tiny-row>
      <!-- 文件视图和操作区域行栏-->
      <tiny-row nospace gutter="0">
        <tiny-col :span="1">
          <ToolTabBar />
        </tiny-col>
        <tiny-col :span="5">
          <SessionView v-show="toolBarStore.nowToolBarModel === 'session'" />
          <SettingView v-show="toolBarStore.nowToolBarModel === 'setting'" />
          <FileView v-show="toolBarStore.nowToolBarModel === 'file'" />
        </tiny-col>
        <tiny-col :span="18">
          <ShellTabBar />
        </tiny-col>
      </tiny-row>
    </tiny-layout>
  </div>
</template>

<script setup>
import { TinyLayout, TinyRow, TinyCol } from '@opentiny/vue'
import FileView from '@renderer/components/view/FileView/index.vue'
import SettingView from '@renderer/components/view/settingView/index.vue'
import SessionView from '@renderer/components/view/SessionView/index.vue'
import ShellTabBar from '@renderer/components/ShellTabBar/index.vue'
import MenuBar from '@renderer/components/MenuBar/index.vue'
import ToolTabBar from '@renderer/components/ToolTabBar/index.vue'
//导入piniaToolBar仓库，用于不同tool栏点击的时候切换视图
import { useToolBarStore } from '@renderer/stores/useToolBarStore'

const toolBarStore = useToolBarStore()
</script>

<style scoped></style>
