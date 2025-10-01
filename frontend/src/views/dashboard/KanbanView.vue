<template>
  <div class="min-h-screen bg-background">
    <div class="container mx-auto px-4 py-8">
      <div class="flex justify-between items-center mb-8">
        <h1 class="text-3xl font-bold text-foreground">项目看板</h1>
        <Button @click="goHome" variant="outline">
          返回首页
        </Button>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <KanbanColumn
          v-for="column in kanbanStore.columns"
          :key="column.id"
          :column="column"
          @select-task="selectTask"
          @delete-task="deleteTask"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useKanbanStore } from '@/stores/kanban'
import Button from '@/components/ui/Button.vue'
import KanbanColumn from '@/components/features/KanbanColumn.vue'

const router = useRouter()
const kanbanStore = useKanbanStore()

const goHome = () => {
  router.push('/')
}

const selectTask = (task: any) => {
  // 点击任务卡片跳转到任务详情页
  router.push(`/tasks/${task.id}`)
}

const deleteTask = (taskId: string) => {
  kanbanStore.deleteTask(taskId)
}
</script>
