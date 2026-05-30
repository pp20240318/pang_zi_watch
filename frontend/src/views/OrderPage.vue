<script setup>
import { ref, computed, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import AppHeader from '../components/AppHeader.vue'
import AppFooter from '../components/AppFooter.vue'
import { products as fallbackProducts } from '../data/catalog.js'
import { fetchProduct, createOrder } from '../api/public.js'
import { savePendingOrder } from '../utils/orderStorage.js'

const route = useRoute()
const router = useRouter()

const product = ref(null)
const form = ref({
  name: '',
  email: '',
  phone: '',
  address: '',
  remark: '',
})
const errors = ref({})
const submitting = ref(false)

const totalPrice = computed(() => product.value?.price ?? 0)

function formatPrice(price) {
  return price.toLocaleString('zh-CN')
}

function validate() {
  const next = {}
  const name = form.value.name.trim()
  const email = form.value.email.trim()
  const phone = form.value.phone.trim()
  const address = form.value.address.trim()

  if (!name) next.name = '请填写收货人姓名'
  if (!email) {
    next.email = '请填写邮箱'
  } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) {
    next.email = '邮箱格式不正确'
  }
  if (!phone) {
    next.phone = '请填写手机号'
  } else if (!/^1\d{10}$/.test(phone)) {
    next.phone = '请输入 11 位有效手机号'
  }
  if (!address) next.address = '请填写收货地址'

  errors.value = next
  return Object.keys(next).length === 0
}

async function onSubmit() {
  if (!product.value || submitting.value) return
  if (!validate()) return

  submitting.value = true
  try {
    const data = await createOrder({
      productId: product.value.id,
      payAmount: totalPrice.value,
      customer: {
        name: form.value.name.trim(),
        email: form.value.email.trim(),
        phone: form.value.phone.trim(),
        address: form.value.address.trim(),
        remark: form.value.remark.trim(),
      },
    })
    const order = {
      orderNo: data.orderNo,
      product: data.product,
      customer: data.customer,
      payAmount: data.payAmount,
      createdAt: data.createdAt,
    }
    savePendingOrder(order)
    router.push({ name: 'payment' })
  } catch (e) {
    alert(e.message || '下单失败，请稍后重试')
  } finally {
    submitting.value = false
  }
}

async function resolveProduct() {
  const raw = route.query.productId
  if (raw == null || raw === '') return null
  try {
    return await fetchProduct(raw)
  } catch {
    return fallbackProducts.find((p) => String(p.id) === String(raw)) ?? null
  }
}

watch(
  () => route.query.productId,
  async () => {
    const found = await resolveProduct()
    if (!found) {
      router.replace({ name: 'home' })
      return
    }
    product.value = found
  },
  { immediate: true },
)
</script>

<template>
  <div class="order-page">
    <AppHeader />

    <main class="page-container order-main">
      <nav class="breadcrumb">
        <router-link to="/">首页</router-link>
        <span class="sep">/</span>
        <span>确认订单</span>
      </nav>

      <h1 class="page-title">填写订单信息</h1>

      <div v-if="product" class="order-layout">
        <section class="card product-summary">
          <h2>商品信息</h2>
          <div class="product-row">
            <img :src="product.image" :alt="product.name" class="thumb" />
            <div class="product-meta">
              <h3>{{ product.name }}</h3>
              <p class="price">¥{{ formatPrice(product.price) }}</p>
            </div>
          </div>
        </section>

        <section class="card form-section">
          <h2>收货与联系信息</h2>
          <form class="order-form" @submit.prevent="onSubmit">
            <div class="field">
              <label for="name">收货人 <span class="required">*</span></label>
              <input
                id="name"
                v-model="form.name"
                type="text"
                placeholder="请输入姓名"
                autocomplete="name"
              />
              <p v-if="errors.name" class="error">{{ errors.name }}</p>
            </div>

            <div class="field">
              <label for="email">邮箱 <span class="required">*</span></label>
              <input
                id="email"
                v-model="form.email"
                type="email"
                placeholder="用于接收订单通知"
                autocomplete="email"
              />
              <p v-if="errors.email" class="error">{{ errors.email }}</p>
            </div>

            <div class="field">
              <label for="phone">手机号 <span class="required">*</span></label>
              <input
                id="phone"
                v-model="form.phone"
                type="tel"
                placeholder="11 位手机号码"
                autocomplete="tel"
                maxlength="11"
              />
              <p v-if="errors.phone" class="error">{{ errors.phone }}</p>
            </div>

            <div class="field">
              <label for="address">收货地址 <span class="required">*</span></label>
              <textarea
                id="address"
                v-model="form.address"
                rows="3"
                placeholder="省市区 + 详细地址"
                autocomplete="street-address"
              />
              <p v-if="errors.address" class="error">{{ errors.address }}</p>
            </div>

            <div class="field">
              <label for="remark">备注（选填）</label>
              <input
                id="remark"
                v-model="form.remark"
                type="text"
                placeholder="配送时间、礼品包装等"
              />
            </div>

            <div class="form-actions">
              <router-link to="/" class="btn-secondary">返回选购</router-link>
              <button type="submit" class="btn-primary" :disabled="submitting">
                {{ submitting ? '提交中…' : '提交订单' }}
              </button>
            </div>
          </form>
        </section>

        <aside class="card order-aside">
          <h2>订单金额</h2>
          <dl class="amount-list">
            <div class="amount-row">
              <dt>商品金额</dt>
              <dd>¥{{ formatPrice(totalPrice) }}</dd>
            </div>
            <div class="amount-row">
              <dt>运费</dt>
              <dd class="free">免运费</dd>
            </div>
            <div class="amount-row total">
              <dt>应付合计</dt>
              <dd>¥{{ formatPrice(totalPrice) }}</dd>
            </div>
          </dl>
        </aside>
      </div>
    </main>

    <AppFooter />
  </div>
