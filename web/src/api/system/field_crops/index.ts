import request from "/@/utils/request"


/* 列表 */
export function getFieldCropListApi(params:any){
    return request({
        url:"/api/v1/system/fieldCrop/list",
        method:"get",
        params
    })
}


/* 详情 */
export function getFieldCropDetailApi(params:any){
    return request({
        url:"/api/v1/system/fieldCrop/detail",
        method:"get",
        params
    })
}


/* 新增/修改 */
export function addFieldCropApi(data:any){
    return request({
        url:"/api/v1/system/fieldCrop/addOrUpdate",
        method:"post",
        data
    })
}


/* 删除 */
export function deleteFieldCropApi(data:any){
    return request({
        url:"/api/v1/system/fieldCrop/delete",
        method:"post",
        data
    })
}

/* 列表 */

export function getMaterialUseListApi(data:any){
    return request({
        url: "/api/v1/system/materialUse/list",
        method: "get",
        params: data
    })
}

/* 新增 */

export function addMaterialUseApi(data:any){
    return request({
        url: "/api/v1/system/materialUse/add",
        method: "post",
        data
    })
}

/* 修改 */

export function updateMaterialUseApi(data:any){
    return request({
        url: "/api/v1/system/materialUse/update",
        method: "post",
        data
    })
}

/* 删除 */

export function deleteMaterialUseApi(data:any){
    return request({
        url: "/api/v1/system/materialUse/delete",
        method: "post",
        data
    })
}