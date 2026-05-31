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
const activeImage = ref(0)
const activeTab = ref('description')
const selectedVariant = ref('default')
const isFavorite = ref(false)

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

const galleryImages = computed(() => {
  if (!product.value) return []
  if (product.value.gallery?.length) return product.value.gallery
  return [product.value.image]
})

const originalPrice = computed(() => {
  if (!product.value) return 0
  return Math.round(product.value.price * 1.12)
})

const discountPercent = computed(() => {
  if (!product.value || !originalPrice.value) return 0
  return Math.round((1 - product.value.price / originalPrice.value) * 100)
})

const deliveryDate = computed(() => {
  const d = new Date()
  d.setDate(d.getDate() + 3)
  return d.toLocaleDateString('zh-CN', { month: 'long', day: 'numeric' })
})

const specs = computed(() => {
  const common = [
    { label: '品牌', value: brandName.value },
    { label: '商品编号', value: `PZ-${String(product.value?.id ?? 0).padStart(6, '0')}` },
    { label: '机芯类型', value: '瑞士自动机械' },
    { label: '表壳材质', value: '904L 精钢' },
    { label: '表镜', value: '蓝宝石水晶，防反光涂层' },
    { label: '防水深度', value: '100 米 / 10 ATM' },
    { label: '表径', value: '41 mm' },
    { label: '表带', value: '蚝式精钢表带' },
    { label: '表扣', value: '折叠 Oysterlock 安全扣' },
    { label: '动力储存', value: '约 70 小时' },
    { label: '保修', value: '国际联保 2 年' },
    { label: '产地', value: '瑞士' },
  ]
  return common
})

const shortSpecs = computed(() => specs.value.slice(2, 7))

const variants = [
  { id: 'default', label: '经典黑面', priceOffset: 0 },
  { id: 'blue', label: '蓝色表盘', priceOffset: 3200 },
  { id: 'green', label: '绿色表盘', priceOffset: 5800 },
]

const displayPrice = computed(() => {
  if (!product.value) return 0
  const variant = variants.find((v) => v.id === selectedVariant.value)
  return product.value.price + (variant?.priceOffset ?? 0)
})

const ratingBreakdown = [
  { stars: 5, percent: 82 },
  { stars: 4, percent: 12 },
  { stars: 3, percent: 4 },
  { stars: 2, percent: 1 },
  { stars: 1, percent: 1 },
]

const reviews = [
  {
    id: 1,
    author: '腕表爱好者',
    rating: 5,
    date: '2026-04-18',
    text: '做工非常精致，走时准确，包装完好。顺丰次日达，验货无误，非常满意的一次购买体验。',
    helpful: 24,
  },
  {
    id: 2,
    author: '李**',
    rating: 5,
    date: '2026-03-02',
    text: '正品无疑，表盘细节到位，夜光效果清晰。客服回复及时，下单流程顺畅。',
    helpful: 18,
  },
  {
    id: 3,
    author: '王**',
    rating: 4,
    date: '2026-02-14',
    text: '整体很好，表带略紧，调整后佩戴舒适。价格相比专柜有优势，推荐。',
    helpful: 11,
  },
]

const tabs = [
  { id: 'description', label: '商品描述' },
  { id: 'specs', label: '规格参数' },
  { id: 'reviews', label: '用户评价' },
]

function formatPrice(price) {
  return price.toLocaleString('zh-CN')
}

function stars(rating) {
  const full = Math.floor(rating)
  const half = rating % 1 >= 0.5
  return { full, half, empty: 5 - full - (half ? 1 : 0) }
}

