import request from '@/utils/request'

// 查询InformationNewsFavorite列表
export function listInformationNewsFavorite(query) {
    return request({
        url: '/api/v1/information-news-favorite',
        method: 'get',
        params: query
    })
}

// 查询InformationNewsFavorite详细
export function getInformationNewsFavorite (id) {
    return request({
        url: '/api/v1/information-news-favorite/' + id,
        method: 'get'
    })
}


// 新增InformationNewsFavorite
export function addInformationNewsFavorite(data) {
    return request({
        url: '/api/v1/information-news-favorite',
        method: 'post',
        data: data
    })
}

// 修改InformationNewsFavorite
export function updateInformationNewsFavorite(data) {
    return request({
        url: '/api/v1/information-news-favorite/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除InformationNewsFavorite
export function delInformationNewsFavorite(data) {
    return request({
        url: '/api/v1/information-news-favorite',
        method: 'delete',
        data: data
    })
}

