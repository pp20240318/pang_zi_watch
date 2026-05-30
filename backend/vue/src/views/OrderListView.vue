<template>
  <AppLayout title="订单管理">
    <div class="toolbar">
      <el-input v-model="keyword" placeholder="订单号/手机/邮箱" style="width:240px" clearable @keyup.enter="load" />
      <el-select v-model="status" placeholder="状态" clearable style="width:140px" @change="load">
        <el-option label="待支付" value="pending" />
        <el-option label="已支付" value="paid" />
        <el-option label="已发货" value="shipped" />
        <el-option label="已完成" value="completed" />
        <el-option label="已取消" value="cancelled" />
      </el-select>
      <el-button type="primary" @click="load">查询</el-button>
    </div>
    <el-table :data="list" v-loading="loading" border>
      <el-table-column prop="order_no" label="订单号" width="180" />
      <el-table-column prop="product_name" label="商品" min-width="140" />
      <el-table-column label="金额" width="110">
        <template #default="{ row }">¥{{ row.pay_amount?.toLocaleString() }}</template>
      </el-table-column>
      <el-table-column prop="customer_name" label="收货人" width="90" />
      <el-table-column prop="customer_phone" label="手机" width="120" />
      <el-table-column prop="status" label="状态" width="90" />
      <el-table-column label="下单时间" width="170">
        <template #default="{ row }">{{ formatTime(row.created_at) }}</template>
      </el-table-column>
      <el-table-column label="操作" width="100">
        <template #default="{ row }">
          <el-button link type="primary" @click="openDetail(row)">详情</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      v-model:current-page="page"
      v-model:page-size="size"
      :total="total"
      layout="total, prev, pager, next"
      style="margin-top:12px"
      @current-change="load"
    />

    <el-dialog v-model="detailVisible" title="订单详情" width="520px">
      <el-descriptions v-if="current" :column="1" border>
        <el-descriptions-item label="订单号">{{ current.order_no }}</el-descriptions-item>
        <el-descriptions-item label="商品">{{ current.product_name }}</el-descriptions-item>
        <el-descriptions-item label="金额">¥{{ current.pay_amount?.toLocaleString() }}</el-descriptions-item>
        <el-descriptions-item label="收货人">{{ current.customer_name }}</el-descriptions-item>
        <el-descriptions-item label="手机">{{ current.customer_phone }}</el-descriptions-item>
        <el-descriptions-item label="邮箱">{{ current.customer_email }}</el-descriptions-item>
        <el-descriptions-item label="地址">{{ current.customer_address }}</el-descriptions-item>
        <el-descriptions-item label="备注">{{ current.customer_remark || '-' }}</el-descriptions-item>
      </el-descriptions>
      <el-form v-if="current" label-width="80px" style="margin-top:16px">
        <el-form-item label="状态">
          <el-select v-model="editStatus" style="width:100%">
            <el-option label="待支付" value="pending" />
            <el-option label="已支付" value="paid" />
            <el-option label="已发货" value="shipped" />
            <el-option label="已完成" value="completed" />
            <el-option label="已取消" value="cancelled" />
          </el-select>
        </el-form-item>
        <el-form-item label="内部备注"><el-input v-model="editRemark" type="textarea" /></el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="detailVisible = false">关闭</el-button>
        <el-button type="primary" @click="saveOrder">保存</el-button>
      </template>
    </el-dialog>
  </AppLayout>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { ElMessage } from 'element-plus'
import AppLayout from '../components/AppLayout.vue'
import api from '../api/client'

interface Order {
  id: number
  order_no: string
  product_name: string
  pay_amount: number
  customer_name: string
  customer_phone: string
  customer_email: string
  customer_address: string
  customer_remark: string
  status: string
  admin_remark: string
  created_at: string
}

const list = ref<Order[]>([])
const loading = ref(false)
const keyword = ref('')
const status = ref('')
const page = ref(1)
const size = ref(20)
const total = ref(0)
const detailVisible = ref(false)
const current = ref<Order | null>(null)
const editStatus = ref('')
const editRemark = ref('')

function formatTime(t: string) {
  return t ? new Date(t).toLocaleString('zh-CN') : ''
}

async function load() {
  loading.value = true
  const { data } = await api.get('/api/orders', {
    params: { page: page.value, size: size.value, q: keyword.value || undefined, status: status.value || undefined }
  })
  list.value = data.data
  total.value = data.total
  loading.value = false
}

function openDetail(row: Order) {
  current.value = row
  editStatus.value = row.status
  editRemark.value = row.admin_remark || ''
  detailVisible.value = true
}

async function saveOrder() {
  if (!current.value) return
  await api.patch(`/api/orders/${current.value.id}`, { status: editStatus.value, admin_remark: editRemark.value })
  ElMessage.success('已更新')
  detailVisible.value = false
  load()
}

onMounted(load)
</script>

<style scoped>
.toolbar { margin-bottom: 12px; display: flex; gap: 8px; flex-wrap: wrap; }
</style>
