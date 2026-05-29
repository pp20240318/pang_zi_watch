<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import AppHeader from '../components/AppHeader.vue'
import AppFooter from '../components/AppFooter.vue'
import { loadPendingOrder, clearPendingOrder } from '../utils/orderStorage.js'

const router = useRouter()
const order = ref(null)
const paymentMethod = ref('wechat')
const paying = ref(false)
const paid = ref(false)

const totalPrice = computed(() => order.value?.product?.price ?? 0)

function formatPrice(price) {
  return price.toLocaleString('zh-CN')
}

function onPay() {
  if (!order.value || paying.value || paid.value) return
  paying.value = true
  setTimeout(() => {
    paying.value = false
    paid.value = true
    clearPendingOrder()
  }, 1200)
}

function goHome() {
  router.push({ name: 'home' })
}

onMounted(() => {
  const pending = loadPendingOrder()
  if (!pending) {
    router.replace({ name: 'home' })
    return
  }
  order.value = pending
})
</script>

<template>
  <div class="payment-page">
    <AppHeader />

    <main class="page-container payment-main">
      <nav class="breadcrumb">
        <router-link to="/">首页</router-link>
        <span class="sep">/</span>
        <router-link :to="{ name: 'order', query: { productId: order?.product?.id } }">
          确认订单
        </router-link>
        <span class="sep">/</span>
        <span>支付</span>
      </nav>

      <h1 class="page-title">{{ paid ? '支付成功' : '订单支付' }}</h1>

      <div v-if="order" class="payment-layout">
        <section v-if="paid" class="card success-card">
          <div class="success-icon">✓</div>
          <h2>感谢您的购买</h2>
          <p>订单号：<strong>{{ order.orderNo }}</strong></p>
          <p class="hint">我们已向 {{ order.customer.email }} 发送订单确认（演示环境未真实发送）。</p>
          <button type="button" class="btn-primary" @click="goHome">返回首页</button>
        </section>

        <template v-else>
          <section class="card pay-methods">
            <h2>选择支付方式</h2>
            <label class="method-option">
              <input v-model="paymentMethod" type="radio" value="wechat" />
              <span class="method-label">
                <span class="method-icon wechat">微</span>
                微信支付
              </span>
            </label>
            <label class="method-option">
              <input v-model="paymentMethod" type="radio" value="alipay" />
              <span class="method-label">
                <span class="method-icon alipay">支</span>
                支付宝
              </span>
            </label>
            <label class="method-option">
              <input v-model="paymentMethod" type="radio" value="card" />
              <span class="method-label">
                <span class="method-icon card">卡</span>
                银行卡支付
              </span>
            </label>
          </section>

          <section class="card qr-panel">
            <h2>扫码支付（演示）</h2>
            <div class="qr-placeholder">
              <div class="qr-box" aria-hidden="true">
                <span class="qr-inner">模拟二维码</span>
              </div>
              <p>请使用{{ paymentMethod === 'wechat' ? '微信' : paymentMethod === 'alipay' ? '支付宝' : '银行 App' }}扫码完成支付</p>
            </div>
            <button
              type="button"
              class="btn-primary btn-pay"
              :disabled="paying"
              @click="onPay"
            >
              {{ paying ? '支付处理中…' : `确认支付 ¥${formatPrice(totalPrice)}` }}
            </button>
          </section>
        </template>

        <aside class="card order-summary">
          <h2>订单摘要</h2>
          <p class="order-no">订单号：{{ order.orderNo }}</p>
          <div class="summary-product">
            <img :src="order.product.image" :alt="order.product.name" />
            <div>
              <h3>{{ order.product.name }}</h3>
              <p class="price">¥{{ formatPrice(order.product.price) }}</p>
            </div>
          </div>
          <dl class="info-list">
            <div>
              <dt>收货人</dt>
              <dd>{{ order.customer.name }}</dd>
            </div>
            <div>
              <dt>手机</dt>
              <dd>{{ order.customer.phone }}</dd>
            </div>
            <div>
              <dt>邮箱</dt>
              <dd>{{ order.customer.email }}</dd>
            </div>
            <div>
              <dt>地址</dt>
              <dd>{{ order.customer.address }}</dd>
            </div>
          </dl>
          <div class="pay-total">
            <span>应付金额</span>
            <strong>¥{{ formatPrice(totalPrice) }}</strong>
          </div>
        </aside>
      </div>
    </main>

    <AppFooter />
  </div>
