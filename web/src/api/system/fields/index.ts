import request from '/@/utils/request';

export function  getFieldListApi(query:Object){
    return request({
        url: '/api/v1/system/field/list',
        method: 'get',
        params:query,
    })
}



export function deleteFieldApi(id:number){
    return request({
        url: '/api/v1/system/field/delete',
        method: 'delete',
        data:id
    })
}

export function addFieldApi(data: any) {
    return request({
        url: '/api/v1/system/field/addOrUp',
        method: 'post',
        data
    })
}

export function updateFieldApi(data: any) {
    return request({
        url: '/api/v1/system/field/addOrUp',
        method: 'post',
        data
    })
}

export function getUserList(query:Object) {
    return request({
        url: '/api/v1/system/user/list',
        method: 'get',
        params:query
    })
}