<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { onShow } from '@dcloudio/uni-app'
import { api } from '../../api'
import { useSessionStore } from '../../store/session'
import { t, locale, toggleLocale } from '../../locales'
import { recordAppLog } from '../../api/diagnostics'
import googleLogo from '../../img/gogeimages.png'
import microsoftLogo from '../../img/Microsoft_Logo_512px.png'
import walletLogo from '../../img/web3.png'
import { oauthRedirectUrl, openOAuthUrl } from '../../native/oauth'
import '../../styles/app.scss'
const store = useSessionStore()
const tab = ref<'login'|'wallet'>('login')
const email = ref('')
const password = ref('')
const twoFactorCode = ref('')
const walletAddress = ref('')
const chainId = ref('56')
const requires2fa = ref(false)
const accepted = ref(false)
const trusted = ref(false)
const loading = ref(false)
const home = ref<any>(null)
const news = ref<any[]>([])
const homeTab = ref<'featured'|'calendar'|'data'>('featured')
const showLogin = ref(true)
async function loadHome() {
  try {
    home.value = await api.home()
    news.value = home.value.news || await api.news()
  } catch (error: any) {
    recordAppLog('home.load', 'failure', error?.message || t('首页数据加载失败'))
  }
}
watch(locale, () => { if (!showLogin.value) loadHome() })
onMounted(async () => {
  const current = getCurrentPages().at(-1) as any
  const options = current?.options || {}
  const query = typeof location !== 'undefined' ? new URLSearchParams(location.search) : null
  const authRoute = String(options.auth || query?.get('auth') || '')
  showLogin.value = authRoute === 'guest' ? false : authRoute === 'login' || !store.token
  const provider = String(options.oauthProvider || query?.get('oauthProvider') || '') as 'google' | 'microsoft'
  const oauthAction = String(options.oauthAction || query?.get('oauthAction') || '')
  const oauthCode = String(options.code || query?.get('code') || '')
  const oauthState = String(options.state || query?.get('state') || '')
  if ((provider === 'google' || provider === 'microsoft') && oauthCode && oauthState) {
    try {
      const identity: any = await api.oauthCallback(provider, oauthCode, oauthState)
      if (oauthAction === 'binding' && store.token) {
        if (provider === 'google') await api.bindGoogle(identity)
        else await api.bindMicrosoft(identity)
        uni.showToast({ title: `${provider === 'google' ? 'Google' : 'Microsoft'} ${t('绑定成功')}`, icon: 'none' })
        return
      }
      const result: any = provider === 'google'
        ? await api.loginGoogle({ ...identity, isRemember: trusted.value })
        : await api.loginMicrosoft({ ...identity, isRemember: trusted.value })
      if (!result?.token) throw new Error(t('登录响应缺少 token'))
      store.login(result.token)
      uni.setStorageSync('cf_remember', trusted.value)
      showLogin.value = false
      uni.showTabBar(); await loadHome()
      continueAfterLogin()
    } catch (error: any) { uni.showToast({ title: error?.message || t('第三方登录失败'), icon: 'none' }) }
  } else if (!showLogin.value) { await loadHome() }
})
onShow(() => { if (showLogin.value) uni.hideTabBar(); else uni.showTabBar() })
function continueAfterLogin() {
  const target = uni.getStorageSync('cf_after_login')
  if (target === 'points') {
    uni.removeStorageSync('cf_after_login')
    uni.switchTab({ url: '/pages/points/list' })
  } else if (typeof target === 'string' && target.startsWith('/pages/detail/news?id=')) {
    uni.removeStorageSync('cf_after_login')
    uni.navigateTo({ url: target })
  }
}
function skipLogin() { uni.removeStorageSync('cf_after_login'); showLogin.value = false; uni.showTabBar(); if (!home.value) loadHome() }
async function login() {
  if (tab.value === 'wallet') return loginWallet()
  if (!email.value.trim()) return uni.showToast({ title: t('请输入邮箱或用户名'), icon: 'none' })
  if (tab.value === 'login' && !password.value) return uni.showToast({ title: t('请输入密码'), icon: 'none' })
  if (requires2fa.value && !twoFactorCode.value.trim()) return uni.showToast({ title: t('请输入 2FA 验证码'), icon: 'none' })
  if (!accepted.value) return uni.showToast({ title: t('请先同意协议'), icon: 'none' })
  if (!requires2fa.value) {
    try {
      const status: any = await api.query2fa(email.value.trim())
      const enabled = status === true || status === 1 || status === '1' || Boolean(status?.enabled ?? status?.isEnabled ?? status?.twoFactorEnabled ?? status?.status)
      if (enabled) {
        requires2fa.value = true
        return uni.showToast({ title: t('请输入 2FA 验证码'), icon: 'none' })
      }
    } catch {}
  }
  loading.value = true
  let loginSucceeded = false
  try {
    const result: any = await api.loginByPassword(email.value.trim(), password.value, twoFactorCode.value.trim() || undefined)
    if (!result?.token) throw new Error(t('登录响应缺少 token'))
    store.login(result.token)
    uni.setStorageSync('cf_remember', trusted.value)
    loginSucceeded = true
    recordAppLog('login.password', 'success', t('账号密码登录成功'))
    requires2fa.value = false
  } catch (error: any) {
    recordAppLog('login.password', 'failure', error?.message || t('登录失败'))
    uni.showToast({ title: error?.message || t('登录失败'), icon: 'none' })
  } finally { loading.value = false }
  if (loginSucceeded) {
    uni.showTabBar()
    showLogin.value = false
    try {
      await loadHome()
      recordAppLog('login.home-load', 'success', t('首页数据加载成功'))
    } catch (error: any) {
      recordAppLog('login.home-load', 'failure', error?.message || t('首页数据加载失败'))
      uni.showToast({ title: t('登录成功，但首页数据加载失败'), icon: 'none' })
    }
    continueAfterLogin()
  }
}
async function loginWallet() {
  if (!walletAddress.value.trim()) return uni.showToast({ title: t('请输入钱包地址'), icon: 'none' })
  if (!chainId.value.trim()) return uni.showToast({ title: t('请输入链 ID'), icon: 'none' })
  if (!accepted.value) return uni.showToast({ title: t('请先同意协议'), icon: 'none' })
  loading.value = true
  try {
    const result: any = await api.loginWallet(walletAddress.value.trim(), chainId.value.trim(), trusted.value)
    if (!result?.token) throw new Error(t('登录响应缺少 token'))
    store.login(result.token)
    uni.setStorageSync('cf_remember', trusted.value)
    showLogin.value = false
    uni.showTabBar(); await loadHome()
    continueAfterLogin()
  } catch (error: any) { uni.showToast({ title: error?.message || t('钱包登录失败'), icon: 'none' }) } finally { loading.value = false }
}
async function openOAuth(provider: 'google' | 'microsoft', action = '') {
  try {
    const redirectUrl = oauthRedirectUrl(provider, action)
    const result: any = await api.oauthAuth(provider, redirectUrl)
    const authUrl = result?.authUrl
    if (!authUrl) throw new Error(t('接口未返回授权地址'))
    await openOAuthUrl(authUrl)
  } catch (error: any) { uni.showToast({ title: error?.message || t('授权发起失败'), icon: 'none' }) }
}
</script>
<template>
  <view class="page" :class="{ 'login-route': showLogin }">
    <view v-if="showLogin" class="login-page">
      <view class="login-top"><view class="network-pill"><text class="status-dot">●</text><text>Mainnet</text></view><view class="login-tools"><text class="locale-switch" @click="skipLogin">{{t('跳过')}}</text><text class="locale-switch" @click="uni.navigateTo({url:'/pages/me/diagnostics'})">{{t('接口日志')}}</text><text class="locale-switch" @click="toggleLocale">{{locale==='zh'?'中 / EN':'EN / 中'}}</text></view></view>
      <view class="login-hero"><view class="login-logo"><svg viewBox="0 0 24 24" fill="currentColor" aria-hidden="true"><path d="M13 2L3 14h9l-1 8 10-12h-9l1-8z"/></svg></view><view class="login-name">Crypto Flash <text class="pro-badge">PRO</text></view><text class="login-subtitle">{{t('专业级加密资讯与高胜率策略终端')}}</text></view>
      <view class="login-panel">
        <view class="login-tabs"><text :class="{active:tab==='login'}" @click="tab='login'">◎　{{t('账号登录')}}</text><text :class="{active:tab==='wallet'}" @click="tab='wallet'">◇　{{t('钱包连接')}}</text></view>
        <view v-if="tab==='login'" class="login-form">
          <view class="field"><text class="field-label">{{t('邮箱 / 用户名')}}</text><view class="field-control"><input v-model="email" :placeholder="t('请输入邮箱或用户名')"/><text class="field-icon">✉</text></view></view>
          <view class="field"><text class="field-label">{{t('密码')}}</text><view class="field-control"><input v-model="password" password :placeholder="t('请输入密码')"/></view></view>
          <view v-if="requires2fa" class="field"><text class="field-label">{{t('2FA 验证码')}}</text><view class="field-control"><input v-model="twoFactorCode" :placeholder="t('请输入身份验证器验证码')" maxlength="6"/></view></view>
          <view class="login-options"><label><checkbox :checked="trusted" @click="trusted=!trusted"/><text>{{t('保持信任登录 (30天)')}}</text></label><view class="account-links"><text @click="uni.navigateTo({url:'/pages/auth/register'})">{{t('注册')}}</text><text class="forgot-link" @click="uni.navigateTo({url:'/pages/auth/forgot'})">{{t('忘记密码?')}}</text></view></view>
        </view>
        <view v-else class="wallet-panel"><image class="wallet-icon" :src="walletLogo" mode="aspectFit"/><text class="wallet-title">{{t('连接钱包登录')}}</text><input v-model="walletAddress" class="wallet-address" :placeholder="t('输入钱包地址')"/><input v-model="chainId" class="wallet-address" :placeholder="t('链 ID，例如 56')"/><button class="wallet-button" @click="tab='login'">{{t('返回账号登录')}}</button></view>
        <button class="login-button" :disabled="loading" @click="login">{{loading ? t('正在进入...') : tab === 'wallet' ? t('连接钱包并登录') : t('进入交易终端 · 探索点位')}} <text>→</text></button>
        <view class="third-party"><view class="divider"><i></i><text>{{t('快捷连接与第三方')}}</text><i></i></view><view class="third-actions"><view><button class="third-button microsoft" @click="openOAuth('microsoft')"><image :src="microsoftLogo" mode="aspectFit"/></button><text>Microsoft</text></view><view><button class="third-button google" @click="openOAuth('google')"><image :src="googleLogo" mode="aspectFit"/></button><text>Google</text></view></view></view>
        <view class="agreement"><checkbox :checked="accepted" @click="accepted=!accepted"/><text>{{t('同意')}}</text><text class="agreement-link" @click="uni.navigateTo({url:'/pages/content/static?type=service'})">{{t('《服务协议》')}}</text><text>{{t('与')}}</text><text class="agreement-link" @click="uni.navigateTo({url:'/pages/content/static?type=privacy'})">{{t('《隐私条款》')}}</text></view>
      </view>
      <text class="login-risk">{{t('资产与市场波动剧烈，交易前请独立研究 (DYOR)')}}</text>
    </view>
    <view v-else class="home-content">
      <view class="row"><view class="brand"><view class="brand-mark"><svg viewBox="0 0 24 24" fill="currentColor" aria-hidden="true"><path d="M13 2L3 14h9l-1 8 10-12h-9l1-8z"/></svg></view><text>Crypto Flash</text></view><image class="bell-icon" src="/static/bell-notice.svg" mode="aspectFit" aria-label="通知" @click="uni.navigateTo({url:'/pages/me/messages'})"/></view>
      <view class="input search" @click="uni.navigateTo({url:'/pages/search/index'})">⌕ {{t('搜索新闻、币种、项目...')}}</view>
      <view class="tabs"><text :class="['tab',{active:homeTab==='featured'}]" @click="homeTab='featured'">{{t('精选')}}</text><text :class="['tab',{active:homeTab==='calendar'}]" @click="homeTab='calendar'">{{t('日历')}}</text><text :class="['tab',{active:homeTab==='data'}]" @click="homeTab='data'">{{t('数据')}}</text><text class="tab" @click="uni.navigateTo({url:'/pages/me/markets'})">{{t('更多')}}</text></view>
      <view v-if="homeTab === 'featured' || homeTab === 'data'" class="card"><view class="row"><text class="title card-title">{{t('多空比 (BTC)')}}</text><text class="muted home-action-link">{{t('更多 ›')}}</text></view><view v-if="home" class="ratio"><view><text class="up">{{t('多方')}}</text><text class="title up">{{home.ratio.long}}%</text></view><view class="gauge"><text class="title">{{home.ratio.value}}</text><text class="muted">{{t('多方主导')}}</text></view><view><text class="down">{{t('空方')}}</text><text class="title down">{{home.ratio.short}}%</text></view></view><view v-if="home" class="bar"><view :style="{width:home.ratio.long+'%'}"></view></view></view>
      <view v-if="(homeTab === 'featured' || homeTab === 'calendar') && home?.calendar?.length" class="calendar-card"><view class="row"><text class="title card-title">{{t('重要财经日历')}}</text><text class="muted home-action-link">{{t('实时数据')}}</text></view><view v-for="(event,index) in home.calendar.slice(0,3)" :key="event.id || index" class="calendar-row"><view class="calendar-date">{{event.date || event.time || event.eventTime || '--'}}</view><view class="calendar-copy"><text class="calendar-title">{{event.title || event.name || event.event || t('财经事件')}}</text><text class="muted">{{event.country || event.region || event.currency || ''}}</text></view><text class="calendar-impact">{{event.impact ? t(event.impact) : (event.level || '')}}</text></view></view>
      <view class="row section-head"><text class="title card-title">{{t('热门快讯')}}</text><text class="muted home-action-link" @click="uni.switchTab({url:'/pages/flash/list'})">{{t('查看全部 ›')}}</text></view>
      <view class="home-news-list"><view v-for="(item,index) in news.slice(0,3)" :key="item.id" class="home-news-card" @click="uni.navigateTo({url:'/pages/detail/news?id='+item.id})"><image class="news-thumb" :src="`/static/home-news-${index+1}.jpg`" mode="aspectFill"/><view class="home-news-copy"><view class="title news-title">{{item.title}}</view><view class="muted">{{item.time}} · {{item.source}}</view></view></view></view>
      <view class="row section-head"><text class="title card-title">{{t('市场行情')}}</text><text class="muted home-action-link" @click="uni.navigateTo({url:'/pages/me/markets'})">{{t('更多 ›')}}</text></view><view class="market-grid"><view v-for="coin in home?.coins" :key="coin[0]" class="market-tile"><text class="muted">{{coin[0]}}</text><text class="title price">{{coin[1]}}</text><text :class="coin[2].startsWith('+')?'up':'down'">{{coin[2]}}</text></view></view>
    </view>
  </view>
