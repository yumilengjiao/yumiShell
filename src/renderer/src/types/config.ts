export interface BasicConfig {
  username: string
  shellOpacity: number
  shellBgColor: string
  shellTextColor: string
  cursorColor: string
  isBlink: boolean
}

export interface Theme {
  backgroundColor: string
  menuBgColor: string
  textColor: string
  borderColor: string
  toolbarBgColor: string
  directoryBgColor: string
  showBoxBgColor: string
}

export interface Config {
  basicConfig: BasicConfig
  theme: Theme
}
