import request from '@/utils/request'

// 查询InformationAppVersion列表
export function listInformationAppVersion(query) {
    return request({
        url: '/api/v1/information-app-version',
        method: 'get',
        params: query
    })
}

// 查询InformationAppVersion详细
export function getInformationAppVersion (id) {
    return request({
        url: '/api/v1/information-app-version/' + id,
        method: 'get'
    })
}


// 新增InformationAppVersion
export function addInformationAppVersion(data) {
    return request({
        url: '/api/v1/information-app-version',
        method: 'post',
        data: data
    })
}

// 修改InformationAppVersion
export function updateInformationAppVersion(data) {
    return request({
        url: '/api/v1/information-app-version/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除InformationAppVersion
export function delInformationAppVersion(data) {
    return request({
        url: '/api/v1/information-app-version',
        method: 'delete',
        data: data
    })
}

