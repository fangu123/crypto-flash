import request from '@/utils/request'

// 查询InformationNotificationUser列表
export function listInformationNotificationUser(query) {
    return request({
        url: '/api/v1/information-notification-user',
        method: 'get',
        params: query
    })
}

// 查询InformationNotificationUser详细
export function getInformationNotificationUser (id) {
    return request({
        url: '/api/v1/information-notification-user/' + id,
        method: 'get'
    })
}


// 新增InformationNotificationUser
export function addInformationNotificationUser(data) {
    return request({
        url: '/api/v1/information-notification-user',
        method: 'post',
        data: data
    })
}

// 修改InformationNotificationUser
export function updateInformationNotificationUser(data) {
    return request({
        url: '/api/v1/information-notification-user/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除InformationNotificationUser
export function delInformationNotificationUser(data) {
    return request({
        url: '/api/v1/information-notification-user',
        method: 'delete',
        data: data
    })
}

