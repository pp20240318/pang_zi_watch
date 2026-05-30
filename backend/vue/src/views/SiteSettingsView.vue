<template>
  <AppLayout title="站点设置">
    <el-form v-loading="loading" label-width="120px" style="max-width:640px">
      <el-divider content-position="left">品牌信息</el-divider>
      <el-form-item label="站点名称">
        <el-input v-model="form.siteName" placeholder="胖子腕表" />
      </el-form-item>
      <el-form-item label="英文标语">
        <el-input v-model="form.siteSlogan" placeholder="PANGZI WATCHES" />
      </el-form-item>
      <el-form-item label="品牌简介">
        <el-input v-model="form.brandDescription" type="textarea" :rows="3" />
      </el-form-item>

      <el-divider content-position="left">客服信息</el-divider>
      <el-form-item label="客服热线">
        <el-input v-model="form.servicePhone" placeholder="400-888-6688" />
      </el-form-item>
      <el-form-item label="服务时间">
        <el-input v-model="form.serviceHours" placeholder="9:00 - 21:00" />
      </el-form-item>
      <el-form-item label="客服邮箱">
        <el-input v-model="form.contactEmail" placeholder="service@example.com" />
      </el-form-item>
      <el-form-item label="商务邮箱">
        <el-input v-model="form.businessEmail" placeholder="business@example.com" />
      </el-form-item>
      <el-form-item label="公司地址">
        <el-input v-model="form.companyAddress" type="textarea" :rows="2" />
      </el-form-item>

      <el-divider content-position="left">页脚信息</el-divider>
      <el-form-item label="版权信息">
        <el-input v-model="form.copyright" placeholder="© 2026 胖子腕表. 保留所有权利." />
      </el-form-item>
      <el-form-item label="ICP 备案号">
        <el-input v-model="form.icpNumber" placeholder="京ICP备xxxxxxxx号" />
      </el-form-item>
      <el-form-item label="ICP 链接">
        <el-input v-model="form.icpLink" placeholder="https://beian.miit.gov.cn/" />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" :loading="saving" @click="onSave">保存设置</el-button>
      </el-form-item>
    </el-form>
  </AppLayout>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import AppLayout from '../components/AppLayout.vue'
import api from '../api/client'

interface SiteSetting {
  siteName: string
  siteSlogan: string
  brandDescription: string
  servicePhone: string
  serviceHours: string
  copyright: string
  icpNumber: string
  icpLink: string
  contactEmail: string
  businessEmail: string
  companyAddress: string
}

const loading = ref(false)
const saving = ref(false)
const form = reactive<SiteSetting>({
  siteName: '',
  siteSlogan: '',
  brandDescription: '',
  servicePhone: '',
  serviceHours: '',
  copyright: '',
  icpNumber: '',
  icpLink: '',
  contactEmail: '',
  businessEmail: '',
  companyAddress: '',
})

async function load() {
  loading.value = true
  try {
    const { data } = await api.get('/api/settings')
    Object.assign(form, data.data)
  } finally {
    loading.value = false
  }
}

async function onSave() {
  if (!form.siteName) {
    ElMessage.warning('请填写站点名称')
    return
  }
  saving.value = true
  try {
    await api.put('/api/settings', form)
    ElMessage.success('已保存')
  } finally {
    saving.value = false
  }
}

onMounted(load)
</script>
