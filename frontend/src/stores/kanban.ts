import { defineStore } from 'pinia'
import { ref } from 'vue'

export interface Task {
  id: string
  title: string
  description?: string
  status: 'todo' | 'in-progress' | 'done'
  priority: 'low' | 'medium' | 'high'
  createdAt: Date
  updatedAt: Date
}

export interface Column {
  id: string
  title: string
  status: Task['status']
  tasks: Task[]
}

export const useKanbanStore = defineStore('kanban', () => {
  const columns = ref<Column[]>([
    {
      id: 'todo',
      title: '待办',
      status: 'todo',
      tasks: [
        {
          id: '1',
          title: '设计用户界面',
          description: '创建看板界面的设计稿',
          status: 'todo',
          priority: 'high',
          createdAt: new Date(),
          updatedAt: new Date()
        },
        {
          id: '2',
          title: '设置项目环境',
          description: '配置开发环境和依赖',
          status: 'todo',
          priority: 'medium',
          createdAt: new Date(),
          updatedAt: new Date()
        }
      ]
    },
    {
      id: 'in-progress',
      title: '进行中',
      status: 'in-progress',
      tasks: [
        {
          id: '3',
          title: '实现拖拽功能',
          description: '添加任务拖拽排序功能',
          status: 'in-progress',
          priority: 'high',
          createdAt: new Date(),
          updatedAt: new Date()
        }
      ]
    },
    {
      id: 'done',
      title: '已完成',
      status: 'done',
      tasks: [
        {
          id: '4',
          title: '项目初始化',
          description: '创建Vue项目并配置基础依赖',
          status: 'done',
          priority: 'low',
          createdAt: new Date(),
          updatedAt: new Date()
        }
      ]
    }
  ])

  const addTask = (task: Omit<Task, 'id' | 'createdAt' | 'updatedAt'>) => {
    const newTask: Task = {
      ...task,
      id: Date.now().toString(),
      createdAt: new Date(),
      updatedAt: new Date()
    }
    
    const column = columns.value.find(col => col.status === task.status)
    if (column) {
      column.tasks.push(newTask)
    }
  }

  const updateTask = (taskId: string, updates: Partial<Task>) => {
    for (const column of columns.value) {
      const task = column.tasks.find(t => t.id === taskId)
      if (task) {
        Object.assign(task, updates, { updatedAt: new Date() })
        break
      }
    }
  }

  const moveTask = (taskId: string, newStatus: Task['status']) => {
    let taskToMove: Task | null = null
    let sourceColumn: Column | null = null

    // 找到要移动的任务
    for (const column of columns.value) {
      const taskIndex = column.tasks.findIndex(t => t.id === taskId)
      if (taskIndex !== -1) {
        taskToMove = column.tasks[taskIndex] as Task
        sourceColumn = column
        column.tasks.splice(taskIndex, 1)
        break
      }
    }

    if (taskToMove && sourceColumn) {
      // 更新任务状态
      taskToMove.status = newStatus
      taskToMove.updatedAt = new Date()

      // 添加到目标列
      const targetColumn = columns.value.find(col => col.status === newStatus)
      if (targetColumn) {
        targetColumn.tasks.push(taskToMove)
      }
    }
  }

  const deleteTask = (taskId: string) => {
    for (const column of columns.value) {
      const taskIndex = column.tasks.findIndex(t => t.id === taskId)
      if (taskIndex !== -1) {
        column.tasks.splice(taskIndex, 1)
        break
      }
    }
  }

  return {
    columns,
    addTask,
    updateTask,
    moveTask,
    deleteTask
  }
})
