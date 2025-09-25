import { BaseApiService, type ApiResponse } from './base-api'
import { getEndpointUrl, getEndpointMethod } from '@/config/api'
import type { Board } from '@/stores/board'

// 看板 API 响应接口
export interface KanbanListResponse extends Array<Board> {}

// 看板 API 服务
export class BoardApiService extends BaseApiService {
  // 获取看板列表
  async getKanbanList(): Promise<ApiResponse<KanbanListResponse>> {
    return this.request<KanbanListResponse>(getEndpointUrl('KANBAN_LIST'), {
      method: getEndpointMethod('KANBAN_LIST')
    })
  }

  // 获取单个看板详情
  async getKanbanDetail(boardId: string): Promise<ApiResponse<Board>> {
    return this.request<Board>(`${getEndpointUrl('KANBAN_DETAIL')}/${boardId}`, {
      method: getEndpointMethod('KANBAN_DETAIL')
    })
  }

  // 创建看板
  async createKanban(board: Omit<Board, 'id' | 'createdAt' | 'updatedAt'>): Promise<ApiResponse<Board>> {
    return this.request<Board>(getEndpointUrl('KANBAN_CREATE'), {
      method: getEndpointMethod('KANBAN_CREATE'),
      body: JSON.stringify(board),
    })
  }

  // 更新看板
  async updateKanban(boardId: string, updates: Partial<Board>): Promise<ApiResponse<Board>> {
    return this.request<Board>(`${getEndpointUrl('KANBAN_UPDATE')}/${boardId}`, {
      method: getEndpointMethod('KANBAN_UPDATE'),
      body: JSON.stringify(updates),
    })
  }

  // 删除看板
  async deleteKanban(boardId: string): Promise<ApiResponse<void>> {
    return this.request<void>(`${getEndpointUrl('KANBAN_DELETE')}/${boardId}`, {
      method: getEndpointMethod('KANBAN_DELETE'),
    })
  }
}

// 导出单例实例
export const boardApiService = new BoardApiService()
