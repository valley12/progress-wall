<template>
  <nav class="bg-white shadow-sm border-b dark:bg-gray-900">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between h-16">
        <!-- 左侧 Logo 和导航 -->
        <div class="flex items-center">
          <router-link to="/" class="flex items-center space-x-2">
            <div class="w-8 h-8 bg-primary rounded-lg flex items-center justify-center">
              <span class="text-primary-foreground font-bold text-sm">P</span>
            </div>
            <span class="text-xl font-bold text-foreground">Progress Wall</span>
          </router-link>
          
          <div class="hidden md:ml-5 md:flex md:space-x-5">
            <router-link
              to="/"
              class="text-gray-500 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white px-3 py-2 rounded-md text-sm font-medium transition-colors"
              :class="{ 'text-primary': $route.path === '/' }"
            >
              首页
            </router-link>
            <router-link
              to="/boards"
              class="text-gray-500 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white px-3 py-2 rounded-md text-sm font-medium transition-colors"
              :class="{ 'text-primary': $route.path === '/boards' }"
            >
              看板列表
            </router-link>
          </div>
        </div>

        <!-- 右侧用户菜单 -->
        <div class="flex items-center space-x-4">
          <template v-if="userStore.isLoggedIn">
            <DropdownMenu>
              <template #trigger>
                <button class="flex items-center space-x-2 text-sm rounded-full focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary">
                  <Avatar
                    :src="userStore.currentUser?.avatar"
                    :fallback="userStore.currentUser?.name?.charAt(0) || 'U'"
                    class="h-8 w-8"
                  />
                  <span class="hidden md:block text-gray-700 dark:text-gray-300">
                    {{ userStore.currentUser?.name }}
                  </span>
                  <svg class="h-4 w-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                  </svg>
                </button>
              </template>
              
              <DropdownMenuItem @click="goToProfile">
                个人资料
              </DropdownMenuItem>
              <DropdownMenuItem @click="goToSettings">
                设置
              </DropdownMenuItem>
              <DropdownMenuItem @click="handleLogout" class="text-red-600 dark:text-red-400">
                退出登录
              </DropdownMenuItem>
            </DropdownMenu>
          </template>
          
          <template v-else>
            <Button @click="showLoginModal = true" variant="outline" size="sm">
              登录
            </Button>
            <Button @click="showRegisterModal = true" size="sm">
              注册
            </Button>
          </template>
        </div>
      </div>
    </div>

    <!-- 移动端菜单 -->
    <div v-if="mobileMenuOpen" class="md:hidden">
      <div class="px-2 pt-2 pb-3 space-y-1 sm:px-3 bg-white dark:bg-gray-900 border-t">
        <router-link
          to="/"
          class="text-gray-500 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white block px-3 py-2 rounded-md text-base font-medium"
          @click="mobileMenuOpen = false"
        >
          首页
        </router-link>
        <router-link
          to="/kanban"
          class="text-gray-500 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white block px-3 py-2 rounded-md text-base font-medium"
          @click="mobileMenuOpen = false"
        >
          看板
        </router-link>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import Button from '@/components/ui/Button.vue'
import Avatar from '@/components/ui/Avatar.vue'
import DropdownMenu from '@/components/ui/DropdownMenu.vue'
import DropdownMenuItem from '@/components/ui/DropdownMenuItem.vue'

const router = useRouter()
const userStore = useUserStore()

const mobileMenuOpen = ref(false)
const showLoginModal = ref(false)
const showRegisterModal = ref(false)

const goToProfile = () => {
  router.push('/profile')
}

const goToSettings = () => {
  router.push('/settings')
}

const handleLogout = () => {
  userStore.logout()
  router.push('/')
}
</script>
