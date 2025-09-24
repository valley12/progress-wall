// 基础 API 服务类
export interface ApiResponse<T> {
  msg: string
  data: T
}

export class BaseApiService {
  protected async request<T>(
    url: string, 
    options: RequestInit = {}
  ): Promise<ApiResponse<T>> {
    const defaultOptions: RequestInit = {
      headers: {
        'Content-Type': 'application/json',
        ...options.headers,
      },
    }

    try {
      const response = await fetch(url, { ...defaultOptions, ...options })
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }

      const result = await response.json()
      return {
        msg: result.msg || '',
        data: result.data,
      }
    } catch (error) {
      console.error('API request failed:', error)
      return {
        msg: error instanceof Error ? error.message : 'Unknown error',
        data: null as T,
      }
    }
  }
}
