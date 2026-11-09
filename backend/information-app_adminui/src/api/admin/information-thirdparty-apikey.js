import request from '@/utils/request'

// 查询InformationThirdpartyApikey列表
export function listInformationThirdpartyApikey(query) {
    return request({
        url: '/api/v1/information-thirdparty-apikey',
        method: 'get',
        params: query
    })
}

// 查询InformationThirdpartyApikey详细
export function getInformationThirdpartyApikey (id) {
    return request({
        url: '/api/v1/information-thirdparty-apikey/' + id,
        method: 'get'
    })
}


// 新增InformationThirdpartyApikey
export function addInformationThirdpartyApikey(data) {
    return request({
        url: '/api/v1/information-thirdparty-apikey',
        method: 'post',
        data: data
    })
}

// 修改InformationThirdpartyApikey
export function updateInformationThirdpartyApikey(data) {
    return request({
        url: '/api/v1/information-thirdparty-apikey/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除InformationThirdpartyApikey
export function delInformationThirdpartyApikey(data) {
    return request({
        url: '/api/v1/information-thirdparty-apikey',
        method: 'delete',
        data: data
    })
}

