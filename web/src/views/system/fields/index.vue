<template>
  <div class="system-field-container">
    <el-card shadow="hover">

      <!-- 查询区域 -->
      <div class="system-field-search mb15">
        <el-form :inline="true">

          <el-form-item label="状态">
            <el-select
                v-model="tableData.param.status"
                placeholder="请选择状态"
                clearable
                style="width: 180px"
            >
              <el-option label="使用中" :value="1" />
              <el-option label="停用" :value="2" />
            </el-select>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" @click="getFieldList">
              查询
            </el-button>

            <el-button type="success" @click="onOpenAdd">
              新增地块
            </el-button>
          </el-form-item>

        </el-form>
      </div>

      <!-- 表格 -->
      <el-table :data="tableData.data" style="width: 100%">

        <el-table-column type="index" label="序号" width="60" />

        <el-table-column
            prop="name"
            label="地块名称"
            show-overflow-tooltip
        />

        <el-table-column
            prop="userInfo.userNickName"
            label="耕种人"
            show-overflow-tooltip
        />

        <el-table-column
            prop="area"
            label="面积(亩)"
            width="100"
        />

        <el-table-column
            prop="location"
            label="位置"
            show-overflow-tooltip
        />

        <el-table-column
            prop="sortOrder"
            label="排序"
            width="80"
        />

        <!-- 状态 -->
        <el-table-column
            label="状态"
            width="100"
        >
          <template #default="scope">
            <el-tag type="success" v-if="scope.row.status === 1">
              使用中
            </el-tag>
            <el-tag type="info"v-if="scope.row.status === 2">
              停用
            </el-tag>
          </template>
        </el-table-column>

        <el-table-column
            prop="createdAt"
            label="创建时间"
            width="180"
        />

        <!-- 操作 -->
        <el-table-column label="操作" width="220">
          <template #default="scope">

            <el-button
                size="small"
                text
                type="primary"
                @click="onOpenEdit(scope.row)"
            >
              修改
            </el-button>

            <el-button
                size="small"
                text
                type="danger"
                @click="onRowDel(scope.row)"
            >
              删除
            </el-button>

          </template>
        </el-table-column>

      </el-table>

      <!-- 分页 -->
      <pagination
          v-show="tableData.total > 0"
          :total="tableData.total"
          v-model:page="tableData.param.pageNum"
          v-model:limit="tableData.param.pageSize"
          @pagination="getFieldList"
      />

    </el-card>

    <!-- 编辑弹窗 -->
    <EditField ref="editRef" @refresh="getFieldList" />

  </div>
</template>

<script setup>
import { reactive, onMounted, ref } from "vue"
import { ElMessage, ElMessageBox } from "element-plus"
import { getFieldListApi, deleteFieldApi } from "/@/api/system/fields"
import { useUserInfo } from "/@/stores/userInfo"
import EditField from "./component/editField.vue"

const editRef = ref()
const userStore = useUserInfo()

const tableData = reactive({
  data: [],
  total: 0,
  param: {
    pageNum: 1,
    pageSize: 10,
    name: "",
    status: "",
    userId: userStore.userInfos.id
  }
})

/**
 * 获取列表
 */
const getFieldList = async () => {
  const res = await getFieldListApi(tableData.param)
  tableData.data = res.data.list
  tableData.total = res.data.total
}

/**
 * 删除
 */
const onRowDel = (row) => {
  ElMessageBox.confirm("确定删除该地块吗？", "提示", {
    type: "warning"
  }).then(async () => {
    await deleteFieldApi({ fieldId: row.id })
    ElMessage.success("删除成功")
    getFieldList()
  })
}

/**
 * 新增
 */
const onOpenAdd = () => {
  editRef.value.open()
}

/**
 * 编辑
 */
const onOpenEdit = (row) => {
  editRef.value.open(row)
}

onMounted(() => {
  getFieldList()
})
</script>