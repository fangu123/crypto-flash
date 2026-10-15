import type { RequestOptions } from './request'
import { locale } from '../locales'

const news = [
  { id: 'n-1001', title: '比特币现货 ETF 连续录得净流入，市场关注资金持续性', summary: '机构资金流向成为短期市场焦点，分析师提醒关注成交量与宏观数据变化。', body: '比特币现货 ETF 近期录得连续净流入，市场参与者正评估资金流入对现货需求的影响。分析人士指出，单日流量不代表长期趋势，投资者仍需结合成交量、流动性和宏观环境判断。', source: 'Crypto Flash', time: '2026-10-11 09:30', tags: 'BTC ETF', level: '要闻', sentiment: '利好' },
  { id: 'n-1002', title: '以太坊网络活跃度回升，Layer 2 生态数据改善', summary: '过去一周链上交互增加，多个扩容网络的日活地址同步上升。', body: '以太坊及其扩容网络的链上交互数据有所改善。网络活动可能受到短期活动和激励影响，数据变化不构成价格方向的保证。', source: '链上观察', time: '2026-10-11 08:45', tags: 'ETH 山寨币', level: '快讯', sentiment: '中性' },
  { id: 'n-1003', title: '主要交易平台公布储备金报告更新', summary: '报告披露了主要资产储备情况，市场继续关注透明度和托管风险。', body: '多家交易平台更新储备金披露。储备金报告的范围、负债核验方式和更新时间各不相同，用户应结合平台说明独立评估托管风险。', source: '市场简报', time: '2026-10-10 21:18', tags: '交易所', level: '快讯', sentiment: '中性' },
  { id: 'n-1004', title: '美国宏观数据公布前，加密市场波动率有所收窄', summary: '交易员降低方向性敞口，等待新一轮经济数据。', body: '宏观数据发布前，多个主要加密资产的短期波动率有所收窄。数据公布后市场仍可能出现快速波动，请注意仓位和风险控制。', source: '全球市场', time: '2026-10-10 18:05', tags: 'BTC 政策', level: '要闻', sentiment: '中性' },
  { id: 'n-1005', title: 'Solana 生态开发者活动增加，应用更新陆续发布', summary: '基础设施与消费应用方向均有新版本上线。', body: 'Solana 生态近期发布多项应用和基础设施更新。开发者活动是生态观察指标之一，但不能单独作为投资依据。', source: '生态周报', time: '2026-10-10 15:42', tags: 'SOL 山寨币', level: '快讯', sentiment: '利好' },
  { id: 'n-1006', title: '监管机构发布数字资产风险提示', summary: '公告提醒用户核实平台资质并警惕高收益承诺。', body: '监管机构提示消费者关注数字资产交易和托管风险，核验服务提供方资质，谨慎对待保本或高收益宣传。', source: '政策速递', time: '2026-10-10 11:20', tags: '政策 交易所', level: '要闻', sentiment: '中性' },
  { id: 'n-1007', title: '比特币链上长期持有者供应保持稳定', summary: '链上指标显示长期持有者群体暂未出现明显变化。', body: '链上数据平台的样本指标显示长期持有者供应总体稳定。链上指标存在口径差异，且历史行为不代表未来表现。', source: '链上观察', time: '2026-10-09 22:06', tags: 'BTC', level: '快讯', sentiment: '中性' },
  { id: 'n-1008', title: '稳定币总供应量小幅增加，流动性指标待观察', summary: '新增供应分布在多个网络，市场关注实际交易需求。', body: '稳定币供应指标小幅增加。供应变化需要结合交易所余额、链上转账和实际使用场景综合理解。', source: '市场简报', time: '2026-10-09 16:31', tags: '交易所 山寨币', level: '快讯', sentiment: '利好' }
]

const points = [
  { id: 'p-2001', coin: 'BTC/USDT', dir: '做多', status: '进行中', range: '62,400 - 63,100', currentPrice: '62,780', tp: '64,200 / 65,500', tpLevels: ['64,200', '65,500'], tpRatios: ['40%', '60%'], sl: '61,600', published: '2026-10-11 08:00', validity: '24H', strategy: '演示点位：关注区间支撑有效性，设置止损并控制仓位。', locked: false },
  { id: 'p-2002', coin: 'ETH/USDT', dir: '做多', status: '进行中', range: '3,280 - 3,340', currentPrice: '3,315', tp: '3,410 / 3,520', tpLevels: ['3,410', '3,520'], tpRatios: ['50%', '50%'], sl: '3,220', published: '2026-10-10 20:30', validity: '24H', strategy: '演示点位：观察区间内的量价变化，避免追高。', locked: false },
  { id: 'p-2003', coin: 'SOL/USDT', dir: '做空', status: '进行中', range: '148 - 152', currentPrice: '150.6', tp: '145 / 141', tpLevels: ['145', '141'], tpRatios: ['50%', '50%'], sl: '155', published: '2026-10-10 16:15', validity: '12H', strategy: '演示点位：仅作界面演示，不构成交易建议。', locked: true },
  { id: 'p-2004', coin: 'BNB/USDT', dir: '做多', status: '已失效', range: '580 - 590', currentPrice: '601', tp: '608', tpLevels: ['608'], tpRatios: ['100%'], sl: '568', published: '2026-10-09 12:10', validity: '24H', strategy: '演示点位已失效。', locked: false }
]

