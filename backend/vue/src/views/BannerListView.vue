<template>
  <AppLayout title="轮播管理">
    <div class="toolbar"><el-button type="primary" @click="openForm()">新增轮播</el-button></div>
    <el-table :data="list" v-loading="loading" border>
      <el-table-column prop="id" label="ID" width="60" />
      <el-table-column label="图片" width="120">
        <template #default="{ row }"><el-image :src="row.image" style="width:80px;height:48px" fit="cover" /></template>
      </el-table-column>
      <el-table-column prop="title" label="标题" />
      <el-table-column prop="subtitle" label="副标题" />
      <el-table-column prop="cta" label="按钮文案" width="100" />
      <el-table-column prop="sort" label="排序" width="70" />
      <el-table-column label="操作" width="160">
        <template #default="{ row }">
          <el-button link type="primary" @click="openForm(row)">编辑</el-button>
          <el-button link type="danger" @click="onDelete(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="dialogVisible" :title="form.id ? '编辑轮播' : '新增轮播'" width="560px">
      <el-form label-width="90px">
        <el-form-item label="标题"><el-input v-model="form.title" /></el-form-item>
        <el-form-item label="副标题"><el-input v-model="form.subtitle" /></el-form-item>
        <el-form-item label="图片">
          <el-input v-model="form.image" placeholder="图片 URL" />
          <el-upload :show-file-list="false" :http-request="onUpload" style="margin-top:8px">
            <el-button size="small">上传图片</el-button>
          </el-upload>
        </el-form-item>
        <el-form-item label="按钮文案"><el-input v-model="form.cta" /></el-form-item>
        <el-form-item label="跳转链接"><el-input v-model="form.link" placeholder="#" /></el-form-item>
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
import { uploadFile } from '../utils/upload'

interface Banner {
  id?: number
  title: string
  subtitle: string
  image: string
  cta: string
  link: string
  sort: number
  enabled: boolean
}

const list = ref<Banner[]>([])
const loading = ref(false)
const dialogVisible = ref(false)
const form = reactive<Banner>({ title: '', subtitle: '', image: '', cta: '', link: '', sort: 0, enabled: true })

async function load() {
  loading.value = true
  const { data } = await api.get('/api/banners')
  list.value = data.data
  loading.value = false
}

function openForm(row?: Banner) {
  Object.assign(form, row || { id: undefined, title: '', subtitle: '', image: '', cta: '立即选购', link: '', sort: 0, enabled: true })
  dialogVisible.value = true
}

async function onUpload(opt: { file: File }) {
  form.image = await uploadFile(opt.file)
  ElMessage.success('上传成功')
}

async function onSave() {
  if (!form.title || !form.image) {
    ElMessage.warning('请填写标题和图片')
    return
  }
  if (form.id) await api.put(`/api/banners/${form.id}`, form)
  else await api.post('/api/banners', form)
  ElMessage.success('已保存')
  dialogVisible.value = false
  load()
}

async function onDelete(row: Banner) {
  await ElMessageBox.confirm('确定删除？')
  await api.delete(`/api/banners/${row.id}`)
  load()
}

onMounted(load)
</script>

<style scoped>
.toolbar { margin-bottom: 12px; }
</style>
