<script setup lang="ts">
import { onBeforeUnmount, onMounted, ref } from 'vue'
import { t } from '../locales'

const sentimentTarget = 72
const sentimentValue = ref(0)
// 从左侧零刻度起步，最终保留原来指向 72 的位置。
const needleStartRotation = -180 - Math.atan2(-47, 44) * 180 / Math.PI
const needleRotation = ref(needleStartRotation)
let animationTimer: ReturnType<typeof setTimeout> | undefined

onMounted(() => {
  const startedAt = Date.now()
  const animate = () => {
    const progress = Math.min((Date.now() - startedAt) / 1500, 1)
    sentimentValue.value = Math.round(sentimentTarget * (1 - Math.pow(1 - progress, 3)))
    const remaining = progress - 1
    const swing = 1 + 2.2 * remaining ** 3 + 1.2 * remaining ** 2
    needleRotation.value = progress === 1 ? 0 : needleStartRotation * (1 - swing)
    if (progress < 1) animationTimer = setTimeout(animate, 16)
  }
  animationTimer = setTimeout(animate, 16)
})

onBeforeUnmount(() => {
  if (animationTimer !== undefined) clearTimeout(animationTimer)
})

const metrics = [
  { label: 'BTC 资金费率', value: '0.0128%', detail: '多头付费', tone: 'positive', icon: '↗' },
  { label: 'ETF 净流入', value: '+$286M', detail: '连续 5 日净流入', tone: 'positive', icon: '$' },
  { label: '全网持仓量', value: '$38.4B', detail: '+2.1%', tone: 'positive', icon: '◉' },
  { label: 'BTC 主导率', value: '54.2%', detail: '-0.3%', tone: 'negative', icon: '₿' },
]

const exchangeFlows = [
  { name: 'Binance', value: '-2,340 BTC', direction: 'outflow', width: '100%' },
  { name: 'Coinbase', value: '+860 BTC', direction: 'inflow', width: '42%' },
  { name: 'OKX', value: '-1,120 BTC', direction: 'outflow', width: '58%' },
]

const gainers = [
  { symbol: 'TON/USDT', price: '7.12', change: '+12.4%' },
  { symbol: 'ARB/USDT', price: '1.08', change: '+9.7%' },
  { symbol: 'INJ/USDT', price: '28.45', change: '+8.1%' },
]
</script>

