<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import { api } from '../api'
import { locale, t } from '../locales'

const props = defineProps<{ visible: boolean }>()
const emit = defineEmits<{ (event: 'close'): void; (event: 'unlocked'): void }>()

const channels = ref<any[]>([])
const selectedChannel = ref(0)
const loadingContacts = ref(false)
const unlocking = ref(false)
const notice = ref('')

async function loadContacts() {
  loadingContacts.value = true
  notice.value = ''
  try {
    channels.value = await api.contacts()
    selectedChannel.value = 0
  } catch (error: any) {
    channels.value = []
    notice.value = error?.message || t('客服信息加载失败')
  } finally {
    loadingContacts.value = false
  }
}

watch(() => props.visible, (visible) => {
  if (visible) loadContacts()
})
watch(locale, () => {
  if (props.visible) loadContacts()
})
onMounted(() => {
  if (props.visible) loadContacts()
})

function close() {
  notice.value = ''
  emit('close')
}

async function unlock() {
  if (unlocking.value) return
  unlocking.value = true
  notice.value = ''
  try {
    await api.unlockAll()
    emit('unlocked')
  } catch (error: any) {
    notice.value = error?.message || t('解锁失败')
  } finally {
    unlocking.value = false
  }
}

async function copyContact() {
  const account = channels.value[selectedChannel.value]?.account || ''
  if (!account) {
    notice.value = t('暂无客服联系方式')
    return
  }
  try {
    await new Promise<void>((resolve, reject) => {
      uni.setClipboardData({ data: account, success: () => resolve(), fail: reject })
    })
  } catch {
    try {
      if (typeof navigator !== 'undefined' && navigator.clipboard?.writeText) await navigator.clipboard.writeText(account)
      else {
        const input = document.createElement('textarea')
        input.value = account
        input.style.position = 'fixed'
        input.style.opacity = '0'
        document.body.appendChild(input)
        input.select()
        if (!document.execCommand('copy')) throw new Error('copy failed')
        input.remove()
      }
    } catch {
      notice.value = t('复制失败')
      return
    }
  }
  notice.value = t('已复制')
}

function isTelegram(channel: any) {
  return /telegram/i.test(String(channel?.type || ''))
}

function isWhatsApp(channel: any) {
  return /whatsapp/i.test(String(channel?.type || ''))
}
</script>

<template>
  <view v-if="visible" class="unlock-overlay" @click="close">
    <view v-if="notice" class="unlock-notice" role="alert">{{ notice }}</view>
    <view class="unlock-stack" @click.stop>
      <view class="unlock-card">
        <text class="dialog-close" @click="close">×</text>
        <view class="lock-badge"><svg viewBox="0 0 24 24" aria-hidden="true"><rect x="4" y="10" width="16" height="11" rx="2"/><path d="M8 10V7a4 4 0 1 1 8 0v3"/></svg></view>
        <text class="dialog-title">{{ t('解锁全部点位') }}</text>
        <text class="dialog-copy">{{ t('持有平台指定 NFT 可解锁全部点位') }}<br/>{{ t('查看全部历史与未来点位') }}</text>
        <button class="unlock-primary" :disabled="unlocking" @click="unlock">{{ unlocking ? t('处理中...') : t('去购买 NFT') }}</button>
        <text class="dialog-link" @click="close">{{ t('如何获得 NFT?') }}</text>
      </view>

      <view class="contact-card">
        <text class="dialog-close" @click="close">×</text>
        <text class="dialog-title">{{ t('联系客服') }}</text>
        <text class="contact-subtitle">{{ t('选择您的联系方式') }}</text>
        <view v-if="loadingContacts" class="contact-state">{{ t('加载中...') }}</view>
        <view v-else-if="!channels.length" class="contact-state">{{ t('暂无客服联系方式') }}</view>
        <view v-else>
          <view v-for="(channel, index) in channels" :key="channel.type || index" class="contact-choice" :class="{ selected: selectedChannel === index }" @click="selectedChannel = index">
            <view class="channel-icon" :class="{ telegram: isTelegram(channel), whatsapp: isWhatsApp(channel) }">
              <svg v-if="isTelegram(channel)" viewBox="0 0 24 24" aria-hidden="true"><path d="m21 3-3.2 18-6.1-5.1-3.5 3.4.6-5.4L21 3Z"/><path d="m8.8 13.9 8.9-7.4"/></svg>
              <svg v-else-if="isWhatsApp(channel)" viewBox="0 0 24 24" aria-hidden="true"><path d="M20 11.5a8 8 0 0 1-11.8 7L4 20l1.5-4A8 8 0 1 1 20 11.5Z"/><path d="M8 8.5c.3 2 1.5 3.5 3.5 4.5 1 .5 1.8.6 2.5.2l.8-.6"/></svg>
              <svg v-else viewBox="0 0 24 24" aria-hidden="true"><path d="M4 5h16v14H4z"/><path d="m4 7 8 6 8-6"/></svg>
            </view>
            <view class="channel-copy"><text>{{ channel.type || t('客服') }}</text><text>{{ channel.account || t('暂无') }}</text></view>
            <text v-if="selectedChannel === index" class="selected-mark">✓</text>
          </view>
          <view v-if="channels[selectedChannel]?.qr" class="contact-qr-wrap"><text>{{ t('扫码加入客服') }}</text><image class="contact-qr" :src="channels[selectedChannel].qr" mode="aspectFit"/></view>
        </view>
        <button class="contact-copy" @click="copyContact">{{ t('复制联系方式') }}</button>
      </view>
    </view>
  </view>
