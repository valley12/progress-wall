import { defineStore } from 'pinia'
import { ref } from 'vue'

export interface User {
  id: string
  name: string
  email: string
  avatar?: string
  role: 'admin' | 'user'
}

export const useUserStore = defineStore('user', () => {
  const currentUser = ref<User | null>(null)
  const isLoggedIn = ref(false)

  // 模拟用户数据
  const mockUser: User = {
    id: '1',
    name: '张三',
    email: 'zhangsan@example.com',
    avatar: '',
    role: 'user'
  }

  const login = (email: string, password: string) => {
    // 模拟登录逻辑
    if (email && password) {
      currentUser.value = mockUser
      isLoggedIn.value = true
      return true
    }
    return false
  }

  // 自动登录用于演示
  const autoLogin = () => {
    currentUser.value = mockUser
    isLoggedIn.value = true
  }

  const logout = () => {
    currentUser.value = null
    isLoggedIn.value = false
  }

  const updateProfile = (updates: Partial<User>) => {
    if (currentUser.value) {
      Object.assign(currentUser.value, updates)
    }
  }

  return {
    currentUser,
    isLoggedIn,
    login,
    logout,
    updateProfile,
    autoLogin
  }
})