const calendar = [
  { id: 'c-1', date: '10/11 20:30', title: '美国初请失业金人数', country: '美国 USD', impact: '高' },
  { id: 'c-2', date: '10/12 02:00', title: '美联储官员讲话', country: '美国 USD', impact: '中' },
  { id: 'c-3', date: '10/12 14:00', title: '欧元区工业产出', country: '欧元区 EUR', impact: '中' }
]

const initialMessages = [
  { id: 'm-1', type: 'news', target: 'n-1001', title: '市场快讯', summary: '比特币现货 ETF 连续录得净流入。', time: '今天 09:30', read: false },
  { id: 'm-2', type: 'point', target: 'p-2001', title: '新交易点位', summary: 'BTC/USDT 演示点位已发布。', time: '今天 08:00', read: false },
  { id: 'm-3', type: 'system', target: '', title: '风险提示', summary: '演示数据不代表实时行情或投资建议。', time: '昨天 18:00', read: true }
]

const get = <T>(key: string, fallback: T): T => {
  try { return uni.getStorageSync(`cf_mock_${key}`) || fallback } catch { return fallback }
}
const set = (key: string, value: any) => uni.setStorageSync(`cf_mock_${key}`, value)
const copy = <T>(value: T): T => JSON.parse(JSON.stringify(value))

function localizedNews(rows: any[]) {
  if (locale.value === 'zh') return rows
  const titles = [
    'Bitcoin spot ETFs record consecutive net inflows',
    'Ethereum network activity rebounds as Layer 2 data improves',
    'Major exchanges publish updated proof of reserves',
    'Crypto market volatility narrows ahead of US macro data',
    'Solana developer activity rises with new app releases',
    'Regulators issue a digital asset risk warning',
    'Bitcoin long-term holder supply remains stable',
    'Stablecoin supply edges higher as liquidity is monitored'
  ]
  return rows.filter(Boolean).map((row) => ({ ...row, title: titles[Number(String(row.id).split('-').pop()) - 1001] || row.title, summary: 'Market participants are monitoring liquidity, volume, and macro conditions.', body: 'This demo article is provided for interface preview only and is not investment advice.', source: 'Crypto Flash', level: row.level === '要闻' ? 'Headline' : 'Flash', sentiment: row.sentiment === '利好' ? 'Bullish' : row.sentiment === '利空' ? 'Bearish' : 'Neutral', tags: String(row.tags).replace('山寨币', 'Altcoins').replace('交易所', 'Exchange').replace('政策', 'Policy') }))
}

function localizedPoints(rows: any[]) {
  if (locale.value === 'zh') return rows
  return rows.map(row => ({ ...row, dir: row.dir === '做多' ? 'Long' : 'Short', status: row.status === '进行中' ? 'Active' : 'Expired', strategy: 'Demo signal: monitor price action, set a stop loss, and manage position size.' }))
}

