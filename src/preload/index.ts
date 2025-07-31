import { contextBridge, ipcRenderer } from 'electron'
import { electronAPI } from '@electron-toolkit/preload'

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
  }
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
