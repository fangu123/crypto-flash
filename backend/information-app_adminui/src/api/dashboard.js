import request from '@/utils/request'

// 获取统计数据
export function getStatistics() {
  return request({
    url: '/api/v1/information-statistics',
    method: 'get'
  })
}
