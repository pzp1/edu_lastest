<template>
  <div class="system-fieldCrop-container">
    <el-card shadow="hover">

      <!-- 查询 -->
      <div class="search mb15">
        <el-form :inline="true">

          <el-form-item label="作物">
            <el-select v-model="tableData.param.cropId" clearable style="width:160px">
              <el-option v-for="item in cropOptions" :key="item.id" :label="item.name" :value="item.id"/>
            </el-select>
          </el-form-item>

          <el-form-item label="地块">
            <el-select v-model="tableData.param.fieldId" clearable style="width:160px">
              <el-option v-for="item in fieldOptions" :key="item.id" :label="item.name" :value="item.id"/>
            </el-select>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" @click="onSearch">查询</el-button>
            <el-button type="success" @click="openAdd">新增种植记录</el-button>
          </el-form-item>

        </el-form>
      </div>

      <!-- 表格 -->
      <el-table :data="tableData.data" v-loading="loading">

        <el-table-column label="序号" width="70">
          <template #default="scope">
            {{
              (tableData.param.pageNum - 1) * tableData.param.pageSize
              + scope.$index + 1
            }}
          </template>
        </el-table-column>

        <el-table-column prop="fieldName" label="地块"/>
        <el-table-column prop="cropName" label="作物"/>
        <el-table-column prop="plantDate" label="种植日期"/>
        <el-table-column prop="expectedHarvestDate" label="预计收获"/>

        <el-table-column label="实际收获">
          <template #default="scope">
            <span v-if="scope.row.status === 2">
              {{ scope.row.harvestDate }}
            </span>
            <span v-else>-</span>
          </template>
        </el-table-column>

        <el-table-column prop="currentStage" label="生长阶段"/>

        <el-table-column label="状态">
          <template #default="scope">
            <el-tag v-if="scope.row.status === 1">种植中</el-tag>
            <el-tag type="success" v-if="scope.row.status === 2">已收获</el-tag>
          </template>
        </el-table-column>

        <el-table-column prop="createdAt" label="创建时间"/>

        <!-- 操作 -->
        <el-table-column label="操作" width="220">
          <template #default="scope">

            <!-- 已收获 -->
            <template v-if="scope.row.status === 2">
              <el-button
                  text
                  type="success"
                  size="small"
                  @click="openMaterial(scope.row)"
              >
                物资使用
              </el-button>
            </template>

            <!-- 未收获 -->
            <template v-else>
              <el-button text type="primary" size="small" @click="openEdit(scope.row)">
                修改
              </el-button>

              <el-button text type="danger" size="small" @click="onDelete(scope.row)">
                删除
              </el-button>

              <el-button text type="success" size="small" @click="openMaterial(scope.row)">
                物资使用
              </el-button>

              <el-button size="small" type="warning" @click="openPestDialog(scope.row)">
                上报虫灾
              </el-button>
            </template>

          </template>
        </el-table-column>

      </el-table>

      <!-- 分页 -->
      <div style="margin-top:20px;text-align:right">
        <el-pagination
            v-model:current-page="tableData.param.pageNum"
            v-model:page-size="tableData.param.pageSize"
            :total="tableData.total"
            :page-sizes="[10,20,50,100]"
            layout="total, sizes, prev, pager, next"
            @current-change="getList"
            @size-change="getList"
        />
      </div>

    </el-card>

    <!-- 新增/修改 -->
    <el-dialog v-model="dialogVisible" :title="form.id ? '修改' : '新增'" width="420px">
      <el-form :model="form" label-width="100px">

        <el-form-item label="地块">
          <el-select v-model="form.fieldId" style="width:100%">
            <el-option v-for="item in fieldOptions" :key="item.id" :label="item.name" :value="item.id"/>
          </el-select>
        </el-form-item>

        <el-form-item label="作物">
          <el-select v-model="form.cropId" style="width:100%">
            <el-option v-for="item in cropOptions" :key="item.id" :label="item.name" :value="item.id"/>
          </el-select>
        </el-form-item>

        <el-form-item label="种植日期">
          <el-date-picker v-model="form.plantDate" value-format="YYYY-MM-DD" style="width:100%"/>
        </el-form-item>

        <el-form-item label="预计收获">
          <el-date-picker v-model="form.expectedHarvestDate" value-format="YYYY-MM-DD" style="width:100%"/>
        </el-form-item>

        <el-form-item label="状态">
          <el-select v-model="form.status" style="width:100%">
            <el-option label="种植中" :value="1"/>
            <el-option label="已收获" :value="2"/>
          </el-select>
        </el-form-item>

        <el-form-item label="生长阶段">
          <el-input v-model="form.currentStage"/>
        </el-form-item>

      </el-form>

      <template #footer>
        <el-button @click="dialogVisible=false">取消</el-button>
        <el-button type="primary" @click="submit">确定</el-button>
      </template>
    </el-dialog>

    <!-- 虫灾 -->
    <el-dialog v-model="pestDialogVisible" title="上报虫灾" width="420px">
      <el-form :model="pestForm" label-width="80px">

        <el-form-item label="地块">
          <el-input v-model="pestForm.fieldName" disabled/>
        </el-form-item>

        <el-form-item label="作物">
          <el-input v-model="pestForm.cropName" disabled/>
        </el-form-item>

        <el-form-item label="虫灾描述">
          <el-input v-model="pestForm.description" type="textarea"/>
        </el-form-item>

      </el-form>

      <template #footer>
        <el-button @click="pestDialogVisible=false">取消</el-button>
        <el-button type="primary" @click="submitPest">提交</el-button>
      </template>
    </el-dialog>

    <!-- ✅ 关键：子组件 -->
    <MaterialUseDialog ref="materialRef" />

  </div>
