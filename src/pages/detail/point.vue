<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { onShareAppMessage } from '@dcloudio/uni-app'
import { api } from '../../api'
import UnlockDialog from '../../components/UnlockDialog.vue'
import { useSessionStore } from '../../store/session'
import { locale, t } from '../../locales'
import '../../styles/app.scss'

const item = ref<any>()
const favorite = ref(false)
const loading = ref(true)
const error = ref('')
const showUnlock = ref(false)
const store = useSessionStore()
const targets = computed(() => Array.isArray(item.value?.tpLevels) ? item.value.tpLevels : String(item.value?.tp || '').split(' / ').filter(Boolean))

async function load() {
  loading.value = true
  error.value = ''
  const page: any = getCurrentPages().at(-1)
  try {
    item.value = await api.pointById(page?.options?.id || '')
  } catch (e: any) {
    error.value = e?.message || t('点位加载失败')
  } finally {
    loading.value = false
  }
  if (item.value) {
    try { favorite.value = await api.isFavorite(item.value.id, 'points') } catch {}
  }
}

onMounted(load)
watch(locale, load)

async function toggle() {
  if (!item.value) return
  if (!store.token) return uni.reLaunch({ url: '/pages/index/index?auth=login' })
  try {
    favorite.value = await api.toggleFavorite(item.value.id, 'points', !favorite.value)
    uni.showToast({ title: favorite.value ? t('已收藏') : t('已取消收藏'), icon: 'none' })
  } catch (e: any) {
    uni.showToast({ title: e?.message || t('操作失败'), icon: 'none' })
  }
}

function closeUnlock() { showUnlock.value = false; uni.navigateBack() }
function unlocked() { store.unlocked = true; uni.setStorageSync('cf_unlocked', true); showUnlock.value = false; uni.showToast({ title: t('点位已解锁'), icon: 'none' }) }

onShareAppMessage(() => {
  if (item.value) api.sharePoint(item.value.id).catch(() => {})
  return { title: item.value?.coin || 'Crypto Flash', path: `/pages/detail/point?id=${item.value?.id || ''}` }
})
</script>
<template>
  <view class="point-detail">
    <view class="point-nav"><text class="nav-back" @click="uni.navigateBack()"><svg viewBox="0 0 24 24" aria-hidden="true"><path d="m15 18-6-6 6-6"/></svg></text><view><text class="nav-share" @click="uni.showShareMenu({})"><svg viewBox="0 0 24 24" aria-hidden="true"><circle cx="18" cy="5" r="3"/><circle cx="6" cy="12" r="3"/><circle cx="18" cy="19" r="3"/><path d="m8.7 10.7 6.6-4.4M8.7 13.3l6.6 4.4"/></svg></text><text class="nav-bookmark" :class="{'is-favorite':favorite}" @click="toggle"><svg viewBox="0 0 24 24" aria-hidden="true"><path d="M6 4h12v17l-6-4-6 4V4Z"/></svg></text></view></view>
    <view v-if="loading" class="point-state">{{t('加载中...')}}</view>
    <view v-else-if="error" class="point-state">{{error}}</view>
    <view v-else-if="item" class="point-body">
      <view class="point-heading"><view><text class="coin-icon">{{item.coin?.startsWith('BTC')?'₿':item.coin?.startsWith('ETH')?'◆':'◎'}}</text><text class="coin-name">{{item.coin || t('未知币种')}}</text><text class="direction" :class="item.dir==='做多'?'up':'down'">{{t(item.dir || '未知方向')}}</text></view><text class="status">{{t(item.status || '未知状态')}}</text></view>
      <view class="price-panel"><view class="price-row"><text>{{t('入场区间')}}</text><text>{{!item.locked||store.unlocked?(item.range || t('暂无')):t('已锁定')}}</text></view><view class="price-row"><text>{{t('当前价格')}}</text><text class="up">{{item.currentPrice || t('暂无')}}</text></view><view class="price-row"><text>{{t('止盈')}}</text><view class="targets"><text v-for="(target,i) in targets" :key="i">TP{{i+1}} {{!item.locked||store.unlocked?target:t('已锁定')}} <text v-if="item.tpRatios?.[i]" class="ratio">({{item.tpRatios[i]}})</text></text><text v-if="!targets.length">{{t('暂无')}}</text></view></view><view class="price-row"><text>{{t('止损')}}</text><text class="down">{{!item.locked||store.unlocked?(item.sl || t('暂无')):t('已锁定')}}</text></view><view class="price-row fine"><text>{{t('发布时间')}}</text><text>{{item.published || t('暂无')}}</text></view><view class="price-row fine"><text>{{t('有效期')}}</text><text>{{item.validity || t('暂无')}}</text></view></view><view class="strategy"><view>{{t('策略说明')}}</view><text>{{item.strategy || t('暂无策略说明')}}</text></view>
    </view>
    <view class="point-actions"><view><button @click="showUnlock=true">{{t('解锁全部点位')}}</button></view><text>{{t('仅供参考，不构成投资建议，加密资产投资存在高风险，请 DYOR。')}}</text></view>
    <UnlockDialog :visible="showUnlock" @close="closeUnlock" @unlocked="unlocked"/>
  </view>
