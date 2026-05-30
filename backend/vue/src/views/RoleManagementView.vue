<template>
  <AppLayout title="角色权限">
    <el-button type="primary" class="mb" @click="openForm()">新增角色</el-button>
    <el-table :data="list" border>
      <el-table-column prop="name" label="角色名" width="140" />
      <el-table-column prop="permissions" label="权限" show-overflow-tooltip />
      <el-table-column label="操作" width="160">
        <template #default="{ row }">
          <el-button link type="primary" @click="openForm(row)">编辑</el-button>
          <el-button link type="danger" @click="onDelete(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="visible" :title="form.id ? '编辑角色' : '新增角色'" width="560px">
      <el-form label-width="80px">
        <el-form-item label="名称"><el-input v-model="form.name" /></el-form-item>
        <el-form-item label="权限">
          <el-checkbox v-model="selectAll" @change="onSelectAll">全部权限 (all)</el-checkbox>
          <div v-if="!selectAll" class="perm-groups">
            <div v-for="g in permissionGroups" :key="g.label" class="group">
              <div class="group-label">{{ g.label }}</div>
              <el-checkbox-group v-model="selected">
                <el-checkbox v-for="p in g.items" :key="p" :label="p">{{ p }}</el-checkbox>
              </el-checkbox-group>
            </div>
          </div>
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

const permissionGroups = [
  { label: '仪表盘', items: ['dashboard:view'] },
  { label: '轮播', items: ['banner:read', 'banner:write', 'banner:delete'] },
  { label: '页面', items: ['page:read', 'page:write', 'page:delete'] },
  { label: '站点', items: ['setting:read', 'setting:write'] },
  { label: '品牌', items: ['brand:read', 'brand:write', 'brand:delete'] },
  { label: '商品', items: ['product:read', 'product:write', 'product:delete'] },
  { label: '订单', items: ['order:read', 'order:write'] },
  { label: '系统', items: ['admin:manage', 'admin:roles'] }
]

const list = ref<any[]>([])
const visible = ref(false)
const form = reactive<any>({ name: '' })
const selected = ref<string[]>([])
const selectAll = ref(false)

async function load() {
  const { data } = await api.get('/api/roles')
  list.value = data.data
}

function openForm(row?: any) {
  if (row) {
    Object.assign(form, { id: row.id, name: row.name })
    if (row.permissions === 'all') {
      selectAll.value = true
      selected.value = []
    } else {
      selectAll.value = false
      selected.value = row.permissions.split(',').map((s: string) => s.trim()).filter(Boolean)
    }
  } else {
    Object.assign(form, { id: undefined, name: '' })
    selectAll.value = false
    selected.value = []
  }
  visible.value = true
}

function onSelectAll(val: boolean) {
  if (val) selected.value = []
}

async function onSave() {
  const permissions = selectAll.value ? 'all' : selected.value.join(',')
  const payload = { name: form.name, permissions }
  if (form.id) await api.put(`/api/roles/${form.id}`, payload)
  else await api.post('/api/roles', payload)
  ElMessage.success('已保存')
  visible.value = false
  load()
}

async function onDelete(row: any) {
  await ElMessageBox.confirm('确定删除？')
  await api.delete(`/api/roles/${row.id}`)
  load()
}

onMounted(load)
</script>

<style scoped>
.mb { margin-bottom: 12px; }
.perm-groups { max-height: 320px; overflow-y: auto; }
.group { margin-bottom: 12px; }
.group-label { font-weight: 600; margin-bottom: 4px; font-size: 13px; }
</style>
