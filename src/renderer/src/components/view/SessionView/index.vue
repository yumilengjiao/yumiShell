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
    <tiny-tree-menu :allow-drop="() => false" :default-expand-all="true" :show-filter="false"
      :data="sessionsStore.sessions" ellipsis draggable @node-click="handleNodeClick" @node-drag-start="handleDragStart"
      empty-text="null sessions"
      style="width: 100%;box-shadow: 0 0 10px var(--base-shadow-color);background-color: var(--base-background-color);">


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
        <tiny-form-item label="labelName" title="labelName">
          <tiny-input type="text" v-model="sessionData.labelName"></tiny-input>
        </tiny-form-item>
        <tiny-form-item label="host" title="host">
          <tiny-ip-address v-if="isIpv4" v-model="sessionData.host"></tiny-ip-address>
          <tiny-ip-address type="ipv6" v-if="!isIpv4" v-model="sessionData.host"></tiny-ip-address>
          <br>
          <tiny-radio v-model="isIpv4" :label="true">IPv4</tiny-radio>
          <tiny-radio v-model="isIpv4" :label="false">IPv6</tiny-radio>
        </tiny-form-item>
        <tiny-form-item label="port" title="port">
          <tiny-numeric min="0" max="65535" v-model="sessionData.port" placeholder="请输入非空数值"
            style="width: 30%;"></tiny-numeric>
        </tiny-form-item>
        <tiny-form-item label="username" title="username">
          <tiny-input type="text" v-model="sessionData.username"></tiny-input>
        </tiny-form-item>
        <tiny-form-item label="authType" title="authType">
          <tiny-radio-group v-model="sessionData.authType">
            <tiny-radio :label="'password'">password</tiny-radio>
            <tiny-radio :label="'privateKey'">privateKey</tiny-radio>
          </tiny-radio-group>
        </tiny-form-item>
        <tiny-form-item label="password" title="password">
          <tiny-input type="password" v-model="sessionData.password"></tiny-input>
        </tiny-form-item>
        <tiny-form-item label="timeOut" title="timeOut">
          <tiny-input type="text" v-model="sessionData.timeout"></tiny-input>
        </tiny-form-item>
        <tiny-form-item label="ptyType" title="ptyType">
          <tiny-base-select v-model="sessionData.terminalType">
            <tiny-option v-for="item in termType" :key="item.value" :label="item.label" :value="item.value">
            </tiny-option>
          </tiny-base-select>
        </tiny-form-item>
      </tiny-form>
      <template #footer>
        <tiny-button @click="sessionBoxVisibility = false" round>取消</tiny-button>
        <tiny-button type="primary" @click="addSession" round>确定</tiny-button>
      </template>
    </tiny-dialog-box>
    <div class="rubbish-btn" title="拖拽session到此处删除" @dragenter="handleDragEnter" @dragover="handleDragOver"
      @dragleave="handleDragLeave" @drop="handleDrop" ref="rubbish">
      <svg t="1754706097154" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
        p-id="4603" width="50" height="50">
        <path
          d="M853.333333 213.333333a42.666667 42.666667 0 0 1 3.2 85.226667L853.333333 298.666667H170.666667a42.666667 42.666667 0 0 1-3.2-85.226667L170.666667 213.333333h682.666666z m-42.666666 149.333334a42.666667 42.666667 0 0 1 42.56 39.466666L853.333333 405.333333v384a170.666667 170.666667 0 0 1-165.333333 170.581334L682.666667 960H341.333333a170.666667 170.666667 0 0 1-170.581333-165.333333L170.666667 789.333333V405.333333a42.666667 42.666667 0 0 1 85.226666-3.2L256 405.333333v384a85.333333 85.333333 0 0 0 81.066667 85.226667L341.333333 874.666667h341.333334a85.333333 85.333333 0 0 0 85.226666-81.066667L768 789.333333V405.333333a42.666667 42.666667 0 0 1 42.666667-42.666666z m-405.333334 64a42.666667 42.666667 0 0 1 42.56 39.466666L448 469.333333v256a42.666667 42.666667 0 0 1-85.226667 3.2L362.666667 725.333333V469.333333a42.666667 42.666667 0 0 1 42.666666-42.666666z m213.333334 0a42.666667 42.666667 0 0 1 42.56 39.466666L661.333333 469.333333v256a42.666667 42.666667 0 0 1-85.226666 3.2L576 725.333333V469.333333a42.666667 42.666667 0 0 1 42.666667-42.666666zM597.333333 64a42.666667 42.666667 0 0 1 3.2 85.226667L597.333333 149.333333h-170.666666a42.666667 42.666667 0 0 1-3.2-85.226666L426.666667 64h170.666666z"
          fill="#2c2c2c" p-id="4604"></path>
      </svg>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { useSessionsStore } from '@renderer/stores/useSessionsStore'
