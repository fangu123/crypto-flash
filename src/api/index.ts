import { request } from './request'

const path = (template: string, id: string | number) => template.replace('{id}', encodeURIComponent(String(id)))
const unwrapList = (value: any): any[] => Array.isArray(value) ? value : value?.list || value?.records || value?.items || value?.data || []
const dateText = (value: any) => {
  if (!value) return ''
  const d = new Date(Number(value) < 1e12 ? Number(value) * 1000 : value)
  return isNaN(d.getTime()) ? '' : d.toLocaleString('zh-CN', { hour12: false })
}

// 后端枚举 -> 前端内部文案（文案 key 已在语言包中）
const NEWS_LEVEL: Record<number, string> = { 0: '普通', 1: '普通', 2: '要闻', 3: '紧急' }
const NEWS_SENTIMENT: Record<number, string> = { 0: '中性', 1: '中性', 2: '利好', 3: '利空' }
const POINT_STATUS: Record<number, string> = { 1: '进行中', 2: '已失效' }

function stripHtml(html: any): string {
  if (!html) return ''
  return String(html)
    .replace(/<br\s*\/?>/gi, '\n')
    .replace(/<\/(p|div|li|h[1-6])>/gi, '\n')
    .replace(/<[^>]+>/g, '')
    .replace(/&nbsp;/gi, ' ')
    .replace(/&amp;/gi, '&')
    .replace(/&lt;/gi, '<')
    .replace(/&gt;/gi, '>')
    .replace(/&quot;/gi, '"')
    .replace(/&#39;/gi, "'")
    .replace(/\n{3,}/g, '\n\n')
    .trim()
}

function truncate(text: string, max = 80): string {
  return text.length > max ? `${text.slice(0, max)}…` : text
}

function normalizeNews(row: any): any {
  const body = stripHtml(row.content || row.summary || row.description || '')
  return {
    ...row,
    id: String(row.id ?? row.newsId ?? ''),
    title: row.title || row.name || '',
    summary: truncate(body) || row.summary || '',
    body,
    source: row.source || row.author || '',
    time: row.time || dateText(row.publishedAt || row.createdAt || row.created_at || row.publishTime),
    tags: String(row.tagName || row.tags || row.tagNames || '').split(',').map((s: string) => s.trim()).filter(Boolean).join(' '),
    level: NEWS_LEVEL[row.level] || row.level || '',
    sentiment: NEWS_SENTIMENT[row.sentiment] || row.sentiment || '',
  }
}

function parseTakeProfits(value: any): { levels: string[]; ratios: string[] } {
  const levels: string[] = []
  const ratios: string[] = []
  String(value || '').split(',').map((s) => s.trim()).filter(Boolean).forEach((part) => {
    const m = part.match(/^([\d.,]+)\s*(?:\(([^)]*)\))?/)
    if (m) {
      levels.push(m[1])
      if (m[2]) ratios.push(m[2])
    }
  })
  return { levels, ratios }
}

function normalizePoint(row: any): any {
  const type = row.type ?? row.direction
  const { levels, ratios } = parseTakeProfits(row.takeProfits ?? row.take_profits ?? row.tp ?? row.description)
  return {
    ...row,
    id: String(row.id ?? row.pointId ?? ''),
    coin: row.coinName || row.coin || row.symbol || row.name || '',
    dir: row.dir || row.typeName || (Number(type) === 2 ? '做空' : Number(type) === 1 ? '做多' : ''),
    status: POINT_STATUS[row.status] || row.statusName || row.status || '',
    range: row.entryPrice ?? row.entry_price ?? row.range ?? '',
    currentPrice: row.currentPrice ?? row.coinPrice ?? '',
    tp: levels.join(' / '),
    tpLevels: levels,
    tpRatios: ratios,
    sl: row.stopLoss ?? row.stop_loss ?? row.sl ?? '',
    published: row.published || dateText(row.createdAt || row.created_at || row.publishTime || row.publishedAt),
    validity: row.validity || (row.validDuration ? `${row.validDuration}H` : ''),
    strategy: row.notes ?? row.strategy ?? '',
    locked: Boolean(row.locked ?? row.isLocked ?? false),
  }
}

function normalizeMarket(x: any): any[] {
  const raw = String(x.symbol || x.name || '').toUpperCase()
  const symbol = raw && !raw.includes('/') ? `${raw}/USDT` : raw
  const price = x.current_price ?? x.price ?? ''
  const change = Number(x.price_change_percentage_24h ?? x.change24h ?? x.change ?? 0)
  return [symbol, price, `${change >= 0 ? '+' : ''}${change.toFixed(2)}%`]
}

function normalizeCalendarEvent(row: any): any {
  return {
    ...row,
    date: row.date || row.publishDate || dateText(row.publishTimestamp || row.createdAt),
    title: row.title || row.name || row.event || '',
    country: row.country || row.region || row.currency || '',
    impact: row.impact || row.level || (row.isImportant ? '要闻' : ''),
  }
}

