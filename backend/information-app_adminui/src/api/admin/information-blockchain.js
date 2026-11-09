import request from '@/utils/request'

// 查询InformationBlockchain列表
export function listInformationBlockchain(query) {
    return request({
        url: '/api/v1/information-blockchain',
        method: 'get',
        params: query
    })
}

// 查询InformationBlockchain详细
export function getInformationBlockchain (id) {
    return request({
        url: '/api/v1/information-blockchain/' + id,
        method: 'get'
    })
}


// 新增InformationBlockchain
export function addInformationBlockchain(data) {
    return request({
        url: '/api/v1/information-blockchain',
        method: 'post',
        data: data
    })
}

// 修改InformationBlockchain
export function updateInformationBlockchain(data) {
    return request({
        url: '/api/v1/information-blockchain/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除InformationBlockchain
export function delInformationBlockchain(data) {
    return request({
        url: '/api/v1/information-blockchain',
        method: 'delete',
        data: data
    })
}

