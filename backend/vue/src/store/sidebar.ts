import { defineStore } from 'pinia'

export const useSidebarStore = defineStore('sidebar', {
  state: () => ({ collapsed: false }),
  actions: {
    toggle() {
      this.collapsed = !this.collapsed
    }
  }
})
