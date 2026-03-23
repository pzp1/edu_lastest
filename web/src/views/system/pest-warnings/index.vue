<template>
  <div class="system-pest-container">
    <el-card shadow="hover">

      <!-- 查询 -->
      <div class="search mb15">
        <el-form :inline="true">

          <el-form-item label="作物">
            <el-select v-model="tableData.param.cropId" clearable style="width:160px">
              <el-option
                  v-for="item in cropOptions"
                  :key="item.id"
                  :label="item.name"
                  :value="item.id"
              />
            </el-select>
          </el-form-item>

          <el-form-item label="地块">
            <el-select v-model="tableData.param.fieldId" clearable style="width:160px">
              <el-option
                  v-for="item in fieldOptions"
                  :key="item.id"
                  :label="item.name"
                  :value="item.id"
              />
            </el-select>
          </el-form-item>

          <el-form-item label="状态">
            <el-select v-model="tableData.param.status" clearable style="width:140px">
              <el-option label="未解决" :value="0"/>
              <el-option label="已解决" :value="1"/>
            </el-select>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" @click="onSearch">查询</el-button>
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

        <el-table-column prop="description" label="虫灾描述"/>

        <el-table-column label="状态" width="120">
          <template #default="scope">
            <el-tag type="danger" v-if="scope.row.status === 0">未解决</el-tag>
            <el-tag type="success" v-if="scope.row.status === 1">已解决</el-tag>
          </template>
        </el-table-column>

        <el-table-column prop="createdAt" label="上报时间" width="180"/>

        <!-- 操作 -->
        <el-table-column label="操作" width="220">
          <template #default="scope">

            <el-button
                text
                type="primary"
                size="small"
                @click="openEdit(scope.row)"
            >
              编辑
            </el-button>

            <el-button
                text
                type="success"
                size="small"
                v-if="scope.row.status === 0"
                @click="openSolveDialog(scope.row)"
            >
              处理
            </el-button>

            <el-button
                text
                type="danger"
                size="small"
                @click="onDelete(scope.row)"
            >
              删除
            </el-button>

          </template>
        </el-table-column>

      </el-table>

      <!-- 分页 -->
      <div style="margin-top:20px;text-align:right">
        <el-pagination
            v-model:current-page="tableData.param.pageNum"
            v-model:page-size="tableData.param.pageSize"
            :total="tableData.total"
            :page-sizes="[10,20,50]"
            layout="total, sizes, prev, pager, next"
            @current-change="getList"
            @size-change="getList"
        />
      </div>

    </el-card>

    <!-- 编辑弹窗 -->
    <el-dialog v-model="dialogVisible" title="编辑虫灾" width="420px">

      <el-form :model="form" label-width="100px">

        <el-form-item label="虫灾描述">
          <el-input v-model="form.description" type="textarea"/>
        </el-form-item>

        <el-form-item label="状态">
          <el-select v-model="form.status" style="width:100%">
            <el-option label="未解决" :value="0"/>
            <el-option label="已解决" :value="1"/>
          </el-select>
        </el-form-item>

      </el-form>

      <template #footer>
        <el-button @click="dialogVisible=false">取消</el-button>
        <el-button type="primary" @click="submit">确定</el-button>
      </template>

    </el-dialog>
    <!-- 处理虫灾弹窗 -->
    <el-dialog v-model="solveDialogVisible" title="处理虫灾" width="420px">

      <el-form :model="solveForm" label-width="100px">

        <el-form-item label="处理结果">
          <el-select v-model="solveForm.status" style="width:100%">
            <el-option label="未解决" :value="0"/>
            <el-option label="已解决" :value="1"/>
          </el-select>
        </el-form-item>

        <el-form-item label="处理说明">
          <el-input
              v-model="solveForm.solveAnswer"
              type="textarea"
              placeholder="请输入处理方案或结果"
          />
        </el-form-item>

      </el-form>

      <template #footer>
        <el-button @click="solveDialogVisible=false">取消</el-button>
        <el-button type="primary" @click="submitSolve">确定</el-button>
      </template>

    </el-dialog>
  </div>
</template>

<script setup lang="ts">

import { ref, reactive, onMounted } from "vue"
import { ElMessage, ElMessageBox } from "element-plus"

import {
  getPestListApi,
  updatePestApi,
  deletePestApi,
  updatePestStatusApi
} from "/@/api/system/pest"

import { getFieldListApi } from "/@/api/system/fields"
import { getCropListApi } from "/@/api/system/crop"
import { useUserInfo } from "/@/stores/userInfo"

const userStore = useUserInfo()
const solveDialogVisible = ref(false)

const solveForm = reactive({
  id: 0,
  status: 1,
  solveAnswer: ""
})
const openSolveDialog = (row:any) => {
  solveForm.id = row.id
  solveForm.status = row.status
  solveForm.solveAnswer = row.solveAnswer || ""
  solveDialogVisible.value = true
}
const submitSolve = async () => {
  if (!solveForm.solveAnswer) {
    ElMessage.warning("请输入处理说明")
    return
  }

  await updatePestApi({
    id: solveForm.id,
    status: solveForm.status,
    sloveAnswer: solveForm.solveAnswer,
    solveBY:userStore.userInfos?.id || 0
  })

  ElMessage.success("处理成功")
  solveDialogVisible.value = false
  getList()
}
const loading = ref(false)
const dialogVisible = ref(false)

const fieldOptions = ref([])
const cropOptions = ref([])

const form = reactive({
  id: 0,
  description: "",
  status: 0
})

const tableData = reactive({
  data: [],
  total: 0,
  param: {
    fieldId: "",
    cropId: "",
    status: "",
    pageNum: 1,
    pageSize: 10,
    userId: userStore.userInfos.id
  }
})

/* 获取列表 */
const getList = async () => {
  loading.value = true
  const res = await getPestListApi(tableData.param)
  tableData.data = res.data.list || []
  tableData.total = res.data.total || 0
  loading.value = false
}

/* 查询 */
const onSearch = () => {
  tableData.param.pageNum = 1
  getList()
}

/* 编辑 */
const openEdit = (row:any) => {
  Object.assign(form, row)
  dialogVisible.value = true
}

/* 提交 */
const submit = async () => {
  await updatePestApi(form)
  ElMessage.success("修改成功")
  dialogVisible.value = false
  getList()
}

/* 标记解决 */
const handleSolve = async (row:any) => {
  await updatePestStatusApi({
    id: row.id,
    status: 1
  })
  ElMessage.success("已标记为解决")
  getList()
}

/* 删除 */
const onDelete = (row:any) => {
  ElMessageBox.confirm("确定删除吗？","提示",{type:"warning"})
      .then(async()=>{
        await deletePestApi({id:row.id})
        ElMessage.success("删除成功")
        getList()
      })
}

/* 下拉数据 */
const getFieldOptions = async () => {
  const res = await getFieldListApi({
    pageNum:1,
    pageSize:1000,
    userId:userStore.userInfos.id
  })
  fieldOptions.value = res.data.list || []
}

const getCropOptions = async () => {
  const res = await getCropListApi({
    pageNum:1,
    pageSize:1000
  })
  cropOptions.value = res.data.list || []
}

onMounted(()=>{
  getList()
  getFieldOptions()
  getCropOptions()
})

</script>