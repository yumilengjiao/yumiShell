<template>
  <div class="session-view">
    <tiny-button @click="boxVisibility = true" title="弹出 Dialog" style="width: 100%;">Add new group</tiny-button>
    <tiny-dialog-box :show-close="false" v-model:visible="boxVisibility" title="消息" width="30%">
      <span>Add new group</span>
      <tiny-input type="text" v-model="groupName" placeholder="Please input group name"></tiny-input>
      <template #footer>
        <tiny-button @click="boxVisibility = false" round>取 消</tiny-button>
        <tiny-button type="primary" @click="addGroup" round>确 定</tiny-button>
      </template>
    </tiny-dialog-box>
    <!-- 展示当前会话 -->
    <tiny-tree-menu :default-expand-all="true" :show-filter="false" :data="sessionsStore.sessions" ellipsis draggable
      @node-drag-start="handleDragStart" empty-text="null sessions" style="width: 100%;">
      <template #default="slotScope">
        <button class="addSession" v-if="slotScope.data.children"
          @click.stop="sessionAddControl(slotScope.data.id)">+</button>
        <a :target="slotScope.data.target" :class="slotScope.data.class">
          {{ slotScope.data.label }}
        </a>
      </template>
    </tiny-tree-menu>
    <!-- 添加session的表单 -->
    <tiny-dialog-box :show-close="false" v-model:visible="sessionBoxVisibility" title="Add new session" width="70%">
      <tiny-form label-width="100px" label-position="left">
        <tiny-form-item label="labelName">
          <tiny-input type="text" v-model="sessionData.labelName"></tiny-input>
        </tiny-form-item>
        <tiny-form-item label="host">
          <tiny-ip-address v-if="isIpv4" v-model="sessionData.host"></tiny-ip-address>
          <tiny-ip-address type="ipv6" v-if="!isIpv4" v-model="sessionData.host"></tiny-ip-address>
          <br>
          <tiny-radio v-model="isIpv4" :label="true">IPv4</tiny-radio>
          <tiny-radio v-model="isIpv4" :label="false">IPv6</tiny-radio>
        </tiny-form-item>
        <tiny-form-item label="port">
          <tiny-numeric min="0" max="65535" v-model="sessionData.port" placeholder="请输入非空数值"
            style="width: 30%;"></tiny-numeric>
        </tiny-form-item>
        <tiny-form-item label="username">
          <tiny-input type="text" v-model="sessionData.username"></tiny-input>
        </tiny-form-item>
        <tiny-form-item label="authType">
          <tiny-radio-group v-model="sessionData.authType">
            <tiny-radio :label="'password'">password</tiny-radio>
            <tiny-radio :label="'privateKey'">privateKey</tiny-radio>
          </tiny-radio-group>
        </tiny-form-item>
        <tiny-form-item label="password">
          <tiny-input type="password" v-model="sessionData.password"></tiny-input>
        </tiny-form-item>
        <tiny-form-item label="timeOut">
          <tiny-input type="text" v-model="sessionData.timeout"></tiny-input>
        </tiny-form-item>
      </tiny-form>
      <template #footer>
        <tiny-button @click="sessionBoxVisibility = false" round>取消</tiny-button>
        <tiny-button type="primary" @click="addSession" round>确定</tiny-button>
      </template>
    </tiny-dialog-box>
    <tiny-button class="deleteModel" ref="rubbishBtn" size="large" :icon="IconDeleteL" circle @click="deleteSessions"
      @drop="handleDrop" @dragover="handleDragOver"></tiny-button>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { useSessionsStore } from '@renderer/stores/useSessionsStore'
import type { Session } from '@renderer/types/session'
import { iconDeleteL } from '@opentiny/vue-icon'
const IconDeleteL = iconDeleteL()
const sessionsStore = useSessionsStore()
//这是添加自定义组的dialog-box  
const boxVisibility = ref(false)
//这是添加自定义会话的dialog-box
const sessionBoxVisibility = ref(false)
//当前会话的组id
const currentGroupId = ref('')
//组名
const groupName = ref('')
//垃圾桶实例
const rubbishBtn = ref()
//表单接收的对象
const sessionData = ref<Session>({
  id: '',
  labelName: '',
  host: '',
  port: 22,
  username: '',
  authType: 'password',
  password: '',
  timeout: 20,
  encoding: 'utf-8',
  terminalType: 'xterm',
})
// 定义单选框的值
let isIpv4 = ref(true);
//定义将要删除的session
const willDeleteSession = ref<any>()
//添加会话的控制函数
const sessionAddControl = (groupId: string) => {
  sessionBoxVisibility.value = true
  currentGroupId.value = groupId
  console.log("传入的groupId", groupId);

  console.log("当前点击的会话的组id", currentGroupId.value);
}

//添加会话
const addSession = () => {
  sessionsStore.addSession(currentGroupId.value, sessionData.value)
  sessionsStore.saveSessions(currentGroupId.value, groupName.value, sessionData.value)
  //手动重置表单数据
  sessionData.value = {
    id: '',
    labelName: '',
    host: '',
    port: 22,
    username: '',
    authType: 'password',
    password: '',
    timeout: 20,
    encoding: 'utf-8',
    terminalType: 'xterm',
  }
  console.log("添加session完成");
  //将显示框关闭
  sessionBoxVisibility.value = false
  //currentGroupId清除
  currentGroupId.value = ''
}
//添加组
const addGroup = () => {
  sessionsStore.addGroup(groupName.value)
  sessionsStore.saveSessions(crypto.randomUUID(), groupName.value, [])
  //将显示框关闭
  boxVisibility.value = false
  //组名清除
  groupName.value = ''
}
//处理拖拽到垃圾桶的事件
const handleDragOver = (e: DragEvent) => {
  console.log(e);
  e.preventDefault()
}
//处理拖拽开始的事件
const handleDragStart = (e: any) => {
  console.log(e);
  willDeleteSession.value = e.data
  console.log(willDeleteSession.value);
}
//处理拖拽松手垃圾桶的事件
const handleDrop = (e: DragEvent) => {
  console.log(willDeleteSession.value);

  console.log(sessionsStore.sessions);

}
//处理删除session
const deleteSessions = (willDeleteSession) => {
  //遍历仓库的数据查看树id是否一致，一致就删除

}
</script>

<style scoped lang="scss">
@use '@renderer/styles/variables.scss' as variables;

.session-view {
  width: 100%;
  height: calc(100vh - variables.$menu-bar-height);
  background-color: var(--base-background-color);
  border-left: 1px solid var(--base-border-color);
  position: relative;

  .addSession {
    width: 20px;
    height: 20px;
    background-color: var(--base-background-color);
    color: var(--base-text-color);
    font-size: 14px;
    font-weight: 500;
    cursor: pointer;
    display: inline-flex;
    align-items: center;
    justify-content: center;
  }

  .deleteModel {
    position: absolute;
    bottom: 10px;
    right: 10px;
  }

}
</style>