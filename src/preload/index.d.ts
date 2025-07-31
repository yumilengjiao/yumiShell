import { ElectronAPI } from '@electron-toolkit/preload'

declare global {
  interface Window {
    electron: ElectronAPI
    api: {
      closeWindow: () => void
      minimizeWindow: () => void
      maximizeWindow: () => void
      windowedWindow: () => void
    }
  }
}