<template>
  <view class="data-panel">
    <view class="sentiment-section">
      <view class="sentiment-gauge" :aria-label="String(sentimentValue)">
          <svg viewBox="0 0 164 98" aria-hidden="true">
            <defs><linearGradient id="sentiment-gradient" x1="0" y1="0" x2="1" y2="0"><stop offset="0%" stop-color="#ef4444"/><stop offset="48%" stop-color="#facc15"/><stop offset="100%" stop-color="#16b981"/></linearGradient></defs>
            <path class="gauge-track" d="M18 82 A64 64 0 0 1 146 82" />
            <path class="gauge-color" d="M18 82 A64 64 0 0 1 146 82" />
            <line class="gauge-needle" x1="82" y1="82" x2="126" y2="35" :transform="`rotate(${needleRotation} 82 82)`" />
            <circle class="gauge-hub" cx="82" cy="82" r="5" />
          </svg>
      </view>
      <view class="sentiment-copy">
        <text class="sentiment-title">{{ t('恐惧贪婪指数') }}</text>
        <view class="sentiment-reading"><text>{{ sentimentValue }}</text><text>{{ t('贪婪') }}</text></view>
        <text class="sentiment-comparison">{{ t('昨日') }} 68 · {{ t('上周') }} 55</text>
      </view>
    </view>

    <view class="metrics-grid">
      <view v-for="metric in metrics" :key="metric.label" class="metric-card">
        <view class="metric-top"><text class="metric-icon" :class="metric.tone">{{ metric.icon }}</text><text class="metric-label">{{ t(metric.label) }}</text></view>
        <text class="metric-value">{{ metric.value }}</text>
        <text class="metric-detail" :class="metric.tone">{{ t(metric.detail) }}</text>
      </view>
    </view>

    <view class="data-section liquidation-section">
      <view class="section-heading"><text class="section-title">{{ t('24H 爆仓') }}</text><text class="updated-label">{{ t('24小时') }}</text></view>
      <view class="liquidation-total">$412.6M</view>
      <view class="liquidation-bar"><view class="long-liquidations"></view><view class="short-liquidations"></view></view>
      <view class="liquidation-legend">
        <view><i class="legend-dot long-dot"></i><text>{{ t('多单') }}</text><text class="legend-amount">$156.8M</text><text class="legend-share">38%</text></view>
        <view><i class="legend-dot short-dot"></i><text>{{ t('空单') }}</text><text class="legend-amount">$255.8M</text><text class="legend-share">62%</text></view>
      </view>
    </view>

    <view class="data-section flow-section">
      <view class="section-heading"><text class="section-title">{{ t('交易所 BTC 净流向') }}</text><text class="updated-label">{{ t('24小时') }}</text></view>
      <view v-for="flow in exchangeFlows" :key="flow.name" class="flow-row">
        <view class="flow-topline">
          <text class="exchange-name">{{ flow.name }}</text>
          <text class="flow-value" :class="flow.direction">{{ flow.value }}</text>
        </view>
        <view class="flow-bar"><view class="flow-bar-fill" :class="flow.direction" :style="{ width: flow.width }"></view></view>
      </view>
    </view>

    <view class="data-section gainers-section">
      <view class="section-heading"><text class="section-title">{{ t('涨幅榜') }}</text><text class="updated-label">{{ t('24小时涨幅') }}</text></view>
      <view class="gainer-heading"><text>{{ t('交易对') }}</text><text>{{ t('价格') }}</text><text>{{ t('涨幅') }}</text></view>
      <view v-for="(coin, index) in gainers" :key="coin.symbol" class="gainer-row">
        <view class="coin-pair"><text class="coin-rank">0{{ index + 1 }}</text><text>{{ coin.symbol }}</text></view>
        <text class="coin-price">{{ coin.price }}</text>
        <text class="coin-change">{{ coin.change }}</text>
      </view>
    </view>
  </view>
</template>

