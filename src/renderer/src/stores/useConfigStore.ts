import { defineStore } from "pinia";
import { onMounted, ref } from "vue";
import { Config } from '@renderer/types/config'

export const useConfigStore = defineStore('configStore', () => {
  const config = ref<Config>({
    basicConfig: {
      username: '',
      shellOpacity: 0.95,
      shellBgColor: "#0e0e0e",
      shellTextColor: "white",
      cursorColor: "white",
      isBlink: true
    },
    theme: {
      backgroundColor: "#efefef",
      menuBgColor: "#b3b3b3",
      textColor: "#000000",
      borderColor: "#000000",
      toolbarBgColor: "#383838",
      directoryBgColor: "#f5e0e0",
      showBoxBgColor: "#f2f2f2",
    }
  })
  const changeStyle = () => {
    document.documentElement.style.setProperty('--theme-background-color', config.value?.theme.backgroundColor || '#efefef')
    document.documentElement.style.setProperty('--theme-menu-bg-color', config.value?.theme.menuBgColor || '#efefef')
    document.documentElement.style.setProperty('--theme-text-color', config.value?.theme.textColor || '#000')
    document.documentElement.style.setProperty('--theme-border-color', config.value?.theme.borderColor || '#000')
    document.documentElement.style.setProperty('--theme-toolbar-bg-color', config.value?.theme.toolbarBgColor || '#efefef')
    document.documentElement.style.setProperty('--theme-directory-bg-color', config.value?.theme.directoryBgColor || '#efefef')
    document.documentElement.style.setProperty('--theme-show-box-bg-color', config.value?.theme.showBoxBgColor || '#efefef')
    document.documentElement.style.setProperty('--base-shell-opacity', config.value?.basicConfig.shellOpacity.toString() || "0.00")
  }
  onMounted(() => {
    window.api.loadConfig().then((res) => {
      if (res) {
        config.value = res
        changeStyle()
      }
    }).catch((err) => {
      console.log(err)
    })
  })
  return {
    config,
    changeStyle,
  }

})