</template>

<style scoped>
.order-page {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background: var(--color-bg);
}

.page-container {
  max-width: 1280px;
  margin: 0 auto;
  width: 100%;
  padding-left: 24px;
  padding-right: 24px;
  box-sizing: border-box;
}

.order-main {
  flex: 1;
  padding: 32px 24px 48px;
}

.breadcrumb {
  font-size: 0.85rem;
  color: var(--color-muted);
  margin-bottom: 16px;
}

.breadcrumb a {
  color: var(--color-text);
  text-decoration: none;
}

.breadcrumb a:hover {
  color: var(--color-accent);
}

.sep {
  margin: 0 8px;
}

.page-title {
  font-size: 1.75rem;
  color: var(--color-primary);
  margin: 0 0 28px;
}

.order-layout {
  display: grid;
  grid-template-columns: 1fr 320px;
  grid-template-rows: auto auto;
  gap: 24px;
}

.product-summary {
  grid-column: 1;
}

.form-section {
  grid-column: 1;
}

.order-aside {
  grid-column: 2;
  grid-row: 1 / span 2;
  align-self: start;
  position: sticky;
  top: 96px;
}

.card {
  background: #fff;
  border: 1px solid var(--border);
  border-radius: 8px;
  padding: 24px;
}

.card h2 {
  font-size: 1.1rem;
  color: var(--color-primary);
  margin: 0 0 20px;
}

.product-row {
  display: flex;
  gap: 16px;
  align-items: center;
}

.thumb {
  width: 96px;
  height: 96px;
  object-fit: cover;
  border-radius: 6px;
  background: var(--color-bg-soft);
}

.product-meta h3 {
  font-size: 1rem;
  font-weight: 500;
  margin: 0 0 8px;
  color: var(--color-primary);
}

.price {
  margin: 0;
  color: #c0392b;
  font-weight: 600;
  font-size: 1.1rem;
}

.order-form {
  display: flex;
  flex-direction: column;
  gap: 18px;
}

.field label {
  display: block;
  font-size: 0.9rem;
  font-weight: 500;
  color: var(--color-primary);
  margin-bottom: 8px;
}

.required {
  color: #c0392b;
}

.field input,
.field textarea {
  width: 100%;
  padding: 10px 12px;
  border: 1px solid var(--border);
  border-radius: 4px;
  font-size: 0.95rem;
  font-family: inherit;
  transition: border-color 0.2s;
}

.field input:focus,
.field textarea:focus {
  outline: none;
  border-color: var(--color-accent);
}

.error {
  margin: 6px 0 0;
  font-size: 0.8rem;
  color: #c0392b;
}

.form-actions {
  display: flex;
  gap: 12px;
  justify-content: flex-end;
  padding-top: 8px;
}

.btn-primary,
.btn-secondary {
  padding: 12px 28px;
  font-size: 0.95rem;
  font-weight: 600;
  border-radius: 4px;
  text-decoration: none;
  cursor: pointer;
  transition: opacity 0.2s, background 0.2s;
}

.btn-primary {
  border: none;
  background: var(--color-primary);
  color: #fff;
}

.btn-primary:hover:not(:disabled) {
  background: var(--color-accent);
  color: #1a2332;
}

.btn-primary:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.btn-secondary {
  display: inline-flex;
  align-items: center;
  border: 1px solid var(--border);
  background: #fff;
  color: var(--color-text);
}

.btn-secondary:hover {
  border-color: var(--color-accent);
  color: var(--color-accent);
}

.amount-list {
  margin: 0;
}

.amount-row {
  display: flex;
  justify-content: space-between;
  padding: 10px 0;
  font-size: 0.9rem;
  border-bottom: 1px solid var(--border);
}

.amount-row dt {
  margin: 0;
  color: var(--color-muted);
}

.amount-row dd {
  margin: 0;
  font-weight: 500;
  color: var(--color-primary);
}

.amount-row .free {
  color: #27ae60;
}

.amount-row.total {
  border-bottom: none;
  padding-top: 16px;
  font-size: 1rem;
}

.amount-row.total dd {
  color: #c0392b;
  font-size: 1.25rem;
  font-weight: 600;
}

@media (max-width: 900px) {
  .order-layout {
    grid-template-columns: 1fr;
  }

  .order-aside {
    grid-column: 1;
    grid-row: auto;
    position: static;
  }
}

@media (max-width: 768px) {
  .order-main {
    padding: 24px 16px 40px;
  }

  .form-actions {
    flex-direction: column-reverse;
  }

  .btn-primary,
  .btn-secondary {
    width: 100%;
    justify-content: center;
    text-align: center;
  }
}
</style>
