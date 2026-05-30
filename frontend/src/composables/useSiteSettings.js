import { ref } from 'vue'
import { fetchSettings } from '../api/public.js'

export const defaultSiteSettings = {
  siteName: '胖子腕表',
  siteSlogan: 'PANGZI WATCHES',
  brandDescription: '专注正品名表销售，提供全国联保与专业售后服务。让每一秒都值得珍藏。',
  servicePhone: '400-888-6688',
  serviceHours: '9:00 - 21:00',
  copyright: '© 2026 胖子腕表 Pangzi Watches. 保留所有权利.',
  icpNumber: '京ICP备xxxxxxxx号',
  icpLink: '',
  contactEmail: 'service@pangziwatch.com',
  businessEmail: 'business@pangziwatch.com',
  companyAddress: '北京市朝阳区建国门外大街 1 号国贸写字楼 A 座 18 层',
}

const settings = ref({ ...defaultSiteSettings })
let loadPromise = null

export function useSiteSettings() {
  if (!loadPromise) {
    loadPromise = fetchSettings()
      .then((data) => {
        settings.value = { ...defaultSiteSettings, ...data }
      })
      .catch(() => {
        settings.value = { ...defaultSiteSettings }
      })
  }
  return { settings, ready: loadPromise }
}
