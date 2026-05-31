<script setup>
import { ref, computed, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import AppHeader from '../components/AppHeader.vue'
import AppFooter from '../components/AppFooter.vue'
import ProductCard from '../components/ProductCard.vue'
import { products, brands } from '../data/catalog.js'

const route = useRoute()
const router = useRouter()

const product = ref(null)

const brandName = computed(() => {
  if (!product.value) return ''
  const brand = brands.find((b) => b.id === product.value.brandId)
  return brand?.name ?? ''
})

const orderTo = computed(() => ({
  path: '/order',
  query: { productId: String(product.value?.id ?? '') },
}))

const relatedProducts = computed(() => {
  if (!product.value) return []
  return products
    .filter((p) => p.brandId === product.value.brandId && p.id !== product.value.id)
    .slice(0, 4)
})

const specs = [
  { label: '机芯类型', value: '瑞士自动机械' },
  { label: '表壳材质', value: '精钢' },
  { label: '防水深度', value: '100 米' },
  { label: '表镜', value: '蓝宝石水晶' },
  { label: '保修', value: '国际联保 2 年' },
]

function formatPrice(price) {
  return price.toLocaleString('zh-CN')
}

function stars(rating) {
  const full = Math.floor(rating)
  const half = rating % 1 >= 0.5
  return { full, half, empty: 5 - full - (half ? 1 : 0) }
}

function resolveProduct() {
  const raw = route.params.id
  if (raw == null || raw === '') return null
  return products.find((p) => String(p.id) === String(raw)) ?? null
}

watch(
  () => route.params.id,
  () => {
    const found = resolveProduct()
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
  <div class="detail-page">
    <AppHeader />

    <main v-if="product" class="page-container detail-main">
      <nav class="breadcrumb">
        <router-link to="/">首页</router-link>
        <span class="sep">/</span>
        <span>{{ brandName }}</span>
        <span class="sep">/</span>
        <span>{{ product.name }}</span>
      </nav>

      <div class="detail-layout">
        <section class="gallery">
          <div class="image-wrap">
            <img :src="product.image" :alt="product.name" class="main-image" />
            <span v-if="product.tag" class="product-tag">{{ product.tag }}</span>
          </div>
        </section>

        <section class="info-panel">
          <p class="brand-label">{{ brandName }}</p>
          <h1 class="product-title">{{ product.name }}</h1>

          <div class="product-rating">
            <span class="stars" :aria-label="`评分 ${product.rating}`">
              <span
                v-for="n in stars(product.rating).full"
                :key="'f' + n"
                class="star full"
              >★</span>
              <span v-if="stars(product.rating).half" class="star half">★</span>
              <span
                v-for="n in stars(product.rating).empty"
                :key="'e' + n"
                class="star empty"
              >★</span>
            </span>
            <span class="rating-value">{{ product.rating }}</span>
            <span class="review-count">({{ product.reviews }} 条评价)</span>
          </div>

          <div class="price-block">
            <span class="price-label">售价</span>
            <div class="price">
              <span class="currency">¥</span>
              <span class="amount">{{ formatPrice(product.price) }}</span>
            </div>
            <p class="price-note">含税价格 · 全国包邮</p>
          </div>

          <div class="actions">
            <router-link :to="orderTo" class="btn-primary">立即下单</router-link>
            <router-link to="/" class="btn-secondary">继续选购</router-link>
          </div>

          <ul class="highlights">
            <li>7 天无理由退换（未拆封）</li>
            <li>顺丰速运，安全配送</li>
          </ul>
        </section>
      </div>

      <section class="detail-tabs card">
        <h2>商品详情</h2>
        <p class="description">
          {{ product.name }} 来自 {{ brandName }}，融合经典设计与现代制表工艺。
          精选材质与精密机芯，适合日常佩戴与重要场合，展现独特品味与格调。
        </p>

        <h3>规格参数</h3>
        <dl class="spec-list">
          <div v-for="item in specs" :key="item.label" class="spec-row">
            <dt>{{ item.label }}</dt>
            <dd>{{ item.value }}</dd>
          </div>
        </dl>
      </section>

      <section v-if="relatedProducts.length" class="related-section">
        <h2>同系列推荐</h2>
        <div class="related-grid">
          <ProductCard
            v-for="item in relatedProducts"
            :key="item.id"
            :product="item"
          />
        </div>
      </section>
    </main>

    <AppFooter />
  </div>
</template>

<style scoped>
.detail-page {
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

.detail-main {
  flex: 1;
  padding: 32px 24px 48px;
}

.breadcrumb {
  font-size: 0.85rem;
  color: var(--color-muted);
  margin-bottom: 24px;
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

.detail-layout {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 48px;
  margin-bottom: 40px;
}

.image-wrap {
  position: relative;
  aspect-ratio: 1;
  border-radius: 8px;
  overflow: hidden;
  background: #fff;
  border: 1px solid var(--border);
}

.main-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.product-tag {
  position: absolute;
  top: 16px;
  left: 16px;
  padding: 6px 12px;
  background: var(--color-accent);
  color: #1a2332;
  font-size: 0.75rem;
  font-weight: 600;
  border-radius: 2px;
}

.brand-label {
  font-size: 0.85rem;
  color: var(--color-muted);
  text-transform: uppercase;
  letter-spacing: 0.05em;
  margin: 0 0 8px;
}

.product-title {
  font-size: 1.75rem;
  color: var(--color-primary);
  margin: 0 0 16px;
  line-height: 1.3;
}

.product-rating {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 4px 8px;
  margin-bottom: 24px;
  font-size: 0.9rem;
}

.stars {
  display: inline-flex;
  letter-spacing: -1px;
}

.star.full {
  color: var(--color-accent);
}

.star.half {
  color: var(--color-accent);
  opacity: 0.5;
}

.star.empty {
  color: #ddd;
}

.rating-value {
  font-weight: 600;
  color: var(--color-primary);
}

.review-count {
  color: var(--color-muted);
}

.price-block {
  padding: 20px 24px;
  background: #fff;
  border: 1px solid var(--border);
  border-radius: 8px;
  margin-bottom: 24px;
}

.price-label {
  display: block;
  font-size: 0.85rem;
  color: var(--color-muted);
  margin-bottom: 8px;
}

.price {
  color: #c0392b;
  font-weight: 600;
  line-height: 1;
}

.currency {
  font-size: 1.1rem;
}

.amount {
  font-size: 2rem;
}

.price-note {
  margin: 10px 0 0;
  font-size: 0.85rem;
  color: var(--color-muted);
}

.actions {
  display: flex;
  gap: 12px;
  margin-bottom: 24px;
}

.btn-primary,
.btn-secondary {
  flex: 1;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 14px 24px;
  font-size: 1rem;
  font-weight: 600;
  border-radius: 4px;
  text-decoration: none;
  transition: background 0.2s, border-color 0.2s;
}

.btn-primary {
  background: var(--color-primary);
  color: #fff;
}

.btn-primary:hover {
  background: var(--color-accent);
  color: #1a2332;
}

.btn-secondary {
  border: 1px solid var(--border);
  background: #fff;
  color: var(--color-text);
}

.btn-secondary:hover {
  border-color: var(--color-accent);
  color: var(--color-accent);
}

.highlights {
  margin: 0;
  padding: 0;
  list-style: none;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.highlights li {
  position: relative;
  padding-left: 20px;
  font-size: 0.9rem;
  color: var(--color-text);
}

.highlights li::before {
  content: '✓';
  position: absolute;
  left: 0;
  color: var(--color-accent);
  font-weight: 600;
}

.card {
  background: #fff;
  border: 1px solid var(--border);
  border-radius: 8px;
  padding: 28px 32px;
  margin-bottom: 40px;
}

.detail-tabs h2,
.related-section h2 {
  font-size: 1.25rem;
  color: var(--color-primary);
  margin: 0 0 20px;
}

.detail-tabs h3 {
  font-size: 1rem;
  color: var(--color-primary);
  margin: 28px 0 16px;
}

.description {
  margin: 0;
  line-height: 1.8;
  color: var(--color-text);
}

.spec-list {
  margin: 0;
}

.spec-row {
  display: grid;
  grid-template-columns: 140px 1fr;
  gap: 16px;
  padding: 12px 0;
  border-bottom: 1px solid var(--border);
  font-size: 0.9rem;
}

.spec-row:last-child {
  border-bottom: none;
}

.spec-row dt {
  margin: 0;
  color: var(--color-muted);
}

.spec-row dd {
  margin: 0;
  color: var(--color-primary);
}

.related-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(220px, 1fr));
  gap: 20px;
}

@media (max-width: 900px) {
  .detail-layout {
    grid-template-columns: 1fr;
    gap: 32px;
  }
}

@media (max-width: 768px) {
  .detail-main {
    padding: 24px 16px 40px;
  }

  .page-container {
    padding-left: 16px;
    padding-right: 16px;
  }

  .product-title {
    font-size: 1.4rem;
  }

  .actions {
    flex-direction: column;
  }

  .spec-row {
    grid-template-columns: 1fr;
    gap: 4px;
  }

  .card {
    padding: 20px;
  }
}
</style>