</template>
<style scoped>
.point-detail{min-height:100vh;background:#fff;color:var(--cf-text)}.point-nav{display:flex;justify-content:space-between;align-items:center;height:80rpx;padding:0 28rpx;border-bottom:1rpx solid var(--cf-line);font-size:42rpx}.point-nav view{display:flex;gap:32rpx;font-size:32rpx;color:var(--cf-muted)}.point-state{padding:100rpx 32rpx;color:var(--cf-muted);text-align:center}.point-body{padding:28rpx 32rpx 150rpx}.point-heading,.point-heading>view,.price-row{display:flex;align-items:center;justify-content:space-between}.coin-name{margin-left:10rpx;font-size:34rpx;font-weight:800}.coin-icon{display:inline-grid;place-items:center;width:38rpx;height:38rpx;border-radius:50%;background:#1E293B;color:#fff}.direction,.status{margin-left:12rpx;padding:4rpx 9rpx;background:#ECFDF5;color:var(--cf-up);font-size:21rpx}.status{margin:0}.price-panel{margin-top:28rpx;padding:20rpx;background:#F8FAFC;border-radius:14rpx}.price-row{padding:15rpx 0;border-bottom:1rpx solid var(--cf-line);font-size:24rpx}.price-row:last-child{border:0}.targets{display:flex;flex-direction:column;align-items:flex-end;gap:5rpx}.fine{font-size:21rpx;color:var(--cf-soft)}.down{color:var(--cf-down)}.strategy{margin-top:24rpx;color:var(--cf-muted);font-size:23rpx;line-height:1.7}.strategy view{margin-bottom:8rpx;font-size:24rpx;font-weight:700;color:var(--cf-text)}.point-actions{position:fixed;bottom:0;left:0;right:0;padding:16rpx 24rpx calc(16rpx + env(safe-area-inset-bottom));background:#fff;border-top:1rpx solid var(--cf-line)}.point-actions>view{display:flex;gap:12rpx}.point-actions button{flex:1;height:72rpx;font-size:25rpx}.point-actions>text{display:block;margin-top:10rpx;color:var(--cf-soft);font-size:18rpx;text-align:center}.overlay{position:fixed;inset:0;z-index:10;display:flex;align-items:center;justify-content:center;padding:30rpx;background:rgba(15,23,42,.45)}.dialog{position:relative;width:100%;padding:30rpx;background:#fff;border-radius:18rpx;text-align:center}.close{position:absolute;top:12rpx;right:18rpx;font-size:32rpx;color:var(--cf-soft)}.dialog-title{font-size:29rpx;font-weight:800}.dialog-copy{margin:18rpx 0;color:var(--cf-muted);font-size:23rpx}.dialog button{width:100%;height:72rpx}.contact-choice{display:flex;justify-content:space-between;padding:18rpx 0;border-bottom:1rpx solid var(--cf-line);text-align:left}.contact-account{display:block;margin-top:4rpx;color:var(--cf-muted);font-size:19rpx}.check{color:var(--cf-brand)}.contact-qr{width:220rpx;height:220rpx;margin:18rpx auto}
</style>
<style scoped>.point-nav{padding-top:var(--cf-safe-top);height:calc(80rpx + var(--cf-safe-top))}</style>
<style scoped>.point-nav .nav-bookmark.is-favorite{color:#2563eb}.point-nav .nav-bookmark.is-favorite svg{fill:currentColor}</style>
<style scoped>
.unlock-dialog{box-sizing:border-box;max-width:690rpx;padding:44rpx 36rpx 34rpx;border:1rpx solid #E5EDF8;border-radius:32rpx;box-shadow:0 18rpx 42rpx rgba(15,23,42,.14)}
.unlock-dialog .close{top:20rpx;right:20rpx;width:56rpx;height:56rpx;font-size:52rpx;line-height:1;color:#94A3B8}
.unlock-dialog .modal-lock{display:flex;align-items:center;justify-content:center;width:112rpx;height:112rpx;margin:4rpx auto 24rpx;border-radius:50%;background:#EFF6FF;color:#2563EB}
.unlock-dialog .modal-lock svg{width:58rpx;height:58rpx;fill:none;stroke:currentColor;stroke-width:1.8;stroke-linecap:round;stroke-linejoin:round}
.unlock-dialog .dialog-title{font-size:34rpx;font-weight:800;color:#0F172A}
.unlock-dialog .dialog-copy{margin:24rpx 0 30rpx;color:#64748B;font-size:25rpx;line-height:1.55}
.unlock-dialog .unlock-primary{width:100%;height:82rpx;border-radius:16rpx;background:#2563EB;color:#fff;font-size:27rpx;font-weight:700}
.unlock-dialog .modal-secondary{display:block;margin-top:22rpx;color:#2563EB;font-size:23rpx;text-align:center}
.contact-dialog{box-sizing:border-box;max-width:690rpx;padding:44rpx 36rpx 34rpx;border:1rpx solid #E5EDF8;border-radius:32rpx;box-shadow:0 18rpx 42rpx rgba(15,23,42,.14)}
.contact-dialog .close{top:20rpx;right:20rpx;width:56rpx;height:56rpx;font-size:52rpx;line-height:1;color:#94A3B8}
.contact-dialog .dialog-title{font-size:34rpx;font-weight:800;color:#0F172A}
.contact-dialog .dialog-copy{margin:8rpx 0 26rpx;color:#94A3B8;font-size:24rpx}
.contact-dialog .contact-choice{min-height:92rpx;padding:16rpx;border:2rpx solid #E2E8F0;border-radius:18rpx}
.contact-dialog .contact-choice+.contact-choice{margin-top:14rpx}
.contact-dialog .contact-choice.selected{border-color:#2563EB;background:#F8FBFF}
.contact-dialog .contact-qr{margin:24rpx auto}
.contact-dialog button{height:82rpx;border-radius:16rpx;background:#2563EB;color:#fff;font-size:27rpx;font-weight:700}
</style>