export async function mockRequest<T = any>(path: string, options: RequestOptions = {}): Promise<T> {
  const query = options.query || {}
  const data = options.data || {}
  let result: any

  switch (path) {
    case '/app/register': result = { success: true }; break
    case '/profile/sendEmailCode': result = { success: true, demoCode: '123456' }; break
    case '/profile/forgetPassword': result = { success: true }; break
    case '/profile/query2fa': result = { enabled: false }; break
    case '/app/login/captcha':
    case '/app/login':
      set('profile', { ...get('profile', { nickname: '演示用户', email: '' }), email: data.email || '' })
      result = { token: 'demo-local-session', userId: 'demo-user' }
      break
    case '/app/logout': result = { success: true }; break
    case '/app/home/market-overview':
      result = { tokens: [['BTC/USDT', '62,780.00', '+2.35%'], ['ETH/USDT', '3,315.00', '+1.25%'], ['SOL/USDT', '150.60', '-0.45%'], ['BNB/USDT', '601.20', '+0.85%']].map(([symbol, price, change24h]) => ({ symbol, price, change24h: change24h.replace('%', '') })) }
      break
    case '/app/home/long-short-ratio': result = { long: 57.62, short: 42.38, value: 1.36 }; break
    case '/app/home/financial-calendar': result = copy(calendar); break
    case '/app/home/news-flash': result = localizedNews(copy(news.slice(0, 5))); break
    case '/app/home/latest-trading-point': result = localizedPoints(copy(points.slice(0, 3))); break
    case '/app/news/getTags': result = locale.value === 'en' ? ['BTC', 'ETH', 'Altcoins', 'Exchange', 'Policy'] : ['BTC', 'ETH', '山寨币', '交易所', '政策']; break
    case '/app/news/list': {
      const keyword = String(query.keyword || '').toLowerCase()
      const tag = String(query.tagId || '')
      const rows = news.filter(row => (!keyword || `${row.title} ${row.summary} ${row.tags}`.toLowerCase().includes(keyword)) && (!tag || row.tags.includes(tag)))
      const page = Number(query.page || 1), size = Number(query.size || 20)
      result = localizedNews(copy(rows.slice((page - 1) * size, page * size)))
      break
    }
    case '/app/trading-point/list': {
      const filter = Number(query.type)
      const rows = points.filter(row => !filter || (filter === 1 ? row.dir === '做多' : row.dir === '做空'))
      const page = Number(query.page || 1), size = Number(query.size || 20)
      result = localizedPoints(copy(rows.slice((page - 1) * size, page * size)))
      break
    }
    case '/app/trading-point/unlock-all':
      set('unlocked', true)
      result = { success: true }
      break
    case '/app/profile/myCollection': result = copy(get('favorites', [])); break
    case '/app/profile/contact': result = [{ type: 'Telegram', account: '@CryptoFlashDemo', qr: '' }, { type: 'Email', account: 'support@cryptoflash.example', qr: '' }]; break
    case '/app/profile/appVersion': result = { version: '1.0.0-demo', description: 'Crypto Flash 演示版\n本版本使用本地模拟数据，不代表实时市场行情。' }; break
    case '/app/profile/messageList': {
      const rows = copy(get('messages', initialMessages))
      result = query.isRead === 0 || query.isRead === '0' ? rows.filter((row: any) => !row.read) : rows
      break
    }
    case '/app/profile/messageSettingList': result = get('notifications', [true, true, true, false]).map((status: boolean) => ({ status })); break
    case '/app/profile/messageSettingUpdate': {
      const values = get<boolean[]>('notifications', [true, true, true, false])
      values[Number(data.type) - 1] = Boolean(data.status)
      set('notifications', values)
      result = { success: true }
      break
    }
    case '/app/profile/get2fa': result = { enabled: Boolean(get('2fa', false)), secret: 'DEMO-ONLY-SECRET' }; break
    case '/app/profile/update2fa': set('2fa', true); result = { success: true }; break
    case '/app/profile/enabled2fa': set('2fa', false); result = { success: true }; break
    case '/app/profile/password': result = { success: true }; break
    case '/app/profile/bindWallet': {
      const profile = get('profile', { nickname: '演示用户', email: '' })
      profile.walletAddress = data.walletAddress
      set('profile', profile)
      result = { success: true }
      break
    }
    default:
      if (path.startsWith('/app/profile/markMessageRead/')) {
        const id = pathId(options, path)
        const rows = get<any[]>('messages', initialMessages)
        const row = rows.find(item => String(item.id) === id)
        if (row) row.read = true
        set('messages', rows)
        result = { success: true }
      } else if (/^\/app\/(news|trading-point)\/share\//.test(path)) result = { success: true }
      else if (path.startsWith('/app/news/detail/')) result = localizedNews(copy([news.find(row => row.id === decodeURIComponent(path.split('/').pop() || ''))])).pop()
      else if (path.startsWith('/app/trading-point/detail/')) result = localizedPoints(copy([points.find(row => row.id === decodeURIComponent(path.split('/').pop() || ''))])).pop()
      else if (/^\/app\/(news|trading-point)\/(favorite|unfavorite)\//.test(path)) {
        const [, kind, action, id] = path.split('/')
        const rows = get<any[]>('favorites', [])
        const collectionKey = kind === 'news' ? 'newsId' : 'pointId'
        const existing = rows.findIndex(row => String(row[collectionKey]) === decodeURIComponent(id))
        if (action === 'favorite' && existing < 0) rows.push({ [collectionKey]: decodeURIComponent(id) })
        if (action === 'unfavorite' && existing >= 0) rows.splice(existing, 1)
        set('favorites', rows)
        result = { success: true }
      } else if (path === '/app/profile/messageSettingUpdate') result = { success: true }
      else if (path === '/app/profile/password') result = { success: true }
      else if (path === '/app/profile/getProfile' || path === '/app/profile/info') result = get('profile', { nickname: '演示用户', email: '' })
      else if (path === '/app/profile') result = get('profile', { nickname: '演示用户', email: '' })
      else if (path === '/app/profile/appVersion') result = { version: '1.0.0-demo', description: 'Crypto Flash 演示版' }
      else if (path === '/app/home/market-overview') result = { tokens: [] }
      else throw new Error(`未配置模拟接口: ${path}`)
  }

  return result as T
}

function pathId(_options: RequestOptions, path: string) {
  const match = path.match(/\/([^/]+)$/)
  return decodeURIComponent(match?.[1] || '')
}
