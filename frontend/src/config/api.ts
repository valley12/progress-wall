// API 配置
export const API_CONFIG = {
  // API 基础域名
  BASE_URL: 'https://2f8c8fdd-a2bc-4443-bebc-97e2c601e632.mock.pstmn.io',
  
  // API 版本
  VERSION: 'v1',
  
  // 端点配置 - 同时定义路径和 HTTP 方法
  ENDPOINTS: {
    // 看板相关
    KANBAN_LIST: {
      path: '/boards',
      method: 'GET'
    },
    KANBAN_DETAIL: {
      path: '/dashboard/kanban',
      method: 'GET'
    },
    KANBAN_CREATE: {
      path: '/dashboard/kanban',
      method: 'POST'
    },
    KANBAN_UPDATE: {
      path: '/dashboard/kanban',
      method: 'PUT'
    },
    KANBAN_DELETE: {
      path: '/dashboard/kanban',
      method: 'DELETE'
    },
  }
} as const

// 构建完整的 API URL
export const buildApiUrl = (endpoint: string): string => {
  return `${API_CONFIG.BASE_URL}/${API_CONFIG.VERSION}${endpoint}`
}

// 获取端点的完整 URL
export const getEndpointUrl = (endpointKey: keyof typeof API_CONFIG.ENDPOINTS): string => {
  return buildApiUrl(API_CONFIG.ENDPOINTS[endpointKey].path)
}

// 获取端点的 HTTP 方法
export const getEndpointMethod = (endpointKey: keyof typeof API_CONFIG.ENDPOINTS): string => {
  return API_CONFIG.ENDPOINTS[endpointKey].method
}
