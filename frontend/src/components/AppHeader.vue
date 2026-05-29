<script setup>
import logoUrl from '../assets/logo.svg'
</script>

<template>
  <header class="site-header">
    <div class="header-inner">
      <router-link to="/" class="logo">
        <img :src="logoUrl" alt="胖子腕表" class="logo-icon" />
        <div class="logo-text">
          <span class="logo-name">胖子腕表</span>
          <span class="logo-slogan">PANGZI WATCHES</span>
        </div>
      </router-link>

      <nav class="header-nav">
        <a href="#" class="nav-link">售后服务</a>
        <a href="#order-query" class="nav-link">订单查询</a>
      </nav>
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

.nav-link:hover {
  color: var(--color-accent);
}

</style>
