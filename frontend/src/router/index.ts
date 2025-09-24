import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Home',
    component: () => import('@/views/HomeView.vue')
  },
  {
    path: '/boards',
    name: 'BoardList',
    component: () => import('@/views/dashboard/BoardListView.vue')
  },
  {
    path: '/kanban/:boardId?',
    name: 'Kanban',
    component: () => import('@/views/dashboard/KanbanView.vue')
  },
  {
    path: '/profile',
    name: 'Profile',
    component: () => import('@/views/user/ProfileView.vue')
  },
  {
    path: '/settings',
    name: 'Settings',
    component: () => import('@/views/user/SettingsView.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