</template>

<style scoped>
.unlock-overlay{position:fixed;inset:0;z-index:10000;display:flex;align-items:flex-start;justify-content:center;overflow-y:auto;padding:24rpx;background:#fff}
.unlock-notice{position:fixed;top:calc(24rpx + var(--cf-safe-top));left:50%;z-index:10002;max-width:calc(100vw - 64rpx);padding:18rpx 28rpx;transform:translateX(-50%);border-radius:14rpx;background:#0f172a;color:#fff;font-size:23rpx;line-height:1.4;text-align:center;box-shadow:0 10rpx 30rpx rgba(15,23,42,.25);pointer-events:none}
.unlock-stack{display:flex;flex-direction:column;gap:24rpx;width:100%;max-width:690rpx;margin:auto;padding:24rpx 0 48rpx}
.unlock-card,.contact-card{position:relative;box-sizing:border-box;width:100%;padding:44rpx 36rpx 34rpx;border:1rpx solid #e5edf8;border-radius:28rpx;background:#fff;text-align:center;box-shadow:0 18rpx 42rpx rgba(15,23,42,.14)}
.dialog-close{position:absolute;top:18rpx;right:20rpx;width:56rpx;height:56rpx;color:#94a3b8;font-size:52rpx;line-height:1;text-align:center}
.lock-badge{display:flex;align-items:center;justify-content:center;width:112rpx;height:112rpx;margin:4rpx auto 24rpx;border-radius:50%;background:#eff6ff;color:#2563eb}
.lock-badge svg{width:58rpx;height:58rpx;fill:none;stroke:currentColor;stroke-width:1.8;stroke-linecap:round;stroke-linejoin:round}
.dialog-title{display:block;color:#0f172a;font-size:34rpx;font-weight:800}
.dialog-copy{display:block;margin:24rpx 0 30rpx;color:#64748b;font-size:25rpx;line-height:1.55}
.unlock-primary,.contact-copy{width:100%;height:82rpx;padding:0;border:0;border-radius:16rpx;background:#2563eb;color:#fff;font-size:27rpx;font-weight:700;line-height:82rpx}
.unlock-primary[disabled]{opacity:.55}
.dialog-link{display:block;margin-top:22rpx;color:#2563eb;font-size:23rpx}
.contact-subtitle{display:block;margin:10rpx 0 26rpx;color:#94a3b8;font-size:24rpx}
.contact-state{padding:30rpx 0;color:#94a3b8;font-size:23rpx}
.contact-choice{display:flex;align-items:center;gap:18rpx;min-height:92rpx;padding:16rpx;border:2rpx solid #e2e8f0;border-radius:18rpx;text-align:left}
.contact-choice+.contact-choice{margin-top:14rpx}.contact-choice.selected{border-color:#2563eb;background:#f8fbff}
.channel-icon{display:flex;align-items:center;justify-content:center;flex:0 0 58rpx;width:58rpx;height:58rpx;border-radius:50%;background:#64748b;color:#fff}.channel-icon.telegram{background:#229ed9}.channel-icon.whatsapp{background:#25d366}.channel-icon svg{width:32rpx;height:32rpx;fill:none;stroke:currentColor;stroke-width:1.6;stroke-linecap:round;stroke-linejoin:round}
.channel-copy{display:flex;flex:1;flex-direction:column;gap:5rpx;min-width:0;color:#0f172a;font-size:25rpx;font-weight:700}.channel-copy text+text{color:#94a3b8;font-size:22rpx;font-weight:400;overflow:hidden;text-overflow:ellipsis;white-space:nowrap}.selected-mark{color:#2563eb;font-size:30rpx}
.contact-qr-wrap{display:flex;align-items:center;flex-direction:column;gap:12rpx;margin:24rpx 0;color:#94a3b8;font-size:22rpx}.contact-qr{width:180rpx;height:180rpx;padding:8rpx;border:1rpx solid #e2e8f0;border-radius:16rpx}
.contact-copy{margin-top:24rpx}
</style>
