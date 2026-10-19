import { API_EQUIPMENT, API_MODE, apiUrl } from '../config/api'
import { locale, t } from '../locales'
import { mockRequest } from './mock'
import { recordApiLog } from './diagnostics'

type RequestMethod = 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH'

export class ApiError extends Error {
  statusCode?: number
  code?: number

  constructor(message: string, statusCode?: number, code?: number) {
    super(message)
    this.name = 'ApiError'
    this.statusCode = statusCode
    this.code = code
  }
}

export type RequestOptions = {
  method?: RequestMethod
  data?: any
  query?: Record<string, string | number | boolean | undefined>
  auth?: boolean
  skipRefresh?: boolean
}

function queryString(query?: RequestOptions['query']) {
  if (!query) return ''
  const entries = Object.entries(query).filter(([, value]) => value !== undefined && value !== '')
  if (!entries.length) return ''
  return `?${entries.map(([key, value]) => `${encodeURIComponent(key)}=${encodeURIComponent(String(value))}`).join('&')}`
}

function messageFrom(data: any, fallback: string) {
  return typeof data?.msg === 'string' && data.msg ? data.msg : fallback
}

function errorDetails(error: any) {
  if (error instanceof Error) return `${error.name}: ${error.message}${error.stack ? `\n${error.stack}` : ''}`
  try { return JSON.stringify(error, null, 2) ?? String(error) } catch { return String(error) }
}

export async function request<T = any>(path: string, options: RequestOptions = {}): Promise<T> {
  const method = options.method || 'GET'
  const startedAt = Date.now()
  let requestUrl = ''
  let header: Record<string, string> = {}
  const requestBody = options.data
  let token = ''
  try {
    requestUrl = `${API_MODE === 'mock' ? '[MOCK] ' : ''}${apiUrl(path)}${queryString(options.query)}`
    token = uni.getStorageSync('cf_token')
    header = {
      'Accept-Language': locale.value === 'zh' ? 'cn' : 'en',
      'Content-Type': 'application/json'
    }
    if (API_EQUIPMENT) header['Equipment-Number'] = API_EQUIPMENT
    if (options.auth !== false && token) header.Authorization = token.startsWith('Bearer ') ? token : `Bearer ${token}`
  } catch (error: any) {
    const detail = errorDetails(error)
    recordApiLog({ method, endpoint: path, outcome: 'failure', failureSource: 'frontend', requestUrl, requestHeaders: header, requestBody, durationMs: Date.now() - startedAt, message: error?.message || String(error), errorDetail: String(detail) })
    throw new ApiError(error?.message || t('请求构造失败'))
  }

  if (API_MODE === 'mock') {
    try {
      const payload = await mockRequest<T>(path, options)
      recordApiLog({ method, endpoint: path, outcome: 'success', requestUrl, requestHeaders: header, requestBody, responseHeaders: {}, responseBody: payload, durationMs: Date.now() - startedAt, message: 'MOCK 响应（未请求后端）' })
      return payload
    } catch (error: any) {
      const detail = errorDetails(error)
      recordApiLog({ method, endpoint: path, outcome: 'failure', failureSource: 'frontend', requestUrl, requestHeaders: header, requestBody, responseHeaders: {}, responseBody: error, durationMs: Date.now() - startedAt, message: error?.message || String(error), errorDetail: String(detail) })
      throw error
    }
  }

  let response: UniApp.RequestSuccessCallbackResult
  try {
    response = await new Promise<UniApp.RequestSuccessCallbackResult>((resolve, reject) => {
      try {
        uni.request({ url: requestUrl, method, data: requestBody, header, success: resolve, fail: reject })
      } catch (error: any) {
        const setupError: any = new Error(error?.message || String(error))
        setupError.cause = error
        setupError.__frontendRequestSetup = true
        reject(setupError)
      }
    })
  } catch (error: any) {
    const frontendFailure = Boolean(error?.__frontendRequestSetup)
    const detail = errorDetails(error)
    recordApiLog({ method, endpoint: path, outcome: frontendFailure ? 'failure' : 'network_error', failureSource: frontendFailure ? 'frontend' : 'network', requestUrl, requestHeaders: header, requestBody, durationMs: Date.now() - startedAt, message: error?.errMsg || error?.message || t('网络请求失败'), errorDetail: String(detail) })
    throw new ApiError(error?.errMsg || error?.message || t('网络请求失败'))
  }

  const payload: any = response.data
  const businessCode = payload && typeof payload === 'object' && 'code' in payload ? payload.code : undefined
  const requestId = payload?.requestId || response.header?.['X-Request-Id'] || response.header?.['x-request-id']
  const outcome = response.statusCode < 200 || response.statusCode >= 300
    ? 'http_error'
    : businessCode !== undefined && businessCode !== 200 ? 'api_error' : 'success'
  const failureSource = response.statusCode < 200 || response.statusCode >= 300 ? 'http' : businessCode !== undefined && businessCode !== 200 ? 'business' : undefined
  recordApiLog({ method, endpoint: path, outcome, failureSource, requestUrl, requestHeaders: header, requestBody, responseHeaders: response.header as Record<string, any>, responseBody: payload, durationMs: Date.now() - startedAt, httpStatus: response.statusCode, code: businessCode, message: messageFrom(payload, ''), requestId: requestId ? String(requestId) : undefined })
  const authFailure = response.statusCode === 401 || (payload && typeof payload === 'object' && payload.code === 401)
  if (authFailure && !options.skipRefresh && token) {
    try {
      const refreshed: any = await request('/app/profile/refreshToken', { method: 'POST', auth: true, data: { isRemember: Boolean(uni.getStorageSync('cf_remember')), }, skipRefresh: true })
      const nextToken = refreshed?.token || refreshed?.currentAuthority
      if (nextToken) {
        uni.setStorageSync('cf_token', nextToken)
        return request<T>(path, { ...options, skipRefresh: true })
      }
    } catch {}
  }
  if (authFailure) {
    uni.removeStorageSync('cf_token')
    const current = getCurrentPages().at(-1) as any
    if (token && !current?.route?.includes('pages/index/index')) uni.reLaunch({ url: '/pages/index/index' })
  }
  if (response.statusCode < 200 || response.statusCode >= 300) {
    throw new ApiError(messageFrom(payload, `${t('请求失败')} (${response.statusCode})`), response.statusCode, payload?.code)
  }
  if (payload && typeof payload === 'object' && 'code' in payload && payload.code !== 200) {
    throw new ApiError(messageFrom(payload, t('接口返回失败')), response.statusCode, payload.code)
  }
  return payload && typeof payload === 'object' && 'data' in payload ? payload.data : payload
}
