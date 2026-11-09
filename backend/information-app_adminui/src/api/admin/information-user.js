import request from '@/utils/request'

// 查询InformationUser列表
export function listInformationUser(query) {
    return request({
        url: '/api/v1/information-user',
        method: 'get',
        params: query
    })
}

// 查询InformationUser详细
export function getInformationUser (id) {
    return request({
        url: '/api/v1/information-user/' + id,
        method: 'get'
    })
}


// 新增InformationUser
export function addInformationUser(data) {
    return request({
        url: '/api/v1/information-user',
        method: 'post',
        data: data
    })
}

// 修改InformationUser
export function updateInformationUser(data) {
    return request({
        url: '/api/v1/information-user/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除InformationUser
export function delInformationUser(data) {
    return request({
        url: '/api/v1/information-user',
        method: 'delete',
        data: data
    })
}

