<script setup lang="ts">
import { t } from '../locales'

const tools = [
  { label: '爆仓数据', icon: 'flame', color: 'red' },
  { label: '资金费率', icon: 'percent', color: 'yellow' },
  { label: 'ETF流向', icon: 'etf', color: 'green' },
  { label: '巨鲸监控', icon: 'whale', color: 'cyan' },
  { label: '链上数据', icon: 'chain', color: 'purple' },
  { label: '解锁日历', icon: 'unlock', color: 'indigo' },
]

const posts = [
  { title: '孙雨辰八卦', readTime: '12', icon: 'file', color: 'blue' },
  { title: '特朗普中期选举', readTime: '9', icon: 'layers', color: 'purple' },
]

const communities = [
  { name: 'Telegram 频道', meta: '@CryptoFlash_CN · 12.8万订阅', action: '加入', icon: 'telegram' },
  { name: 'X (Twitter)', meta: '@CryptoFlashPro · 8.6万关注', action: '关注', icon: 'x' },
  { name: 'Discord 社区', meta: '与交易者实时交流', action: '加入', icon: 'discord' },
]
</script>

<template>
  <view class="more-panel">
    <view class="more-section tools-section">
      <text class="more-section-title">{{ t('实用工具') }}</text>
      <view class="tool-grid">
        <view v-for="tool in tools" :key="tool.label" class="tool-item">
          <view class="tool-icon" :class="tool.color">
            <svg v-if="tool.icon === 'scale'" viewBox="0 0 24 24" aria-hidden="true"><path d="M12 3v17m-7 1h14M5 6h14M12 4l-2 2m2-2 2 2M5 6l-3 7h6L5 6Zm14 0-3 7h6l-3-7Z"/></svg>
            <svg v-else-if="tool.icon === 'flame'" viewBox="0 0 24 24" aria-hidden="true"><path d="M13 2c1 4-2 5-1 8 1-1 2-2 2-4 4 3 6 6 5 10a7 7 0 0 1-14 0c0-3 2-6 5-8 0 3 1 4 2 5 1-4 0-7 1-11Z"/></svg>
            <svg v-else-if="tool.icon === 'percent'" viewBox="0 0 24 24" aria-hidden="true"><circle cx="6.5" cy="6.5" r="2.5"/><circle cx="17.5" cy="17.5" r="2.5"/><path d="m19 4-14 16"/></svg>
            <svg v-else-if="tool.icon === 'etf'" viewBox="0 0 24 24" aria-hidden="true"><rect x="3" y="6" width="18" height="13" rx="2"/><path d="M7 6V4h10v2m-5 4v5m-2.5-2.5H14.5M7 19v2m10-2v2"/></svg>
            <svg v-else-if="tool.icon === 'whale'" viewBox="0 0 24 24" aria-hidden="true"><path d="M3 13c2-4 7-6 12-4l4-3 1 6c1 1 1 3 0 4l-2-1c-3 3-8 3-12 1l-3 2v-5Zm9-3v.1"/><circle cx="15" cy="11" r=".7"/></svg>
            <svg v-else-if="tool.icon === 'chain'" viewBox="0 0 24 24" aria-hidden="true"><path d="m9 15-2 2a4 4 0 0 1-6-6l4-4a4 4 0 0 1 6 0m4 2 2-2a4 4 0 0 1 6 6l-4 4a4 4 0 0 1-6 0m-5-1 6-6"/></svg>
            <svg v-else-if="tool.icon === 'unlock'" viewBox="0 0 24 24" aria-hidden="true"><rect x="4" y="10" width="16" height="11" rx="2"/><path d="M8 10V7a4 4 0 0 1 7.5-2M12 14v3"/></svg>
            <svg v-else viewBox="0 0 24 24" aria-hidden="true"><rect x="5" y="2" width="14" height="20" rx="2"/><path d="M8 6h8M8 10h2m4 0h2m-8 4h2m4 0h2m-8 4h2m4 0h2"/></svg>
          </view>
          <text class="tool-label">{{ t(tool.label) }}</text>
        </view>
      </view>
    </view>

    <view class="posts-section">
      <view class="more-section-heading"><text class="more-section-title">{{ t('精选帖子') }}</text><text class="more-link">{{ t('更多 ›') }}</text></view>
      <view class="post-list">
        <view v-for="post in posts" :key="post.title" class="post-card">
          <view class="post-icon" :class="post.color">
            <svg v-if="post.icon === 'file'" viewBox="0 0 24 24" aria-hidden="true"><path d="M6 2h8l5 5v15H6zM14 2v6h5M9 13h7m-7 3h7"/></svg>
            <svg v-else viewBox="0 0 24 24" aria-hidden="true"><path d="m12 2 9 5-9 5-9-5 9-5Zm-9 10 9 5 9-5M3 17l9 5 9-5"/></svg>
          </view>
          <view class="post-copy"><text class="post-title">{{ t(post.title) }}</text><text class="post-meta">{{ post.readTime }} {{ t('分钟阅读') }}</text></view>
        </view>
      </view>
    </view>

    <view class="more-section community-section">
      <text class="more-section-title community-heading">{{ t('官方社区') }}</text>
      <view v-for="community in communities" :key="community.icon" class="community-row">
        <view class="community-icon" :class="community.icon">
          <svg v-if="community.icon === 'telegram'" viewBox="0 0 24 24" aria-hidden="true"><path d="m21 3-3.2 18-6.1-5.1-3.5 3.4.6-5.4L21 3ZM8.8 13.9l8.9-7.4"/></svg>
          <text v-else-if="community.icon === 'x'">X</text>
          <svg v-else viewBox="0 0 24 24" aria-hidden="true"><path d="M7 7c-1-1-2-1-3-1l-1 4c-1 4 1 8 4 9l2-2h6l2 2c3-1 5-5 4-9l-1-4c-1 0-2 0-3 1-3-1-7-1-10 0Z"/><circle cx="9" cy="12" r="1"/><circle cx="15" cy="12" r="1"/></svg>
        </view>
        <view class="community-copy"><text class="community-name">{{ t(community.name) }}</text><text class="community-meta">{{ t(community.meta) }}</text></view>
        <text class="community-action">{{ t(community.action) }}</text>
      </view>
    </view>
  </view>
