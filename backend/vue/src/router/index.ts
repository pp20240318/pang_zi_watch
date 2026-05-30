import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../store/auth'
import LoginView from '../views/LoginView.vue'
import DashboardView from '../views/DashboardView.vue'
import BrandListView from '../views/BrandListView.vue'
import ProductListView from '../views/ProductListView.vue'
import BannerListView from '../views/BannerListView.vue'
import OrderListView from '../views/OrderListView.vue'
import AdminManagementView from '../views/AdminManagementView.vue'
import RoleManagementView from '../views/RoleManagementView.vue'
import ChangePasswordView from '../views/ChangePasswordView.vue'

const routes = [
  { path: '/login', name: 'login', component: LoginView },
  { path: '/', name: 'dashboard', component: DashboardView, meta: { requiresAuth: true } },
  { path: '/brands', name: 'brands', component: BrandListView, meta: { requiresAuth: true } },
  { path: '/products', name: 'products', component: ProductListView, meta: { requiresAuth: true } },
  { path: '/banners', name: 'banners', component: BannerListView, meta: { requiresAuth: true } },
  { path: '/orders', name: 'orders', component: OrderListView, meta: { requiresAuth: true } },
  { path: '/admins', name: 'admins', component: AdminManagementView, meta: { requiresAuth: true } },
  { path: '/roles', name: 'roles', component: RoleManagementView, meta: { requiresAuth: true } },
  { path: '/change-password', name: 'change-password', component: ChangePasswordView, meta: { requiresAuth: true } }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to) => {
  const auth = useAuthStore()
  if (to.meta.requiresAuth && !auth.token) {
    return { name: 'login' }
  }
  if (to.name === 'login' && auth.token) {
    return { name: 'dashboard' }
  }
})

export default router
