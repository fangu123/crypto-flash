import request from '@/utils/request'

// 查询InformationCustomerService列表
export function listInformationCustomerService(query) {
    return request({
        url: '/api/v1/information-customer-service',
        method: 'get',
        params: query
    })
}

// 查询InformationCustomerService详细
export function getInformationCustomerService (id) {
    return request({
        url: '/api/v1/information-customer-service/' + id,
        method: 'get'
    })
}


// 新增InformationCustomerService
export function addInformationCustomerService(data) {
    return request({
        url: '/api/v1/information-customer-service',
        method: 'post',
        data: data
    })
}

// 修改InformationCustomerService
export function updateInformationCustomerService(data) {
    return request({
        url: '/api/v1/information-customer-service/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除InformationCustomerService
export function delInformationCustomerService(data) {
    return request({
        url: '/api/v1/information-customer-service',
        method: 'delete',
        data: data
    })
}

