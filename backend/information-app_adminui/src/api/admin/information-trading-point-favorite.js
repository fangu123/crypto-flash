import request from '@/utils/request'

// 查询InformationTradingPointFavorite列表
export function listInformationTradingPointFavorite(query) {
    return request({
        url: '/api/v1/information-trading-point-favorite',
        method: 'get',
        params: query
    })
}

// 查询InformationTradingPointFavorite详细
export function getInformationTradingPointFavorite (id) {
    return request({
        url: '/api/v1/information-trading-point-favorite/' + id,
        method: 'get'
    })
}


// 新增InformationTradingPointFavorite
export function addInformationTradingPointFavorite(data) {
    return request({
        url: '/api/v1/information-trading-point-favorite',
        method: 'post',
        data: data
    })
}

// 修改InformationTradingPointFavorite
export function updateInformationTradingPointFavorite(data) {
    return request({
        url: '/api/v1/information-trading-point-favorite/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除InformationTradingPointFavorite
export function delInformationTradingPointFavorite(data) {
    return request({
        url: '/api/v1/information-trading-point-favorite',
        method: 'delete',
        data: data
    })
}