</template>
<style scoped>.login-brand{justify-content:center;margin-top:90rpx}.center{text-align:center}.pro{font-size:20rpx;color:var(--cf-brand)}.agreement{margin:24rpx 0;font-size:22rpx;color:var(--cf-muted)}.auth-links{display:flex;justify-content:space-between;margin-top:16rpx}.wallet{text-align:center;padding:36rpx 0}.wallet-icon{font-size:52rpx;color:var(--cf-brand)}.small-title{font-size:28rpx;margin-top:10rpx}.bell{font-size:42rpx;color:var(--cf-brand)}.search{margin-top:28rpx}.card-title{font-size:28rpx}.section-head{margin:28rpx 0 14rpx}.news-title{font-size:27rpx;margin-bottom:10rpx}.market-card{padding:22rpx 26rpx;margin-bottom:14rpx}.price{font-size:26rpx}.input{height:78rpx;border:1rpx solid var(--cf-line);border-radius:16rpx;padding:0 22rpx;color:var(--cf-text)}.ratio{display:flex;justify-content:space-between;align-items:end;text-align:center;padding:45rpx 10rpx 24rpx}.ratio view{display:flex;flex-direction:column;gap:8rpx}.gauge{width:200rpx;height:100rpx;border:14rpx solid var(--cf-line);border-bottom:0;border-radius:200rpx 200rpx 0 0;justify-content:end!important}.bar{height:14rpx;background:var(--cf-down);border-radius:8rpx;overflow:hidden}.bar view{height:100%;background:var(--cf-up)}</style>
<style scoped>.calendar-card{margin-top:18rpx;padding:20rpx;border:1rpx solid var(--cf-line);border-radius:16rpx;background:var(--cf-card)}.calendar-row{display:flex;align-items:center;gap:14rpx;padding:16rpx 0;border-top:1rpx solid #F1F5F9}.calendar-date{flex:0 0 112rpx;color:var(--cf-soft);font-size:19rpx}.calendar-copy{display:flex;flex:1;flex-direction:column;gap:4rpx;min-width:0}.calendar-title{overflow:hidden;text-overflow:ellipsis;white-space:nowrap;font-size:22rpx;font-weight:700}.calendar-copy .muted{font-size:18rpx}.calendar-impact{color:var(--cf-brand);font-size:18rpx}</style>
<style scoped>
.login-page{min-height:calc(100vh - 64rpx);padding:8rpx 20rpx 24rpx;background:var(--cf-bg);color:var(--cf-text)}
.login-top{display:flex;justify-content:space-between;align-items:center;height:52rpx}.login-tools{display:flex;align-items:center;gap:20rpx}.network-pill{display:flex;align-items:center;gap:8rpx;padding:6rpx 14rpx;border-radius:24rpx;background:var(--cf-surface);color:var(--cf-muted);font-size:20rpx}.status-dot{color:var(--cf-up);font-size:15rpx}.locale-switch{color:var(--cf-muted);font-size:20rpx}
.login-hero{text-align:center;padding:44rpx 0 30rpx}.login-logo{width:104rpx;height:104rpx;margin:0 auto 18rpx;border-radius:32rpx;background:linear-gradient(to top right,#2563EB,#6366F1);color:#fff;display:flex;align-items:center;justify-content:center;box-shadow:0 12rpx 28rpx rgba(37,99,235,.2)}.login-logo svg{display:block;width:56rpx;height:56rpx}.login-name{font-size:34rpx;line-height:1.2;font-weight:800;color:var(--cf-text)}.pro-badge{display:inline-block;margin-left:8rpx;padding:3rpx 8rpx;border-radius:6rpx;background:var(--cf-surface);color:var(--cf-brand);font-size:18rpx;vertical-align:middle}.login-subtitle{display:block;margin-top:10rpx;color:var(--cf-muted);font-size:22rpx}
.login-panel{padding:24rpx 24rpx 22rpx;border:1rpx solid var(--cf-line);border-radius:24rpx;background:var(--cf-card);box-shadow:0 8rpx 24rpx rgba(15,23,42,.04)}.login-tabs{display:grid;grid-template-columns:1fr 1fr;padding:4rpx;border-radius:14rpx;background:var(--cf-surface);color:var(--cf-muted);font-size:22rpx;text-align:center}.login-tabs text{padding:16rpx 8rpx;border-radius:10rpx}.login-tabs .active{background:var(--cf-card);color:var(--cf-brand);font-weight:700;box-shadow:0 2rpx 8rpx rgba(15,23,42,.08)}.login-form{padding-top:26rpx}.field{margin-bottom:20rpx}.field-label-row{display:flex;justify-content:space-between;align-items:center;margin-bottom:8rpx}.field-label{display:block;color:var(--cf-muted);font-size:20rpx;font-weight:600}.resend-hint{color:var(--cf-soft);font-size:18rpx}.field-control{display:flex;align-items:center;gap:10rpx;height:76rpx;padding:0 18rpx;border:1rpx solid var(--cf-line);border-radius:14rpx;background:var(--cf-input)}.field-control:focus-within{border-color:var(--cf-brand);background:var(--cf-card)}.field-icon{width:24rpx;color:var(--cf-soft);font-size:22rpx;font-weight:700;text-align:center}.field-control input{flex:1;height:100%;font-size:22rpx;color:var(--cf-text)}.code-row{display:flex;gap:12rpx}.code-row .field-control{flex:1}.resend-button{width:176rpx;height:76rpx;padding:0;border:1rpx solid var(--cf-line);border-radius:14rpx;background:var(--cf-surface);color:var(--cf-brand);font-size:19rpx}.resend-button[disabled]{color:var(--cf-soft);background:var(--cf-input);border-color:var(--cf-line)}.login-options{display:flex;justify-content:space-between;align-items:center;margin-top:2rpx;color:var(--cf-muted);font-size:19rpx}.login-options label{display:flex;align-items:center;gap:6rpx}.login-options checkbox,.agreement checkbox{transform:scale(.72);margin:0}.forgot-link{color:var(--cf-brand)}
.login-button{display:flex;align-items:center;justify-content:center;gap:12rpx;width:100%;height:82rpx;margin-top:26rpx;border-radius:14rpx;background:var(--cf-brand);color:#fff;font-size:23rpx;font-weight:700;box-shadow:0 8rpx 18rpx rgba(37,99,235,.2)}.login-button[disabled]{opacity:.65}.login-button text{font-size:32rpx}.third-party{margin-top:28rpx}.divider{display:flex;align-items:center;gap:12rpx;color:var(--cf-soft);font-size:18rpx}.divider i{height:1rpx;flex:1;background:var(--cf-line)}.third-actions{display:flex;justify-content:center;gap:54rpx;margin-top:20rpx}.third-actions view{display:flex;flex-direction:column;align-items:center;gap:8rpx;color:var(--cf-muted);font-size:17rpx}.third-button{display:flex;align-items:center;justify-content:center;width:68rpx;height:68rpx;padding:0;border:1rpx solid var(--cf-line);border-radius:50%;background:var(--cf-card)}.third-button image{display:block;width:34rpx;height:34rpx;object-fit:contain}.microsoft{color:var(--cf-brand)}.google{color:#4285F4}.agreement{display:flex;align-items:center;justify-content:center;gap:4rpx;margin:26rpx 0 0;color:var(--cf-muted);font-size:18rpx}.agreement-link{color:var(--cf-brand)}.login-risk{display:block;margin-top:24rpx;color:var(--cf-soft);font-size:18rpx;text-align:center}.wallet-panel{display:flex;align-items:center;flex-direction:column;padding:54rpx 0 36rpx}.wallet-panel .wallet-icon{display:block;width:70rpx;height:70rpx;object-fit:contain}.wallet-title{margin-top:14rpx;color:var(--cf-text);font-size:25rpx;font-weight:700}.wallet-copy{margin-top:8rpx;color:var(--cf-muted);font-size:20rpx}.wallet-button{width:80%;height:72rpx;margin-top:28rpx;border-radius:14rpx;background:var(--cf-surface);color:var(--cf-brand);font-size:22rpx}
.page.login-route{box-sizing:border-box;width:100%;min-height:100vh;padding:0;background:var(--cf-bg);overflow-x:hidden}
.login-page{box-sizing:border-box;width:100%;min-height:100vh;padding:calc(8rpx + var(--status-bar-height, env(safe-area-inset-top, 0px))) 40rpx 24rpx;background:var(--cf-bg)}
.login-hero{padding:38rpx 0 20rpx}
.login-panel{box-sizing:border-box;width:100%;min-width:0;padding:0;border:0;border-radius:0;background:transparent;box-shadow:none}
.login-tabs{grid-template-columns:minmax(0,1fr) minmax(0,1fr);border-radius:10rpx}
.login-tabs text{min-width:0;white-space:nowrap}
.field-control{box-sizing:border-box;min-width:0;padding:0 18rpx}
.field-control input{min-width:0}
.code-row .resend-button{flex:0 0 176rpx}
.field-icon{flex:0 0 24rpx}
.login-risk{margin-top:18rpx}
.home-content{padding-bottom:20rpx}
.home-content .brand{font-size:32rpx;color:var(--cf-text)}
.home-content .brand-mark{width:42rpx;height:42rpx;border-radius:12rpx;font-size:0}
.home-content .brand-mark svg{display:block;width:27rpx;height:27rpx}
.home-content .search{height:64rpx;display:flex;align-items:center;font-size:22rpx;color:var(--cf-soft)}
.home-content .tabs{margin:20rpx 0;gap:12rpx}
.home-content .tab{padding:8rpx 18rpx;background:var(--cf-card);border:1rpx solid var(--cf-line);font-size:22rpx}
.home-content .tab.active{background:var(--cf-brand);border-color:var(--cf-brand)}
.home-content>.card{padding:20rpx;border-radius:16rpx;margin-bottom:18rpx}
.home-content .ratio{padding:28rpx 0 16rpx}
.home-content .gauge{border-color:var(--cf-up)}
.home-content .section-head{margin:20rpx 0 12rpx}
.home-news-list{display:flex;flex-direction:column;gap:10rpx}
.home-news-card{display:flex;align-items:center;gap:14rpx;min-height:112rpx;padding:10rpx;background:var(--cf-card);border:1rpx solid var(--cf-line);border-radius:12rpx}
.news-thumb{flex:0 0 80rpx;width:80rpx;height:80rpx;border-radius:10rpx}
.home-news-copy{min-width:0;flex:1}
.home-news-copy .news-title{overflow:hidden;text-overflow:ellipsis;white-space:nowrap;font-size:23rpx;margin-bottom:4rpx}
.home-news-copy .muted{font-size:18rpx}
.market-grid{display:grid;grid-template-columns:repeat(4,minmax(0,1fr));gap:8rpx}
.market-tile{display:flex;align-items:center;flex-direction:column;justify-content:center;min-width:0;min-height:92rpx;padding:8rpx 2rpx;border:1rpx solid var(--cf-line);border-radius:12rpx;background:var(--cf-card);white-space:nowrap}
.market-tile .muted,.market-tile .up,.market-tile .down{font-size:17rpx}
.market-tile .price{font-size:19rpx;margin:3rpx 0}
</style>
<style scoped>
.login-page{padding-top:calc(16px + var(--cf-safe-top))}
.login-top{min-height:52rpx}
</style>
<style scoped>.account-links{display:flex;align-items:center;gap:18rpx;white-space:nowrap}.account-links text{color:var(--cf-brand)}</style>
<style scoped>.wallet-address{box-sizing:border-box;width:80%;height:72rpx;margin-top:18rpx;padding:0 18rpx;border:1rpx solid var(--cf-line);border-radius:14rpx;background:var(--cf-input);font-size:22rpx}</style>
<style scoped>
.login-page{display:flex;flex-direction:column;justify-content:space-between}.bell-icon{display:block;width:48rpx;height:48rpx;flex:0 0 48rpx}.login-subtitle{font-size:25rpx}.field-label{font-size:23rpx}.resend-hint{font-size:21rpx}.field-control{height:84rpx}.field-control input{font-size:26rpx}.resend-button{display:flex;align-items:center;justify-content:center;box-sizing:border-box;height:84rpx;padding:0;font-size:22rpx;line-height:1;text-align:center}.login-options{font-size:22rpx}.login-button{height:84rpx;font-size:26rpx}.divider{font-size:21rpx}.third-actions view{font-size:20rpx}.agreement{font-size:21rpx}.login-risk{font-size:21rpx}.home-content .search{height:72rpx;font-size:25rpx}.home-content .tab{font-size:25rpx}.home-content .card-title{font-size:31rpx}.home-action-link{color:var(--cf-brand)}.home-news-copy .news-title{font-size:26rpx}.home-news-copy .muted{font-size:21rpx}.market-tile .muted,.market-tile .up,.market-tile .down{font-size:21rpx}.market-tile .price{font-size:23rpx}
</style>
