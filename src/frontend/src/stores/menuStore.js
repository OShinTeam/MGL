import { defineStore } from 'pinia'
 
export const useMenuStore = defineStore('menu', {
  state: () => ({
    activeMenu: localStorage.getItem('activeMenu') || '1'
  }),
  actions: {
    setActiveMenu(index) {
      this.activeMenu = index
      localStorage.setItem('activeMenu', index)
    }
  }
})