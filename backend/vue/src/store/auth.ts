import { defineStore } from 'pinia'
import axios from 'axios'
import { parseJWT } from '../utils/jwt'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('watch_admin_token') || ''
  }),
  getters: {
    userInfo: (state) => (state.token ? parseJWT(state.token) : null),
    hasPermission: (state) => (permission: string) => {
      if (!state.token) return false
      const info = parseJWT(state.token)
      if (!info?.permissions) return false
      if (info.permissions === 'all') return true
      return info.permissions.split(',').map((p) => p.trim()).includes(permission)
    },
    hasAnyPermission: (state) => (permissions: string[]) => {
      if (!state.token) return false
      const info = parseJWT(state.token)
      if (!info?.permissions) return false
      if (info.permissions === 'all') return true
      const userPerms = info.permissions.split(',').map((p) => p.trim())
      return permissions.some((p) => userPerms.includes(p))
    }
  },
  actions: {
    async login(username: string, password: string) {
      const { data } = await axios.post('/api/login', { username, password })
      this.token = data.token
      localStorage.setItem('watch_admin_token', data.token)
    },
    logout() {
      this.token = ''
      localStorage.removeItem('watch_admin_token')
    }
  }
})
