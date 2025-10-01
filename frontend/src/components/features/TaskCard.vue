<template>
  <Card
    class="p-4 cursor-pointer hover:shadow-md transition-all hover:scale-[1.02]"
    :data-task-id="task.id"
    @click="$emit('select', task)"
  >
    <div class="space-y-2">
      <div class="flex justify-between items-start">
        <h4 class="font-medium text-sm">{{ task.title }}</h4>
        <Button
          @click.stop="$emit('delete', task.id)"
          variant="ghost"
          size="sm"
          class="h-6 w-6 p-0 text-muted-foreground hover:text-destructive"
        >
          ×
        </Button>
      </div>
      
      <p v-if="task.description" class="text-xs text-muted-foreground line-clamp-2">
        {{ task.description }}
      </p>
      
      <div class="flex justify-between items-center">
        <span
          :class="[
            'px-2 py-1 rounded-full text-xs font-medium',
            {
              'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-200': task.priority === 'high',
              'bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-200': task.priority === 'medium',
              'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200': task.priority === 'low'
            }
          ]"
        >
          {{ getPriorityText(task.priority) }}
        </span>
        
        <span class="text-xs text-muted-foreground">
          {{ formatDate(task.updatedAt) }}
        </span>
      </div>
    </div>
  </Card>
</template>

<script setup lang="ts">
import type { Task } from '@/stores/kanban'
import Card from '@/components/ui/Card.vue'
import Button from '@/components/ui/Button.vue'

interface Props {
  task: Task
}

defineProps<Props>()

defineEmits<{
  select: [task: Task]
  delete: [taskId: string]
}>()

const getPriorityText = (priority: string) => {
  const priorityMap = {
    high: '高',
    medium: '中',
    low: '低'
  }
  return priorityMap[priority as keyof typeof priorityMap] || priority
}

const formatDate = (date: Date) => {
  return new Intl.DateTimeFormat('zh-CN', {
    month: 'short',
    day: 'numeric'
  }).format(new Date(date))
}
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
