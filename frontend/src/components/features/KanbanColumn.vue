<template>
  <div class="bg-muted/50 rounded-lg p-4 min-h-[500px]" :data-column-id="column.id">
    <div class="flex justify-between items-center mb-4">
      <h3 class="text-lg font-semibold">
        {{ column.title }}
      </h3>
      <span class="text-sm text-muted-foreground bg-muted px-2 py-1 rounded-full">
        {{ column.tasks.length }}
      </span>
    </div>
    
    <VueDraggable
      v-model="column.tasks"
      :group="{ name: 'tasks', pull: true, put: true }"
      :animation="200"
      ghost-class="ghost-task"
      chosen-class="chosen-task"
      drag-class="drag-task"
      class="space-y-3 min-h-[200px]"
      @end="onDragEnd"
    >
      <TaskCard
        v-for="task in column.tasks"
        :key="task.id"
        :task="task"
        @select="$emit('select-task', task)"
        @delete="$emit('delete-task', task.id)"
      />
      
      <div
        v-if="column.tasks.length === 0"
        class="text-center text-muted-foreground py-8 pointer-events-none"
      >
        暂无任务
      </div>
    </VueDraggable>
  </div>
</template>

<script setup lang="ts">
import { VueDraggable } from 'vue-draggable-plus'
import type { Column } from '@/stores/kanban'
import { useKanbanStore } from '@/stores/kanban'
import TaskCard from '@/components/features/TaskCard.vue'

interface Props {
  column: Column
}

const props = defineProps<Props>()

const emit = defineEmits<{
  'select-task': [task: any]
  'delete-task': [taskId: string]
}>()

const kanbanStore = useKanbanStore()

// 拖拽结束事件处理
const onDragEnd = async (event: any) => {
  const { item, to, from, newIndex, oldIndex } = event
  
  // 获取被拖拽的任务ID
  const taskId = item.querySelector('[data-task-id]')?.getAttribute('data-task-id')
  if (!taskId) return

  // 获取目标列ID
  const targetColumnElement = to.closest('[data-column-id]')
  const targetColumnId = targetColumnElement?.getAttribute('data-column-id')
  if (!targetColumnId) return

  // 如果是同一列内的重新排序，或者跨列移动
  if (from !== to || newIndex !== oldIndex) {
    try {
      await kanbanStore.moveTaskWithDrag(taskId, targetColumnId, newIndex)
    } catch (error) {
      console.error('拖拽移动任务失败:', error)
      // 这里可以添加错误提示
    }
  }
}
</script>

<style scoped>
/* 拖拽样式 */
:deep(.ghost-task) {
  opacity: 0.5;
  background: #f0f0f0;
  border: 2px dashed #ccc;
}

:deep(.chosen-task) {
  transform: rotate(5deg);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.15);
}

:deep(.drag-task) {
  transform: rotate(5deg);
  opacity: 0.8;
}

/* 拖拽区域样式 */
.sortable-ghost {
  opacity: 0.5;
}

.sortable-chosen {
  opacity: 0.8;
}

.sortable-drag {
  opacity: 0.6;
}
</style>
