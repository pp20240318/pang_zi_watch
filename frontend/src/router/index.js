import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '../views/HomePage.vue'
import OrderPage from '../views/OrderPage.vue'
import PaymentPage from '../views/PaymentPage.vue'

const routes = [
  { path: '/', name: 'home', component: HomePage },
  { path: '/order', name: 'order', component: OrderPage },
  { path: '/payment', name: 'payment', component: PaymentPage },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior() {
    return { top: 0 }
  },
})

export default router
