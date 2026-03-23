<template>
  <div class="system-material-container">
    <el-card shadow="hover">

      <!-- 查询 -->
      <div class="search mb15">
        <el-form :inline="true">

          <el-form-item label="物料名称">
            <el-input
                v-model="tableData.param.name"
                placeholder="请输入物料名称"
                clearable
                style="width:160px"
            />
          </el-form-item>

          <el-form-item label="类型">
            <el-select
                v-model="tableData.param.type"
                placeholder="请选择类型"
                clearable
                style="width:140px"
            >
              <el-option label="农药" :value="1"/>
              <el-option label="化肥" :value="2"/>
            </el-select>
          </el-form-item>

          <el-form-item label="分类">
            <el-select
                v-model="tableData.param.category"
                placeholder="请选择分类"
                clearable
                style="width:160px"
            >
              <el-option label="杀虫剂" :value="1"/>
              <el-option label="杀菌剂" :value="2"/>
              <el-option label="除草剂" :value="3"/>
              <el-option label="氮肥" :value="4"/>
              <el-option label="磷肥" :value="5"/>
              <el-option label="钾肥" :value="6"/>
              <el-option label="复合肥" :value="7"/>
              <el-option label="有机肥" :value="8"/>
            </el-select>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" @click="onSearch">
              查询
            </el-button>

            <el-button type="success" @click="openAdd">
              新增物料
            </el-button>
          </el-form-item>

        </el-form>
      </div>

      <!-- 表格 -->
      <el-table
          :data="tableData.data"
          style="width:100%"
          v-loading="loading"
      >

        <el-table-column label="序号" width="70">
          <template #default="scope">
            {{
              (tableData.param.pageNum - 1) * tableData.param.pageSize
              + scope.$index + 1
            }}
          </template>
        </el-table-column>

        <el-table-column prop="name" label="物料名称"/>

        <el-table-column label="类型" width="100">
          <template #default="scope">
            <el-tag v-if="scope.row.type === 1">
              农药
            </el-tag>

            <el-tag type="success" v-if="scope.row.type === 2">
              化肥
            </el-tag>
          </template>
        </el-table-column>

        <el-table-column label="分类" width="120">
          <template #default="scope">

            <span v-if="scope.row.category === 1">杀虫剂</span>
            <span v-if="scope.row.category === 2">杀菌剂</span>
            <span v-if="scope.row.category === 3">除草剂</span>
            <span v-if="scope.row.category === 4">氮肥</span>
            <span v-if="scope.row.category === 5">磷肥</span>
            <span v-if="scope.row.category === 6">钾肥</span>
            <span v-if="scope.row.category === 7">复合肥</span>
            <span v-if="scope.row.category === 8">有机肥</span>

          </template>
        </el-table-column>

        <el-table-column
            prop="usageMethod"
            label="使用方式"
            width="120"
        />

        <el-table-column
            prop="nutrientContent"
            label="养分含量"
        />

        <el-table-column
            prop="createdAt"
            label="创建时间"
            width="180"
        />

        <el-table-column label="操作" width="180">

          <template #default="scope">

            <el-button
                text
                type="primary"
                size="small"
                @click="openEdit(scope.row)"
            >
              修改
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
            :page-sizes="[10,20,50,100]"
            background
            layout="total, sizes, prev, pager, next, jumper"
            @size-change="getList"
            @current-change="getList"
        />

      </div>

    </el-card>

    <!-- 新增 / 修改 -->
    <el-dialog
        v-model="dialogVisible"
        :title="form.id ? '修改物料' : '新增物料'"
        width="420px"
    >

      <el-form
          :model="form"
          ref="formRef"
          label-width="100px"
      >

        <el-form-item label="物料名称">
          <el-input v-model="form.name"/>
        </el-form-item>

        <el-form-item label="类型">

          <el-select
              v-model="form.type"
              style="width:100%"
          >
            <el-option label="农药" :value="1"/>
            <el-option label="化肥" :value="2"/>
          </el-select>

        </el-form-item>

        <el-form-item label="分类">

          <el-select
              v-model="form.category"
              style="width:100%"
          >
            <el-option label="杀虫剂" :value="1"/>
            <el-option label="杀菌剂" :value="2"/>
            <el-option label="除草剂" :value="3"/>
            <el-option label="氮肥" :value="4"/>
            <el-option label="磷肥" :value="5"/>
            <el-option label="钾肥" :value="6"/>
            <el-option label="复合肥" :value="7"/>
            <el-option label="有机肥" :value="8"/>
          </el-select>

        </el-form-item>

        <el-form-item label="使用方式">
          <el-input v-model="form.usageMethod"/>
        </el-form-item>

        <el-form-item label="养分含量">
          <el-input v-model="form.nutrientContent"/>
        </el-form-item>

      </el-form>

      <template #footer>

        <el-button @click="dialogVisible=false">
          取消
        </el-button>

        <el-button type="primary" @click="submit">
          确定
        </el-button>

      </template>

    </el-dialog>

  </div>
</template>

<script setup>

import { reactive, ref, onMounted } from "vue"
import { ElMessage, ElMessageBox } from "element-plus"

import {
  getMaterialListApi,
  addOrUpdateMaterialApi,
  deleteMaterialApi
} from "/@/api/system/material"

const dialogVisible = ref(false)
const formRef = ref()
const loading = ref(false)

const form = reactive({
  id: 0,
  name: "",
  type: 1,
  category: 1,
  usageMethod: "",
  nutrientContent: ""
})

const tableData = reactive({
  data: [],
  total: 0,
  param: {
    name: "",
    type: "",
    category: "",
    pageNum: 1,
    pageSize: 10
  }
})

const getList = async () => {

  loading.value = true

  const res = await getMaterialListApi(tableData.param)

  tableData.data = res.data.list || []
  tableData.total = res.data.total || 0

  loading.value = false

}

const onSearch = () => {

  tableData.param.pageNum = 1
  getList()

}

const openAdd = () => {

  form.id = 0
  form.name = ""
  form.type = 1
  form.category = 1
  form.usageMethod = ""
  form.nutrientContent = ""

  dialogVisible.value = true

}

const openEdit = (row) => {

  form.id = row.id
  form.name = row.name
  form.type = row.type
  form.category = row.category
  form.usageMethod = row.usageMethod
  form.nutrientContent = row.nutrientContent

  dialogVisible.value = true

}

const submit = async () => {

  await addOrUpdateMaterialApi(form)

  ElMessage.success(form.id ? "修改成功" : "新增成功")

  dialogVisible.value = false

  getList()

}

const onDelete = (row) => {

  ElMessageBox.confirm("确定删除该物料吗？", "提示", {
    type: "warning"
  }).then(async () => {

    await deleteMaterialApi({ id: row.id })

    ElMessage.success("删除成功")

    getList()

  })

}

onMounted(() => {
  getList()
})

</script>