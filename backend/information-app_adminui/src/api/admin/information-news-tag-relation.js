import request from '@/utils/request'

// 查询InformationNewsTagRelation列表
export function listInformationNewsTagRelation(query) {
    return request({
        url: '/api/v1/information-news-tag-relation',
        method: 'get',
        params: query
    })
}

// 查询InformationNewsTagRelation详细
export function getInformationNewsTagRelation (id) {
    return request({
        url: '/api/v1/information-news-tag-relation/' + id,
        method: 'get'
    })
}


// 新增InformationNewsTagRelation
export function addInformationNewsTagRelation(data) {
    return request({
        url: '/api/v1/information-news-tag-relation',
        method: 'post',
        data: data
    })
}

// 修改InformationNewsTagRelation
export function updateInformationNewsTagRelation(data) {
    return request({
        url: '/api/v1/information-news-tag-relation/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除InformationNewsTagRelation
export function delInformationNewsTagRelation(data) {
    return request({
        url: '/api/v1/information-news-tag-relation',
        method: 'delete',
        data: data
    })
}

