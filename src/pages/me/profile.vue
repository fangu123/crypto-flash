<script setup lang="ts">
import { computed, ref, onMounted, watch } from 'vue'
import { api } from '../../api'
import { locale, t } from '../../locales'
import { API_BASE_URL } from '../../config/api'
import '../../styles/app.scss'

const profile = ref<any>({})
const username = ref('')
const email = ref('')
const head = ref('')
const wallet = ref('')
const emailCode = ref('')
const loading = ref(false)
const uploading = ref(false)
const avatarSrc = computed(() => {
  const value = head.value.trim()
  if (!value || /^[a-z][a-z\d+.-]*:/i.test(value)) return value
  return `${API_BASE_URL}${value.startsWith('/') ? '' : '/'}${value}`
})

function imageUrl(value: any): string {
  if (typeof value === 'string') return value.trim()
  if (!value || typeof value !== 'object') return ''
  for (const key of ['url', 'fileUrl', 'file_url', 'imageUrl', 'image', 'head', 'path', 'src']) {
    const found = imageUrl(value[key])
    if (found) return found
  }
  for (const key of ['data', 'result', 'file']) {
    const found = imageUrl(value[key])
    if (found) return found
  }
  return ''
}

async function load() {
  try {
    profile.value = await api.profile()
    username.value = profile.value.username || profile.value.nickname || ''
    email.value = profile.value.email || ''
    head.value = profile.value.head || profile.value.avatar || ''
    wallet.value = profile.value.walletAddress || profile.value.wallet_address || profile.value.wallet || ''
  } catch (error: any) {
    uni.showToast({ title: error?.message || t('资料加载失败'), icon: 'none' })
  }
}

onMounted(load)
watch(locale, load)

async function chooseAvatar() {
  if (uploading.value || loading.value) return
  try {
    const result = await uni.chooseImage({ count: 1, sizeType: ['compressed'], sourceType: ['album'] })
    const filePath = result.tempFilePaths?.[0]
    if (!filePath) return
    uploading.value = true
    const uploaded = await api.uploadProfileImage(filePath)
    const url = imageUrl(uploaded)
    if (!url) throw new Error(t('头像上传失败'))
    head.value = url
    uni.showToast({ title: t('头像上传成功'), icon: 'none' })
  } catch (error: any) {
    if (String(error?.errMsg || '').toLowerCase().includes('cancel')) return
    uni.showToast({ title: error?.message || t('头像上传失败'), icon: 'none' })
  } finally {
    uploading.value = false
  }
}

async function save() {
  if (!username.value.trim() || !email.value.trim()) return uni.showToast({ title: t('请输入昵称和邮箱'), icon: 'none' })
  loading.value = true
  try {
    await api.updatePersonalInfo({ username: username.value.trim(), email: email.value.trim(), head: head.value.trim(), walletAddress: wallet.value.trim(), emailCode: emailCode.value.trim() })
    try { profile.value = await api.profile() } catch {}
    uni.showToast({ title: t('资料已更新'), icon: 'none' })
  } catch (error: any) {
    uni.showToast({ title: error?.message || t('资料更新失败'), icon: 'none' })
  } finally {
    loading.value = false
  }
}

async function bindWallet() {
  if (!wallet.value.trim()) return uni.showToast({ title: t('请输入钱包地址'), icon: 'none' })
  try {
    await api.bindWallet(wallet.value.trim())
    uni.showToast({ title: t('钱包绑定成功'), icon: 'none' })
  } catch (error: any) {
    uni.showToast({ title: error?.message || t('钱包绑定失败'), icon: 'none' })
  }
}
</script>

<template>
  <view class="inner-page">
    <view class="inner-nav">
      <text class="back" @click="uni.navigateBack()"><svg viewBox="0 0 24 24" aria-hidden="true"><path d="m15 18-6-6 6-6"/></svg></text>
      <text class="nav-title">{{ t('个人资料') }}</text>
    </view>
    <view class="inner-main">
      <view class="inner-card profile-card">
        <view class="avatar-picker" :class="{ 'is-uploading': uploading }" @click="chooseAvatar">
          <image v-if="avatarSrc" class="profile-avatar" :src="avatarSrc" mode="aspectFill" />
          <view v-else class="profile-avatar avatar-placeholder">♙</view>
          <text class="avatar-action">{{ uploading ? t('上传中...') : t('更换头像') }}</text>
        </view>
        <view class="form-field">
          <text class="form-label">{{ t('昵称') }}</text>
          <input v-model="username" class="form-input" :placeholder="t('输入昵称')" />
        </view>
        <view class="form-field">
          <text class="form-label">{{ t('邮箱') }}</text>
          <input v-model="email" class="form-input" :placeholder="t('输入邮箱')" />
        </view>
        <view class="form-field">
          <text class="form-label">{{ t('邮箱验证码') }}</text>
          <input v-model="emailCode" class="form-input" :placeholder="t('修改邮箱时填写')" />
        </view>
        <view class="form-field">
          <text class="form-label">{{ t('钱包地址') }}</text>
          <input v-model="wallet" class="form-input" :placeholder="t('输入钱包地址')" />
          <button class="action-button wallet-action" @click="bindWallet">{{ t('绑定钱包') }}</button>
        </view>
        <button class="action-button" :disabled="loading || uploading" @click="save">
          {{ loading ? t('保存中...') : t('保存资料') }}
        </button>
      </view>
    </view>
  </view>
</template>

<style scoped>
.profile-card { margin-top: 4rpx; }
.avatar-picker { display: flex; flex-direction: column; align-items: center; gap: 10rpx; margin: 0 auto 26rpx; }
.profile-avatar { display: block; width: 112rpx; height: 112rpx; border-radius: 50%; background: var(--cf-brand); }
.avatar-placeholder { display: grid; place-items: center; color: #fff; font-size: 54rpx; }
.avatar-action { color: var(--cf-brand); font-size: 24rpx; }
.is-uploading { opacity: .65; }
.wallet-action { margin-top: 14rpx; }
.action-button[disabled] { opacity: .6; }
</style>
