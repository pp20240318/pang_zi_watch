<template>
  <AppLayout title="修改密码">
    <el-card style="max-width:480px">
      <el-form label-width="100px">
        <el-form-item label="当前密码"><el-input v-model="form.old_password" type="password" show-password /></el-form-item>
        <el-form-item label="新密码"><el-input v-model="form.new_password" type="password" show-password /></el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="loading" @click="onSubmit">保存</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </AppLayout>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import AppLayout from '../components/AppLayout.vue'
import api from '../api/client'

const loading = ref(false)
const form = reactive({ old_password: '', new_password: '' })

async function onSubmit() {
  if (!form.old_password || form.new_password.length < 6) {
    ElMessage.warning('请填写当前密码，新密码至少 6 位')
    return
  }
  loading.value = true
  try {
    await api.post('/api/change-password', form)
    ElMessage.success('密码已修改')
    form.old_password = ''
    form.new_password = ''
  } catch (e: any) {
    ElMessage.error(e?.response?.data?.error || '修改失败')
  } finally {
    loading.value = false
  }
}
</script>
