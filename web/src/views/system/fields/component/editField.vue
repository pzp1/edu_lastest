<template>
  <el-dialog
      v-model="dialogVisible"
      :title="form.id ? '编辑地块' : '新增地块'"
      width="500px"
      destroy-on-close
  >
    <el-form
        ref="formRef"
        :model="form"
        :rules="rules"
        label-width="100px"
    >

      <el-form-item label="耕种人" prop="userId">
        <el-select
            v-model="form.userId"
            placeholder="请选择耕种人"
            style="width: 100%"
            filterable
        >
          <el-option
              v-for="item in userList"
              :key="item.id"
              :label="item.userNickname"
              :value="item.id"
          />
        </el-select>
      </el-form-item>

      <el-form-item label="地块名称" prop="name">
        <el-input v-model="form.name" placeholder="请输入地块名称" />
      </el-form-item>

      <el-form-item label="面积(亩)" prop="area">
        <el-input-number
            v-model="form.area"
            :min="0"
            :precision="2"
            style="width: 100%"
        />
      </el-form-item>

      <el-form-item label="位置">
        <el-input v-model="form.location" />
      </el-form-item>

      <el-form-item label="状态">
        <el-select v-model="form.status" style="width: 100%">
          <el-option label="使用中" :value="1" />
          <el-option label="停用" :value="2" />
        </el-select>
      </el-form-item>

      <el-form-item label="排序">
        <el-input-number
            v-model="form.sortOrder"
            :min="0"
            style="width: 100%"
        />
      </el-form-item>

    </el-form>

    <template #footer>
      <el-button @click="dialogVisible = false">取消</el-button>
      <el-button type="primary" @click="handleSubmit">
        确定
      </el-button>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, reactive } from "vue"
import { ElMessage } from "element-plus"
import { addFieldApi, updateFieldApi } from "/@/api/system/fields"
import { getUserList } from "/@/api/system/user"

const emit = defineEmits(["refresh"])

const dialogVisible = ref(false)
const formRef = ref()

/** 用户列表 */
const userList = ref([])

/** 表单 */
const form = reactive({
  id: null,
  userId: null,
  name: "",
  area: 0,
  location: "",
  status: 1,
  sortOrder: 0,
  createdAt:""
})

const rules = {
  userId: [{ required: true, message: "请选择耕种人", trigger: "change" }],
  name: [{ required: true, message: "请输入地块名称", trigger: "blur" }]
}

/**
 * 获取用户列表
 */
const loadUsers = async () => {
  const res = await getUserList({
    pageNum: 1,
    pageSize: 1000
  })
  userList.value =  res.data.userList
}

/**
 * 打开弹窗
 */
const open = async (row = null) => {
  await loadUsers()

  if (row) {
    console.log("id:::::============",row.id)
    Object.assign(form, {
      id: row.id,
      userId: row.userId,
      name: row.name,
      area: row.area,
      location: row.location,
      status: row.status,
      sortOrder: row.sortOrder,
      createdAt:row.create,

    })
  } else {
    resetForm()
  }

  dialogVisible.value = true
}

/**
 * 重置表单
 */
const resetForm = () => {
  form.id = null
  form.userId = null
  form.name = ""
  form.area = 0
  form.location = ""
  form.status = 1
  form.sortOrder = 0
  form.createdAt = ""
}

/**
 * 提交
 */
const handleSubmit = () => {
  formRef.value.validate(async (valid) => {
    if (!valid) return

    if (form.id) {
      await updateFieldApi(form)
      ElMessage.success("修改成功")
    } else {
      await addFieldApi(form)
      ElMessage.success("新增成功")
    }

    dialogVisible.value = false
    emit("refresh")
  })
}

defineExpose({open})
</script>