<template>
  <AppLayout title="仪表盘">
    <el-row :gutter="16">
      <el-col :span="6" v-for="item in cards" :key="item.label">
        <el-card shadow="hover">
          <div class="stat-label">{{ item.label }}</div>
          <div class="stat-value">{{ item.value }}</div>
        </el-card>
      </el-col>
    </el-row>
  </AppLayout>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import AppLayout from '../components/AppLayout.vue'
import api from '../api/client'

const stats = ref<Record<string, number>>({})

onMounted(async () => {
  const { data } = await api.get('/api/dashboard')
  stats.value = data
})

const cards = computed(() => [
  { label: '商品数', value: stats.value.product_count ?? 0 },
  { label: '品牌数', value: stats.value.brand_count ?? 0 },
  { label: '待处理订单', value: stats.value.pending_orders ?? 0 },
  { label: '今日订单', value: stats.value.today_orders ?? 0 }
])
</script>

<style scoped>
.stat-label { color: #888; font-size: 14px; }
.stat-value { font-size: 28px; font-weight: 600; margin-top: 8px; }
</style>