function selectImage(index) {
  activeImage.value = index
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
    activeImage.value = 0
    selectedVariant.value = 'default'
    activeTab.value = 'description'
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
        <span class="current">{{ product.name }}</span>
      </nav>

      <div class="detail-layout">
        <!-- 左侧图库 -->
        <section class="gallery">
          <div class="gallery-inner">
            <div class="thumbnails">
              <button
                v-for="(img, idx) in galleryImages"
                :key="idx"
                type="button"
                class="thumb-btn"
                :class="{ active: activeImage === idx }"
                :aria-label="`查看图片 ${idx + 1}`"
                @click="selectImage(idx)"
              >
                <img :src="img" :alt="`${product.name} 图 ${idx + 1}`" />
              </button>
            </div>
            <div class="main-image-wrap">
              <img
                :src="galleryImages[activeImage]"
                :alt="product.name"
                class="main-image"
              />
              <span v-if="product.tag" class="product-tag">{{ product.tag }}</span>
              <button
                type="button"
                class="favorite-btn"
                :class="{ active: isFavorite }"
                :aria-label="isFavorite ? '取消收藏' : '加入收藏'"
                @click="isFavorite = !isFavorite"
              >
                {{ isFavorite ? '♥' : '♡' }}
              </button>
            </div>
          </div>
        </section>

        <!-- 右侧购买区 -->
        <section class="buy-panel">
          <p class="brand-label">{{ brandName }}</p>
          <h1 class="product-title">{{ product.name }}</h1>

          <div class="meta-row">
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
              <button type="button" class="link-btn" @click="activeTab = 'reviews'">
                {{ product.reviews }} 条评价
              </button>
            </div>
            <span class="meta-divider">|</span>
            <span class="sku">货号 PZ-{{ String(product.id).padStart(6, '0') }}</span>
          </div>

          <!-- 价格 -->
          <div class="price-block">
            <div class="price-row">
              <span class="price-current">
                <span class="currency">¥</span>
                <span class="amount">{{ formatPrice(displayPrice) }}</span>
              </span>
              <span v-if="discountPercent > 0" class="price-original">
                ¥{{ formatPrice(originalPrice) }}
              </span>
              <span v-if="discountPercent > 0" class="discount-badge">
                -{{ discountPercent }}%
              </span>
            </div>
            <p class="price-note">含税价格 · 会员专享价 · 全国包邮</p>
          </div>

          <!-- 表盘款式 -->
          <div class="variant-block">
            <p class="block-label">表盘款式</p>
            <div class="variant-list">
              <button
                v-for="v in variants"
                :key="v.id"
                type="button"
                class="variant-btn"
                :class="{ active: selectedVariant === v.id }"
                @click="selectedVariant = v.id"
              >
                {{ v.label }}
                <span v-if="v.priceOffset" class="variant-extra">
                  +¥{{ formatPrice(v.priceOffset) }}
                </span>
              </button>
            </div>
          </div>

          <!-- 配送 -->
          <div class="delivery-block">
            <div class="delivery-row">
              <span class="delivery-icon">🚚</span>
              <div>
                <p class="delivery-title">预计 {{ deliveryDate }} 送达</p>
                <p class="delivery-sub">顺丰速运 · 保价配送 · 支持验货签收</p>
              </div>
            </div>
            <div class="delivery-row">
              <span class="delivery-icon">🛡</span>
              <div>
                <p class="delivery-title">7 天无理由退换（未拆封）</p>
                <p class="delivery-sub">正品保障 · 全国联保</p>
              </div>
            </div>
          </div>

          <!-- 操作按钮 -->
          <div class="actions">
            <router-link :to="orderTo" class="btn-primary">立即下单</router-link>
            <router-link to="/" class="btn-secondary">继续选购</router-link>
          </div>

          <!-- 卖家 -->
          <div class="seller-block">
            <div class="seller-avatar">PZ</div>
            <div class="seller-info">
              <p class="seller-name">胖子腕表官方旗舰</p>
              <p class="seller-meta">
                <span class="seller-rating">★ 4.9</span>
                <span>已售 2,800+ 件</span>
              </p>
            </div>
            <router-link to="/" class="seller-link">进店逛逛</router-link>
          </div>

          <!-- 简要规格 -->
          <div class="short-specs">
            <h3>主要参数</h3>
            <dl class="spec-mini-list">
              <div v-for="item in shortSpecs" :key="item.label" class="spec-mini-row">
                <dt>{{ item.label }}</dt>
                <dd>{{ item.value }}</dd>
              </div>
            </dl>
            <button type="button" class="link-btn" @click="activeTab = 'specs'">
              查看全部参数 →
            </button>
          </div>
        </section>
      </div>

      <!-- Tab 内容区 -->
      <section class="content-tabs">
        <div class="tab-nav">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            type="button"
            class="tab-btn"
            :class="{ active: activeTab === tab.id }"
            @click="activeTab = tab.id"
          >
            {{ tab.label }}
            <span v-if="tab.id === 'reviews'" class="tab-count">{{ product.reviews }}</span>
          </button>
        </div>

        <div class="tab-body card">
          <!-- 描述 -->
          <div v-show="activeTab === 'description'" class="tab-panel">
            <h2>关于此商品</h2>
            <div class="description-rich">
              <p>
                <strong>{{ product.name }}</strong> 来自 {{ brandName }}，采用瑞士自动机械机芯，
                融合经典设计与现代制表工艺。精钢表壳搭配蓝宝石水晶表镜，100 米防水性能，
                适合日常佩戴与重要场合。
              </p>
              <ul>
                <li>瑞士制造，COSC 天文台认证机芯</li>
                <li>904L 蚝式精钢表壳，耐腐蚀性优异</li>
                <li>Chromalight 夜光显示，暗光环境下清晰可读</li>
                <li>Oysterlock 安全扣，佩戴稳固可靠</li>
                <li>附赠官方保修卡、说明书及精美礼盒</li>
              </ul>
              <div class="desc-image">
                <img
                  :src="galleryImages[1] ?? product.image"
                  :alt="`${product.name} 细节展示`"
                />
              </div>
            </div>
          </div>

          <!-- 规格 -->
          <div v-show="activeTab === 'specs'" class="tab-panel">
            <h2>规格参数</h2>
            <dl class="spec-list">
              <div v-for="item in specs" :key="item.label" class="spec-row">
                <dt>{{ item.label }}</dt>
                <dd>{{ item.value }}</dd>
              </div>
            </dl>
          </div>

          <!-- 评价 -->
          <div v-show="activeTab === 'reviews'" class="tab-panel">
            <div class="reviews-header">
              <div class="reviews-summary">
                <span class="reviews-score">{{ product.rating }}</span>
                <div>
                  <div class="stars large">
                    <span
                      v-for="n in stars(product.rating).full"
                      :key="'rf' + n"
                      class="star full"
                    >★</span>
                    <span v-if="stars(product.rating).half" class="star half">★</span>
                    <span
                      v-for="n in stars(product.rating).empty"
                      :key="'re' + n"
                      class="star empty"
                    >★</span>
                  </div>
                  <p class="reviews-total">共 {{ product.reviews }} 条评价</p>
                </div>
              </div>
              <div class="rating-bars">
                <div
                  v-for="row in ratingBreakdown"
                  :key="row.stars"
                  class="rating-bar-row"
                >
                  <span class="bar-label">{{ row.stars }} 星</span>
                  <div class="bar-track">
                    <div class="bar-fill" :style="{ width: row.percent + '%' }" />
                  </div>
                  <span class="bar-pct">{{ row.percent }}%</span>
                </div>
              </div>
            </div>

            <div class="review-list">
              <article v-for="review in reviews" :key="review.id" class="review-card">
                <div class="review-top">
                  <span class="review-author">{{ review.author }}</span>
                  <span class="review-stars">
                    <span v-for="n in review.rating" :key="n" class="star full">★</span>
                  </span>
                  <span class="review-date">{{ review.date }}</span>
                </div>
                <p class="review-text">{{ review.text }}</p>
                <button type="button" class="helpful-btn">
                  有帮助 ({{ review.helpful }})
                </button>
              </article>
            </div>
          </div>
        </div>
      </section>

      <!-- 相关推荐 -->
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

    <!-- 移动端底部购买栏 -->
    <div v-if="product" class="mobile-buy-bar">
      <div class="mobile-price">
        <span class="mobile-amount">¥{{ formatPrice(displayPrice) }}</span>
      </div>
      <router-link :to="orderTo" class="mobile-buy-btn">立即下单</router-link>
    </div>

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
  padding: 24px 24px 80px;
}

