import request from "/@/utils/request"

/**
 * 虫灾列表
 */
export function getPestListApi(params:any) {
    return request({
        url: "/api/v1/system/pest/list",
        method: "get",
        params
    })
}

/**
 * 新增虫灾（上报）
 */
export function addPestApi(data:any) {
    return request({
        url: "/api/v1/system/pest/add",
        method: "post",
        data
    })
}

/**
 * 修改虫灾（编辑描述）
 */
export function updatePestApi(data:any) {
    return request({
        url: "/api/v1/system/pest/updateStatus",
        method: "post",
        data
    })
}

/**
 * 删除虫灾
 */
export function deletePestApi(params:any) {
    return request({
        url: "/api/v1/system/pest/delete",
        method: "get",
        params
    })
}

/**
 * 修改状态（解决 / 未解决）
 */
export function updatePestStatusApi(data:any) {
    return request({
        url: "/api/v1/system/pest/updateStatus",
        method: "post",
        data
    })
}