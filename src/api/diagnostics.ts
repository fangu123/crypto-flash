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
}

const storageKey = 'cf_api_diagnostics'
const maxEntries = 200

function safeMessage(value: unknown) {
  if (typeof value !== 'string' || !value) return ''
  return value
    .replace(/\b[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}\b/gi, '[email]')
    .replace(/\bBearer\s+[^\s,;]+/gi, 'Bearer [redacted]')
    .replace(/\beyJ[A-Za-z0-9_-]+\.[A-Za-z0-9_-]+\.[A-Za-z0-9_-]+\b/g, '[token]')
    .replace(/\b(password|passwd|token|authorization|验证码|code)\s*[:=：]\s*[^\s,;，；]+/gi, '$1=[redacted]')
    .slice(0, 160)
}

function safeEndpoint(value: string) {
  const pathname = String(value || '').split('?')[0]
  return pathname
    .split('/')
    .map(segment => /^\d+$/.test(segment) || /^[0-9a-f]{8}-[0-9a-f-]{27,}$/i.test(segment) ? ':id' : segment)
    .join('/')
    .replace(/\b[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}\b/gi, '[email]')
    .slice(0, 180)
}

export function getDiagnosticLogs(): DiagnosticLog[] {
  try {
    const value = uni.getStorageSync(storageKey)
    return Array.isArray(value) ? value : []
  } catch {
    return []
  }
}

function saveLog(entry: Omit<DiagnosticLog, 'id' | 'time'>) {
  const item: DiagnosticLog = { ...entry, message: safeMessage(entry.message), endpoint: safeEndpoint(entry.endpoint), id: `${Date.now()}-${Math.random().toString(36).slice(2, 8)}`, time: Date.now() }
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