.breadcrumb {
  font-size: 0.82rem;
  color: var(--color-muted);
  margin-bottom: 20px;
}

.breadcrumb a {
  color: var(--color-text);
  text-decoration: none;
}

.breadcrumb a:hover {
  color: var(--color-accent);
}

.breadcrumb .current {
  color: var(--color-muted);
}

.sep {
  margin: 0 6px;
}

/* ---- 主布局 ---- */
.detail-layout {
  display: grid;
  grid-template-columns: 1.1fr 1fr;
  gap: 40px;
  margin-bottom: 36px;
  align-items: start;
}

/* ---- 图库 ---- */
.gallery-inner {
  display: flex;
  gap: 12px;
  background: #fff;
  border: 1px solid var(--border);
  border-radius: 12px;
  padding: 16px;
}

.thumbnails {
  display: flex;
  flex-direction: column;
  gap: 8px;
  flex-shrink: 0;
}

.thumb-btn {
  width: 64px;
  height: 64px;
  padding: 0;
  border: 2px solid transparent;
  border-radius: 8px;
  overflow: hidden;
  cursor: pointer;
  background: #f8f7f5;
  transition: border-color 0.2s;
}

.thumb-btn.active {
  border-color: var(--color-accent);
}

.thumb-btn img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.main-image-wrap {
  position: relative;
  flex: 1;
  aspect-ratio: 1;
  border-radius: 8px;
  overflow: hidden;
  background: #f8f7f5;
}

