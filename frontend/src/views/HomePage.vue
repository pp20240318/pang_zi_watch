<script setup>
import { ref, computed } from 'vue'
import AppHeader from '../components/AppHeader.vue'
import BannerCarousel from '../components/BannerCarousel.vue'
import BrandSidebar from '../components/BrandSidebar.vue'
import ProductCard from '../components/ProductCard.vue'
import AppFooter from '../components/AppFooter.vue'
import { banners, brands, products } from '../data/catalog.js'

const activeBrand = ref('all')

const filteredProducts = computed(() => {
  if (activeBrand.value === 'all') return products
  return products.filter((p) => p.brandId === activeBrand.value)
})

function onOrder(product) {
  alert(`已加入下单流程：${product.name}\n价格：¥${product.price.toLocaleString('zh-CN')}`)
}
</script>

<template>
  <div class="home-page">
    <AppHeader />
    <BannerCarousel :slides="banners" />

    <main class="main-content">
      <section class="catalog-section">
        <div class="section-header">
          <h2>精选腕表</h2>
          <p>正品行货 · 顺丰包邮 · 7天无理由退换</p>
        </div>

        <div class="catalog-layout">
          <BrandSidebar
            :brands="brands"
            :active-id="activeBrand"
            @select="activeBrand = $event"
          />

          <div class="product-area">
            <p class="result-hint">
              共 <strong>{{ filteredProducts.length }}</strong> 款商品
            </p>
            <div class="product-grid">
              <ProductCard
                v-for="product in filteredProducts"
                :key="product.id"
                :product="product"
                @order="onOrder"
              />
            </div>
          </div>
        </div>
      </section>
    </main>

    <AppFooter />
  </div>
</template>

<style scoped>
.home-page {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background: var(--color-bg);
}

.main-content {
  flex: 1;
  max-width: 1280px;
  margin: 0 auto;
  width: 100%;
  padding: 40px 24px 0;
  box-sizing: border-box;
}

.section-header {
  text-align: center;
  margin-bottom: 32px;
}

.section-header h2 {
  font-family: var(--font-serif);
  font-size: 1.75rem;
  color: var(--color-primary);
  margin: 0 0 8px;
}

.section-header p {
  color: var(--color-muted);
  font-size: 0.9rem;
  margin: 0;
}

.catalog-layout {
  display: grid;
  grid-template-columns: 220px 1fr;
  gap: 24px;
  align-items: start;
}

.product-area {
  min-width: 0;
}

.result-hint {
  font-size: 0.85rem;
  color: var(--color-muted);
  margin: 0 0 16px;
}

.result-hint strong {
  color: var(--color-accent);
}

.product-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(220px, 1fr));
  gap: 20px;
}

@media (max-width: 768px) {
  .catalog-layout {
    grid-template-columns: 1fr;
  }

  .main-content {
    padding: 24px 16px 0;
  }
}
</style>
