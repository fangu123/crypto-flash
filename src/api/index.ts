import { request } from './request'

const path = (template: string, id: string | number) => template.replace('{id}', encodeURIComponent(String(id)))
const unwrapList = (value: any): any[] => Array.isArray(value) ? value : value?.list || value?.records || value?.items || value?.data || []
const dateText = (value: any) => value ? new Date(Number(value) < 1e12 ? Number(value) * 1000 : value).toLocaleString('zh-CN', { hour12: false }) : ''
function normalizeNews(row: any): any { const tags = Array.isArray(row.tags) ? row.tags.join(' ') : (row.tags || row.tagNames || ''); return { ...row, id: String(row.id ?? row.newsId ?? ''), title: row.title || row.name || '', summary: row.summary || row.content || row.description || '', source: row.source || row.author || '', time: row.time || dateText(row.publishTime || row.createdAt), tags: String(tags) } }
function normalizePoint(row: any): any { return { ...row, id: String(row.id ?? row.pointId ?? ''), coin: row.coin || row.symbol || row.name || '', dir: row.dir || row.typeName || (String(row.type).includes('2') ? '做空' : '做多'), status: row.status || row.statusName || '', published: row.published || dateText(row.createdAt || row.publishTime) } }

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
  home: async () => { const [overview, ratio, calendar, news, points] = await Promise.all([request<any>('/app/home/market-overview'), request<any>('/app/home/long-short-ratio'), request<any>('/app/home/financial-calendar'), request<any>('/app/home/news-flash'), request<any>('/app/home/latest-trading-point')]); const r = ratio?.data || ratio || {}; return { ratio: { long: Number(r.long ?? r.longRatio ?? r.longPercent ?? 0), short: Number(r.short ?? r.shortRatio ?? r.shortPercent ?? 0), value: Number(r.value ?? r.ratio ?? 0) }, coins: unwrapList(overview).map((x: any) => { const change = Number(x.price_change_percentage_24h ?? x.change24h ?? x.change ?? 0); return [x.symbol || x.name || '', x.current_price ?? x.price ?? '', `${change >= 0 ? '+' : ''}${change}%`] }), calendar: unwrapList(calendar), news: unwrapList(news).map(normalizeNews), points: unwrapList(points).map(normalizePoint) } },
  financialCalendar: () => request('/app/home/financial-calendar'),
  newsTags: async () => unwrapList(await request('/app/news/getTags')),
  news: async (category?: string, page = 1, size = 20) => { const data = await request<any>('/app/news/list', { query: { tagId: category && category !== '全部' && /^\d+$/.test(category) ? category : undefined, keyword: category && category !== '全部' && !/^\d+$/.test(category) ? category : undefined, page, size } }); return unwrapList(data).map(normalizeNews) },
  newsById: async (id: string) => normalizeNews(await request(`/app/news/detail/${encodeURIComponent(id)}`)),
  shareNews: (id: string) => request(path('/app/news/share/{id}', id), { method: 'POST' }),
  favoriteNews: (id: string, value: boolean) => request(path(value ? '/app/news/favorite/{id}' : '/app/news/unfavorite/{id}', id), { method: 'POST' }),
  points: async (filter?: string, page = 1, size = 20) => { const data = await request<any>('/app/trading-point/list', { query: { type: filter === '做多' ? 1 : filter === '做空' ? 2 : undefined, page, size } }); return unwrapList(data).map(normalizePoint) },
  pointById: async (id: string) => normalizePoint(await request(`/app/trading-point/detail/${encodeURIComponent(id)}`)),
  sharePoint: (id: string) => request(path('/app/trading-point/share/{id}', id), { method: 'POST' }),
  favoritePoint: (id: string, value: boolean) => request(path(value ? '/app/trading-point/favorite/{id}' : '/app/trading-point/unfavorite/{id}', id), { method: 'POST' }),
  unlockAll: () => request('/app/trading-point/unlock-all'),
  messages: async (query: any = {}) => unwrapList(await request('/app/profile/messageList', { query })).map((x: any) => ({ ...x, id: String(x.id), title: x.title || x.subject || '', summary: x.summary || x.content || '', time: x.time || dateText(x.createdAt), read: Boolean(x.read ?? x.isRead ?? x.readStatus) })),
  markAllMessagesRead: async () => { const rows = await api.messages({ isRead: 0, page: 1, size: 100 }); await Promise.all(rows.map((x: any) => api.markMessageRead(x.id))); return rows },
  markMessageRead: (id: string) => request(path('/app/profile/markMessageRead/{id}', id), { method: 'POST' }),
  myCollection: () => request('/app/profile/myCollection'), contacts: async () => unwrapList(await request('/app/profile/contact')),
  appVersion: () => request('/app/profile/appVersion'), changePassword: (data: any) => request('/app/profile/password', { method: 'POST', data }),
  notifications: async () => unwrapList(await request('/app/profile/messageSettingList')).map((x: any) => Boolean(x.status ?? x.enabled)),
  saveNotifications: async (values: boolean[]) => { await Promise.all(values.map((value, index) => api.saveNotification(index + 1, value))); return values },
  saveNotification: (type: number, status: boolean) => request('/app/profile/messageSettingUpdate', { method: 'POST', data: { type, status: status ? 1 : 0 } }),
  bindWallet: (walletAddress: string) => request('/app/profile/bindWallet', { method: 'POST', data: { walletAddress } }), get2fa: () => request('/app/profile/get2fa'),
  update2fa: (data: any) => request('/app/profile/update2fa', { method: 'POST', data }), disable2fa: (data: any) => request('/app/profile/enabled2fa', { method: 'POST', data }),
  favorites: async (type: 'news' | 'points') => { const rows = unwrapList(await api.myCollection()); return rows.filter((x: any) => type === 'news' ? x.newsId || x.title : x.pointId || x.coin).map(type === 'news' ? normalizeNews : normalizePoint) },
  isFavorite: async (id: string, type: 'news' | 'points' = 'news') => (await api.favorites(type)).some((x: any) => x.id === id),
  toggleFavorite: async (id: string, type: 'news' | 'points' = 'news', value?: boolean) => { const next = value ?? !(await api.isFavorite(id, type)); await (type === 'news' ? api.favoriteNews(id, next) : api.favoritePoint(id, next)); return next },
  profile: async () => request('/app/profile/personalInfo'),
  updatePersonalInfo: (data: any) => request('/app/profile/updatePersonalInfo', { method: 'POST', data }),
  markets: async () => { const data: any = await request('/app/home/market-overview'); return unwrapList(data).map((x: any) => { const change = Number(x.price_change_percentage_24h ?? x.change24h ?? x.change ?? 0); return [x.symbol || x.name || '', x.current_price ?? x.price ?? '', `${change >= 0 ? '+' : ''}${change}%`, '24h'] }) },
  refreshToken: (isRemember = true) => request<any>('/app/profile/refreshToken', { method: 'POST', data: { isRemember }, skipRefresh: true }),
  staticContent: async (kind: string) => {
    if (kind === 'privacy') return { title: '隐私条款', body: '隐私条款内容尚未配置，请联系服务方获取正式文本。' }
    if (kind === 'service') return { title: '服务协议', body: '服务协议内容尚未配置，请联系服务方获取正式文本。' }
    const x: any = await api.appVersion()
    return { title: '关于我们', body: x?.description || x?.content || '' }
  }
}
