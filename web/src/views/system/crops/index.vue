<template>
  <div class="system-crop-container">
    <el-card shadow="hover">

      <!-- 查询 -->
      <div class="search mb15">
        <el-form :inline="true">

          <el-form-item label="作物名称">
            <el-input
                v-model="tableData.param.name"
                placeholder="请输入作物名称"
                clearable
                style="width:160px"
            />
          </el-form-item>

          <el-form-item label="分类">
            <el-select
                v-model="tableData.param.category"
                placeholder="请选择分类"
                clearable
                style="width:160px"
            >
              <el-option label="粮食" :value="1"/>
              <el-option label="水果" :value="2"/>
              <el-option label="蔬菜" :value="3"/>
            </el-select>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" @click="onSearch">
              查询
            </el-button>

            <el-button type="success" @click="openAdd">
              新增作物
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

        <!-- 序号 -->
        <el-table-column label="序号" width="70">
          <template #default="scope">
            {{
              (tableData.param.pageNum - 1) * tableData.param.pageSize
              + scope.$index + 1
            }}
          </template>
        </el-table-column>

        <el-table-column prop="name" label="作物名称"/>

        <el-table-column
            prop="growthCycleDays"
            label="生长周期(天)"
            width="120"
        />

        <!-- 分类 -->
        <el-table-column label="分类" width="120">

          <template #default="scope">

            <el-tag v-if="scope.row.category === 1">
              粮食
            </el-tag>

            <el-tag type="success" v-if="scope.row.category === 2">
              水果
            </el-tag>

            <el-tag type="warning" v-if="scope.row.category === 3">
              蔬菜
            </el-tag>

          </template>

        </el-table-column>

        <el-table-column
            prop="createdAt"
            label="创建时间"
            width="180"
        />

        <!-- 操作 -->
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

    <!-- 新增/修改弹窗 -->
    <el-dialog
        v-model="dialogVisible"
        :title="form.id ? '修改作物' : '新增作物'"
        width="400px"
    >

      <el-form
          :model="form"
          ref="formRef"
          label-width="100px"
      >

        <el-form-item label="作物名称">
          <el-input v-model="form.name"/>
        </el-form-item>

        <el-form-item label="生长周期">

          <el-input-number
              v-model="form.growthCycleDays"
              :min="1"
              style="width:100%"
          />

        </el-form-item>

        <el-form-item label="分类">

          <el-select
              v-model="form.category"
              style="width:100%"
          >
            <el-option label="粮食" :value="1"/>
            <el-option label="水果" :value="2"/>
            <el-option label="蔬菜" :value="3"/>
          </el-select>

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

import {reactive, ref, onMounted} from "vue"
import {ElMessage, ElMessageBox} from "element-plus"
import {
  getCropListApi,
  addCropApi,
  deleteCropApi
} from "/@/api/system/crop"

const dialogVisible = ref(false)
const formRef = ref()
const loading = ref(false)

const form = reactive({
  id: 0,
  name: "",
  growthCycleDays: 90,
  category: 1
})

const tableData = reactive({
  data: [],
  total: 0,
  param: {
    name: "",
    category: "",
    pageNum: 1,
    pageSize: 10
  }
})

/* 获取列表 */
const getList = async () => {

  loading.value = true

  const res = await getCropListApi(tableData.param)

  tableData.data = res.data.list || []
  tableData.total = res.data.total || 0

  loading.value = false

}

/* 查询 */
const onSearch = () => {

  tableData.param.pageNum = 1

  getList()

}

/* 新增 */
const openAdd = () => {

  form.id = 0
  form.name = ""
  form.growthCycleDays = 90
  form.category = 1

  dialogVisible.value = true

}

/* 编辑 */
const openEdit = (row) => {

  form.id = row.id
  form.name = row.name
  form.growthCycleDays = row.growthCycleDays
  form.category = row.category

  dialogVisible.value = true

}

/* 提交 */
const submit = async () => {

  await addCropApi(form)

  ElMessage.success(form.id ? "修改成功" : "新增成功")

  dialogVisible.value = false

  getList()

}

/* 删除 */
const onDelete = (row) => {

  ElMessageBox.confirm("确定删除该作物吗？", "提示", {
    type: "warning"
  }).then(async () => {

    await deleteCropApi({id: row.id})

    ElMessage.success("删除成功")

    getList()

  })

}

onMounted(() => {
  getList()
})

</script>