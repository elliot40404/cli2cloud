import { defineStore } from 'pinia'

export const useIdStore = defineStore({
  id: 'main',
  state: () => ({
    id: ""
  }),
  actions: {
    setId(value) {
      this.id = value
    }
  },
  persist: true
})
