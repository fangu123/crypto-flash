<script setup lang="ts">
import { locale, setLocale, t } from '../../locales'
import '../../styles/app.scss'
</script>

<template>
  <view class="inner-page language-page">
    <view class="inner-nav">
      <text class="back" @click="uni.navigateBack()">‹</text>
      <text class="nav-title">{{ t('语言设置') }}</text>
    </view>
    <view class="inner-main">
      <view class="language-options" role="group" :aria-label="t('语言设置')">
        <button class="language-option" :class="{ active: locale === 'zh' }" @click="setLocale('zh')">
          <text>{{ t('简体中文') }}</text>
          <text v-if="locale === 'zh'" class="language-check">✓</text>
        </button>
        <button class="language-option" :class="{ active: locale === 'en' }" @click="setLocale('en')">
          <text>English</text>
          <text v-if="locale === 'en'" class="language-check">✓</text>
        </button>
      </view>
    </view>
  </view>
</template>

<style scoped>
.language-page{min-height:100vh}
.language-options{overflow:hidden;border:1rpx solid var(--cf-line);border-radius:12rpx;background:var(--cf-card)}
.language-option{display:flex;align-items:center;justify-content:space-between;width:100%;height:92rpx;padding:0 22rpx;margin:0;border-bottom:1rpx solid var(--cf-line);border-radius:0;background:var(--cf-card);color:var(--cf-text);font-size:26rpx;text-align:left}
.language-option:last-child{border-bottom:0}
.language-option.active{color:var(--cf-brand)}
.language-check{font-size:27rpx;font-weight:700}
</style>
