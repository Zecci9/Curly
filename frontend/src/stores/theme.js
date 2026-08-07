import { defineStore } from 'pinia'
import { getThemes, hasTheme } from '../themes/registry'

const THEME_KEY = 'curly.theme'
const MODE_KEY = 'curly.color-mode'

export const useThemeStore = defineStore('theme', {
  state: () => ({
    themeId: 'default',
    mode: 'system',
    initialized: false,
  }),

  getters: {
    themes: () => getThemes(),
  },

  actions: {
    init() {
      const savedTheme = localStorage.getItem(THEME_KEY)
      const savedMode = localStorage.getItem(MODE_KEY)

      if (savedTheme && hasTheme(savedTheme)) this.themeId = savedTheme
      if (['light', 'dark', 'system'].includes(savedMode)) this.mode = savedMode

      this.apply()

      const media = window.matchMedia('(prefers-color-scheme: dark)')
      media.addEventListener?.('change', () => {
        if (this.mode === 'system') this.apply()
      })

      this.initialized = true
    },

    setTheme(id) {
      if (!hasTheme(id)) return
      this.themeId = id
      localStorage.setItem(THEME_KEY, id)
      this.apply()
    },

    setMode(mode) {
      if (!['light', 'dark', 'system'].includes(mode)) return
      this.mode = mode
      localStorage.setItem(MODE_KEY, mode)
      this.apply()
    },

    cycleMode() {
      const order = ['system', 'light', 'dark']
      const next = order[(order.indexOf(this.mode) + 1) % order.length]
      this.setMode(next)
    },

    apply() {
      const root = document.documentElement
      root.dataset.theme = this.themeId

      const resolved =
        this.mode === 'system'
          ? (window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light')
          : this.mode

      root.dataset.colorMode = resolved
      root.style.colorScheme = resolved
    },
  },
})
