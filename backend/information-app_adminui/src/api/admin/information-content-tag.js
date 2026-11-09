import request from '@/utils/request'

// 查询InformationContentTag列表
export function listInformationContentTag(query) {
    return request({
        url: '/api/v1/information-content-tag',
        method: 'get',
        params: query
    })
}

// 查询InformationContentTag详细
export function getInformationContentTag (id) {
    return request({
        url: '/api/v1/information-content-tag/' + id,
        method: 'get'
    })
}


// 新增InformationContentTag
export function addInformationContentTag(data) {
    return request({
        url: '/api/v1/information-content-tag',
        method: 'post',
        data: data
    })
}

// 修改InformationContentTag
export function updateInformationContentTag(data) {
    return request({
        url: '/api/v1/information-content-tag/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除InformationContentTag
export function delInformationContentTag(data) {
    return request({
        url: '/api/v1/information-content-tag',
        method: 'delete',
        data: data
    })
}