</template>

<style scoped>
.payment-page {
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

.payment-main {
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

.payment-layout {
  display: grid;
  grid-template-columns: 1fr 340px;
  gap: 24px;
  align-items: start;
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

.pay-methods {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.method-option {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 14px 16px;
  border: 1px solid var(--border);
  border-radius: 6px;
  cursor: pointer;
  transition: border-color 0.2s, background 0.2s;
}

.method-option:has(input:checked) {
  border-color: var(--color-accent);
  background: rgba(201, 162, 39, 0.08);
}

.method-option input {
  accent-color: var(--color-accent);
}

.method-label {
  display: flex;
  align-items: center;
  gap: 12px;
  font-weight: 500;
}

.method-icon {
  width: 36px;
  height: 36px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 0.85rem;
  font-weight: 700;
  color: #fff;
}

.method-icon.wechat {
  background: #07c160;
}

.method-icon.alipay {
  background: #1677ff;
}

.method-icon.card {
  background: var(--color-primary);
}

.qr-panel {
  margin-top: 24px;
}

.qr-placeholder {
  text-align: center;
  margin-bottom: 24px;
}

.qr-box {
  width: 200px;
  height: 200px;
  margin: 0 auto 16px;
  border: 2px dashed var(--border);
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: repeating-linear-gradient(
    45deg,
    var(--color-bg-soft),
    var(--color-bg-soft) 8px,
    #fff 8px,
    #fff 16px
  );
}

.qr-inner {
  font-size: 0.85rem;
  color: var(--color-muted);
  background: #fff;
  padding: 8px 12px;
  border-radius: 4px;
}

.qr-placeholder p {
  margin: 0;
  font-size: 0.9rem;
  color: var(--color-muted);
}

.btn-pay {
  width: 100%;
}

.btn-primary {
  padding: 14px 28px;
  font-size: 1rem;
  font-weight: 600;
  border: none;
  border-radius: 4px;
  background: var(--color-primary);
  color: #fff;
  cursor: pointer;
  transition: background 0.2s;
}

.btn-primary:hover:not(:disabled) {
  background: var(--color-accent);
  color: #1a2332;
}

.btn-primary:disabled {
  opacity: 0.65;
  cursor: not-allowed;
}

.order-summary {
  position: sticky;
  top: 96px;
}

.order-no {
  font-size: 0.8rem;
  color: var(--color-muted);
  margin: 0 0 16px;
}

.summary-product {
  display: flex;
  gap: 12px;
  padding-bottom: 16px;
  border-bottom: 1px solid var(--border);
  margin-bottom: 16px;
}

.summary-product img {
  width: 72px;
  height: 72px;
  object-fit: cover;
  border-radius: 6px;
}

.summary-product h3 {
  font-size: 0.9rem;
  margin: 0 0 6px;
  font-weight: 500;
}

.summary-product .price {
  margin: 0;
  color: #c0392b;
  font-weight: 600;
}

.info-list {
  margin: 0 0 20px;
}

.info-list > div {
  display: flex;
  gap: 12px;
  padding: 8px 0;
  font-size: 0.85rem;
  border-bottom: 1px solid var(--border);
}

.info-list dt {
  flex-shrink: 0;
  width: 56px;
  margin: 0;
  color: var(--color-muted);
}

.info-list dd {
  margin: 0;
  color: var(--color-primary);
  word-break: break-all;
}

.pay-total {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 8px;
  font-size: 0.95rem;
}

.pay-total strong {
  color: #c0392b;
  font-size: 1.35rem;
}

.success-card {
  grid-column: 1 / -1;
  text-align: center;
  padding: 48px 24px;
}

.success-icon {
  width: 72px;
  height: 72px;
  margin: 0 auto 20px;
  border-radius: 50%;
  background: #27ae60;
  color: #fff;
  font-size: 2rem;
  line-height: 72px;
  font-weight: 700;
}

.success-card h2 {
  margin-bottom: 12px;
}

.success-card p {
  color: var(--color-muted);
  margin: 0 0 8px;
}

.success-card .hint {
  margin-bottom: 28px;
}

@media (max-width: 900px) {
  .payment-layout {
    grid-template-columns: 1fr;
  }

  .order-summary {
    position: static;
    order: -1;
  }
}

@media (max-width: 768px) {
  .payment-main {
    padding: 24px 16px 40px;
  }
}
</style>
