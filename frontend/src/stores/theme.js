import { defineStore } from 'pinia'
import { getThemes, hasTheme } from '../themes/registry'

const THEME_KEY='curly.theme'
const MODE_KEY='curly.color-mode'

export const useThemeStore=defineStore('theme',{
  state:()=>({ themeId:'default', mode:'system' }),
  getters:{ themes:()=>getThemes() },
  actions:{
    init(){
      const t=localStorage.getItem(THEME_KEY)
      const m=localStorage.getItem(MODE_KEY)
      if(t&&hasTheme(t)) this.themeId=t
      if(['system','light','dark'].includes(m)) this.mode=m
      this.apply()
      matchMedia('(prefers-color-scheme: dark)').addEventListener?.('change',()=>this.mode==='system'&&this.apply())
    },
    setTheme(id){ if(hasTheme(id)){ this.themeId=id; localStorage.setItem(THEME_KEY,id); this.apply() } },
    setMode(mode){ if(['system','light','dark'].includes(mode)){ this.mode=mode; localStorage.setItem(MODE_KEY,mode); this.apply() } },
    cycleMode(){ const arr=['system','light','dark']; this.setMode(arr[(arr.indexOf(this.mode)+1)%arr.length]) },
    apply(){
      const resolved=this.mode==='system'?(matchMedia('(prefers-color-scheme: dark)').matches?'dark':'light'):this.mode
      document.documentElement.dataset.theme=this.themeId
      document.documentElement.dataset.colorMode=resolved
      document.documentElement.style.colorScheme=resolved
    }
  }
})
