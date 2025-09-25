// 基础 API 服务类
export interface ApiResponse<T> {
  msg: string
  data: T
}

// 错误类型定义
export interface ApiError {
  message: string
  status?: number
  code?: string
}

// 请求超时配置
const REQUEST_TIMEOUT = 10000 // 10秒超时

export class BaseApiService {
  /**
   * 创建带超时的 AbortController
   */
  private createTimeoutController(timeout: number = REQUEST_TIMEOUT): AbortController {
    const controller = new AbortController()
    setTimeout(() => controller.abort(), timeout)
    return controller
  }

  /**
   * 统一错误处理
   */
  private handleError(error: unknown): ApiError {
    if (error instanceof Error) {
      if (error.name === 'AbortError') {
        return {
          message: '请求超时，请检查网络连接',
          code: 'TIMEOUT'
        }
      }
      return {
        message: error.message,
        code: 'NETWORK_ERROR'
      }
    }
    return {
      message: '未知错误',
      code: 'UNKNOWN_ERROR'
    }
  }

  /**
   * 统一请求方法
   */
  protected async request<T>(
    url: string, 
    options: RequestInit = {}
  ): Promise<ApiResponse<T>> {
    const controller = this.createTimeoutController()
    
    const defaultOptions: RequestInit = {
      headers: {
        'Content-Type': 'application/json',
        ...options.headers,
      },
      signal: controller.signal,
    }

    try {
      const response = await fetch(url, { ...defaultOptions, ...options })
      
      // 处理HTTP状态码错误
      if (!response.ok) {
        const errorMessage = this.getHttpErrorMessage(response.status)
        throw new Error(errorMessage)
      }

      const result = await response.json()
      
      // 验证响应格式
      if (typeof result !== 'object' || result === null) {
        throw new Error('无效的响应格式')
      }

      return {
        msg: result.msg || '',
        data: result.data,
      }
    } catch (error) {
      const apiError = this.handleError(error)
      console.error('API request failed:', apiError)
      
      return {
        msg: apiError.message,
        data: null as T,
      }
    }
  }

  /**
   * 根据HTTP状态码返回错误信息
   */
  private getHttpErrorMessage(status: number): string {
    const errorMessages: Record<number, string> = {
      400: '请求参数错误',
      401: '未授权，请重新登录',
      403: '禁止访问',
      404: '请求的资源不存在',
      408: '请求超时',
      429: '请求过于频繁，请稍后再试',
      500: '服务器内部错误',
      502: '网关错误',
      503: '服务不可用',
      504: '网关超时'
    }
    
    return errorMessages[status] || `HTTP错误: ${status}`
  }
}
