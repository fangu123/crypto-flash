import request from '@/utils/request'

// 查询InformationComment列表
export function listInformationComment(query) {
    return request({
        url: '/api/v1/information-comment',
        method: 'get',
        params: query
    })
}

// 查询InformationComment详细
export function getInformationComment (id) {
    return request({
        url: '/api/v1/information-comment/' + id,
        method: 'get'
    })
}


// 新增InformationComment
export function addInformationComment(data) {
    return request({
        url: '/api/v1/information-comment',
        method: 'post',
        data: data
    })
}

// 修改InformationComment
export function updateInformationComment(data) {
    return request({
        url: '/api/v1/information-comment/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除InformationComment
export function delInformationComment(data) {
    return request({
        url: '/api/v1/information-comment',
        method: 'delete',
        data: data
    })
}

