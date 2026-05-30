<template>
  <AppLayout title="页面管理">
    <div class="toolbar">
      <el-select v-model="filterCategory" placeholder="全部分类" clearable style="width:160px;margin-right:12px" @change="load">
        <el-option v-for="c in categories" :key="c" :label="c" :value="c" />
      </el-select>
      <el-button type="primary" @click="openForm()">新增页面</el-button>
    </div>
    <el-table :data="list" v-loading="loading" border>
      <el-table-column prop="id" label="ID" width="60" />
      <el-table-column prop="title" label="标题" width="140" />
      <el-table-column prop="slug" label="Slug" width="160" />
      <el-table-column prop="category" label="分类" width="120" />
      <el-table-column prop="sort" label="排序" width="70" />
      <el-table-column label="启用" width="70">
        <template #default="{ row }">
          <el-tag :type="row.enabled ? 'success' : 'info'" size="small">{{ row.enabled ? '是' : '否' }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="160">
        <template #default="{ row }">
          <el-button link type="primary" @click="openForm(row)">编辑</el-button>
          <el-button link type="danger" @click="onDelete(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="dialogVisible" :title="form.id ? '编辑页面' : '新增页面'" width="720px">
      <el-form label-width="80px">
        <el-form-item label="标题"><el-input v-model="form.title" /></el-form-item>
        <el-form-item label="Slug">
          <el-input v-model="form.slug" placeholder="如 shopping-flow" />
        </el-form-item>
        <el-form-item label="分类">
          <el-select v-model="form.category" allow-create filterable style="width:100%">
            <el-option v-for="c in categories" :key="c" :label="c" :value="c" />
          </el-select>
        </el-form-item>
        <el-form-item label="内容">
          <el-input v-model="form.body" type="textarea" :rows="14" placeholder="支持 HTML 格式" />
        </el-form-item>
        <el-form-item label="排序"><el-input-number v-model="form.sort" :min="0" /></el-form-item>
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

interface ContentPage {
  id?: number
  slug: string
  title: string
  category: string
  body: string
  sort: number
  enabled: boolean
}

const categories = ['购物指南', '售后服务', '关于我们', '法律条款']
const list = ref<ContentPage[]>([])
const loading = ref(false)
const filterCategory = ref('')
const dialogVisible = ref(false)
const form = reactive<ContentPage>({ slug: '', title: '', category: '购物指南', body: '', sort: 0, enabled: true })

async function load() {
  loading.value = true
  const params = filterCategory.value ? { category: filterCategory.value } : {}
  const { data } = await api.get('/api/pages', { params })
  list.value = data.data
  loading.value = false
}

function openForm(row?: ContentPage) {
  Object.assign(form, row || { id: undefined, slug: '', title: '', category: '购物指南', body: '', sort: 0, enabled: true })
  dialogVisible.value = true
}

async function onSave() {
  if (!form.title || !form.slug || !form.category) {
    ElMessage.warning('请填写标题、Slug 和分类')
    return
  }
  if (form.id) await api.put(`/api/pages/${form.id}`, form)
  else await api.post('/api/pages', form)
  ElMessage.success('已保存')
  dialogVisible.value = false
  load()
}

async function onDelete(row: ContentPage) {
  await ElMessageBox.confirm('确定删除？')
  await api.delete(`/api/pages/${row.id}`)
  load()
}

onMounted(load)
</script>

<style scoped>
.toolbar { margin-bottom: 12px; display: flex; align-items: center; }
</style>
