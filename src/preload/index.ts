import { contextBridge, ipcRenderer } from 'electron'
import { electronAPI } from '@electron-toolkit/preload'
import { SessionGroup } from '../main/types/session'
import { Config } from '../main/types/configType'

// Custom APIs for renderer
const api = {
  // 暴露给渲染进程的方法
  closeWindow: () => {
    // 推荐使用ipcRenderer发送事件给主进程，而不是直接调用window.electron.app.quit()
    ipcRenderer.send('window-close')
  },
  minimizeWindow: () => {
    ipcRenderer.send('window-minimize')
  },
  maximizeWindow: () => {
    ipcRenderer.send('window-maximize')
  },
  windowedWindow: () => {
    ipcRenderer.send('window-windowed')
  },
  //保存sessions的方法
  saveSessions: (sessions: SessionGroup[]) => {
    ipcRenderer.send('save-sessions', sessions)
  },
  //读取sessions的方法
  readSessions: () => ipcRenderer.invoke('read-sessions'),
  //保存配置的方法
  saveConfig: (config: Config) => ipcRenderer.send('save-config', config),
  // 加载配置的方法
  loadConfig: () => ipcRenderer.invoke('load-config'),
}

// Use `contextBridge` APIs to expose Electron APIs to
// renderer only if context isolation is enabled, otherwise
// just add to the DOM global.
// 如果启用了上下文隔离则通过contextBridge.exposeInMainWorld()暴露API
if (process.contextIsolated) {
  try {
    contextBridge.exposeInMainWorld('electron', electronAPI)
    contextBridge.exposeInMainWorld('api', api)
  } catch (error) {
    console.error(error)
  }
} else {
  // @ts-ignore (define in dts)
  window.electron = electronAPI
  // @ts-ignore (define in dts)
  window.api = api
}
