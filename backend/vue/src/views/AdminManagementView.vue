<template>
  <AppLayout title="管理员">
    <el-button type="primary" class="mb" @click="openForm()">新增管理员</el-button>
    <el-table :data="list" border>
      <el-table-column prop="id" label="ID" width="60" />
      <el-table-column prop="username" label="用户名" />
      <el-table-column label="角色">
        <template #default="{ row }">{{ row.role?.name }}</template>
      </el-table-column>
      <el-table-column label="操作" width="160">
        <template #default="{ row }">
          <el-button link type="primary" @click="openForm(row)">编辑</el-button>
          <el-button link type="danger" @click="onDelete(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="visible" :title="form.id ? '编辑' : '新增'" width="420px">
      <el-form label-width="80px">
        <el-form-item label="用户名"><el-input v-model="form.username" /></el-form-item>
        <el-form-item label="密码"><el-input v-model="form.password" type="password" :placeholder="form.id ? '留空不改' : ''" /></el-form-item>
        <el-form-item label="角色">
          <el-select v-model="form.role_id" style="width:100%">
            <el-option v-for="r in roles" :key="r.id" :label="r.name" :value="r.id" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="visible = false">取消</el-button>
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

const list = ref<any[]>([])
const roles = ref<any[]>([])
const visible = ref(false)
const form = reactive<any>({ username: '', password: '', role_id: undefined })

async function load() {
  const [a, r] = await Promise.all([api.get('/api/admins'), api.get('/api/roles')])
  list.value = a.data.data
  roles.value = r.data.data
}

function openForm(row?: any) {
  Object.assign(form, row ? { id: row.id, username: row.username, password: '', role_id: row.role_id } : { id: undefined, username: '', password: '', role_id: roles.value[0]?.id })
  visible.value = true
}

async function onSave() {
  const payload: any = { username: form.username, role_id: form.role_id }
  if (form.password) payload.password = form.password
  if (form.id) await api.put(`/api/admins/${form.id}`, payload)
  else {
    if (!form.password) {
      ElMessage.warning('请设置密码')
      return
    }
    await api.post('/api/admins', { ...payload, password: form.password })
  }
  ElMessage.success('已保存')
  visible.value = false
  load()
}

async function onDelete(row: any) {
  await ElMessageBox.confirm('确定删除？')
  await api.delete(`/api/admins/${row.id}`)
  load()
}

onMounted(load)
</script>

<style scoped>.mb { margin-bottom: 12px; }</style>
