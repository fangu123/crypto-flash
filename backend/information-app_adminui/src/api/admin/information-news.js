import request from '@/utils/request'

// 查询InformationNews列表
export function listInformationNews(query) {
    return request({
        url: '/api/v1/information-news',
        method: 'get',
        params: query
    })
}

// 查询InformationNews详细
export function getInformationNews (id) {
    return request({
        url: '/api/v1/information-news/' + id,
        method: 'get'
    })
}


// 新增InformationNews
export function addInformationNews(data) {
    return request({
        url: '/api/v1/information-news',
        method: 'post',
        data: data
    })
}

// 修改InformationNews
export function updateInformationNews(data) {
    return request({
        url: '/api/v1/information-news/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除InformationNews
export function delInformationNews(data) {
    return request({
        url: '/api/v1/information-news',
        method: 'delete',
        data: data
    })
}