.main-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.product-tag {
  position: absolute;
  top: 12px;
  left: 12px;
  padding: 5px 10px;
  background: var(--color-accent);
  color: #1a2332;
  font-size: 0.72rem;
  font-weight: 600;
  border-radius: 4px;
}

.favorite-btn {
  position: absolute;
  top: 12px;
  right: 12px;
  width: 40px;
  height: 40px;
  border: 1px solid var(--border);
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.92);
  font-size: 1.2rem;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--color-muted);
  transition: color 0.2s, border-color 0.2s;
}

.favorite-btn.active {
  color: #e74c3c;
  border-color: #e74c3c;
}

/* ---- 购买面板 ---- */
.buy-panel {
  background: #fff;
  border: 1px solid var(--border);
  border-radius: 12px;
  padding: 28px;
}

.brand-label {
  font-size: 0.8rem;
  color: var(--color-muted);
  text-transform: uppercase;
  letter-spacing: 0.06em;
  margin: 0 0 6px;
}

.product-title {
  font-size: 1.6rem;
  color: var(--color-primary);
  margin: 0 0 12px;
  line-height: 1.35;
}

.meta-row {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 20px;
  font-size: 0.88rem;
}

.product-rating {
  display: flex;
  align-items: center;
  gap: 4px 6px;
}

.stars {
  display: inline-flex;
  letter-spacing: -1px;
}

