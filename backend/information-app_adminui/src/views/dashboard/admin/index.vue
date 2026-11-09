<template>
  <div class="dashboard-editor-container">
    <el-row :gutter="12">
      <!-- 快讯统计卡片 -->
      <el-col :sm="24" :xs="24" :md="6" :xl="6" :lg="6" :style="{ marginBottom: '12px' }">
        <chart-card title="快讯总数量" :total="statistics.news.total">
          <el-tooltip slot="action" class="item" effect="dark" content="快讯数据统计" placement="top-start">
            <i class="el-icon-warning-outline" />
          </el-tooltip>
<!--          <div>
            <trend v-if="statistics.news.today" flag="top" style="margin-right: 16px;" :rate="calculateRate(statistics.news.today, 0)">
              <span slot="term">日新增</span>
            </trend>
            <mini-area v-if="statistics.news.trend && statistics.news.trend.length" :data="statistics.news.trend" />
          </div>-->
          <template slot="footer">日快讯数量<span>{{ statistics.news.today || '0' }}</span></template>
        </chart-card>
      </el-col>

      <!-- 点位统计卡片 -->
      <el-col :sm="24" :xs="24" :md="6" :xl="6" :lg="6" :style="{ marginBottom: '12px' }">
        <chart-card title="点位总数量" :total="statistics.tradingPoint.total">
          <el-tooltip slot="action" class="item" effect="dark" content="点位数据统计" placement="top-start">
            <i class="el-icon-warning-outline" />
          </el-tooltip>
<!--          <div>-->
<!--            <trend v-if="statistics.tradingPoint.today" flag="top" style="margin-right: 16px;" :rate="calculateRate(statistics.tradingPoint.today, 0)">-->
<!--              <span slot="term">日新增</span>-->
<!--            </trend>-->
<!--            <mini-area v-if="statistics.tradingPoint.trend && statistics.tradingPoint.trend.length" :data="statistics.tradingPoint.trend" />-->
<!--          </div>-->
          <template slot="footer">日点位数量<span>{{ statistics.tradingPoint.today || '0' }}</span></template>
        </chart-card>
      </el-col>

      <!-- 用户统计卡片 -->
      <el-col :sm="24" :xs="24" :md="6" :xl="6" :lg="6" :style="{ marginBottom: '12px' }">
        <chart-card title="用户总数量" :total="statistics.user.total">
          <el-tooltip slot="action" class="item" effect="dark" content="用户数据统计" placement="top-start">
            <i class="el-icon-warning-outline" />
          </el-tooltip>
<!--          <div>-->
<!--            <trend v-if="statistics.user.today" flag="top" style="margin-right: 16px;" :rate="calculateRate(statistics.user.today, 0)">-->
<!--              <span slot="term">日新增</span>-->
<!--            </trend>-->
<!--          </div>-->
          <template slot="footer">日用户数量<span>{{ statistics.user.today || '0' }}</span></template>
        </chart-card>
      </el-col>

      <!-- 阅读统计卡片 -->
      <el-col :sm="24" :xs="24" :md="6" :xl="6" :lg="6" :style="{ marginBottom: '12px' }">
        <chart-card title="阅读总数量" :total="statistics.readCount.total">
          <el-tooltip slot="action" class="item" effect="dark" content="阅读数据统计" placement="top-start">
            <i class="el-icon-warning-outline" />
          </el-tooltip>
<!--          <div>-->
<!--            <trend v-if="statistics.readCount.today" flag="top" style="margin-right: 16px;" :rate="calculateRate(statistics.readCount.today, 0)">-->
<!--              <span slot="term">日新增</span>-->
<!--            </trend>-->
<!--          </div>-->
          <template slot="footer">日阅读数量<span>{{ statistics.readCount.today || '0' }}</span></template>
        </chart-card>
      </el-col>
    </el-row>
Q
    <!-- 数据趋势图表 -->
    <el-card :bordered="false" :body-style="{padding: '0'}" style="margin-top: 12px;">
      <div class="salesCard">
        <!-- 快讯趋势 -->
        <el-row style="margin-bottom: 16px;">
          <el-col :xl="24" :lg="24" :md="24" :sm="24" :xs="24">
            <line-chart :data="newsTrendData" title="近7日快讯趋势" />
          </el-col>
        </el-row>
        <!-- 点位趋势 -->
        <el-row>
          <el-col :xl="24" :lg="24" :md="24" :sm="24" :xs="24">
            <line-chart :data="tradingPointTrendData" title="近7日点位趋势" />
          </el-col>
        </el-row>
      </div>
    </el-card>

  </div>
</template>

<script>
import ChartCard from '@/components/ChartCard'
import Trend from '@/components/Trend'
import MiniArea from '@/components/MiniArea'
import MiniBar from '@/components/MiniBar'
import MiniProgress from '@/components/MiniProgress'
import RankList from '@/components/RankList/index'
import Bar from '@/components/Bar.vue'
import LineChart from '@/components/LineChart.vue'
import { getStatistics } from '@/api/dashboard'

export default {
  name: 'DashboardAdmin',
  components: {
    ChartCard,
    Trend,
    MiniArea,
    MiniBar,
    MiniProgress,
    RankList,
    Bar,
    LineChart
  },
  data() {
    return {
      statistics: {
        news: { total: 0, today: 0, trend: [] },
        tradingPoint: { total: 0, today: 0, trend: [] },
        user: { total: 0, today: 0 },
        readCount: { total: 0, today: 0 }
      },
      newsTrendData: [],
      tradingPointTrendData: []
    }
  },
  created() {
    this.fetchStatistics()
  },
  methods: {
    async fetchStatistics() {
      try {
        const response = await getStatistics()
        if (response && response.data) {
          this.statistics = response.data

          // 处理趋势数据
          if (this.statistics.news.trend && this.statistics.news.trend.length) {
            this.newsTrendData = this.statistics.news.trend.map(item => ({
              name: item.date,
              value: item.count
            }))
          }

          if (this.statistics.tradingPoint.trend && this.statistics.tradingPoint.trend.length) {
            this.tradingPointTrendData = this.statistics.tradingPoint.trend.map(item => ({
              name: item.date,
              value: item.count
            }))
          }
        }
      } catch (error) {
        console.error('获取统计数据失败:', error)
        this.$message.error('获取统计数据失败')
      }
    },

    calculateRate(today, yesterday) {
      // 计算增长率，如果没有昨天的数据则返回0
      if (!yesterday) return 0
      return Math.round(((today - yesterday) / yesterday) * 100)
    }
  }
}
</script>

<style lang="scss" scoped>
.dashboard-editor-container {
  padding: 12px;
  background-color: rgb(240, 242, 245);
  position: relative;

  .github-corner {
    position: absolute;
    top: 0;
    border: 0;
    right: 0;
  }

  .chart-wrapper {
    background: #fff;
    padding: 16px 16px 0;
    margin-bottom: 32px;
  }
}

::v-deep .el-tabs__item{
  padding-left: 16px!important;
  height: 50px;
  line-height: 50px;
}

@media (max-width:1024px) {
  .chart-wrapper {
    padding: 8px;
  }
}
</style>
