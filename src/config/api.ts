/** Production builds must provide VITE_API_BASE_URL (use an HTTPS origin). */
export const API_BASE_URL = (import.meta.env?.VITE_API_BASE_URL || '').replace(/\/+$/, '')
export const API_MODE = import.meta.env?.VITE_API_MODE === 'real' ? 'real' : 'mock'
export const API_PREFIX = '/api/v1'
export const API_LANGUAGE = import.meta.env?.VITE_API_LANGUAGE || 'cn'
export const API_EQUIPMENT = import.meta.env?.VITE_API_EQUIPMENT || ''

export function apiUrl(path: string) {
  const normalized = path.startsWith('/') ? path : `/${path}`
  return `${API_BASE_URL}${API_PREFIX}${normalized}`
}
