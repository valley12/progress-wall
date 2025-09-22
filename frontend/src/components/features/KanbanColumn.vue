<template>
  <div class="bg-muted/50 rounded-lg p-4 min-h-[500px]">
    <div class="flex justify-between items-center mb-4">
      <h3 class="text-lg font-semibold">
        {{ column.title }}
      </h3>
      <span class="text-sm text-muted-foreground bg-muted px-2 py-1 rounded-full">
        {{ column.tasks.length }}
      </span>
    </div>
    
    <div class="space-y-3">
      <TaskCard
        v-for="task in column.tasks"
        :key="task.id"
        :task="task"
        @select="$emit('select-task', task)"
        @delete="$emit('delete-task', task.id)"
      />
      
      <div
        v-if="column.tasks.length === 0"
        class="text-center text-muted-foreground py-8"
      >
        暂无任务
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Column } from '@/stores/kanban'
import TaskCard from '@/components/features/TaskCard.vue'

interface Props {
  column: Column
}

defineProps<Props>()

defineEmits<{
  'select-task': [task: any]
  'delete-task': [taskId: string]
}>()
</script>
