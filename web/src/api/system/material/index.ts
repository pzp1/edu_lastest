import request from "/@/utils/request"

export function getMaterialListApi(params) {
    return request({
        url: "/api/v1/system/material/list",
        method: "get",
        params
    })
}

export function addOrUpdateMaterialApi(data) {
    return request({
        url: "/api/v1/system/material/addOrUpdate",
        method: "post",
        data
    })
}

export function deleteMaterialApi(data) {
    return request({
        url: "/api/v1/system/material/delete",
        method: "post",
        data
    })
}