import { defineStore } from 'pinia'
import { ref } from 'vue'
import { boardApiService } from '@/services/board-api'

export interface Board {
  id: string
  name: string
  description?: string
  color: string
  createdAt: Date
  updatedAt: Date
}

export const useBoardStore = defineStore('board', () => {
  const boards = ref<Board[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  // 从 API 获取看板列表
  const fetchBoards = async () => {
    loading.value = true
    error.value = null
    
    try {
      const response = await boardApiService.getKanbanList()
      if (response.data && Array.isArray(response.data)) {
        // 转换日期字符串为 Date 对象，并确保 id 是字符串
        boards.value = response.data.map(board => ({
          ...board,
          id: board.id.toString(), // 确保 id 是字符串
          createdAt: new Date(board.createdAt),
          updatedAt: new Date(board.updatedAt)
        }))
      } else {
        error.value = response.msg || '获取看板列表失败'
      }
    } catch (err) {
      error.value = err instanceof Error ? err.message : '获取看板列表失败'
    } finally {
      loading.value = false
    }
  }

  // 创建看板
  const addBoard = async (board: Omit<Board, 'id' | 'createdAt' | 'updatedAt'>) => {
    loading.value = true
    error.value = null
    
    try {
      const response = await boardApiService.createKanban(board)
      if (response.data) {
        const newBoard = {
          ...response.data,
          id: response.data.id.toString(), // 确保 id 是字符串
          createdAt: new Date(response.data.createdAt),
          updatedAt: new Date(response.data.updatedAt)
        }
        boards.value.push(newBoard)
        return newBoard
      } else {
        error.value = response.msg || '创建看板失败'
        return null
      }
    } catch (err) {
      error.value = err instanceof Error ? err.message : '创建看板失败'
      return null
    } finally {
      loading.value = false
    }
  }

  // 更新看板
  const updateBoard = async (boardId: string, updates: Partial<Board>) => {
    loading.value = true
    error.value = null
    
    try {
      const response = await boardApiService.updateKanban(boardId, updates)
      if (response.data) {
        const updatedBoard = {
          ...response.data,
          id: response.data.id.toString(), // 确保 id 是字符串
          createdAt: new Date(response.data.createdAt),
          updatedAt: new Date(response.data.updatedAt)
        }
        const index = boards.value.findIndex(b => b.id === boardId)
        if (index !== -1) {
          boards.value[index] = updatedBoard
        }
        return updatedBoard
      } else {
        error.value = response.msg || '更新看板失败'
        return null
      }
    } catch (err) {
      error.value = err instanceof Error ? err.message : '更新看板失败'
      return null
    } finally {
      loading.value = false
    }
  }

  // 删除看板
  const deleteBoard = async (boardId: string) => {
    loading.value = true
    error.value = null
    
    try {
      const response = await boardApiService.deleteKanban(boardId)
      if (response.data !== null) {
        const index = boards.value.findIndex(b => b.id === boardId)
        if (index !== -1) {
          boards.value.splice(index, 1)
        }
        return true
      } else {
        error.value = response.msg || '删除看板失败'
        return false
      }
    } catch (err) {
      error.value = err instanceof Error ? err.message : '删除看板失败'
      return false
    } finally {
      loading.value = false
    }
  }

  // 根据 ID 获取看板
  const getBoardById = (boardId: string) => {
    return boards.value.find(b => b.id === boardId)
  }

  // 清除错误状态
  const clearError = () => {
    error.value = null
  }

  return {
    boards,
    loading,
    error,
    fetchBoards,
    addBoard,
    updateBoard,
    deleteBoard,
    getBoardById,
    clearError
  }
})