</template>

<style scoped lang="scss">
.more-panel{display:flex;flex-direction:column;gap:20rpx;padding-bottom:22rpx;color:var(--cf-text)}
.more-section,.post-card{min-width:0;border:1rpx solid var(--cf-line);border-radius:22rpx;background:var(--cf-card)}
.more-section{padding:20rpx 22rpx}
.more-section-title{font-size:31rpx;font-weight:700;line-height:1.4}
.tool-grid{display:grid;grid-template-columns:repeat(4,minmax(0,1fr));row-gap:24rpx;margin-top:20rpx}
.tool-item{display:flex;align-items:center;flex-direction:column;gap:8rpx;min-width:0}
.tool-icon{display:grid;place-items:center;width:62rpx;height:62rpx;border-radius:20rpx}
.tool-icon svg{width:30rpx;height:30rpx;fill:none;stroke:currentColor;stroke-width:1.8;stroke-linecap:round;stroke-linejoin:round}
.tool-icon.blue{background:#edf5ff;color:#3978ed}.tool-icon.red{background:#fff0f0;color:#ef5b5b}.tool-icon.yellow{background:#fff9e8;color:#e5a914}.tool-icon.green{background:#eafaf3;color:#13a878}
.tool-icon.cyan{background:#eafaff;color:#16a9d2}.tool-icon.purple{background:#f7efff;color:#9c54e8}.tool-icon.indigo{background:#eef1ff;color:#5265d8}.tool-icon.gray{background:#f1f3f6;color:#677184}
.tool-label{max-width:100%;color:var(--cf-text);font-size:26rpx;line-height:1.4;text-align:center;overflow-wrap:anywhere}
.posts-section{min-width:0}
.more-section-heading{display:flex;align-items:center;justify-content:space-between;margin:2rpx 0 12rpx}
.more-link{color:var(--cf-brand);font-size:23rpx}
.post-list{display:flex;flex-direction:column;gap:12rpx}
.post-card{display:flex;align-items:center;gap:14rpx;min-height:100rpx;padding:14rpx 16rpx}
.post-icon{display:grid;flex:none;width:66rpx;height:66rpx;place-items:center;border-radius:14rpx;color:#fff}
.post-icon.blue{background:linear-gradient(135deg,#2563eb,#17bfd3)}.post-icon.purple{background:linear-gradient(135deg,#8b35d2,#e650c2)}
.post-icon svg{width:31rpx;height:31rpx;fill:none;stroke:currentColor;stroke-width:1.8;stroke-linecap:round;stroke-linejoin:round}
.post-copy{display:flex;flex:1;min-width:0;flex-direction:column;gap:7rpx}
.post-title{color:var(--cf-text);font-size:27rpx;font-weight:700;line-height:1.4;overflow-wrap:anywhere}
.post-meta{color:var(--cf-soft);font-size:23rpx;line-height:1.4}
.community-section{overflow:hidden;padding:0}
.community-heading{display:block;padding:18rpx 22rpx;border-bottom:1rpx solid var(--cf-divider)}
.community-row{display:flex;align-items:center;gap:12rpx;min-height:100rpx;padding:14rpx 18rpx;border-bottom:1rpx solid var(--cf-divider)}
.community-row:last-child{border-bottom:0}
.community-icon{display:grid;flex:none;width:50rpx;height:50rpx;place-items:center;border-radius:50%;background:#229ed9;color:#fff}
.community-icon svg{width:26rpx;height:26rpx;fill:none;stroke:currentColor;stroke-width:1.8;stroke-linecap:round;stroke-linejoin:round}
.community-icon.x{background:#161b22;font-size:26rpx;font-weight:700}.community-icon.discord{background:#5865f2}
.community-copy{display:flex;flex:1;min-width:0;flex-direction:column;gap:4rpx}
.community-name{color:var(--cf-text);font-size:27rpx;line-height:1.4}
.community-meta{color:var(--cf-soft);font-size:23rpx;line-height:1.4;overflow-wrap:anywhere}
.community-action{display:grid;flex:none;min-width:76rpx;height:52rpx;padding:0 12rpx;place-items:center;border-radius:26rpx;background:#eef6ff;color:var(--cf-brand);font-size:23rpx}
@media(max-width:360px){.more-section{padding-right:16rpx;padding-left:16rpx}.community-row{gap:9rpx;padding-right:12rpx;padding-left:12rpx}.community-action{padding:0 8rpx}}
</style>
