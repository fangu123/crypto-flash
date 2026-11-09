<template>
  <div class="chart-wrapper">
    <div class="chart-title">{{ title }}</div>
    <div class="chart-content" ref="chart"></div>
  </div>
</template>

<script>
import * as echarts from 'echarts'

export default {
  name: 'LineChart',
  props: {
    title: {
      type: String,
      default: ''
    },
    data: {
      type: Array,
      default: () => []
    }
  },
  watch: {
    data: {
      handler(newVal) {
        this.initChart(newVal)
      },
      deep: true
    }
  },
  mounted() {
    this.$nextTick(() => {
      this.initChart(this.data)
    })
  },
  methods: {
    initChart(data) {
      if (!this.$refs.chart) return

      // 销毁之前的图表实例
      if (this.chart) {
        this.chart.dispose()
      }

      // 初始化图表
      this.chart = echarts.init(this.$refs.chart)

      // 设置图表配置
      const option = {
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'shadow'
          }
        },
        grid: {
          left: '3%',
          right: '4%',
          bottom: '20%',
          top: '10%',
          containLabel: true
        },
        xAxis: {
          type: 'category',
          data: data.map(item => item.name),
          axisLine: {
            lineStyle: {
              color: '#999',
              width: 2
            }
          },
          // 确保X轴日期间隔正常显示
          axisLabel: {
            interval: 0,
            rotate: 30,
            formatter: function(value) {
              return value
            },
            margin: 15,
            show: true,
            textStyle: {
              color: '#333',
              fontSize: 12
            },
            // 强制显示所有标签
            showMaxLabel: true,
            showMinLabel: true,
            // 增加标签距离轴线的距离
            verticalAlign: 'top',
            position: 'top'
          },
          // 修改boundaryGap为false，使标签显示在数据点下方
          boundaryGap: false,
          axisTick: {
            show: true,
            alignWithLabel: true,
            interval: 0,
            length: 8
          },
          // 设置最小间隔，确保每个日期有足够空间
          minInterval: 1,
          // 设置X轴的类目间隔，确保布满容器
          splitNumber: Math.min(10, data.length),
          // 强制设置坐标轴样式
          axisPointer: {
            type: 'shadow'
          }
        },
        yAxis: {
          type: 'value',
          axisLine: {
            lineStyle: {
              color: '#999'
            }
          },
          splitLine: {
            lineStyle: {
              color: '#eee'
            }
          }
        },
        series: [{
          name: this.title,
          type: 'line',
          data: data.map(item => item.value),
          smooth: true,
          symbol: 'circle',
          symbolSize: 8,
          lineStyle: {
            width: 5
          },
          areaStyle: {
            opacity: 0.1
          }
        }]
      }

      // 设置配置项并渲染图表
      this.chart.setOption(option)

      // 监听窗口大小变化，调整图表大小
      window.addEventListener('resize', this.handleResize)
    },
    handleResize() {
      if (this.chart) {
        this.chart.resize()
      }
    }
  },
  beforeDestroy() {
    // 移除事件监听
    window.removeEventListener('resize', this.handleResize)
    // 销毁图表实例
    if (this.chart) {
      this.chart.dispose()
    }
  }
}
</script>

<style lang="scss" scoped>
.chart-wrapper {
  background: #fff;
  padding: 16px;
  margin-bottom: 32px;

  .chart-title {
    font-size: 16px;
    font-weight: 500;
    margin-bottom: 12px;
    color: #333;
  }

  .chart-content {
    height: 300px;
    width: 100%;
  }
}
</style>