</template>

<script setup lang="ts">
import { reactive, ref, onMounted } from "vue"
import { ElMessage, ElMessageBox } from "element-plus"
import MaterialUseDialog from "./component/editmaterialUse.vue"

import {
  getFieldCropListApi,
  addFieldCropApi,
  deleteFieldCropApi
} from "/@/api/system/field_crops"

import { addPestApi } from "/@/api/system/pest"
import { getFieldListApi } from "/@/api/system/fields"
import { getCropListApi } from "/@/api/system/crop"
import { useUserInfo } from "/@/stores/userInfo"
import dayjs from "dayjs"

const userStore = useUserInfo()

const loading = ref(false)
const dialogVisible = ref(false)
const pestDialogVisible = ref(false)

const materialRef = ref()

const fieldOptions = ref([])
const cropOptions = ref([])

const form = reactive({
  id: 0,
  fieldId: "",
  cropId: "",
  plantDate: "",
  expectedHarvestDate: "",
  harvestDate: "",
  currentStage: "",
  status: 1
})

const pestForm = reactive({
  fieldId: 0,
  fieldName: "",
  cropId: 0,
  cropName: "",
  description: ""
})

const tableData = reactive({
  data: [],
  total: 0,
  param: {
    fieldId: "",
    cropId: "",
    pageNum: 1,
    pageSize: 10,
    userId: userStore.userInfos?.id || 0
  }
})

const getList = async () => {
  loading.value = true
  const res = await getFieldCropListApi(tableData.param)
  tableData.data = res.data.list || []
  tableData.total = res.data.total || 0
  loading.value = false
}

const getFieldOptions = async () => {
  const res = await getFieldListApi({ pageNum:1,pageSize:1000,userId:userStore.userInfos?.id || 0 })
  fieldOptions.value = res.data.list || []
}

const getCropOptions = async () => {
  const res = await getCropListApi({ pageNum:1,pageSize:1000 })
  cropOptions.value = res.data.list || []
}

const openAdd = () => {
  Object.assign(form,{
    id:0,fieldId:"",cropId:"",plantDate:"",
    expectedHarvestDate:"",harvestDate:"",
    currentStage:"",status:1
  })
  dialogVisible.value = true
}

const openEdit = (row:any) => {
  Object.assign(form,{...row,harvestDate:""})
  dialogVisible.value = true
}

const submit = async () => {
  form.harvestDate = dayjs().format("YYYY-MM-DD")
  await addFieldCropApi(form)
  ElMessage.success("成功")
  dialogVisible.value = false
  getList()
}

const onDelete = (row:any) => {
  ElMessageBox.confirm("确认删除?")
      .then(async()=>{
        await deleteFieldCropApi({id:row.id})
        ElMessage.success("删除成功")
        getList()
      })
}

/* ✅ 正确物资调用 */
const openMaterial = (row:any)=>{
  if (!materialRef.value) return
  materialRef.value.open(row.id)
}

/* 虫灾 */
const openPestDialog = (row:any) => {
  pestForm.fieldId = row.fieldId
  pestForm.fieldName = row.fieldName
  pestForm.cropId = row.cropId
  pestForm.cropName = row.cropName
  pestForm.description = ""
  pestDialogVisible.value = true
}

const submitPest = async () => {
  if (!pestForm.description) return ElMessage.warning("请输入描述")

  await addPestApi({
    ...pestForm,
    status:0,
    createdBy:userStore.userInfos.id,
    solveBy:0
  })

  ElMessage.success("上报成功")
  pestDialogVisible.value = false
}

const onSearch = () => {
  tableData.param.pageNum = 1
  getList()
}

onMounted(()=>{
  getList()
  getFieldOptions()
  getCropOptions()
})
</script>