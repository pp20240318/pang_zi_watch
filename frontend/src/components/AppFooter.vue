<script setup>
import { ref, computed, onMounted } from 'vue'
import { fetchPages } from '../api/public.js'

const FOOTER_CATEGORIES = ['购物指南', '售后服务', '关于我们']
const LEGAL_CATEGORY = '法律条款'

const fallbackPages = [
  { slug: 'shopping-flow', title: '购物流程', category: '购物指南' },
  { slug: 'payment-methods', title: '支付方式', category: '购物指南' },
  { slug: 'order-query', title: '订单查询', category: '购物指南' },
  { slug: 'delivery', title: '配送说明', category: '购物指南' },
  { slug: 'return-policy', title: '退换政策', category: '售后服务' },
  { slug: 'warranty', title: '保修条款', category: '售后服务' },
  { slug: 'service-centers', title: '维修网点', category: '售后服务' },
  { slug: 'faq', title: '常见问题', category: '售后服务' },
  { slug: 'brand-story', title: '品牌故事', category: '关于我们' },
  { slug: 'stores', title: '门店地址', category: '关于我们' },
  { slug: 'business', title: '商务合作', category: '关于我们' },
  { slug: 'contact', title: '联系我们', category: '关于我们' },
  { slug: 'privacy', title: '隐私政策', category: '法律条款' },
  { slug: 'terms', title: '用户协议', category: '法律条款' },
]

const pages = ref([...fallbackPages])

onMounted(async () => {
  try {
    const data = await fetchPages()
    if (data.length) pages.value = data
  } catch {
    /* 使用 fallbackPages */
  }
})

const footerGroups = computed(() =>
  FOOTER_CATEGORIES.map((category) => ({
    category,
    items: pages.value.filter((p) => p.category === category),
  })).filter((g) => g.items.length > 0)
)

const legalPages = computed(() =>
  pages.value.filter((p) => p.category === LEGAL_CATEGORY)
)
</script>

<template>
  <footer class="site-footer">
    <div class="footer-inner">
      <div class="footer-grid">
        <div class="footer-brand">
          <h4>胖子腕表</h4>
          <p>专注正品名表销售，提供全国联保与专业售后服务。让每一秒都值得珍藏。</p>
          <div class="footer-contact">
            <span>客服热线：400-888-6688</span>
            <span>服务时间：9:00 - 21:00</span>
          </div>
        </div>

        <div v-for="group in footerGroups" :key="group.category" class="footer-col">
          <h5>{{ group.category }}</h5>
          <ul>
            <li v-for="item in group.items" :key="item.slug">
              <router-link v-if="item.slug !== 'order-query'" :to="`/page/${item.slug}`">
                {{ item.title }}
              </router-link>
              <a v-else href="#order-query">{{ item.title }}</a>
            </li>
          </ul>
        </div>
      </div>

      <div id="order-query" class="order-query-block">
        <h5>订单查询</h5>
        <form class="query-form" @submit.prevent>
          <input type="text" placeholder="请输入邮箱" />
          <input type="tel" placeholder="预留手机号" />
          <button type="submit">查询</button>
        </form>
      </div>

      <div class="footer-bottom">
        <p>© 2026 胖子腕表 Pangzi Watches. 保留所有权利.</p>
        <p v-if="legalPages.length" class="footer-legal">
          <template v-for="(item, idx) in legalPages" :key="item.slug">
            <router-link :to="`/page/${item.slug}`">{{ item.title }}</router-link>
            <span v-if="idx < legalPages.length - 1">|</span>
          </template>
          <span>|</span>
          <span>ICP备案号：京ICP备xxxxxxxx号</span>
        </p>
        <p v-else class="footer-legal">
          <span>ICP备案号：京ICP备xxxxxxxx号</span>
        </p>
      </div>
    </div>
  </footer>
</template>

<style scoped>
.site-footer {
  background: var(--color-primary);
  color: rgba(255, 255, 255, 0.85);
  margin-top: 64px;
}

.footer-inner {
  max-width: 1280px;
  margin: 0 auto;
  padding: 48px 24px 24px;
}

.footer-grid {
  display: grid;
  grid-template-columns: 1.5fr 1fr 1fr 1fr;
  gap: 40px;
  padding-bottom: 40px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.12);
}

.footer-brand h4 {
  font-family: var(--font-serif);
  font-size: 1.35rem;
  color: var(--color-accent);
  margin: 0 0 12px;
}

.footer-brand p {
  font-size: 0.875rem;
  line-height: 1.7;
  margin: 0 0 16px;
  opacity: 0.8;
}

.footer-contact {
  display: flex;
  flex-direction: column;
  gap: 6px;
  font-size: 0.85rem;
  color: var(--color-accent);
}

.footer-col h5 {
  font-size: 0.95rem;
  color: #fff;
  margin: 0 0 16px;
}

.footer-col ul {
  list-style: none;
  margin: 0;
  padding: 0;
}

.footer-col li {
  margin-bottom: 10px;
}

.footer-col a {
  color: rgba(255, 255, 255, 0.7);
  text-decoration: none;
  font-size: 0.85rem;
  transition: color 0.2s;
}

.footer-col a:hover {
  color: var(--color-accent);
}

.order-query-block {
  padding: 32px 0;
  border-bottom: 1px solid rgba(255, 255, 255, 0.12);
}

.order-query-block h5 {
  margin: 0 0 16px;
  color: #fff;
}

.query-form {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
  max-width: 560px;
}

.query-form input {
  flex: 1;
  min-width: 160px;
  padding: 10px 14px;
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 4px;
  background: rgba(255, 255, 255, 0.08);
  color: #fff;
  font-size: 0.9rem;
}

.query-form input::placeholder {
  color: rgba(255, 255, 255, 0.45);
}

.query-form button {
  padding: 10px 28px;
  border: none;
  background: var(--color-accent);
  color: #1a2332;
  font-weight: 600;
  border-radius: 4px;
  cursor: pointer;
  transition: opacity 0.2s;
}

.query-form button:hover {
  opacity: 0.9;
}

.footer-bottom {
  padding-top: 24px;
  text-align: center;
  font-size: 0.8rem;
  opacity: 0.6;
}

.footer-bottom p {
  margin: 0 0 8px;
}

.footer-legal {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  gap: 8px;
}

.footer-legal a {
  color: inherit;
  text-decoration: none;
}

.footer-legal a:hover {
  color: var(--color-accent);
}

@media (max-width: 900px) {
  .footer-grid {
    grid-template-columns: 1fr 1fr;
  }
}

@media (max-width: 560px) {
  .footer-grid {
    grid-template-columns: 1fr;
  }
}
</style>
