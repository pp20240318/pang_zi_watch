<template>
  <div class="login-page">
    <el-card class="login-card">
      <h1>胖子腕表管理后台</h1>
      <p class="sub">默认账号 admin / admin123</p>
      <el-form @submit.prevent="onSubmit">
        <el-form-item>
          <el-input v-model="form.username" placeholder="用户名" size="large" />
        </el-form-item>
        <el-form-item>
          <el-input v-model="form.password" type="password" placeholder="密码" size="large" show-password @keyup.enter="onSubmit" />
        </el-form-item>
        <el-button type="primary" size="large" :loading="loading" style="width:100%" @click="onSubmit">登录</el-button>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { useAuthStore } from '../store/auth'

const router = useRouter()
const auth = useAuthStore()
const loading = ref(false)
const form = reactive({ username: '', password: '' })

async function onSubmit() {
  if (!form.username || !form.password) {
    ElMessage.warning('请输入用户名和密码')
    return
  }
  try {
    loading.value = true
    await auth.login(form.username, form.password)
    ElMessage.success('登录成功')
    router.push({ name: 'dashboard' })
  } catch (e: any) {
    ElMessage.error(e?.response?.data?.error || '登录失败')
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-page { min-height: 100vh; display: flex; align-items: center; justify-content: center; background: linear-gradient(135deg, #1a1a2e, #16213e); }
.login-card { width: 400px; padding: 8px; }
h1 { margin: 0 0 8px; font-size: 22px; text-align: center; }
.sub { text-align: center; color: #888; font-size: 13px; margin-bottom: 24px; }
</style>
