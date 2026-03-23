import request from '/@/utils/request'

// 获取列表
export function getCropListApi(query) {
    return request({
        url: '/api/v1/system/crop/list',
        method: 'get',
        params: query
    })
}

// 新增
export function addCropApi(data) {
    return request({
        url: '/api/v1/system/crop/addOrUp',
        method: 'post',
        data
    })
}

// 新增
export function updateCropApi(data) {
    return request({
        url: '/api/v1/system/crop/addOrUp',
        method: 'post',
        data
    })
}

// 删除
export function deleteCropApi(id) {
    return request({
        url: '/api/v1/system/crop/delete',
        method: 'delete',
        data:id
    })
}