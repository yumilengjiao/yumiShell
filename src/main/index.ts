import { app, shell, BrowserWindow, ipcMain } from 'electron'
import { join } from 'path'
import { electronApp, optimizer, is } from '@electron-toolkit/utils'
import icon from '../../resources/icon.png?asset'

let win: BrowserWindow | null = null
function createWindow(): void {
  // Create the browser window.
  win = new BrowserWindow({
    width: 1000,
    height: 670,
    show: false,
    frame: false,
    transparent: true,
    title: 'YumiShell',
    backgroundColor: '#00000000',
    ...(process.platform === 'linux' ? { icon } : {}),
    webPreferences: {
      preload: join(__dirname, '../preload/index.js'),
      sandbox: false
    }
  })

  win.on('ready-to-show', () => {
    win?.show()
  })

  win.webContents.setWindowOpenHandler((details) => {
    shell.openExternal(details.url)
    return { action: 'deny' }
  })

  // HMR for renderer base on electron-vite cli.
  // Load the remote URL for development or the local html file for production.
  if (is.dev && process.env['ELECTRON_RENDERER_URL']) {
    win.loadURL(process.env['ELECTRON_RENDERER_URL'])
  } else {
    win.loadFile(join(__dirname, '../renderer/index.html'))
  }
}
// 监听渲染进程发送的回调函数
const closeWindow = () => {
  win?.close()
}
const minimizeWindow = () => {
  win?.minimize()
}
const maximizeWindow = () => {
  win?.maximize()
}
const windowedWindow = () => {
  win?.setSize(1000, 670)
  //窗口在屏幕中央
  win?.setPosition(100, 50)
}
// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.on('ready', () => {
  // Set app user model id for windows
  electronApp.setAppUserModelId('com.electron')

  // Default open or close DevTools by F12 in development
  // and ignore CommandOrControl + R in production.
  // see https://github.com/alex8088/electron-toolkit/tree/master/packages/utils
  app.on('browser-window-created', (_, window) => {
    optimizer.watchWindowShortcuts(window)
  })

  //ipc
  // 监听渲染进程发送的关闭窗口事件
  ipcMain.on('window-close', closeWindow)
  // 监听渲染进程发送的最小化窗口事件
  ipcMain.on('window-minimize', minimizeWindow)
  // 监听渲染进程发送的最大化窗口事件
  ipcMain.on('window-maximize', maximizeWindow)
  //监听渲染进程发送的窗口化事件
  ipcMain.on('window-windowed', windowedWindow)

  createWindow()

  app.on('activate', function () {
    // On macOS it's common to re-create a window in the app when the
    // dock icon is clicked and there are no other windows open.
    if (BrowserWindow.getAllWindows().length === 0) createWindow()
  })
})

// Quit when all windows are closed, except on macOS. There, it's common
// for applications and their menu bar to stay active until the user quits
// explicitly with Cmd + Q.
app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') {
    app.quit()
  }
})

