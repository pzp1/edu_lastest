<template>
  <div class="dashboard-container">

    <el-row :gutter="20">

      <!-- 虫灾 -->
      <el-col :span="12">
        <el-card shadow="hover">
          <template #header>
            <span>🐛 虫灾预警</span>
          </template>

          <el-table :data="pestList" height="300" :row-class-name="rowClassName">

            <el-table-column prop="fieldName" label="地块"/>
            <el-table-column prop="cropName" label="作物"/>

            <el-table-column prop="description" label="描述" show-overflow-tooltip/>

            <el-table-column label="状态" width="100">
              <template #default="scope">
                <el-tag type="danger" v-if="scope.row.status === 0">未解决</el-tag>
                <el-tag type="success" v-else>已解决</el-tag>
              </template>
            </el-table-column>

            <el-table-column prop="solveName" label="处理人" width="120"/>

            <el-table-column label="处理说明" width="200">
              <template #default="scope">

                <el-tooltip
                    v-if="scope.row.sloveAnswer"
                    :content="scope.row.sloveAnswer"
                    placement="top"
                >
                  <span>
                    {{ formatSolve(scope.row.sloveAnswer) }}
                  </span>
                </el-tooltip>

                <span v-else>-</span>

                <el-button
                    v-if="scope.row.sloveAnswer"
                    type="primary"
                    text
                    size="small"
                    @click="openSolveDetail(scope.row)"
                >
                  查看
                </el-button>

              </template>
            </el-table-column>

            <el-table-column prop="createdAt" label="时间" width="160"/>

          </el-table>

        </el-card>
      </el-col>

      <!-- 即将收获 -->
      <el-col :span="12">
        <el-card shadow="hover">
          <template #header>
            <span>🌾 30天内即将收获</span>
          </template>

          <el-table :data="harvestList" height="300">

            <el-table-column prop="fieldName" label="地块"/>
            <el-table-column prop="cropName" label="作物"/>

            <el-table-column prop="expectedHarvestDate" label="预计收获"/>

            <el-table-column label="剩余天数">
              <template #default="scope">
                <el-tag :type="getRemainDays(scope.row.expectedHarvestDate) <= 3 ? 'danger' : 'warning'">
                  {{ getRemainDays(scope.row.expectedHarvestDate) }} 天
                </el-tag>
              </template>
            </el-table-column>

          </el-table>

        </el-card>
      </el-col>

    </el-row>

    <!-- 处理详情弹窗 -->
    <el-dialog v-model="solveDetailVisible" title="处理详情" width="500px">

      <el-form label-width="80px">

        <el-form-item label="处理人">
          <span>{{ solveDetail.solveName || "-" }}</span>
        </el-form-item>

        <el-form-item label="处理说明">
          <el-input
              v-model="solveDetail.sloveAnswer"
              type="textarea"
              :rows="5"
              disabled
          />
        </el-form-item>

      </el-form>

      <template #footer>
        <el-button @click="solveDetailVisible=false">关闭</el-button>
      </template>

    </el-dialog>

  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue"
import dayjs from "dayjs"
import { getPestListApi } from "/@/api/system/pest"
import { getFieldCropListApi } from "/@/api/system/field_crops"
import { useUserInfo } from "/@/stores/userInfo"

const userStore = useUserInfo()

const pestList = ref<any[]>([])
const harvestList = ref<any[]>([])

/* 弹窗 */
const solveDetailVisible = ref(false)
const solveDetail = ref({
  solveName: "",
  sloveAnswer: ""
})

/* 获取虫灾 */
const getPestList = async () => {
  const res = await getPestListApi({
    pageNum: 1,
    pageSize: 5,
    status: 0,
    createdBy: userStore.userInfos?.id
  })

  pestList.value = res.data?.list || []
}

/* 获取30天内收获（核心改动） */
const getHarvestList = async () => {
  const res = await getFieldCropListApi({
    pageNum: 1,
    pageSize: 5,
    userId: userStore.userInfos?.id,
    expectedFlag: 1 // ✅ 关键
  })

  harvestList.value = res.data?.list || []
}

/* 剩余天数 */
const getRemainDays = (dateStr: string) => {
  if (!dateStr) return 0
  return dayjs(dateStr).diff(dayjs(), "day")
}

/* 省略处理说明 */
const formatSolve = (text: string) => {
  if (!text) return ""
  return text.length > 6 ? text.slice(0, 6) + "..." : text
}

/* 查看详情 */
const openSolveDetail = (row: any) => {
  solveDetail.value.solveName = row.solveName
  solveDetail.value.sloveAnswer = row.sloveAnswer
  solveDetailVisible.value = true
}

/* 行高亮 */
const rowClassName = ({ row }: any) => {
  return row.status === 0 ? "danger-row" : ""
}

onMounted(() => {
  getPestList()
  getHarvestList()
})
</script>

<style scoped>
.dashboard-container {
  padding: 20px;
}

.danger-row {
  background: #fff2f0;
}
</style>