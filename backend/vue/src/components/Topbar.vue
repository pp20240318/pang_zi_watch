<template>
  <header class="topbar">
    <div class="left">
      <el-button text @click="sidebar.toggle()">
        <el-icon><Fold /></el-icon>
      </el-button>
      <span class="title">{{ title }}</span>
    </div>
    <div class="right">
      <span class="user">{{ auth.userInfo?.username }}</span>
      <el-button type="danger" text @click="onLogout">退出</el-button>
    </div>
  </header>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useAuthStore } from '../store/auth'
import { useSidebarStore } from '../store/sidebar'

defineProps<{ title?: string }>()
const auth = useAuthStore()
const sidebar = useSidebarStore()
const router = useRouter()

function onLogout() {
  auth.logout()
  router.push({ name: 'login' })
}
</script>

<style scoped>
.topbar {
  height: 56px; display: flex; align-items: center; justify-content: space-between;
  padding: 0 16px; background: #fff; border-bottom: 1px solid #e5e6eb;
}
.left { display: flex; align-items: center; gap: 8px; }
.title { font-weight: 600; font-size: 16px; }
.right { display: flex; align-items: center; gap: 12px; }
.user { color: #666; font-size: 14px; }
</style>
