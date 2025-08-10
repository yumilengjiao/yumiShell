import { ElectronAPI } from '@electron-toolkit/preload'

declare global {
  interface Window {
    electron: ElectronAPI
    api: {
      closeWindow: () => void
      minimizeWindow: () => void
      maximizeWindow: () => void
      windowedWindow: () => void
      saveSessions: (sessions: SessionGroup[]) => void
      readSessions: () => Promise<SessionGroup[]>
      saveConfig: (config: Config) => void
      loadConfig: () => Promise<Config>
    }
  }
}
