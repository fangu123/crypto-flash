import request from '@/utils/request'

// 查询InformationNotification列表
export function listInformationNotification(query) {
    return request({
        url: '/api/v1/information-notification',
        method: 'get',
        params: query
    })
}

// 查询InformationNotification详细
export function getInformationNotification (id) {
    return request({
        url: '/api/v1/information-notification/' + id,
        method: 'get'
    })
}


// 新增InformationNotification
export function addInformationNotification(data) {
    return request({
        url: '/api/v1/information-notification',
        method: 'post',
        data: data
    })
}

// 修改InformationNotification
export function updateInformationNotification(data) {
    return request({
        url: '/api/v1/information-notification/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除InformationNotification
export function delInformationNotification(data) {
    return request({
        url: '/api/v1/information-notification',
        method: 'delete',
        data: data
    })
}

