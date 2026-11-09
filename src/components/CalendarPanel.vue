<script setup lang="ts">
import { computed, ref } from 'vue'
import { locale, t } from '../locales'
import { calendarDateKey, calendarDemoStart, getDemoCalendarEvents, type CalendarCategory, type CalendarEvent } from '../data/calendar-demo'

const selectedDate = ref(new Date(calendarDemoStart.year, calendarDemoStart.month, calendarDemoStart.day))
const visibleMonth = ref(new Date(calendarDemoStart.year, calendarDemoStart.month, 1))
const activeCategory = ref<CalendarCategory | 'all'>('all')
const scrollTop = ref(0)
const reminderKey = 'cf_calendar_demo_reminders'
const reminders = ref<string[]>(readReminders())
const filters: { id: CalendarCategory | 'all'; label: string }[] = [
  { id: 'all', label: '全部' }, { id: 'macro', label: '宏观数据' },
  { id: 'policy', label: '政策监管' }, { id: 'crypto', label: '加密事件' },
]
const weekdays = computed(() => locale.value === 'zh' ? ['一', '二', '三', '四', '五', '六', '日'] : ['M', 'T', 'W', 'T', 'F', 'S', 'S'])
const monthTitle = computed(() => locale.value === 'zh'
  ? `${visibleMonth.value.getFullYear()}年${visibleMonth.value.getMonth() + 1}月`
  : visibleMonth.value.toLocaleDateString('en-US', { month: 'long', year: 'numeric' }))
const events = computed(() => getDemoCalendarEvents(visibleMonth.value.getFullYear(), visibleMonth.value.getMonth()))
const eventDates = computed(() => new Set(events.value.map(event => event.date)))
const importantDates = computed(() => new Set(events.value.filter(event => event.importance === 3).map(event => event.date)))
const selectedKey = computed(() => calendarDateKey(selectedDate.value))
const selectedEvents = computed(() => events.value
  .filter(event => event.date === selectedKey.value && (activeCategory.value === 'all' || event.category === activeCategory.value))
  .sort((a, b) => a.time.localeCompare(b.time)))
const dateHeading = computed(() => {
  const date = selectedDate.value
  const today = calendarDateKey(new Date()) === selectedKey.value
  const label = locale.value === 'zh'
    ? `${date.getMonth() + 1}月${date.getDate()}日 · ${['周日', '周一', '周二', '周三', '周四', '周五', '周六'][date.getDay()]}`
    : date.toLocaleDateString('en-US', { month: 'short', day: 'numeric', weekday: 'short' })
  return `${label}${today ? ` · ${t('今天')}` : ''}`
})
const weeks = computed(() => {
  const year = visibleMonth.value.getFullYear()
  const month = visibleMonth.value.getMonth()
  const offset = (new Date(year, month, 1).getDay() + 6) % 7
  const count = Math.ceil((offset + new Date(year, month + 1, 0).getDate()) / 7)
  return Array.from({ length: count }, (_, week) => Array.from({ length: 7 }, (_, day) => {
    const date = new Date(year, month, 1 - offset + week * 7 + day)
    return { date, key: calendarDateKey(date), day: date.getDate(), outside: date.getMonth() !== month, weekend: date.getDay() === 0 || date.getDay() === 6 }
  }))
})
const weekIndex = computed(() => weeks.value.findIndex(week => week.some(day => day.key === selectedKey.value)))

function readReminders(): string[] {
  try {
    const saved = uni.getStorageSync(reminderKey)
    return Array.isArray(saved) ? saved.filter((id): id is string => typeof id === 'string') : []
  } catch { return [] }
}

function resetScroll() {
  scrollTop.value = scrollTop.value === 0 ? 1 : 0
}

function selectDate(date: Date) {
  selectedDate.value = date
  visibleMonth.value = new Date(date.getFullYear(), date.getMonth(), 1)
  resetScroll()
}

function changeMonth(step: number) {
  const month = new Date(visibleMonth.value.getFullYear(), visibleMonth.value.getMonth() + step, 1)
  const day = Math.min(selectedDate.value.getDate(), new Date(month.getFullYear(), month.getMonth() + 1, 0).getDate())
  selectDate(new Date(month.getFullYear(), month.getMonth(), day))
}

