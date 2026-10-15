<script setup lang="ts">
import { computed, ref } from 'vue'
import { onShow } from '@dcloudio/uni-app'
import { clearDiagnosticLogs, getDiagnosticLogs, type DiagnosticLog } from '../../api/diagnostics'
import { t } from '../../locales'
import '../../styles/app.scss'
const logs = ref<DiagnosticLog[]>([])
const filter = ref<'all' | 'failure'>('all')
const visibleLogs = computed(() => filter.value === 'all' ? logs.value : logs.value.filter(x => x.outcome !== 'success'))
function refresh() { logs.value = getDiagnosticLogs() }
function clear() { clearDiagnosticLogs(); refresh() }
function time(value: number) { return new Date(value).toLocaleString('zh-CN', { hour12: false }) }
function statusLabel(entry: DiagnosticLog) { return entry.outcome === 'success' ? t('成功') : entry.outcome === 'network_error' ? t('网络错误') : entry.outcome === 'http_error' ? t('HTTP 错误') : t('接口失败') }
onShow(refresh)
</script>
<template><view class="inner-page diagnostics-page"><view class="inner-nav"><text class="back" @click="uni.navigateBack()">‹</text><text class="nav-title">{{t('接口日志')}}</text></view><view class="inner-main"><view class="diagnostics-toolbar"><view class="inner-tabs"><text class="inner-tab" :class="{active:filter==='all'}" @click="filter='all'">{{t('全部')}}</text><text class="inner-tab" :class="{active:filter==='failure'}" @click="filter='failure'">{{t('仅失败')}}</text></view><button class="clear-logs" @click="clear">{{t('清空')}}</button></view><view v-if="!visibleLogs.length" class="inner-card empty-state">{{t('暂无接口日志')}}</view><view v-for="entry in visibleLogs" :key="entry.id" class="diagnostic-row"><view class="diagnostic-top"><text class="diagnostic-endpoint">{{entry.method}} {{entry.endpoint}}</text><text class="diagnostic-status" :class="entry.outcome==='success'?'ok':'bad'">{{statusLabel(entry)}}</text></view><view class="diagnostic-meta"><text>{{time(entry.time)}}</text><text v-if="entry.durationMs !== undefined">{{entry.durationMs}}ms</text><text v-if="entry.httpStatus">HTTP {{entry.httpStatus}}</text><text v-if="entry.code !== undefined">code {{entry.code}}</text></view><text v-if="entry.message" class="diagnostic-message">{{entry.message}}</text></view></view></view></template>
<style scoped>.diagnostics-toolbar{display:flex;align-items:center;justify-content:space-between;margin-bottom:18rpx}.diagnostics-toolbar .inner-tabs{margin:0}.clear-logs{height:58rpx;padding:0 18rpx;border:1rpx solid var(--cf-line);border-radius:10rpx;background:#fff;color:var(--cf-down);font-size:21rpx;line-height:58rpx}.diagnostic-row{padding:20rpx;background:#fff;border:1rpx solid var(--cf-line);border-radius:12rpx}.diagnostic-row+.diagnostic-row{margin-top:12rpx}.diagnostic-top{display:flex;align-items:center;justify-content:space-between;gap:12rpx}.diagnostic-endpoint{min-width:0;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;font-size:22rpx;font-weight:700}.diagnostic-status{flex:0 0 auto;font-size:21rpx}.diagnostic-status.ok{color:var(--cf-up)}.diagnostic-status.bad{color:var(--cf-down)}.diagnostic-meta{display:flex;gap:14rpx;margin-top:9rpx;color:var(--cf-muted);font-size:19rpx}.diagnostic-message{display:block;margin-top:9rpx;color:var(--cf-down);font-size:21rpx;line-height:1.4}</style>