<style scoped lang="scss">
.data-panel{display:flex;flex-direction:column;gap:20rpx;padding-bottom:18rpx;color:var(--cf-text)}
.sentiment-section,.metric-card,.data-section{min-width:0;border:1rpx solid var(--cf-line);border-radius:16rpx;background:var(--cf-card)}
.sentiment-section{display:grid;grid-template-columns:minmax(0,1fr) minmax(0,1fr);align-items:center;gap:12rpx;min-height:184rpx;overflow:hidden;padding:18rpx 20rpx;background:#f5f5ff}
.section-heading{display:flex;align-items:center;justify-content:space-between;gap:12rpx;min-width:0}
.section-heading>view{display:flex;align-items:baseline;gap:12rpx;min-width:0}
.section-title{font-size:28rpx;font-weight:700;line-height:1.35}
.sentiment-gauge{width:100%;aspect-ratio:164/98}
.sentiment-gauge svg{display:block;width:100%;height:100%;overflow:visible}
.gauge-track,.gauge-color{fill:none;stroke-width:6;stroke-linecap:round}
.gauge-track{stroke:#e0e2f5}
.gauge-color{stroke:url(#sentiment-gradient)}
.gauge-needle{stroke:#172238;stroke-width:3.5;stroke-linecap:round}
.gauge-hub{fill:#172238}
.sentiment-copy{display:flex;min-width:0;flex-direction:column;align-items:flex-start;gap:6rpx}
.sentiment-title{color:var(--cf-muted);font-size:22rpx;white-space:nowrap}
.sentiment-reading{display:flex;align-items:baseline;gap:10rpx;white-space:nowrap}
.sentiment-reading text:first-child{color:#172238;font-size:62rpx;font-weight:800;line-height:1.05}
.sentiment-reading text:last-child{color:#0f9b70;font-size:31rpx;font-weight:750}
.sentiment-comparison{color:var(--cf-soft);font-size:21rpx;white-space:nowrap}
.metrics-grid{display:grid;grid-template-columns:repeat(2,minmax(0,1fr));gap:12rpx}
.metric-card{display:flex;flex-direction:column;gap:8rpx;padding:17rpx 18rpx}
.metric-top{display:flex;align-items:center;gap:9rpx;min-width:0}
.metric-icon{display:flex;align-items:center;justify-content:center;flex:none;width:34rpx;height:34rpx;border-radius:50%;background:#eef3ff;color:#4776db;font-size:19rpx;font-weight:700}
.metric-icon.negative{background:#fff0ef;color:#e45c57}
.metric-label{overflow:hidden;color:var(--cf-muted);font-size:19rpx;text-overflow:ellipsis;white-space:nowrap}
.metric-value{font-size:29rpx;font-weight:750;line-height:1.15;white-space:nowrap}
.metric-detail{font-size:18rpx;line-height:1.25}
.positive{color:#13a87a}.negative{color:#e45c57}
.data-section{padding:20rpx 22rpx}
.liquidation-total{margin:17rpx 0 13rpx;font-size:33rpx;font-weight:800}
.liquidation-bar{display:flex;gap:4rpx;width:100%;height:12rpx;overflow:hidden;border-radius:8rpx;background:#f2f3f7}
.long-liquidations{width:38%;background:#f1a340}.short-liquidations{flex:1;background:#e76568}
.liquidation-legend{display:flex;flex-direction:column;gap:13rpx;margin-top:17rpx}
.liquidation-legend>view{display:flex;align-items:center;gap:9rpx;color:var(--cf-muted);font-size:19rpx}
.legend-dot{width:13rpx;height:13rpx;border-radius:50%}.long-dot{background:#f1a340}.short-dot{background:#e76568}
.legend-amount{margin-left:auto;color:var(--cf-text);font-size:20rpx;font-weight:650}
.legend-share{flex:none;width:55rpx;color:var(--cf-soft);text-align:right}
.flow-row{display:flex;min-width:0;flex-direction:column;gap:7rpx;padding:14rpx 0 12rpx;border-bottom:1rpx solid var(--cf-divider)}
.flow-row:last-child{border-bottom:0}
.flow-topline{display:flex;align-items:center;justify-content:space-between;gap:8rpx;min-width:0}
.exchange-name{overflow:hidden;color:var(--cf-text);font-size:21rpx;text-overflow:ellipsis;white-space:nowrap}
.flow-value{font-size:21rpx;font-weight:650;white-space:nowrap}.flow-value.outflow{color:#159d78}.flow-value.inflow{color:#e16b65}
.flow-bar{width:78%;height:13rpx;overflow:hidden;border-radius:10rpx;background:transparent}
.flow-bar-fill{height:100%;border-radius:10rpx}.flow-bar-fill.outflow{background:#12b886}.flow-bar-fill.inflow{background:#ef4444}
.gainer-heading,.gainer-row{display:grid;grid-template-columns:minmax(0,1fr) 100rpx 106rpx;align-items:center;gap:8rpx}
.gainer-heading{padding:14rpx 0 8rpx;color:var(--cf-soft);font-size:17rpx}
.gainer-heading text:nth-child(2),.gainer-heading text:nth-child(3),.coin-price,.coin-change{text-align:right}
.gainer-row{min-height:59rpx;border-top:1rpx solid var(--cf-divider);font-size:20rpx}
.coin-pair{display:flex;align-items:center;gap:10rpx;min-width:0;font-weight:650}
.coin-rank{color:#9ba2b2;font-size:17rpx}
.coin-price{font-variant-numeric:tabular-nums;color:var(--cf-muted)}
.coin-change{font-weight:700;color:#13a87a}
@media(max-width:360px){.sentiment-section{gap:8rpx;padding-right:14rpx;padding-left:14rpx}.sentiment-title{font-size:20rpx}.sentiment-reading{gap:7rpx}.sentiment-reading text:first-child{font-size:56rpx}.sentiment-reading text:last-child{font-size:27rpx}.sentiment-comparison{font-size:18rpx}.data-section{padding-right:17rpx;padding-left:17rpx}.metric-card{padding-right:12rpx;padding-left:12rpx}.metric-value{font-size:26rpx}.metric-label{font-size:17rpx}.gainer-heading,.gainer-row{grid-template-columns:minmax(0,1fr) 84rpx 92rpx}}
</style>
