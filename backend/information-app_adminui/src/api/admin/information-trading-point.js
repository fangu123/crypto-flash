import request from '@/utils/request'

// 查询InformationTradingPoint列表
export function listInformationTradingPoint(query) {
    return request({
        url: '/api/v1/information-trading-point',
        method: 'get',
        params: query
    })
}

// 查询InformationTradingPoint详细
export function getInformationTradingPoint (id) {
    return request({
        url: '/api/v1/information-trading-point/' + id,
        method: 'get'
    })
}


// 新增InformationTradingPoint
export function addInformationTradingPoint(data) {
    return request({
        url: '/api/v1/information-trading-point',
        method: 'post',
        data: data
    })
}

// 修改InformationTradingPoint
export function updateInformationTradingPoint(data) {
    return request({
        url: '/api/v1/information-trading-point/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除InformationTradingPoint
export function delInformationTradingPoint(data) {
    return request({
        url: '/api/v1/information-trading-point',
        method: 'delete',
        data: data
    })
}

