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
          <FileView />
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
import FileView from '@renderer/components/FileView/index.vue'
import ShellTabBar from '@renderer/components/ShellTabBar/index.vue'
import MenuBar from '@renderer/components/MenuBar/index.vue'
import ToolTabBar from '@renderer/components/ToolTabBar/index.vue'
</script>

<style scoped></style>