import type { Session } from '@renderer/types/session'
import { useShellViewTabStore } from '@renderer/stores/useShellViewTabStore'
// 存储上次点击的时间戳
let lastClickTime = 0;
// 双击阈值（毫秒）
const DOUBLE_CLICK_THRESHOLD = 300;
const shellViewTabStore = useShellViewTabStore()
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
const rubbish = ref<HTMLDivElement>()
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
  terminalType: 'xterm-256color',
})
//定义伪终端类型的数组
const termType = ref([
  {
    value: 'xterm',
    label: 'xterm',
  },
  {
    value: 'xterm-256color',
    label: 'xterm-256color',
  },
  {
    value: 'screen',
    label: 'screen',
  },
  {
    value: 'screen-256color',
    label: 'screen-256color',
  },
  {
    value: 'vt100',
    label: 'vt100',
  },
  {
    value: 'vt220',
    label: 'vt220',
  },
  {
    value: 'ansi',
    label: 'ansi',
  },
  {
    value: 'linux',
    label: 'linux',
  },
  {
    value: 'linux-16color',
    label: 'linux-16color',
  },
  {
    value: 'rxvt',
    label: 'rxvt',
  },
  {
    value: 'rxvt-256color',
    label: 'rxvt-256color',
  },
  {
    value: 'konsole',
    label: 'konsole',
  },
  {
    value: 'iterm2',
    label: 'iterm2',
  },
  {
    value: 'wezterm',
    label: 'wezterm',
  },
  {
    value: 'xterm-direct',
    label: 'xterm-direct',
  },
  {
    value: 'tmux-direct',
    label: 'tmux-direct',
  },
  {
    value: 'cygwin',
    label: 'cygwin',
  },
  {
    value: 'putty',
    label: 'putty',
  },
  {
    value: 'Eterm',
    label: 'Eterm',
  },
  {
    value: 'aterm',
    label: 'aterm',
  },
  {
    value: 'dtterm',
    label: 'dtterm',
  },
  {
    value: 'xterm-new',
    label: 'xterm-new',
  },
  {
    value: 'gnome',
    label: 'gnome',
  },
  {
    value: 'st-256color',
    label: 'st-256color',
  },
  {
    value: 'alacritty',
    label: 'alacritty',
  }
])

//定义将要删除的session
const willDeleteSession = ref<any>()
// 定义单选框的值
let isIpv4 = ref(true);
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
  // 阻止默认行为，允许 Drop
  e.preventDefault()
}
//处理拖拽开始的事件
const handleDragStart = (e: any) => {
  // console.log(e);
  willDeleteSession.value = e.data
  console.log(willDeleteSession.value);
}
//处理拖拽松手垃圾桶的事件
const handleDrop = () => {
  console.log(willDeleteSession.value);
  deleteSessions(willDeleteSession.value)
  //清空willDeleteSession
  willDeleteSession.value = null
  //清空垃圾桶的样式
  rubbish.value?.classList.remove('dragover')
}
//处理删除session
const deleteSessions = (willDeleteSession) => {
  //遍历仓库的数据查看树id是否一致，一致就删除
  sessionsStore.deleteSession(willDeleteSession.id)
}
//处理拖拽进入垃圾桶的事件
const handleDragEnter = (e: DragEvent) => {
  console.log("拖拽进入");
  // 阻止默认行为，允许 Drop
  e.preventDefault()
  // 显示垃圾桶
  rubbish.value?.classList.add('dragover')
}
//处理拖拽离开垃圾桶的事件
const handleDragLeave = (e: DragEvent) => {
  console.log(e);
  // 阻止默认行为，允许 Drop
  e.preventDefault()
  // 显示垃圾桶
  rubbish.value?.classList.remove('dragover')
}
const handleNodeClick = (nodeData, node) => {
  //父节点点击无效
  if (!node.isLeaf) {
    return
  }
  // 判断是否双击
  const currentTime = Date.now();
  const isDoubleClick = (currentTime - lastClickTime) < DOUBLE_CLICK_THRESHOLD;
  lastClickTime = currentTime;
  // 仅在双击时执行操作
  if (isDoubleClick) {
    //生成一个随机数表示一个会话的唯一标识
    const uniqId = crypto.randomUUID()
    // 增加一个标签页
    shellViewTabStore.tabs.push({
      title: "session" + shellViewTabStore.index,
      name: uniqId,
      sessionId: nodeData.id,
    });
    shellViewTabStore.index++;
  }
}
</script>

<style scoped lang="scss">
@use '@renderer/styles/variables.scss' as variables;

.session-view {
  width: 100%;
  height: calc(100vh - variables.$menu-bar-height);
  background-color: var(--base-background-color);
  border-left: 1px solid var(--base-border-color);
  border-right: 1px solid var(--base-border-color);
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

  .rubbish-btn {
    position: absolute;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 50px;
    height: 50px;
    bottom: 10px;
    right: 10px;
    border-radius: 50%;

    .icon {
      path {
        fill: var(--base-text-color);
      }
    }

    &::after {
      content: '';
      position: absolute;
      width: 100%;
      height: 100%;
      background-color: rgba(223, 32, 86, 0.5);
      z-index: 1;
      //默认隐藏
      opacity: 0;
      transition: all 0.3s ease-in-out;
    }

    &.dragover::after {
      opacity: 1 !important;
      /* 拖拽时显示 */
    }

  }

}
</style>