function changeWeek(event: { detail: { current: number; source?: string } }) {
  if (event.detail.source !== 'touch') return
  const week = weeks.value[event.detail.current]
  const weekday = (selectedDate.value.getDay() + 6) % 7
  const next = week[weekday].outside ? week.find(day => !day.outside) : week[weekday]
  if (next) selectDate(next.date)
}

function chooseCategory(category: CalendarCategory | 'all') {
  activeCategory.value = category
  resetScroll()
}

function eventLabel(event: CalendarEvent) {
  return t({ economic: '高影响', unlock: '代币解锁', project: '项目事件', policy: '政策监管' }[event.kind])
}

function toggleReminder(id: string) {
  const marked = reminders.value.includes(id)
  const next = marked ? reminders.value.filter(value => value !== id) : [...reminders.value, id]
  try {
    uni.setStorageSync(reminderKey, next)
    reminders.value = next
    uni.showToast({ title: t(marked ? '已取消本地标记' : '已标记在本机，演示不发送通知'), icon: 'none' })
  } catch {
    uni.showToast({ title: t('本地标记保存失败'), icon: 'none' })
  }
}
</script>

<template>
  <view class="calendar-panel">
    <view class="calendar-picker">
      <view class="month-navigation">
        <button class="month-arrow" :aria-label="t('上个月')" @click="changeMonth(-1)"><svg viewBox="0 0 24 24" aria-hidden="true"><path d="m14 6-6 6 6 6"/></svg></button>
        <text class="month-title">{{ monthTitle }}</text>
        <button class="month-arrow" :aria-label="t('下个月')" @click="changeMonth(1)"><svg viewBox="0 0 24 24" aria-hidden="true"><path d="m10 6 6 6-6 6"/></svg></button>
      </view>
      <view class="weekdays"><text v-for="(day, index) in weekdays" :key="index">{{ day }}</text></view>
      <swiper class="week-swiper" :current="weekIndex" :duration="200" @change="changeWeek">
        <swiper-item v-for="(week, index) in weeks" :key="index">
          <view class="week-dates">
            <button v-for="day in week" :key="day.key" class="date-button" :class="{ selected: day.key === selectedKey, outside: day.outside, weekend: day.weekend }" :aria-label="day.key" :aria-pressed="day.key === selectedKey" @click="selectDate(day.date)">
              <text>{{ day.day }}</text>
              <view v-if="eventDates.has(day.key) && !day.outside" class="event-dot" :class="{ important: importantDates.has(day.key) }"></view>
            </button>
          </view>
        </swiper-item>
      </swiper>
    </view>

    <view class="calendar-filters">
      <button v-for="filter in filters" :key="filter.id" class="category-button" :class="{ active: activeCategory === filter.id }" :aria-pressed="activeCategory === filter.id" @click="chooseCategory(filter.id)">{{ t(filter.label) }}</button>
    </view>

    <scroll-view class="event-scroll" scroll-y :scroll-top="scrollTop" :show-scrollbar="false">
      <view class="selected-date-heading"><text>{{ dateHeading }}</text><text class="demo-label">UTC+8</text></view>
      <view v-if="!selectedEvents.length" class="calendar-empty">
        <svg viewBox="0 0 48 48" aria-hidden="true"><rect x="9" y="12" width="30" height="28" rx="5"/><path d="M16 8v8M32 8v8M9 22h30m-23 8h5m6 0h5"/></svg>
        <text>{{ t('当日暂无此类事件') }}</text>
        <text class="empty-hint">{{ t('切换日期或分类查看其他事件') }}</text>
      </view>
      <view v-for="event in selectedEvents" :key="event.id" class="timeline-row">
        <view class="event-time"><text>{{ event.time }}</text><view class="importance-stars" :aria-label="`${event.importance}/3`"><svg v-for="star in event.importance" :key="star" viewBox="0 0 24 24" aria-hidden="true"><path d="m12 2.5 2.9 5.9 6.5.9-4.7 4.6 1.1 6.5-5.8-3.1-5.8 3.1 1.1-6.5-4.7-4.6 6.5-.9z"/></svg></view></view>
        <view class="event-card" :class="event.kind">
          <view class="event-card-heading">
            <view class="event-title-wrap"><view v-if="event.kind === 'unlock'" class="token-icon">A</view><text class="event-title">{{ locale === 'en' ? event.enTitle : event.title }}</text></view>
            <text class="event-badge">{{ eventLabel(event) }}</text>
          </view>
          <view v-if="event.kind === 'economic'" class="economic-values">
            <view><text>{{ t('前值') }}</text><text>{{ event.previous || '--' }}</text></view>
            <view><text>{{ t('预期') }}</text><text>{{ event.forecast || '--' }}</text></view>
            <view><text>{{ t('公布') }}</text><text :class="{ pending: !event.actual || event.actual === '--' }">{{ event.actual || '--' }}</text></view>
          </view>
          <text v-else class="event-description">{{ locale === 'en' ? event.enDescription : event.description }}</text>
          <button class="reminder-button" :class="{ marked: reminders.includes(event.id) }" :aria-pressed="reminders.includes(event.id)" @click="toggleReminder(event.id)"><svg viewBox="0 0 24 24" aria-hidden="true"><path d="M18 8a6 6 0 0 0-12 0c0 7-3 7-3 9h18c0-2-3-2-3-9M10 21h4"/></svg><text>{{ t(reminders.includes(event.id) ? '已标记' : '提醒我') }}</text></button>
        </view>
      </view>
      <view class="timeline-bottom-space"></view>
    </scroll-view>
  </view>
