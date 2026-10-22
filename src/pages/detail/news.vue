<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import { onShareAppMessage } from '@dcloudio/uni-app'
import { api } from '../../api'
import { useSessionStore } from '../../store/session'
import { locale, t } from '../../locales'
import '../../styles/app.scss'

const store = useSessionStore()
const item = ref<any>()
const favorite = ref(false)
const loading = ref(true)
const error = ref('')

async function load() {
  loading.value = true
  error.value = ''
  const page: any = getCurrentPages().at(-1)
  const id = String(page?.options?.id || '')
  if (!store.token) {
    uni.setStorageSync('cf_after_login', `/pages/detail/news?id=${encodeURIComponent(id)}`)
    uni.reLaunch({ url: '/pages/index/index?auth=login' })
    return
  }
  try {
    item.value = await api.newsById(id)
  } catch (e: any) {
    error.value = e?.message || t('资讯加载失败')
  } finally {
    loading.value = false
  }
  if (item.value) {
    try { favorite.value = await api.isFavorite(item.value.id, 'news') } catch {}
  }
}

onMounted(load)
watch(locale, load)

async function toggle() {
  if (!item.value) return
  if (!store.token) return uni.reLaunch({ url: '/pages/index/index?auth=login' })
  favorite.value = await api.toggleFavorite(item.value.id, 'news', !favorite.value)
  uni.showToast({ title: favorite.value ? t('已收藏') : t('已取消收藏'), icon: 'none' })
}

async function share() {
  if (item.value) await api.shareNews(item.value.id).catch(() => {})
  uni.showShareMenu({})
}

onShareAppMessage(() => {
  if (item.value) api.shareNews(item.value.id).catch(() => {})
  return { title: item.value?.title || 'Crypto Flash', path: `/pages/detail/news?id=${item.value?.id || ''}` }
})
</script>
<template>
  <view class="news-detail">
    <view class="detail-nav"><text @click="uni.navigateBack()">‹</text><view><text @click="share">↗</text><text @click="toggle">{{favorite ? '★' : '♧'}}</text></view></view>
    <view v-if="loading" class="detail-state">{{t('加载中...')}}</view>
    <view v-else-if="error" class="detail-state">{{error}}</view>
    <view v-else-if="item" class="detail-content">
      <view class="detail-tags"><text v-if="item.level && item.level!=='普通'" class="urgent-tag">{{t(item.level)}}</text><text v-for="tag in item.tags.split(' ').filter(Boolean)" :key="tag">{{tag}}</text><text v-if="item.sentiment && item.sentiment!=='中性'" class="sentiment" :class="item.sentiment==='利好'?'up':'down'">{{t(item.sentiment)}}</text></view>
      <view class="detail-title">{{item.title || t('未命名资讯')}}</view>
      <view class="detail-meta"><text>{{t('来源：')}}{{item.source || t('未知来源')}}</text><text>{{item.time || ''}}</text></view>
      <view class="detail-body">{{item.body || item.content || item.summary || t('暂无正文')}}</view>
    </view>
    <view class="detail-actions"><button @click="share">↗ {{t('分享')}}</button><button @click="toggle">{{favorite ? `★ ${t('已收藏')}` : `♧ ${t('收藏')}`}}</button></view>
  </view>
</template>
<style scoped>
.news-detail{min-height:100vh;background:#fff;color:var(--cf-text)}.detail-nav{display:flex;justify-content:space-between;align-items:center;height:80rpx;padding:0 28rpx;border-bottom:1rpx solid var(--cf-line);font-size:42rpx}.detail-nav view{display:flex;gap:32rpx;font-size:32rpx;color:var(--cf-muted)}.detail-content{padding:28rpx 32rpx 150rpx}.detail-state{padding:100rpx 32rpx;color:var(--cf-muted);text-align:center}.detail-tags{display:flex;align-items:center;gap:8rpx;color:var(--cf-brand);font-size:20rpx}.detail-tags text{padding:3rpx 8rpx;background:#EFF6FF;border-radius:5rpx}.detail-tags .urgent-tag{background:#FEF2F2;color:var(--cf-down)}.detail-tags .sentiment{color:var(--cf-down)}.detail-tags .sentiment.up{color:var(--cf-up);background:#ECFDF5}.detail-title{margin:20rpx 0 10rpx;font-size:35rpx;line-height:1.4;font-weight:800}.detail-meta{display:flex;justify-content:space-between;color:var(--cf-soft);font-size:22rpx}.detail-body{margin:48rpx 0;font-size:29rpx;line-height:1.8;color:var(--cf-muted);white-space:pre-line}.detail-actions{position:fixed;bottom:0;left:0;right:0;display:grid;grid-template-columns:1fr 1fr;gap:12rpx;padding:16rpx 24rpx calc(16rpx + env(safe-area-inset-bottom));border-top:1rpx solid var(--cf-line);background:#fff}.detail-actions button{height:72rpx;margin:0;border-radius:14rpx;background:#F1F5F9;color:var(--cf-muted);font-size:25rpx}
</style>
<style scoped>.detail-nav{padding-top:var(--cf-safe-top);height:calc(80rpx + var(--cf-safe-top))}</style>
