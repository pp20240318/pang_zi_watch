<script setup>
import { ref, computed } from 'vue'
import AppHeader from '../components/AppHeader.vue'
import BannerCarousel from '../components/BannerCarousel.vue'
import BrandSidebar from '../components/BrandSidebar.vue'
import ProductCard from '../components/ProductCard.vue'
import AppFooter from '../components/AppFooter.vue'
import { banners, brands, products } from '../data/catalog.js'

const activeBrand = ref('all')

const isCustomOrder = computed(() => activeBrand.value === 'custom-order')

const filteredProducts = computed(() => {
  if (isCustomOrder.value) return []
  if (activeBrand.value === 'all') return products
  return products.filter((p) => p.brandId === activeBrand.value)
})

</script>

<template>
  <div class="home-page">
    <AppHeader />
    <div class="page-container">
      <BannerCarousel :slides="banners" />
    </div>

    <main class="main-content page-container">
      <section class="catalog-section">
      

        <div class="catalog-layout">
          <BrandSidebar
            :brands="brands"
            :active-id="activeBrand"
            @select="activeBrand = $event"
          />

          <div class="product-area">
            <div v-if="isCustomOrder" class="custom-order-panel">
              <h3>其他品牌定制下单</h3>
              <p>如需劳力士、欧米茄以外的品牌或特殊款式，请联系在线客服为您报价并安排下单。</p>
              <a href="tel:4008886688" class="custom-order-cta">联系客服：400-888-6688</a>
            </div>
            <template v-else>
              <p class="result-hint">
                共 <strong>{{ filteredProducts.length }}</strong> 款商品
              </p>
              <div class="product-grid">
                <ProductCard
                  v-for="product in filteredProducts"
                  :key="product.id"
                  :product="product"
                />
              </div>
            </template>
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

.page-container {
  max-width: 1280px;
  margin: 0 auto;
  width: 100%;
  padding-left: 24px;
  padding-right: 24px;
  box-sizing: border-box;
}

.main-content {
  flex: 1;
  padding-top: 40px;
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

.custom-order-panel {
  padding: 48px 32px;
  text-align: center;
  background: #fff;
  border: 1px solid var(--border);
  border-radius: 8px;
}

.custom-order-panel h3 {
  font-family: var(--font-serif);
  font-size: 1.35rem;
  color: var(--color-primary);
  margin: 0 0 12px;
}

.custom-order-panel p {
  color: var(--color-muted);
  font-size: 0.95rem;
  line-height: 1.7;
  max-width: 480px;
  margin: 0 auto 24px;
}

.custom-order-cta {
  display: inline-block;
  padding: 12px 28px;
  background: var(--color-accent);
  color: #1a2332;
  text-decoration: none;
  font-weight: 600;
  border-radius: 4px;
  transition: opacity 0.2s;
}

.custom-order-cta:hover {
  opacity: 0.9;
}

@media (max-width: 768px) {
  .catalog-layout {
    grid-template-columns: 1fr;
  }

  .page-container {
    padding-left: 16px;
    padding-right: 16px;
  }

  .main-content {
    padding-top: 24px;
  }
}
</style>
