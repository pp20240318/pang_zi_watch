<template>
  <AppLayout title="品牌管理">
    <div class="toolbar">
      <el-button type="primary" @click="openForm()">新增品牌</el-button>
    </div>
    <el-table :data="list" v-loading="loading" border>
      <el-table-column prop="code" label="编码" width="140" />
      <el-table-column prop="name" label="名称" />
      <el-table-column prop="sort" label="排序" width="80" />
      <el-table-column label="联系客服" width="100">
        <template #default="{ row }">{{ row.contact ? '是' : '否' }}</template>
      </el-table-column>
      <el-table-column label="启用" width="80">
        <template #default="{ row }"><el-tag :type="row.enabled ? 'success' : 'info'">{{ row.enabled ? '是' : '否' }}</el-tag></template>
      </el-table-column>
      <el-table-column label="操作" width="160">
        <template #default="{ row }">
          <el-button link type="primary" @click="openForm(row)">编辑</el-button>
          <el-button link type="danger" @click="onDelete(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="dialogVisible" :title="form.id ? '编辑品牌' : '新增品牌'" width="480px">
      <el-form label-width="90px">
        <el-form-item label="编码"><el-input v-model="form.code" :disabled="!!form.id" /></el-form-item>
        <el-form-item label="名称"><el-input v-model="form.name" /></el-form-item>
        <el-form-item label="排序"><el-input-number v-model="form.sort" :min="0" /></el-form-item>
        <el-form-item label="联系客服"><el-switch v-model="form.contact" /></el-form-item>
        <el-form-item label="启用"><el-switch v-model="form.enabled" /></el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="onSave">保存</el-button>
      </template>
    </el-dialog>
  </AppLayout>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import AppLayout from '../components/AppLayout.vue'
import api from '../api/client'

interface Brand {
  id?: number
  code: string
  name: string
  sort: number
  contact: boolean
  enabled: boolean
}

const list = ref<Brand[]>([])
const loading = ref(false)
const dialogVisible = ref(false)
const form = reactive<Brand>({ code: '', name: '', sort: 0, contact: false, enabled: true })

async function load() {
  loading.value = true
  try {
    const { data } = await api.get('/api/brands')
    list.value = data.data
  } finally {
    loading.value = false
  }
}

function openForm(row?: Brand) {
  if (row) {
    Object.assign(form, row)
  } else {
    Object.assign(form, { id: undefined, code: '', name: '', sort: 0, contact: false, enabled: true })
  }
  dialogVisible.value = true
}

async function onSave() {
  if (!form.code || !form.name) {
    ElMessage.warning('请填写编码和名称')
    return
  }
  if (form.id) {
    await api.put(`/api/brands/${form.id}`, form)
  } else {
    await api.post('/api/brands', form)
  }
  ElMessage.success('已保存')
  dialogVisible.value = false
  load()
}

async function onDelete(row: Brand) {
  await ElMessageBox.confirm('确定删除该品牌？', '提示')
  await api.delete(`/api/brands/${row.id}`)
  ElMessage.success('已删除')
  load()
}

onMounted(load)
</script>

<style scoped>
.toolbar { margin-bottom: 12px; }
</style>
