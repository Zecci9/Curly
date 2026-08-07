import { defineStore } from 'pinia'
import { login as loginApi, getMe } from '../api/auth'

const TOKEN_KEY = 'curly.token'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: '',
    user: null,
    loading: false,
  }),

  getters: {
    isLoggedIn: (state) => Boolean(state.token),
  },

  actions: {
    restore() {
      this.token = localStorage.getItem(TOKEN_KEY) || ''
    },

    async login(username, password) {
      this.loading = true
      try {
        const data = await loginApi({ username, password })
        const token = data?.data?.token ?? data?.token

        if (!token) throw new Error('后端没有返回 token')

        this.token = token
        localStorage.setItem(TOKEN_KEY, token)

        try {
          const me = await getMe()
          this.user = me?.data ?? me
        } catch {
          this.user = { username }
        }

        return true
      } finally {
        this.loading = false
      }
    },

    logout() {
      this.token = ''
      this.user = null
      localStorage.removeItem(TOKEN_KEY)
    },
  },
})
