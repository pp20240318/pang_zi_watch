<script setup>
import logoUrl from '../assets/logo.svg'
</script>

<template>
  <header class="site-header">
    <div class="header-inner">
      <a href="/" class="logo">
        <img :src="logoUrl" alt="胖子腕表" class="logo-icon" />
        <div class="logo-text">
          <span class="logo-name">胖子腕表</span>
          <span class="logo-slogan">PANGZI WATCHES</span>
        </div>
      </a>

      <nav class="header-nav">
        <a href="/" class="nav-link active">首页</a>
        <a href="#" class="nav-link">品牌馆</a>
        <a href="#" class="nav-link">新品上市</a>
        <a href="#" class="nav-link">售后服务</a>
      </nav>

      <div class="header-actions">
        <a href="#order-query" class="order-query">
          <svg viewBox="0 0 24 24" width="18" height="18" fill="none" stroke="currentColor" stroke-width="2">
            <circle cx="11" cy="11" r="8" />
            <path d="M21 21l-4.35-4.35" />
          </svg>
          订单查询
        </a>
        <button type="button" class="btn-cart" aria-label="购物车">
          <svg viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M6 6h15l-1.5 9h-12z" />
            <circle cx="9" cy="20" r="1" />
            <circle cx="18" cy="20" r="1" />
            <path d="M6 6L5 3H2" />
          </svg>
        </button>
      </div>
    </div>
  </header>
</template>

<style scoped>
.site-header {
  position: sticky;
  top: 0;
  z-index: 100;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(12px);
  border-bottom: 1px solid var(--border);
  box-shadow: 0 1px 12px rgba(26, 35, 50, 0.06);
}

.header-inner {
  max-width: 1280px;
  margin: 0 auto;
  padding: 0 24px;
  height: 72px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 24px;
}

.logo {
  display: flex;
  align-items: center;
  gap: 12px;
  text-decoration: none;
  flex-shrink: 0;
}

.logo-icon {
  width: 44px;
  height: 44px;
}

.logo-text {
  display: flex;
  flex-direction: column;
  line-height: 1.2;
}

.logo-name {
  font-family: var(--font-serif);
  font-size: 1.25rem;
  font-weight: 600;
  color: var(--color-primary);
  letter-spacing: 0.08em;
}

.logo-slogan {
  font-size: 0.65rem;
  color: var(--color-muted);
  letter-spacing: 0.2em;
}

.header-nav {
  display: flex;
  gap: 28px;
}

.nav-link {
  color: var(--color-text);
  text-decoration: none;
  font-size: 0.9rem;
  transition: color 0.2s;
}

.nav-link:hover,
.nav-link.active {
  color: var(--color-accent);
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 16px;
}

.order-query {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 8px 16px;
  border: 1px solid var(--color-accent);
  border-radius: 4px;
  color: var(--color-accent);
  text-decoration: none;
  font-size: 0.875rem;
  transition: background 0.2s, color 0.2s;
}

.order-query:hover {
  background: var(--color-accent);
  color: #fff;
}

.btn-cart {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  border: none;
  background: var(--color-bg-soft);
  border-radius: 50%;
  color: var(--color-primary);
  cursor: pointer;
  transition: background 0.2s;
}

.btn-cart:hover {
  background: var(--border);
}

@media (max-width: 900px) {
  .header-nav {
    display: none;
  }
}
</style>
