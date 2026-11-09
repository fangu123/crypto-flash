import request from '@/utils/request'

// 查询InformationCoin列表
export function listInformationCoin(query) {
    return request({
        url: '/api/v1/information-coin',
        method: 'get',
        params: query
    })
}

// 查询InformationCoin详细
export function getInformationCoin (id) {
    return request({
        url: '/api/v1/information-coin/' + id,
        method: 'get'
    })
}


// 新增InformationCoin
export function addInformationCoin(data) {
    return request({
        url: '/api/v1/information-coin',
        method: 'post',
        data: data
    })
}

// 修改InformationCoin
export function updateInformationCoin(data) {
    return request({
        url: '/api/v1/information-coin/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除InformationCoin
export function delInformationCoin(data) {
    return request({
        url: '/api/v1/information-coin',
        method: 'delete',
        data: data
    })
}

