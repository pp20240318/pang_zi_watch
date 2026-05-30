import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '../views/HomePage.vue'
import ProductDetailPage from '../views/ProductDetailPage.vue'
import OrderPage from '../views/OrderPage.vue'
import PaymentPage from '../views/PaymentPage.vue'
import ContentPage from '../views/ContentPage.vue'

const routes = [
  { path: '/', name: 'home', component: HomePage },
  { path: '/product/:id', name: 'product', component: ProductDetailPage },
  { path: '/order', name: 'order', component: OrderPage },
  { path: '/payment', name: 'payment', component: PaymentPage },
  { path: '/page/:slug', name: 'page', component: ContentPage },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to) {
    if (to.hash) {
      return { el: to.hash, behavior: 'smooth' }
    }
    return { top: 0 }
  },
})

export default router
