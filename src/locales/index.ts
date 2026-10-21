import { ref } from 'vue'

export type Locale = 'zh' | 'en'
const storageKey = 'cf_locale'
export const locale = ref<Locale>('zh')

const translations: Record<string, string> = Object.assign({}, {
  '首页': 'Home', '快讯': 'Flash', '点位': 'Signals', '我的': 'Profile',
  '精选': 'Featured', '日历': 'Calendar', '数据': 'Data', '更多': 'More',
  '市场行情': 'Market', 'market': 'Market', 'ratio': 'Long/Short (BTC)', 'featured': 'Featured',
  '多空比 (BTC)': 'Long/Short (BTC)', '解锁全部点位': 'Unlock all signals', '联系客服': 'Contact support',
  '查看全部 ›': 'View all ›', '更多 ›': 'More ›', '实时数据': 'Live data', '热门快讯': 'Top stories',
  '多方': 'Long', '空方': 'Short', '多方主导': 'Longs dominate', '重要财经日历': 'Economic calendar',
  '账号登录': 'Account login', '钱包连接': 'Connect wallet', '跳过': 'Skip', '登录': 'Sign in', '注册': 'Sign up', '游客模式': 'Guest mode', '登录后同步收藏与个人设置': 'Sign in to sync saved items and settings', '专业级加密资讯与高胜率策略终端': 'Professional crypto insights and trading signals',
  '语言设置': 'Language', '简体中文': 'Simplified Chinese',
  '密码': 'Password', '请输入密码': 'Enter your password', '账号密码登录成功': 'Password sign-in succeeded',
  '邮箱 / 用户名': 'Email / username', '请输入邮箱或用户名': 'Enter your email or username', '动态验证码': 'Verification code',
  '未获取验证码': 'No code requested', '重发验证码': 'Resend code', '请输入 6 位验证码': 'Enter 6-digit code',
  '2FA 验证码': '2FA code', '请输入身份验证器验证码': 'Enter your authenticator code', '保持信任登录 (30天)': 'Keep me signed in (30 days)',
  '忘记密码?': 'Forgot password?', '连接钱包登录': 'Connect wallet to sign in', '输入钱包地址': 'Enter wallet address',
  '链 ID，例如 56': 'Chain ID, e.g. 56', '返回账号登录': 'Back to account login', '正在进入...': 'Signing in...',
  '连接钱包并登录': 'Connect wallet and sign in', '进入交易终端 · 探索点位': 'Enter terminal · Explore signals',
  '快捷连接与第三方': 'Quick connect and third-party sign-in', '同意': 'I agree to the', '与': 'and',
  '《服务协议》': 'Terms of Service', '《隐私条款》': 'Privacy Policy', '服务协议': 'Terms of Service', '隐私条款': 'Privacy Policy',
  '资产与市场波动剧烈，交易前请独立研究 (DYOR)': 'Markets are volatile. Do your own research before trading (DYOR).',
  '搜索新闻、币种、项目...': 'Search news, assets, projects...', '搜索新闻、币种、项目': 'Search news, assets, projects',
  '加载中...': 'Loading...', '加载失败，请重试': 'Failed to load. Try again.', '重试': 'Retry', '暂无快讯': 'No stories',
  '市场快讯与行情动态。': 'Market news and updates.', '全部': 'All', '山寨币': 'Altcoins', '交易所': 'Exchange', '政策': 'Policy',
  '做多': 'Long', '做空': 'Short', '自选': 'Favorites', '交易点位': 'Trading signals', '实时策略': 'Live strategy',
  '已解锁': 'Unlocked', '暂无点位': 'No signals', '点位详情已锁定，点击解锁': 'Signal details locked. Tap to unlock',
  '入场区间': 'Entry range', '当前价格': 'Current price', '止盈': 'Take profit', '止损': 'Stop loss', '发布时间': 'Published',
  '有效期': 'Validity', '进行中': 'Active', '已失效': 'Expired', '利好': 'Bullish', '利空': 'Bearish', '中性': 'Neutral',
  '要闻': 'Headline', '紧急': 'Urgent', '普通': 'Normal', '分享': 'Share', '收藏': 'Save', '已收藏': 'Saved', '已取消收藏': 'Removed from saved',
  '暂无收藏内容': 'No saved items', '点位已解锁': 'Signal unlocked', '解锁失败': 'Unable to unlock', '确认解锁': 'Confirm unlock',
  '调用平台解锁接口获取全部点位': 'Unlock all signals through the platform', '选择您的联系方式': 'Choose a contact method',
  '复制联系方式': 'Copy contact', '已复制': 'Copied', '扫码加入客服': 'Scan to contact support', '找回密码': 'Reset password',
  '按步骤验证身份并设置新密码': 'Verify your identity and set a new password', '步骤': 'Step', '获取验证码': 'Get code', '验证': 'Verify',
  '完成': 'Done', '邮箱或手机号': 'Email or phone', '验证码': 'Verification code', '新密码': 'New password', '设置新密码': 'Set a new password',
  '输入账号': 'Enter account', '输入验证码': 'Enter verification code', '创建账户': 'Create account', '加入 Crypto Flash 专业交易终端': 'Join the Crypto Flash trading terminal',
  '输入邮箱或手机号': 'Enter email or phone', '发送验证码': 'Send code', '设置密码': 'Set password', '确认密码': 'Confirm password', '再次输入密码': 'Enter password again',
  '同意服务协议与隐私条款': 'I agree to the terms of service and privacy policy', '注册并登录': 'Create account and sign in',
  '请输入邮箱': 'Enter email', '验证码发送失败': 'Could not send verification code', '请完整填写注册信息': 'Complete all required fields',
  '密码至少 6 位': 'Password must be at least 6 characters', '两次密码不一致': 'Passwords do not match', '请先同意协议': 'Please accept the agreements first',
  '注册成功，请登录': 'Account created. Please sign in.', '注册失败': 'Registration failed', '登录失败': 'Login failed', '钱包登录失败': 'Wallet login failed',
  '第三方登录失败': 'Third-party sign-in failed', '授权发起失败': 'Could not start authorization', '请在浏览器中完成授权': 'Complete authorization in your browser',
  '主题已切换': 'Theme changed', '深色模式': 'Dark mode', '浅色模式': 'Light mode', '跟随系统': 'System', '我的收藏': 'Saved items', '客服联系方式': 'Support contacts',
  '消息通知设置': 'Notification settings', '主题模式': 'Theme', '风险提示': 'Risk notice', '关于我们': 'About us', '安全设置': 'Security settings',
  '修改密码 / 2FA　›': 'Password / 2FA　›', '版本信息　›': 'Version info　›', '消息列表': 'Messages', '完整行情': 'Markets', '退出登录': 'Sign out',
  '个人资料': 'Profile', '未设置资料': 'Profile not set', '未绑定邮箱': 'Email not linked', '昵称': 'Nickname', '邮箱': 'Email', '邮箱验证码': 'Email verification code',
  '头像地址': 'Avatar URL', '钱包地址': 'Wallet address', '绑定钱包': 'Link wallet', '保存资料': 'Save profile', '保存中...': 'Saving...', '输入昵称': 'Enter nickname',
  '输入邮箱': 'Enter email', '修改邮箱时填写': 'Required to change email', '输入头像 URL': 'Enter avatar URL', '请输入昵称和邮箱': 'Enter a nickname and email',
  '资料已更新': 'Profile updated', '资料更新失败': 'Could not update profile', '资料加载失败': 'Could not load profile', '钱包绑定成功': 'Wallet linked', '钱包绑定失败': 'Could not link wallet',
  '暂无消息': 'No messages', '消息': 'Messages', '全部已读': 'Mark all as read', '已全部读': 'All messages marked read', '保存失败': 'Could not save',
  '重要新闻快讯': 'Important news', '新交易点位': 'New trading signal', '点位状态更新': 'Signal status update', '系统通知': 'System notifications',
  '开启后将通过系统推送通知送达': 'Enabled notifications are delivered by system push', '实时追踪主要交易对的 24 小时变化': 'Track 24-hour changes for major trading pairs',
  '暂无行情': 'No market data', '24 小时': '24 hours', '修改密码': 'Change password', '2FA 双重验证': 'Two-factor authentication', '已开启': 'Enabled', '未开启': 'Disabled',
  '解绑': 'Unlink', '处理中': 'Working...', '绑定': 'Link', '原密码': 'Current password', '确认新密码': 'Confirm new password', '保存修改': 'Save changes',
  '请输入身份验证器中的 6 位动态码': 'Enter the 6-digit code from your authenticator', '密钥': 'Secret key', '请输入邮箱验证码': 'Enter email code', '关闭 2FA': 'Turn off 2FA', '开启 2FA': 'Turn on 2FA',
  '请检查密码输入': 'Check your password entries', '密码已修改，请重新登录': 'Password changed. Please sign in again.', '修改失败': 'Could not update', '2FA 已开启': '2FA enabled', '2FA 已关闭': '2FA disabled',
  '操作失败': 'Action failed', '接口未返回授权地址': 'The service did not return an authorization URL', '发起绑定失败': 'Could not start linking', '解绑成功': 'Unlinked', '解绑失败': 'Could not unlink',
  '说明': 'Information', '暂无正文': 'No content', '未知来源': 'Unknown source', '未命名资讯': 'Untitled story', '来源：': 'Source: ', '策略说明': 'Strategy', '暂无策略说明': 'No strategy details',
  '暂无': 'Unavailable', '已锁定': 'Locked', '搜索': 'Search', '热门搜索': 'Popular searches', '搜索失败': 'Search failed', '网络请求失败': 'Network request failed', '请求失败': 'Request failed', '接口返回失败': 'The request was rejected',
  '接口日志': 'API logs', '仅失败': 'Failures only', '清空': 'Clear', '暂无接口日志': 'No API logs', '成功': 'Success', '网络错误': 'Network error', 'HTTP 错误': 'HTTP error', '接口失败': 'API failure', '查看请求结果': 'View request results', '首页数据加载成功': 'Home data loaded', '首页数据加载失败': 'Home data failed to load', '登录成功，但首页数据加载失败': 'Signed in, but home data failed to load', '验证码登录成功': 'Verification-code sign-in succeeded' }, {
  '请输入验证码': 'Enter verification code', '请输入链 ID': 'Enter chain ID', '请输入邮箱': 'Enter email', '后可重发': 'until resend', '绑定成功': 'linked successfully', '登录响应缺少 token': 'The login response did not include a token', '密码已重置': 'Password reset',
  '资讯加载失败': 'Could not load story', '点位加载失败': 'Could not load signal', '未知币种': 'Unknown asset', '未知方向': 'Unknown direction', '未知状态': 'Unknown status', '版本信息': 'Version info', '美联储': 'Federal Reserve', '说明': 'Information', '加密资产价格波动剧烈，历史表现不代表未来收益。本文及应用内信息仅供参考，不构成任何投资建议。请在充分了解风险的前提下独立决策。': 'Crypto assets are highly volatile. Past performance does not predict future results. The information here is for reference only and is not investment advice. Make independent decisions after understanding the risks.' }, {
  '日历': 'Calendar', '数据': 'Data', '风险提示': 'Risk notice', '财经事件': 'Economic event', '实时数据': 'Live data', '热门快讯': 'Top stories', '查看全部 ›': 'View all ›', '更多 ›': 'More ›', '多方': 'Long', '空方': 'Short', '多方主导': 'Longs dominate', '重要财经日历': 'Economic calendar' }, {
  '交易点位': 'Trading signals', '实时策略': 'Live strategy', '全部解锁': 'Unlock all', '请输入链 ID，例如 56': 'Enter chain ID, e.g. 56', '后可重发': 'until resend', '客服': 'Support', '绑定成功': 'Linked successfully', '未设置资料': 'Profile not set', '未绑定邮箱': 'Email not linked', '主题已切换': 'Theme changed', '浅色模式': 'Light mode', '深色模式': 'Dark mode' }, {
  '消息通知设置': 'Notification settings', '开启后将通过系统推送通知送达': 'Enabled notifications are delivered by system push', '全部已读': 'Mark all as read', '已全部读': 'All messages marked read', '个人资料': 'Profile', '昵称': 'Nickname', '头像地址': 'Avatar URL', '修改邮箱时填写': 'Required to change email', '绑定钱包': 'Link wallet', '保存资料': 'Save profile', '保存中...': 'Saving...', '资料已更新': 'Profile updated', '资料加载失败': 'Could not load profile', '资料更新失败': 'Could not update profile', '请输入昵称和邮箱': 'Enter a nickname and email', '输入昵称': 'Enter nickname', '输入邮箱': 'Enter email', '输入头像 URL': 'Enter avatar URL' }, {
  '安全设置': 'Security settings', '修改密码': 'Change password', '2FA 双重验证': 'Two-factor authentication', 'Google 绑定': 'Google linking', 'Microsoft 绑定': 'Microsoft linking', '解绑': 'Unlink', '处理中': 'Working...', '绑定': 'Link', '原密码': 'Current password', '新密码': 'New password', '确认新密码': 'Confirm new password', '请输入身份验证器中的 6 位动态码': 'Enter the 6-digit code from your authenticator', '输入 6 位验证码': 'Enter 6-digit code', '开启 2FA': 'Turn on 2FA', '关闭 2FA': 'Turn off 2FA', '修改失败': 'Could not update', '操作失败': 'Action failed', '发起绑定失败': 'Could not start linking', '解绑成功': 'Unlinked', '解绑失败': 'Could not unlink' }, {
  '联系客服': 'Contact support', '选择您的联系方式': 'Choose a contact method', '复制联系方式': 'Copy contact', '扫码加入客服': 'Scan to contact support', '已复制': 'Copied', '我的收藏': 'Saved items', '暂无收藏内容': 'No saved items', '完整行情': 'Markets', '市场行情': 'Market', '实时追踪主要交易对的 24 小时变化': 'Track 24-hour changes for major trading pairs', '24 小时': '24 hours', '暂无行情': 'No market data', '消息': 'Messages', '暂无消息': 'No messages', '消息列表': 'Messages', '退出登录': 'Sign out', '主题模式': 'Theme', '关于我们': 'About us' }, {
  '请输入邮箱或用户名': 'Enter your email or username', '未获取验证码': 'No code requested', '重发验证码': 'Resend code', '请输入身份验证器验证码': 'Enter your authenticator code', '保持信任登录 (30天)': 'Keep me signed in (30 days)', '忘记密码?': 'Forgot password?', '连接钱包登录': 'Connect wallet to sign in', '输入钱包地址': 'Enter wallet address', '链 ID，例如 56': 'Chain ID, e.g. 56', '返回账号登录': 'Back to account login', '正在进入...': 'Signing in...', '连接钱包并登录': 'Connect wallet and sign in', '快捷连接与第三方': 'Quick connect and third-party sign-in', '同意': 'I agree to the', '与': 'and', '《服务协议》': 'Terms of Service', '《隐私条款》': 'Privacy Policy',
  '仅供参考，不构成投资建议，加密资产投资存在高风险，请 DYOR。': 'For reference only. Not financial advice. Crypto assets are high risk. DYOR.', '持有平台指定 NFT 可解锁全部历史与未来点位与策略信息': 'Hold the platform NFT to unlock all past and future signals and strategy details', '去购买 NFT': 'Buy NFT', '如何获得 NFT?': 'How do I get an NFT?', '钱包绑定成功': 'Wallet linked', '钱包绑定失败': 'Could not link wallet', '邮箱验证码': 'Email verification code', '未开启': 'Disabled', '密码已修改，请重新登录': 'Password changed. Please sign in again.'
  })

export function t(key: string) { return locale.value === 'zh' ? key : (translations[key] || key) }

export function setLocale(value: Locale) {
  locale.value = value
  uni.setStorageSync(storageKey, value)
  const labels = value === 'zh' ? ['首页', '快讯', '点位', '我的'] : ['Home', 'Flash', 'Signals', 'Profile']
  const paths = ['pages/index/index', 'pages/flash/list', 'pages/points/list', 'pages/me/index']
  paths.forEach((pagePath, index) => uni.setTabBarItem({ index, pagePath, text: labels[index] }))
}

export function restoreLocale() { setLocale(uni.getStorageSync(storageKey) === 'en' ? 'en' : 'zh') }
export function toggleLocale() { setLocale(locale.value === 'zh' ? 'en' : 'zh') }
