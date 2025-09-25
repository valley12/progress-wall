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

  /**
   * 转换API数据为本地格式
   * 避免重复的对象转换，提高性能
   */
  const transformBoardData = (apiBoard: any): Board => {
    return {
      ...apiBoard,
      id: String(apiBoard.id), // 确保 id 是字符串
      createdAt: new Date(apiBoard.createdAt),
      updatedAt: new Date(apiBoard.updatedAt)
    }
  }

  /**
   * 从 API 获取看板列表
   * 优化数据处理，减少不必要的对象转换
   */
  const fetchBoards = async () => {
    loading.value = true
    error.value = null
    
    try {
      const response = await boardApiService.getKanbanList()
      
      // 验证响应数据
      if (response.data && Array.isArray(response.data)) {
        // 批量转换数据，避免多次对象创建
        boards.value = response.data.map(transformBoardData)
      } else {
        error.value = response.msg || '获取看板列表失败'
        boards.value = [] // 确保空状态处理
      }
    } catch (err) {
      error.value = err instanceof Error ? err.message : '获取看板列表失败'
      boards.value = [] // 错误时清空列表
    } finally {
      loading.value = false
    }
  }

  /**
   * 创建看板
   * 使用统一的数据转换函数，提高代码复用性
   */
  const addBoard = async (board: Omit<Board, 'id' | 'createdAt' | 'updatedAt'>) => {
    loading.value = true
    error.value = null
    
    try {
      const response = await boardApiService.createKanban(board)
      
      if (response.data) {
        const newBoard = transformBoardData(response.data)
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

  /**
   * 更新看板
   * 使用统一的数据转换函数，提高代码复用性
   */
  const updateBoard = async (boardId: string, updates: Partial<Board>) => {
    loading.value = true
    error.value = null
    
    try {
      const response = await boardApiService.updateKanban(boardId, updates)
      
      if (response.data) {
        const updatedBoard = transformBoardData(response.data)
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
