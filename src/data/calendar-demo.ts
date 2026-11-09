export type CalendarCategory = 'macro' | 'policy' | 'crypto'

export type CalendarEvent = {
  id: string
  date: string
  time: string
  category: CalendarCategory
  kind: 'economic' | 'unlock' | 'project' | 'policy'
  title: string
  enTitle: string
  description?: string
  enDescription?: string
  importance: 1 | 2 | 3
  previous?: string
  forecast?: string
  actual?: string
}

// Frontend demonstration only. Replace this source when the event API is ready.
export const calendarDemoStart = { year: 2026, month: 6, day: 13 }

export function calendarDateKey(date: Date) {
  return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, '0')}-${String(date.getDate()).padStart(2, '0')}`
}

export function getDemoCalendarEvents(year: number, month: number): CalendarEvent[] {
  const event = (day: number, id: string, data: Omit<CalendarEvent, 'id' | 'date'>): CalendarEvent => {
    const date = calendarDateKey(new Date(year, month, day))
    return { ...data, id: `${date}-${id}`, date }
  }
  return [
    event(13, 'us-cpi', {
      time: '20:30', category: 'macro', kind: 'economic', importance: 3,
      title: '美国 6 月 CPI 年率', enTitle: 'US June CPI YoY',
      previous: '3.3%', forecast: '3.1%', actual: '--',
    }),
    event(13, 'apt-unlock', {
      time: '22:00', category: 'crypto', kind: 'unlock', importance: 2,
      title: 'APT 代币解锁', enTitle: 'APT token unlock',
      description: '解锁11.31M APT,约占流通量2.4%,价值约$9.8M',
      enDescription: '11.31M APT unlock · 2.4% of circulating supply · approx. $9.8M',
    }),
    event(13, 'arb-vote', {
      time: '23:30', category: 'crypto', kind: 'project', importance: 1,
      title: 'Arbitrum DAO 治理投票截止', enTitle: 'Arbitrum DAO voting deadline',
      description: '提案 AIP-42：调整排序器费用分配机制',
      enDescription: 'AIP-42: adjust the sequencer fee allocation mechanism',
    }),
    event(14, 'policy-hearing', {
      time: '16:00', category: 'policy', kind: 'policy', importance: 2,
      title: '数字资产监管听证会', enTitle: 'Digital asset regulatory hearing',
      description: '关注稳定币发行与数字资产服务商监管议题',
      enDescription: 'Stablecoin issuance and oversight of digital asset service providers',
    }),
    event(15, 'us-ppi', {
      time: '20:30', category: 'macro', kind: 'economic', importance: 3,
      title: '美国 PPI 年率', enTitle: 'US PPI YoY',
      previous: '2.6%', forecast: '2.3%', actual: '--',
    }),
    event(17, 'eth-community', {
      time: '19:00', category: 'crypto', kind: 'project', importance: 1,
      title: '以太坊生态社区会议', enTitle: 'Ethereum ecosystem community call',
      description: '讨论网络升级进展与生态项目路线图',
      enDescription: 'Network upgrade progress and ecosystem project roadmaps',
    }),
    event(25, 'policy-update', {
      time: '10:00', category: 'policy', kind: 'policy', importance: 2,
      title: '加密资产政策进展说明会', enTitle: 'Crypto asset policy briefing',
      description: '关注交易平台合规与投资者保护相关进展',
      enDescription: 'Updates on exchange compliance and investor protection',
    }),
  ]
}
