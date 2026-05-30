<template>
  <AppLayout title="商品管理">
    <div class="toolbar">
      <el-select v-model="brandFilter" placeholder="品牌筛选" clearable style="width:180px" @change="load">
        <el-option v-for="b in brands" :key="b.code" :label="b.name" :value="b.code" />
      </el-select>
      <el-button type="primary" @click="openForm()">新增商品</el-button>
    </div>
    <el-table :data="list" v-loading="loading" border>
      <el-table-column prop="id" label="ID" width="60" />
      <el-table-column label="图" width="80">
        <template #default="{ row }"><el-image :src="row.image" style="width:56px;height:56px" fit="cover" /></template>
      </el-table-column>
      <el-table-column prop="name" label="名称" min-width="160" />
      <el-table-column prop="brandId" label="品牌" width="100" />
      <el-table-column label="价格" width="110">
        <template #default="{ row }">¥{{ row.price?.toLocaleString() }}</template>
      </el-table-column>
      <el-table-column prop="tag" label="标签" width="80" />
      <el-table-column label="操作" width="160">
        <template #default="{ row }">
          <el-button link type="primary" @click="openForm(row)">编辑</el-button>
          <el-button link type="danger" @click="onDelete(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="dialogVisible" :title="form.id ? '编辑商品' : '新增商品'" width="640px">
      <el-form label-width="90px">
        <el-form-item label="品牌">
          <el-select v-model="form.brandId" style="width:100%">
            <el-option v-for="b in brands" :key="b.code" :label="b.name" :value="b.code" />
          </el-select>
        </el-form-item>
        <el-form-item label="名称"><el-input v-model="form.name" /></el-form-item>
        <el-form-item label="价格"><el-input-number v-model="form.price" :min="0" style="width:100%" /></el-form-item>
        <el-form-item label="主图">
          <el-input v-model="form.image" />
          <el-upload :show-file-list="false" :http-request="(o) => uploadMain(o)" style="margin-top:8px">
            <el-button size="small">上传主图</el-button>
          </el-upload>
        </el-form-item>
        <el-form-item label="标签"><el-input v-model="form.tag" placeholder="热销/新品" /></el-form-item>
        <el-form-item label="评分"><el-input-number v-model="form.rating" :min="0" :max="5" :step="0.1" /></el-form-item>
        <el-form-item label="评价数"><el-input-number v-model="form.reviews" :min="0" /></el-form-item>
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

interface Brand { code: string; name: string }
interface Product {
  id?: number
  brandId: string
  name: string
  price: number
  image: string
  tag?: string
  rating: number
  reviews: number
  enabled: boolean
  gallery?: string[]
}

const list = ref<Product[]>([])
const brands = ref<Brand[]>([])
const loading = ref(false)
const brandFilter = ref('')
const dialogVisible = ref(false)
const form = reactive<Product>({ brandId: '', name: '', price: 0, image: '', rating: 5, reviews: 0, enabled: true })

async function loadBrands() {
  const { data } = await api.get('/api/brands')
  brands.value = (data.data as any[]).filter((b) => !b.contact)
}

async function load() {
  loading.value = true
  const params = brandFilter.value ? { brand: brandFilter.value } : {}
  const { data } = await api.get('/api/products', { params })
  list.value = data.data
  loading.value = false
}

function openForm(row?: Product) {
  if (row) Object.assign(form, { ...row, enabled: row.enabled !== false })
  else Object.assign(form, { id: undefined, brandId: brands.value[0]?.code || '', name: '', price: 0, image: '', tag: '', rating: 4.8, reviews: 0, enabled: true })
  dialogVisible.value = true
}

async function uploadMain(opt: { file: File }) {
  form.image = await uploadFile(opt.file)
  ElMessage.success('上传成功')
}

async function onSave() {
  if (!form.name || !form.brandId || !form.image) {
    ElMessage.warning('请完善必填项')
    return
  }
  const payload = { ...form, gallery: form.gallery || [form.image] }
  if (form.id) await api.put(`/api/products/${form.id}`, payload)
  else await api.post('/api/products', payload)
  ElMessage.success('已保存')
  dialogVisible.value = false
  load()
}

async function onDelete(row: Product) {
  await ElMessageBox.confirm('确定删除？')
  await api.delete(`/api/products/${row.id}`)
  load()
}

onMounted(async () => {
  await loadBrands()
  await load()
})
</script>

<style scoped>
.toolbar { margin-bottom: 12px; display: flex; gap: 12px; }
</style>
