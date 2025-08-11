import { app, shell, BrowserWindow, ipcMain, globalShortcut } from 'electron'
import { join, resolve } from 'path'
import { electronApp, optimizer, is } from '@electron-toolkit/utils'
import fs from 'fs'
import icon from '../../resources/icon.png?asset'
import { SessionGroup } from './types/session'
import { installExtension, VUEJS_DEVTOOLS } from 'electron-devtools-installer';
import { Config } from './types/configType'
import { spawn } from 'child_process'

let win: BrowserWindow | null = null
function createWindow(): void {
  // Create the browser window.
  win = new BrowserWindow({
    width: 1000,
    height: 658,
    show: false,
    frame: false,
    transparent: true,
    title: 'YumiShell',
    backgroundColor: '#00000000',
    ...(process.platform === 'linux' ? { icon } : {}),
    webPreferences: {
      preload: join(__dirname, '../preload/index.js'),
      sandbox: false,
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
//这四个负责窗口的关闭、最小化、最大化、窗口化
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
const saveConfigFile = (config: Config) => {
  fs.writeFileSync(join(app.getPath('userData'), 'config.json'), JSON.stringify(config))
}
// 加载配置文件
const loadConfigFile = () => {
  try {
    const config = fs.readFileSync(join(app.getPath('userData'), 'config.json'), 'utf-8')
    return JSON.parse(config) as Config
  } catch (error) {
    console.error('读取文件失败:', error)
    return null
  }
}

//保存用户session的回调
const saveSessions = (sessions: SessionGroup[]) => {
  //获取用户数据目录
  const userDataPath = app.getPath('userData')
  console.log(userDataPath);

  // 确保目录存在
  if (!fs.existsSync(userDataPath)) {
    fs.mkdirSync(userDataPath)
  }
  // 写入文件
  fs.promises.writeFile(join(userDataPath, 'sessions.json'), JSON.stringify(sessions)).then(() => {
    console.log("写入成功");
  }).catch((err) => {
    console.log("写入失败", err);
  })

}
//读取用户session的回调
const readSessions = () => {
  const userDataPath = app.getPath('userData');
  try {
    // 确保目录存在
    if (!fs.existsSync(userDataPath)) {
      fs.mkdirSync(userDataPath, { recursive: true });
    }
    const data = fs.readFileSync(join(app.getPath('userData'), 'sessions.json'), 'utf-8')
    return JSON.parse(data) as SessionGroup[]
  } catch (error) {
    console.error('读取文件失败:', error)
    return []
  }
}
// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.on('ready', () => {
  // 注册全局快捷键
  globalShortcut.register('CommandOrControl+Shift+I', () => {
    const focusedWindow = BrowserWindow.getFocusedWindow()
    if (focusedWindow) {
      focusedWindow.webContents.openDevTools()
    }
  })
  // 启动后端
  startBackend()
  installExtension(VUEJS_DEVTOOLS)
    .then((name) => console.log(`Added Extension:  ${name}`))
    .catch((err) => console.log('An error occurred: ', err));
  // Set app user model id for windows
  electronApp.setAppUserModelId('com.electron')

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
  //监听渲染进程发送的保存session事件
  ipcMain.on('save-sessions', (_, sessions) => {
    saveSessions(sessions)
  })
  //监听渲染进程发送的读取session事件
  ipcMain.handle('read-sessions', (_) => {
    return readSessions()
  })
  //监听渲染进程发送的保存配置事件
  ipcMain.on('save-config', (_, config) => {
    saveConfigFile(config)
  })
  // 监听渲染进程发送的加载配置事件
  ipcMain.handle('load-config', (_) => {
    return loadConfigFile()
  })

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

//启动后端的方法
function startBackend() {
  const exe = process.platform === 'win32' ? 'yumi-shell-conn.exe' : 'yumi-shell-conn';
  const backendPath = resolve(process.resourcesPath, exe);


  // 让后端跟随 Electron 主进程生命周期
  const backend = spawn(backendPath, [], { stdio: 'inherit', windowsHide: true });

  backend.on('error', (err) => {
    console.error('Backend spawn error:', err);
  });

  // 主进程退出时一起杀掉后端
  app.on('before-quit', () => {
    backend.kill();
  });
}