function normalizeMessage(row: any): any {
  const type = Number(row.type)
  return {
    ...row,
    id: String(row.notificationId ?? row.id ?? ''),
    type: type === 3 ? 'news' : type === 2 || type === 4 ? 'point' : 'system',
    target: String(row.link ?? row.target ?? ''),
    title: row.title || row.subject || '',
    summary: truncate(row.summary || stripHtml(row.content) || '', 60),
    time: row.time || dateText(row.createdAt || row.created_at || row.publishTime),
    read: row.isRead != null ? Number(row.isRead) === 1 : Boolean(row.read ?? row.readStatus === 1),
  }
}

export const api = {
  register: (data: any) => request('/app/register', { method: 'POST', auth: false, data }),
  sendLoginCode: (email: string, type = 1) => request('/profile/sendEmailCode', { method: 'POST', auth: false, data: { email, type } }),
  forgetPassword: (data: any) => request('/profile/forgetPassword', { method: 'POST', auth: false, data }),
  query2fa: (email: string) => request('/profile/query2fa', { auth: false, query: { email } }),
  loginByCode: (email: string, code: string, twoFactorCode = '') => request<any>('/app/login/captcha', { method: 'POST', auth: false, data: { email, code, twoFactorCode } }),
  loginByPassword: (email: string, password: string, twoFactorCode?: string) => request<any>('/app/login', { method: 'POST', auth: false, data: { email, password, twoFactorCode } }),
  loginWallet: (walletAddress: string, chainId: string, isRemember = true) => request<any>('/app/login/wallet', { method: 'POST', auth: false, data: { walletAddress, chainId, isRemember } }),
  loginGoogle: (data: any) => request<any>('/app/login/google', { method: 'POST', auth: false, data }),
  loginMicrosoft: (data: any) => request<any>('/app/login/microsoft', { method: 'POST', auth: false, data }),
  oauthAuth: (provider: 'google' | 'microsoft', redirectUrl: string) => request<any>(`/oauth/${provider}/auth`, { method: 'POST', auth: false, data: { redirectUrl } }),
  oauthCallback: (provider: 'google' | 'microsoft', code: string, state: string) => request<any>(`/oauth/${provider}/callback`, { method: 'POST', auth: false, data: { code, state } }),
  bindGoogle: (data: any) => request('/bind/google/binding', { method: 'POST', data }),
  bindMicrosoft: (data: any) => request('/bind/microsoft/binding', { method: 'POST', data }),
  unbindGoogle: () => request('/bind/google/unbinding', { method: 'DELETE' }),
  unbindMicrosoft: () => request('/bind/microsoft/unbinding', { method: 'DELETE' }),
  bindStatus: () => request<any>('/bind/status'),
  logoutRemote: () => request('/app/logout', { method: 'POST' }),
  home: async () => {
    const safe = (p: Promise<any>) => p.catch(() => null)
    const [overview, ratio, calendar, news, latestPoint] = await Promise.all([
      safe(request<any>('/app/home/market-overview')),
      safe(request<any>('/app/home/long-short-ratio')),
      safe(request<any>('/app/home/financial-calendar')),
      safe(request<any>('/app/home/news-flash')),
      safe(request<any>('/app/home/latest-trading-point')),
    ])
    const r = ratio?.data || ratio || {}
    const buy = Number(r.buyRatio ?? r.longRatio ?? r.longPercent ?? r.long ?? 0)
    const sell = Number(r.sellRatio ?? r.shortRatio ?? r.shortPercent ?? r.short ?? 0)
    const long = buy > 0 && buy < 1 ? buy * 100 : buy
    const short = sell > 0 && sell < 1 ? sell * 100 : sell
    const lp = latestPoint?.data || latestPoint
    return {
      ratio: {
        long: Number(long.toFixed(2)),
        short: Number(short.toFixed(2)),
        value: sell > 0 ? Number((buy / sell).toFixed(2)) : 0,
      },
      coins: unwrapList(overview).map(normalizeMarket),
      calendar: unwrapList(calendar).map(normalizeCalendarEvent),
      news: unwrapList(news).map(normalizeNews),
      points: lp ? [normalizePoint(lp)] : [],
    }
  },
  financialCalendar: async () => unwrapList(await request('/app/home/financial-calendar')).map(normalizeCalendarEvent),
  newsTags: async () => unwrapList(await request('/app/news/getTags')).map((x: any) => typeof x === 'string' ? { id: x, name: x, code: x, enName: x } : ({ id: x.id ?? x.tagId, name: x.name ?? x.tagName ?? x.title, code: x.code, enName: x.enName })).filter((x: any) => x.name),
  news: async (category?: string, page = 1, size = 20) => {
    const data = await request<any>('/app/news/list', { query: { tagId: category && category !== '全部' && /^\d+$/.test(category) ? category : undefined, keyword: category && category !== '全部' && !/^\d+$/.test(category) ? category : undefined, page, size } })
    return unwrapList(data).map(normalizeNews)
  },
  newsById: async (id: string) => normalizeNews(await request(`/app/news/detail/${encodeURIComponent(id)}`)),
  shareNews: (id: string) => request(path('/app/news/share/{id}', id), { method: 'POST' }),
  favoriteNews: (id: string, value: boolean) => request(path(value ? '/app/news/favorite/{id}' : '/app/news/unfavorite/{id}', id), { method: 'POST' }),
  points: async (filter?: string, page = 1, size = 20) => {
    const type = filter === '做多' ? 1 : filter === '做空' ? 2 : filter === '自选' ? 3 : undefined
    const data = await request<any>('/app/trading-point/list', { query: { type, page, size } })
    return unwrapList(data).map(normalizePoint)
  },
  pointById: async (id: string) => normalizePoint(await request(`/app/trading-point/detail/${encodeURIComponent(id)}`)),
  sharePoint: (id: string) => request(path('/app/trading-point/share/{id}', id), { method: 'POST' }),
  favoritePoint: (id: string, value: boolean) => request(path(value ? '/app/trading-point/favorite/{id}' : '/app/trading-point/unfavorite/{id}', id), { method: 'POST' }),
  unlockAll: () => request('/app/trading-point/unlock-all'),
  messages: async (query: any = {}) => unwrapList(await request('/app/profile/messageList', { query })).map(normalizeMessage),
  markAllMessagesRead: async () => { const rows = await api.messages({ isRead: 2, page: 1, size: 100 }); await Promise.all(rows.map((x: any) => api.markMessageRead(x.id))); return rows },
  markMessageRead: (id: string) => request(path('/app/profile/markMessageRead/{id}', id), { method: 'POST' }),
  myCollection: async () => request<any>('/app/profile/myCollection'),
  contacts: async () => unwrapList(await request('/app/profile/contact')).map((x: any) => ({ type: x.name || x.customerName || x.type || '客服', account: x.email || x.phone || x.account || '', qr: x.qr || x.qrCode || '' })),
  appVersion: () => request('/app/profile/appVersion'),
  changePassword: (data: any) => request('/app/profile/password', { method: 'POST', data }),
  notifications: async () => { const data: any = await request('/app/profile/messageSettingList'); return [1, 2, 3, 4].map((i) => Boolean(data?.[i] ?? data?.[String(i)])) },
  saveNotifications: async (values: boolean[]) => { await Promise.all(values.map((value, index) => api.saveNotification(index + 1, value))); return values },
  saveNotification: (type: number, status: boolean) => request('/app/profile/messageSettingUpdate', { method: 'POST', data: { type, status: status ? 1 : 2 } }),
  bindWallet: (walletAddress: string) => request('/app/profile/bindWallet', { method: 'POST', data: { walletAddress } }),
  get2fa: () => request<any>('/app/profile/get2fa'),
  update2fa: (data: any) => request('/app/profile/update2fa', { method: 'POST', data }),
  disable2fa: (data: any) => request('/app/profile/enabled2fa', { method: 'POST', data }),
  favorites: async (type: 'news' | 'points') => {
    const data: any = await api.myCollection()
    if (type === 'news') return unwrapList(data?.newsList).map(normalizeNews)
    return unwrapList(data?.pointList).map(normalizePoint)
  },
  isFavorite: async (id: string, type: 'news' | 'points' = 'news') => (await api.favorites(type)).some((x: any) => x.id === id),
  toggleFavorite: async (id: string, type: 'news' | 'points' = 'news', value?: boolean) => { const next = value ?? !(await api.isFavorite(id, type)); await (type === 'news' ? api.favoriteNews(id, next) : api.favoritePoint(id, next)); return next },
  profile: async () => request('/app/profile/personalInfo'),
  updatePersonalInfo: (data: any) => request('/app/profile/updatePersonalInfo', { method: 'POST', data }),
  markets: async () => unwrapList(await request('/app/home/market-overview')).map(normalizeMarket),
  refreshToken: (isRemember = true) => request<any>('/app/profile/refreshToken', { method: 'POST', data: { isRemember }, skipRefresh: true }),
  staticContent: async (kind: string) => {
    if (kind === 'privacy') return { title: '隐私条款', body: '隐私条款内容尚未配置，请联系服务方获取正式文本。' }
    if (kind === 'service') return { title: '服务协议', body: '服务协议内容尚未配置，请联系服务方获取正式文本。' }
    const x: any = await api.appVersion()
    return { title: '关于我们', body: x?.description || x?.content || '' }
  }
}
