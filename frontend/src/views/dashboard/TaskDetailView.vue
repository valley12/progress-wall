<template>
  <div class="min-h-screen bg-background">
    <!-- 加载状态 -->
    <div v-if="loading" class="flex items-center justify-center min-h-screen">
      <div class="text-center">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-primary mx-auto mb-4"></div>
        <p class="text-muted-foreground">加载任务详情中...</p>
      </div>
    </div>

    <!-- 错误状态 -->
    <div v-else-if="error" class="flex items-center justify-center min-h-screen">
      <div class="text-center">
        <p class="text-destructive mb-4">{{ error }}</p>
        <Button @click="loadTaskDetail" variant="outline">
          重新加载
        </Button>
      </div>
    </div>

    <!-- 任务详情内容 -->
    <div v-else-if="task" class="container mx-auto px-4 py-8 max-w-4xl">
      <!-- 头部导航 -->
      <div class="flex items-center gap-4 mb-8">
        <Button @click="goBack" variant="ghost" size="sm" class="p-2">
          ← 返回
        </Button>
        <div class="h-6 w-px bg-border"></div>
        <nav class="text-sm text-muted-foreground">
          <span>项目看板</span>
          <span class="mx-2">/</span>
          <span>任务详情</span>
        </nav>
      </div>

      <!-- 任务详情卡片 -->
      <Card class="p-8">
        <!-- 任务标题和状态 -->
        <div class="flex justify-between items-start mb-6">
          <div class="flex-1">
            <h1 class="text-3xl font-bold text-foreground mb-2">{{ task.title }}</h1>
            <div class="flex items-center gap-4">
              <span
                :class="[
                  'px-3 py-1 rounded-full text-sm font-medium',
                  getStatusClass(task.status)
                ]"
              >
                {{ getStatusText(task.status) }}
              </span>
              <span
                :class="[
                  'px-3 py-1 rounded-full text-sm font-medium',
                  getPriorityClass(task.priority)
                ]"
              >
                {{ getPriorityText(task.priority) }}
              </span>
            </div>
          </div>
        </div>

        <!-- 任务描述 -->
        <div class="mb-8">
          <h2 class="text-xl font-semibold mb-4">任务描述</h2>
          <div class="bg-muted/50 rounded-lg p-4">
            <p v-if="task.description" class="text-foreground leading-relaxed">
              {{ task.description }}
            </p>
            <p v-else class="text-muted-foreground italic">
              暂无描述
            </p>
          </div>
        </div>

        <!-- 任务信息 -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
          <div class="space-y-4">
            <h3 class="text-lg font-semibold">基本信息</h3>
            <div class="space-y-3">
              <div class="flex justify-between">
                <span class="text-muted-foreground">任务ID:</span>
                <span class="font-mono text-sm">{{ task.id }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-muted-foreground">状态:</span>
                <span
                  :class="[
                    'px-2 py-1 rounded text-sm font-medium',
                    getStatusClass(task.status)
                  ]"
                >
                  {{ getStatusText(task.status) }}
                </span>
              </div>
              <div class="flex justify-between">
                <span class="text-muted-foreground">优先级:</span>
                <span
                  :class="[
                    'px-2 py-1 rounded text-sm font-medium',
                    getPriorityClass(task.priority)
                  ]"
                >
                  {{ getPriorityText(task.priority) }}
                </span>
              </div>
            </div>
          </div>

          <div class="space-y-4">
            <h3 class="text-lg font-semibold">时间信息</h3>
            <div class="space-y-3">
              <div class="flex justify-between">
                <span class="text-muted-foreground">创建时间:</span>
                <span class="text-sm">{{ formatDateTime(task.createdAt) }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-muted-foreground">最后更新:</span>
                <span class="text-sm">{{ formatDateTime(task.updatedAt) }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- 操作按钮 -->
        <div class="flex gap-3 pt-6 border-t">
          <Button @click="goBack" variant="outline">
            返回看板
          </Button>
        </div>
      </Card>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useKanbanStore } from '@/stores/kanban'
import type { Task } from '@/stores/kanban'
import Button from '@/components/ui/Button.vue'
import Card from '@/components/ui/Card.vue'

// 路由和状态
const route = useRoute()
const router = useRouter()
const kanbanStore = useKanbanStore()

// 响应式数据
const loading = ref(false)
const error = ref<string | null>(null)
const task = ref<Task | null>(null)

// 获取任务ID
const taskId = route.params.taskId as string

// 加载任务详情
const loadTaskDetail = async () => {
  if (!taskId) {
    error.value = '任务ID不存在'
    return
  }

  loading.value = true
  error.value = null

  try {
    const response = await kanbanStore.fetchTaskDetail(taskId)
    if (response.success) {
      task.value = response.data
    } else {
      error.value = response.message || '加载任务详情失败'
    }
  } catch (err) {
    error.value = '网络错误，请稍后重试'
    console.error('Failed to load task detail:', err)
  } finally {
    loading.value = false
  }
}

// 页面操作
const goBack = () => {
  router.push('/kanban')
}


// 样式和文本转换函数
const getStatusClass = (status: Task['status']) => {
  const statusClasses = {
    'todo': 'bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200',
    'in-progress': 'bg-orange-100 text-orange-800 dark:bg-orange-900 dark:text-orange-200',
    'done': 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200'
  }
  return statusClasses[status] || 'bg-gray-100 text-gray-800 dark:bg-gray-900 dark:text-gray-200'
}

const getStatusText = (status: Task['status']) => {
  const statusMap = {
    'todo': '待办',
    'in-progress': '进行中',
    'done': '已完成'
  }
  return statusMap[status] || status
}

const getPriorityClass = (priority: Task['priority']) => {
  const priorityClasses = {
    'high': 'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-200',
    'medium': 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-200',
    'low': 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200'
  }
  return priorityClasses[priority] || 'bg-gray-100 text-gray-800 dark:bg-gray-900 dark:text-gray-200'
}

const getPriorityText = (priority: Task['priority']) => {
  const priorityMap = {
    high: '高优先级',
    medium: '中优先级',
    low: '低优先级'
  }
  return priorityMap[priority] || priority
}

// 格式化日期时间
const formatDateTime = (date: Date) => {
  return new Intl.DateTimeFormat('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  }).format(new Date(date))
}

// 生命周期
onMounted(() => {
  loadTaskDetail()
})
</script>
