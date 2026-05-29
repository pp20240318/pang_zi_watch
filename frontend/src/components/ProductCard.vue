<script setup>
import { computed } from 'vue'

const props = defineProps({
  product: { type: Object, required: true },
})

const orderTo = computed(() => ({
  path: '/order',
  query: { productId: String(props.product.id) },
}))

function formatPrice(price) {
  return price.toLocaleString('zh-CN')
}

function stars(rating) {
  const full = Math.floor(rating)
  const half = rating % 1 >= 0.5
  return { full, half, empty: 5 - full - (half ? 1 : 0) }
}

const starData = stars(props.product.rating)
</script>

<template>
  <article class="product-card">
    <div class="product-image-wrap">
      <img :src="product.image" :alt="product.name" class="product-image" loading="lazy" />
      <span v-if="product.tag" class="product-tag">{{ product.tag }}</span>
    </div>

    <div class="product-body">
      <h4 class="product-name">{{ product.name }}</h4>

      <div class="product-rating">
        <span class="stars" aria-label="评分 {{ product.rating }}">
          <span v-for="n in starData.full" :key="'f' + n" class="star full">★</span>
          <span v-if="starData.half" class="star half">★</span>
          <span v-for="n in starData.empty" :key="'e' + n" class="star empty">★</span>
        </span>
        <span class="rating-value">{{ product.rating }}</span>
        <span class="review-count">({{ product.reviews }} 条评价)</span>
      </div>

      <div class="product-footer">
        <div class="price">
          <span class="currency">¥</span>
          <span class="amount">{{ formatPrice(product.price) }}</span>
        </div>
        <router-link :to="orderTo" class="btn-order">
          立即下单
        </router-link>
      </div>
    </div>
  </article>
</template>

<style scoped>
.product-card {
  background: #fff;
  border: 1px solid var(--border);
  border-radius: 8px;
  overflow: hidden;
  transition: box-shadow 0.3s, transform 0.3s;
}

.product-card:hover {
  box-shadow: 0 12px 32px rgba(26, 35, 50, 0.1);
  transform: translateY(-4px);
}

.product-image-wrap {
  position: relative;
  aspect-ratio: 1;
  overflow: hidden;
  background: #f8f7f5;
}

.product-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.4s;
}

.product-card:hover .product-image {
  transform: scale(1.05);
}

.product-tag {
  position: absolute;
  top: 12px;
  left: 12px;
  padding: 4px 10px;
  background: var(--color-accent);
  color: #1a2332;
  font-size: 0.7rem;
  font-weight: 600;
  border-radius: 2px;
}

.product-body {
  padding: 16px;
}

.product-name {
  font-size: 0.95rem;
  font-weight: 500;
  color: var(--color-primary);
  margin: 0 0 10px;
  line-height: 1.4;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.product-rating {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 4px 6px;
  margin-bottom: 14px;
  font-size: 0.8rem;
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

.product-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  padding-top: 12px;
  border-top: 1px solid var(--border);
}

.price {
  color: #c0392b;
  font-weight: 600;
}

.currency {
  font-size: 0.85rem;
}

.amount {
  font-size: 1.15rem;
}

.btn-order {
  flex-shrink: 0;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 8px 14px;
  border: none;
  background: var(--color-primary);
  color: #fff;
  font-size: 0.8rem;
  border-radius: 4px;
  cursor: pointer;
  text-decoration: none;
  transition: background 0.2s;
}

.btn-order:hover {
  background: var(--color-accent);
  color: #1a2332;
}
</style>
