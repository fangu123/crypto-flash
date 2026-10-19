export type DiagnosticLog = {
  id: string
  time: number
  kind: 'api' | 'app'
  method: string
  endpoint: string
  outcome: 'success' | 'http_error' | 'api_error' | 'network_error' | 'failure'
  durationMs?: number
  httpStatus?: number
  code?: string | number
  message?: string
  requestId?: string
  failureSource?: 'frontend' | 'network' | 'http' | 'business'
  requestUrl?: string
  requestHeaders?: Record<string, any>
  requestBody?: unknown
  responseHeaders?: Record<string, any>
  responseBody?: unknown
  errorDetail?: string
}

const storageKey = 'cf_api_diagnostics'
const maxEntries = 200

export function getDiagnosticLogs(): DiagnosticLog[] {
  try {
    const value = uni.getStorageSync(storageKey)
    return Array.isArray(value) ? value : []
  } catch {
    return []
  }
}

function saveLog(entry: Omit<DiagnosticLog, 'id' | 'time'>) {
  const item: DiagnosticLog = { ...entry, id: `${Date.now()}-${Math.random().toString(36).slice(2, 8)}`, time: Date.now() }
  try {
    const logs = [item, ...getDiagnosticLogs()].slice(0, maxEntries)
    uni.setStorageSync(storageKey, logs)
  } catch {
    // Console output remains available when local storage is unavailable.
  }
  try { console.info('[diagnostics]', item) } catch {}
}

export function recordApiLog(entry: Omit<DiagnosticLog, 'id' | 'time' | 'kind' | 'endpoint'> & { endpoint: string }) {
  saveLog({ ...entry, kind: 'api' })
}

export function recordAppLog(endpoint: string, outcome: 'success' | 'failure', message = '') {
  saveLog({ kind: 'app', method: 'FLOW', endpoint, outcome, message })
}

export function clearDiagnosticLogs() {
  try { uni.removeStorageSync(storageKey) } catch {}
}