</template>

<style scoped lang="scss">
.calendar-panel{display:flex;flex:1;flex-direction:column;min-height:0;color:var(--cf-text)}
.calendar-picker{flex:none;padding:12rpx 18rpx 24rpx;border:1rpx solid var(--cf-divider);border-radius:30rpx;background:var(--cf-input)}
.month-navigation{display:flex;align-items:center;justify-content:space-between;height:68rpx}
.month-title{font-size:29rpx;font-weight:700}
.month-arrow{width:48rpx;height:56rpx;margin:0;padding:0;background:transparent;color:var(--cf-soft)}
.month-arrow::after,.date-button::after,.category-button::after,.reminder-button::after{border:0}
.month-arrow svg{width:28rpx;height:28rpx;fill:none;stroke:currentColor;stroke-width:1.8;stroke-linecap:round;stroke-linejoin:round}
.weekdays,.week-dates{display:grid;grid-template-columns:repeat(7,minmax(0,1fr));text-align:center}
.weekdays{align-items:center;height:46rpx;color:var(--cf-soft);font-size:22rpx}
.week-swiper{height:80rpx}
.week-dates{height:100%;align-items:center}
.date-button{position:relative;display:flex;flex-direction:column;width:76rpx;max-width:100%;height:76rpx;margin:0 auto;padding:0 0 5rpx;border-radius:24rpx;background:transparent;color:var(--cf-text);font-size:28rpx;font-weight:500}
.date-button.outside,.date-button.weekend{color:var(--cf-soft)}
.date-button.selected{background:var(--cf-brand);color:#fff;font-weight:700}
.event-dot{position:absolute;bottom:11rpx;left:calc(50% - 4rpx);width:8rpx;height:8rpx;border-radius:50%;background:#60a5fa}
.event-dot.important{background:#f87171}
.date-button.selected .event-dot{background:#fff}
.calendar-filters{display:flex;flex:none;gap:28rpx;min-height:84rpx;margin:0 -40rpx;padding:12rpx 40rpx 0;border-bottom:1rpx solid var(--cf-divider)}
.category-button{position:relative;flex:none;min-height:68rpx;margin:0;padding:0;background:transparent;color:var(--cf-muted);font-size:24rpx;white-space:nowrap}
.category-button.active{color:var(--cf-brand);font-weight:700}
.category-button.active::before{position:absolute;right:0;bottom:0;left:0;height:4rpx;background:var(--cf-brand);content:''}
.event-scroll{flex:1;height:0;min-height:0;width:100%}
.selected-date-heading{display:flex;align-items:center;justify-content:space-between;gap:8rpx;padding:32rpx 0 26rpx;color:var(--cf-soft);font-size:22rpx}
.demo-label{flex:none;font-size:17rpx;color:var(--cf-soft)}
.timeline-row{display:flex;align-items:flex-start;gap:22rpx;margin-bottom:24rpx}
.event-time{display:flex;flex:0 0 96rpx;flex-direction:column;align-items:flex-end;gap:10rpx;padding-top:1rpx;font-family:ui-monospace,SFMono-Regular,Consolas,monospace;font-size:25rpx;font-weight:700}
.importance-stars{display:flex;justify-content:flex-end;gap:1rpx}
.importance-stars svg{width:20rpx;height:20rpx;fill:#fbbf24}
.event-card{flex:1;min-width:0;min-height:188rpx;padding:24rpx 24rpx 22rpx;border:1rpx solid var(--cf-line);border-left:8rpx solid var(--cf-up);border-radius:30rpx;background:var(--cf-card)}
.event-card.economic{min-height:222rpx;border-left-color:var(--cf-down)}
.event-card.unlock{border-left-color:#3b82f6}
.event-card.policy{border-left-color:#8b5cf6}
.event-card-heading{display:flex;align-items:flex-start;justify-content:space-between;gap:10rpx}
.event-title-wrap{display:flex;flex:1;align-items:center;gap:10rpx;min-width:0}
.event-title{font-size:27rpx;font-weight:700;line-height:1.45;overflow-wrap:anywhere}
.token-icon{display:flex;align-items:center;justify-content:center;flex:0 0 40rpx;width:40rpx;height:40rpx;border-radius:50%;background:#06b6d4;color:#fff;font-size:25rpx}
.event-badge{flex:none;margin-top:2rpx;padding:5rpx 9rpx;border-radius:5rpx;background:#ecfdf5;color:#10b981;font-size:17rpx;line-height:1.4}
.economic .event-badge{background:#fef2f2;color:#ef4444}
.unlock .event-badge{background:#eff6ff;color:#3b82f6}
.policy .event-badge{background:#f5f3ff;color:#8b5cf6}
.economic-values{display:grid;grid-template-columns:repeat(3,minmax(0,1fr));gap:8rpx;margin:22rpx 0 18rpx;text-align:center}
.economic-values>view{display:flex;flex-direction:column;gap:8rpx}
.economic-values>view>text:first-child{color:var(--cf-soft);font-size:19rpx}
.economic-values>view>text:last-child{font-family:ui-monospace,SFMono-Regular,Consolas,monospace;font-size:25rpx;font-weight:700}
.economic-values .pending{color:var(--cf-line)}
.event-description{display:block;margin:15rpx 0 18rpx;color:var(--cf-muted);font-size:22rpx;line-height:1.7;overflow-wrap:anywhere}
.reminder-button{display:flex;align-items:center;justify-content:flex-start!important;gap:5rpx;width:max-content;max-width:100%;min-height:32rpx;margin:0;padding:0;background:transparent;color:var(--cf-brand);font-size:18rpx}
.reminder-button svg{width:19rpx;height:19rpx;fill:none;stroke:currentColor;stroke-width:1.6;stroke-linecap:round;stroke-linejoin:round}
.reminder-button.marked{color:var(--cf-up)}
.timeline-bottom-space{height:32rpx}
.calendar-empty{display:flex;flex-direction:column;align-items:center;gap:18rpx;padding:70rpx 12rpx;color:var(--cf-muted);font-size:25rpx;text-align:center}
.calendar-empty svg{width:72rpx;height:72rpx;margin-bottom:4rpx;fill:none;stroke:var(--cf-soft);stroke-width:1.8;stroke-linecap:round}
.empty-hint{color:var(--cf-soft);font-size:21rpx}
@media(max-width:360px){.calendar-filters{gap:22rpx}.event-card{padding-right:18rpx;padding-left:18rpx}.event-title{font-size:25rpx}.timeline-row{gap:16rpx}.event-time{flex-basis:84rpx}.event-badge{font-size:16rpx}}
</style>
