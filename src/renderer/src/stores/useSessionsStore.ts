import { ref, onMounted } from 'vue'
import { defineStore } from 'pinia'
import { Session, SessionGroup } from '@renderer/types/session'
import { useOtherConnStore } from './useOtherConnStore'
import { MessageType } from '@renderer/types/others';

interface DisplayGroup {
  id: string;
  label: string;
  children: DisplaySession[];
}
interface DisplaySession {
  id: string;
  label: string;
}

//定义一个非响应式sessionGroup组用来存放所有有关session的信息
let sessionGroups: SessionGroup[] = []
export const useSessionsStore = defineStore('sessions', () => {
  //用于处理其他东西的仓库
  const otherConnStore = useOtherConnStore()
  // 定义一个状态变量，用于存储用于展示的会话组
  const sessions = ref<DisplayGroup[]>([
    {
      id: '',
      label: 'default group',
      children: [
      ],
    }
  ])
  //添加一个用于显示的会话组
  const addGroup = (label: string) => {
    //生成一个随机的uuid
    let randomId = crypto.randomUUID()
    //用于展示的group处理
    const newGroup: DisplayGroup = {
      id: randomId,
      label,
      children: []
    }
    sessions.value.push(newGroup)
    console.log("添加组成功并触发更新");
  }
  //用于添加一个用于显示的会话连接
  const addSession = (groupId: string, session: Session) => {
    console.log(groupId);
    sessions.value = sessions.value.map(group => {
      if (group.id === groupId) {
        // 创建新的 children 数组
        const newChildren = [...group.children, {
          id: session.id, // 使用传入的 session.id 确保唯一性
          label: session.labelName
        }]
        return { ...group, children: newChildren }
      }
      return group
    })
    console.log("添加session成功并触发更新");
  }
  //用于持久化存储sessionGroups
  const saveSessions = (sessionGroupId: string, groupName: string, sessionData: any) => {
    console.log("sessionData要储存的privateKey是", sessionData.privateKey);
    console.log("传入的sessionGroupId", sessionGroupId, "groupName", groupName, "sessionData", sessionData);
    console.log("当前sessionGroups为", sessionGroups);
    const group = sessionGroups.find(group => group.id == sessionGroupId)
    //存在group则将当前session添加进，没有则新添加一个group
    if (group) {
      console.log("存在group,id为", sessionGroupId);
      sessionData.id = crypto.randomUUID()
      group.sessionList.push({
        id: sessionData.id,
        labelName: sessionData.labelName,
        host: sessionData.host,
        port: sessionData.port,
        username: sessionData.username,
        authType: sessionData.authType,
        password: sessionData.password,
        privateKey: sessionData.privateKey,
        passphrase: sessionData.passphrase,
        timeout: sessionData.timeout,
        encoding: sessionData.encoding,
        terminalType: sessionData.terminalType,
      })
    } else {
      console.log("不存在group,id为", sessionGroupId);
      sessionGroups.push({
        id: sessionGroupId,
        label: groupName,
        sessionList: [],
      })
    }
    console.log("即将存储的sessionGroups", sessionGroups);
    window.api.saveSessions(sessionGroups)
    console.log("存储session成功");
    //更新用于显示的sessionGroups
    sessions.value = sessionGroups.map(group => ({
      id: group.id,
      label: group.label,
      children: group.sessionList.map(session => ({
        id: session.id,
        label: session.labelName,
      }))
    }))
    //告知后端重新读取session.json
    otherConnStore.ws!.send(JSON.stringify({
      reqType: MessageType.READ_SESSIONS,
      sessionContent: sessionGroups
    }))
    //通知后端
    console.log("后端通知成功");
  }
  //用于删除本地session
  const deleteSession = (sessionId: string) => {
    console.log("删除的sessionId", sessionId);
    console.log(sessionGroups);
    let deleteParent = sessionGroups.some(group => sessionId === group.id)
    console.log("删除的是否是父组", deleteParent);
    if (deleteParent) {
      sessionGroups = sessionGroups.filter(group => group.id !== sessionId)
      console.log("删除的是父组", sessionGroups);
    } else {
      sessionGroups.forEach(group => {
        group.sessionList = group.sessionList.filter(session => session.id !== sessionId)
      })
    }
    console.log("删除后的sessionGroups", sessionGroups);
    //再次存储sessionGroups
    window.api.saveSessions(sessionGroups)
    console.log("删除成功");

    //更新用于显示的sessionGroups
    sessions.value = sessionGroups.map(group => ({
      id: group.id,
      label: group.label,
      children: group.sessionList.map(session => ({
        id: session.id,
        label: session.labelName,
      }))
    }))
    //告知后端重新读取session.json
    otherConnStore.ws!.send(JSON.stringify({
      reqType: MessageType.READ_SESSIONS,
    }))

  }
  //使用该仓库的组件挂载时
  onMounted(async () => {
    sessionGroups = await window.api.readSessions()
    //将读取到的sessionGroups转换为用于展示的格式
    sessions.value = sessionGroups.map(group => ({
      id: group.id,
      label: group.label,
      children: group.sessionList.map(session => ({
        id: session.id,
        label: session.labelName,
      }))
    }))
    console.log(sessionGroups);
  })

  return {
    sessions,
    addGroup,
    addSession,
    saveSessions,
    deleteSession
  }
})
