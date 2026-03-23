<template>
  <div>
  <el-dialog
      v-model="visible"
      title="物资使用记录"
      width="900px"
  >

    <div style="margin-bottom:15px">

      <el-button type="primary" @click="openAdd">
        新增
      </el-button>

    </div>

    <!-- 列表 -->

    <el-table
        :data="tableData"
        v-loading="loading"
    >

      <el-table-column prop="materialName" label="物料"/>

      <el-table-column prop="operationType" label="操作类型"/>

      <el-table-column prop="amount" label="数量"/>

      <el-table-column prop="useTime" label="使用时间"/>

      <el-table-column prop="remark" label="备注"/>

      <el-table-column label="操作" width="180">

        <template #default="scope">

          <el-button
              text
              type="primary"
              @click="openEdit(scope.row)"
          >
            编辑
          </el-button>

          <el-button
              text
              type="danger"
              @click="onDelete(scope.row)"
          >
            删除
          </el-button>

        </template>

      </el-table-column>

    </el-table>


    <!-- 新增/编辑 -->

    <el-dialog
        v-model="formVisible"
        :title="isEdit ? '编辑物资使用' : '新增物资使用'"
        width="450px"
    >

      <el-form :model="form" label-width="100px">

        <el-form-item label="物料">

          <el-select
              v-model="form.materialId"
              filterable
              placeholder="搜索物料"
              style="width:100%"
          >

            <el-option
                v-for="item in materialOptions"
                :key="item.id"
                :label="item.name"
                :value="item.id"
            />

          </el-select>

        </el-form-item>

        <el-form-item label="操作类型">
          <el-input v-model="form.operationType"/>
        </el-form-item>

        <el-form-item label="数量">
          <el-input-number v-model="form.amount" :min="0"/>
        </el-form-item>

        <el-form-item label="使用时间">

          <el-date-picker
              v-model="form.useTime"
              type="datetime"
              value-format="YYYY-MM-DD HH:mm:ss"
              style="width:100%"
          />

        </el-form-item>

        <el-form-item label="备注">
          <el-input v-model="form.remark"/>
        </el-form-item>

      </el-form>

      <template #footer>

        <el-button @click="formVisible=false">
          取消
        </el-button>

        <el-button type="primary" @click="submit">
          确定
        </el-button>

      </template>

    </el-dialog>

  </el-dialog>
  </div>
</template>


<script setup lang="ts">

import {ref,reactive,onMounted} from "vue"
import {ElMessage,ElMessageBox} from "element-plus"

import {
  getMaterialUseListApi,
  addMaterialUseApi,
  updateMaterialUseApi,
  deleteMaterialUseApi
} from "/@/api/system/field_crops"

import {getMaterialListApi} from "/@/api/system/material"

/* 弹窗 */

const visible = ref(false)

const formVisible = ref(false)

/* 列表 */

const loading = ref(false)

const tableData = ref([])

/* 当前种植记录 */

const fieldCropId = ref(0)

/* 是否编辑 */

const isEdit = ref(false)

/* 物料 */

const materialOptions = ref([])

/* 表单 */

const form = reactive({

  id:0,
  fieldCropId:0,
  materialId:"",
  operationType:"",
  amount:0,
  useTime:"",
  remark:""

})

/* 打开弹窗 */

const open = (id:number)=>{

  visible.value=true

  fieldCropId.value=id

  getList()

}

/* 获取列表 */

const getList = async()=>{

  loading.value=true

  const res = await getMaterialUseListApi({
    fieldCropId:fieldCropId.value
  })

  tableData.value=res.data.list || []

  loading.value=false

}

/* 获取物料 */

const getMaterial = async()=>{

  const res = await getMaterialListApi({
    pageNum:1,
    pageSize:1000
  })

  materialOptions.value=res.data.list || []

}

/* 新增 */

const openAdd = ()=>{

  isEdit.value=false

  form.id=0
  form.fieldCropId=fieldCropId.value
  form.materialId=""
  form.operationType=""
  form.amount=0
  form.useTime=""
  form.remark=""

  formVisible.value=true

}

/* 编辑 */

const openEdit = (row:any)=>{

  isEdit.value=true

  Object.assign(form,row)

  form.fieldCropId = fieldCropId.value

  formVisible.value=true

}

/* 删除 */

const onDelete = (row:any)=>{

  ElMessageBox.confirm("确定删除？","提示")
      .then(async()=>{

        await deleteMaterialUseApi({id:row.id})

        ElMessage.success("删除成功")

        getList()

      })

}

/* 提交 */

const submit = async()=>{

  if(isEdit.value){

    await updateMaterialUseApi(form)

    ElMessage.success("修改成功")

  }else{

    await addMaterialUseApi(form)

    ElMessage.success("新增成功")

  }

  formVisible.value=false

  getList()

}

onMounted(()=>{

  getMaterial()

})

defineExpose({
  open
})

</script>