.stars.large {
  font-size: 1.1rem;
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

.link-btn {
  background: none;
  border: none;
  padding: 0;
  color: var(--color-accent);
  font-size: inherit;
  cursor: pointer;
  text-decoration: underline;
  text-underline-offset: 2px;
}

.meta-divider {
  color: var(--border);
}

.sku {
  color: var(--color-muted);
  font-size: 0.82rem;
}

/* 价格 */
.price-block {
  padding: 16px 20px;
  background: #faf9f7;
  border-radius: 8px;
  margin-bottom: 20px;
}

.price-row {
  display: flex;
  align-items: baseline;
  flex-wrap: wrap;
  gap: 10px;
}

.price-current {
  color: #c0392b;
  font-weight: 700;
  line-height: 1;
}

.currency {
  font-size: 1rem;
}

.amount {
  font-size: 2rem;
}

.price-original {
  color: var(--color-muted);
  text-decoration: line-through;
  font-size: 0.95rem;
}

.discount-badge {
  padding: 2px 8px;
  background: #fdecea;
  color: #c0392b;
  font-size: 0.78rem;
  font-weight: 600;
  border-radius: 4px;
}

.price-note {
  margin: 8px 0 0;
  font-size: 0.82rem;
  color: var(--color-muted);
}

/* 款式 */
.variant-block {
  margin-bottom: 20px;
}

.block-label {
  font-size: 0.85rem;
  color: var(--color-muted);
  margin: 0 0 10px;
}

.variant-list {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.variant-btn {
  padding: 8px 14px;
  border: 1px solid var(--border);
  border-radius: 6px;
  background: #fff;
  font-size: 0.85rem;
  cursor: pointer;
  transition: border-color 0.2s, background 0.2s;
}

.variant-btn.active {
  border-color: var(--color-accent);
  background: rgba(201, 162, 39, 0.08);
  color: var(--color-primary);
  font-weight: 600;
}

.variant-extra {
  display: block;
  font-size: 0.75rem;
  color: var(--color-muted);
  margin-top: 2px;
}

/* 配送 */
.delivery-block {
  display: flex;
  flex-direction: column;
  gap: 14px;
  padding: 16px 0;
  border-top: 1px solid var(--border);
  border-bottom: 1px solid var(--border);
  margin-bottom: 20px;
}

.delivery-row {
  display: flex;
  gap: 12px;
  align-items: flex-start;
}

.delivery-icon {
  font-size: 1.1rem;
  line-height: 1.4;
}

.delivery-title {
  margin: 0;
  font-size: 0.9rem;
  font-weight: 600;
  color: var(--color-primary);
}

.delivery-sub {
  margin: 2px 0 0;
  font-size: 0.8rem;
  color: var(--color-muted);
}

/* 按钮 */
.actions {
  display: flex;
  gap: 12px;
  margin-bottom: 20px;
}

.btn-primary,
.btn-secondary {
  flex: 1;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 14px 20px;
  font-size: 1rem;
  font-weight: 600;
  border-radius: 8px;
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

/* 卖家 */
.seller-block {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 14px 16px;
  background: #faf9f7;
  border-radius: 8px;
  margin-bottom: 20px;
}

.seller-avatar {
  width: 40px;
  height: 40px;
  border-radius: 8px;
  background: var(--color-primary);
  color: var(--color-accent);
  font-weight: 700;
  font-size: 0.85rem;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.seller-info {
  flex: 1;
  min-width: 0;
}

.seller-name {
  margin: 0;
  font-size: 0.9rem;
  font-weight: 600;
  color: var(--color-primary);
}

.seller-meta {
  margin: 2px 0 0;
  font-size: 0.78rem;
  color: var(--color-muted);
  display: flex;
  gap: 10px;
}

.seller-rating {
  color: var(--color-accent);
  font-weight: 600;
}

.seller-link {
  font-size: 0.82rem;
  color: var(--color-accent);
  text-decoration: none;
  white-space: nowrap;
}

.seller-link:hover {
  text-decoration: underline;
}

/* 简要规格 */
.short-specs h3 {
  font-size: 0.95rem;
  color: var(--color-primary);
  margin: 0 0 12px;
}

.spec-mini-list {
  margin: 0 0 10px;
}

.spec-mini-row {
  display: grid;
  grid-template-columns: 90px 1fr;
  gap: 8px;
  padding: 6px 0;
  font-size: 0.85rem;
  border-bottom: 1px dashed var(--border);
}

.spec-mini-row:last-child {
  border-bottom: none;
}

.spec-mini-row dt {
  margin: 0;
  color: var(--color-muted);
}

.spec-mini-row dd {
  margin: 0;
  color: var(--color-primary);
}

/* ---- Tabs ---- */
.content-tabs {
  margin-bottom: 40px;
}

.tab-nav {
  display: flex;
  gap: 4px;
  border-bottom: 2px solid var(--border);
  margin-bottom: 0;
}

.tab-btn {
  padding: 12px 24px;
  border: none;
  background: none;
  font-size: 0.95rem;
  font-weight: 500;
  color: var(--color-muted);
  cursor: pointer;
  position: relative;
  transition: color 0.2s;
}

.tab-btn.active {
  color: var(--color-primary);
  font-weight: 600;
}

.tab-btn.active::after {
  content: '';
  position: absolute;
  bottom: -2px;
  left: 0;
  right: 0;
  height: 2px;
  background: var(--color-accent);
}

.tab-count {
  margin-left: 4px;
  font-size: 0.78rem;
  color: var(--color-muted);
}

.card {
  background: #fff;
  border: 1px solid var(--border);
  border-top: none;
  border-radius: 0 0 12px 12px;
  padding: 28px 32px;
}

.tab-panel h2 {
  font-size: 1.15rem;
  color: var(--color-primary);
  margin: 0 0 20px;
}

.description-rich {
  line-height: 1.85;
  color: var(--color-text);
}

.description-rich p {
  margin: 0 0 16px;
}

.description-rich ul {
  margin: 0 0 24px;
  padding-left: 20px;
}

.description-rich li {
  margin-bottom: 8px;
}

.desc-image {
  border-radius: 8px;
  overflow: hidden;
  max-width: 640px;
}

.desc-image img {
  width: 100%;
  display: block;
}

.spec-list {
  margin: 0;
}

.spec-row {
  display: grid;
  grid-template-columns: 160px 1fr;
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

/* 评价 */
.reviews-header {
  display: grid;
  grid-template-columns: auto 1fr;
  gap: 40px;
  padding-bottom: 24px;
  border-bottom: 1px solid var(--border);
  margin-bottom: 24px;
}

.reviews-summary {
  display: flex;
  align-items: center;
  gap: 16px;
}

.reviews-score {
  font-size: 2.5rem;
  font-weight: 700;
  color: var(--color-primary);
  line-height: 1;
}

.reviews-total {
  margin: 6px 0 0;
  font-size: 0.85rem;
  color: var(--color-muted);
}

.rating-bars {
  display: flex;
  flex-direction: column;
  gap: 6px;
  max-width: 280px;
}

.rating-bar-row {
  display: grid;
  grid-template-columns: 36px 1fr 36px;
  align-items: center;
  gap: 8px;
  font-size: 0.78rem;
  color: var(--color-muted);
}

.bar-track {
  height: 6px;
  background: #eee;
  border-radius: 3px;
  overflow: hidden;
}

.bar-fill {
  height: 100%;
  background: var(--color-accent);
  border-radius: 3px;
}

.review-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.review-card {
  padding-bottom: 20px;
  border-bottom: 1px solid var(--border);
}

.review-card:last-child {
  border-bottom: none;
  padding-bottom: 0;
}

.review-top {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 8px;
  font-size: 0.85rem;
}

.review-author {
  font-weight: 600;
  color: var(--color-primary);
}

.review-stars {
  letter-spacing: -1px;
  font-size: 0.8rem;
}

.review-date {
  color: var(--color-muted);
  margin-left: auto;
}

.review-text {
  margin: 0 0 10px;
  line-height: 1.7;
  font-size: 0.9rem;
}

.helpful-btn {
  padding: 4px 12px;
  border: 1px solid var(--border);
  border-radius: 4px;
  background: #fff;
  font-size: 0.78rem;
  color: var(--color-muted);
  cursor: pointer;
}

.helpful-btn:hover {
  border-color: var(--color-accent);
  color: var(--color-accent);
}

/* 相关推荐 */
.related-section h2 {
  font-size: 1.25rem;
  color: var(--color-primary);
  margin: 0 0 20px;
}

.related-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(220px, 1fr));
  gap: 20px;
}

/* 移动端购买栏 */
.mobile-buy-bar {
  display: none;
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  z-index: 200;
  padding: 12px 16px;
  background: #fff;
  border-top: 1px solid var(--border);
  box-shadow: 0 -4px 16px rgba(26, 35, 50, 0.08);
  align-items: center;
  gap: 16px;
}

.mobile-price {
  flex-shrink: 0;
}

.mobile-amount {
  font-size: 1.25rem;
  font-weight: 700;
  color: #c0392b;
}

.mobile-buy-btn {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 12px;
  background: var(--color-primary);
  color: #fff;
  font-weight: 600;
  border-radius: 8px;
  text-decoration: none;
}

.mobile-buy-btn:hover {
  background: var(--color-accent);
  color: #1a2332;
}

@media (max-width: 900px) {
  .detail-layout {
    grid-template-columns: 1fr;
    gap: 20px;
  }

  .gallery-inner {
    flex-direction: column-reverse;
  }

  .thumbnails {
    flex-direction: row;
    overflow-x: auto;
  }

  .thumb-btn {
    flex-shrink: 0;
  }

  .reviews-header {
    grid-template-columns: 1fr;
    gap: 20px;
  }
}

@media (max-width: 768px) {
  .detail-main {
    padding: 16px 16px 72px;
  }

  .page-container {
    padding-left: 16px;
    padding-right: 16px;
  }

  .buy-panel {
    padding: 20px;
  }

  .product-title {
    font-size: 1.3rem;
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

  .tab-btn {
    padding: 10px 14px;
    font-size: 0.85rem;
  }

  .mobile-buy-bar {
    display: flex;
  }
}
</style>
