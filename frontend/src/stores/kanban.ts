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

// API 响应类型
export interface TaskDetailResponse {
  success: boolean
  data: Task
  message?: string
}

export interface MoveTaskRequest {
  newColumnId: string
  newOrder: number
}

export interface MoveTaskResponse {
  success: boolean
  message?: string
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

  // 拖拽移动任务（乐观更新）
  const moveTaskWithDrag = async (taskId: string, newColumnId: string, newOrder: number) => {
    // 保存原始状态用于回滚
    const originalColumns = JSON.parse(JSON.stringify(columns.value))
    
    try {
      // 1. 乐观更新：立即更新本地状态
      let taskToMove: Task | null = null
      let sourceColumnIndex = -1
      let sourceTaskIndex = -1

      // 找到要移动的任务
      for (let i = 0; i < columns.value.length; i++) {
        const taskIndex = columns.value[i].tasks.findIndex(t => t.id === taskId)
        if (taskIndex !== -1) {
          taskToMove = columns.value[i].tasks[taskIndex]
          sourceColumnIndex = i
          sourceTaskIndex = taskIndex
          break
        }
      }

      if (!taskToMove) {
        throw new Error('任务不存在')
      }

      // 从源列移除任务
      columns.value[sourceColumnIndex].tasks.splice(sourceTaskIndex, 1)

      // 找到目标列
      const targetColumn = columns.value.find(col => col.id === newColumnId)
      if (!targetColumn) {
        throw new Error('目标列不存在')
      }

      // 更新任务状态
      taskToMove.status = targetColumn.status
      taskToMove.updatedAt = new Date()

      // 插入到目标列的指定位置
      if (newOrder >= targetColumn.tasks.length) {
        targetColumn.tasks.push(taskToMove)
      } else {
        targetColumn.tasks.splice(newOrder, 0, taskToMove)
      }

      // 2. 调用API
      const response = await moveTaskAPI(taskId, { newColumnId, newOrder })
      
      if (!response.success) {
        throw new Error(response.message || '移动任务失败')
      }

    } catch (error) {
      // 3. 如果API调用失败，回滚到原始状态
      columns.value = originalColumns
      console.error('Move task failed:', error)
      throw error
    }
  }

  // 模拟API调用移动任务
  const moveTaskAPI = async (taskId: string, request: MoveTaskRequest): Promise<MoveTaskResponse> => {
    // 模拟网络延迟
    await new Promise(resolve => setTimeout(resolve, 200))
    
    // 模拟API调用
    console.log(`PATCH /api/tasks/${taskId}/move`, request)
    
    // 模拟成功响应
    return {
      success: true
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


  // 模拟API调用获取任务详情
  const fetchTaskDetail = async (taskId: string): Promise<TaskDetailResponse> => {
    // 模拟网络延迟
    await new Promise(resolve => setTimeout(resolve, 300))
    
    // 在所有列中查找任务
    let foundTask: Task | null = null
    for (const column of columns.value) {
      const task = column.tasks.find(t => t.id === taskId)
      if (task) {
        foundTask = task
        break
      }
    }

    if (foundTask) {
      return {
        success: true,
        data: foundTask
      }
    } else {
      return {
        success: false,
        data: {} as Task,
        message: '任务不存在'
      }
    }
  }

  return {
    columns,
    addTask,
    updateTask,
    moveTask,
    moveTaskWithDrag,
    deleteTask,
    fetchTaskDetail
  }
})
