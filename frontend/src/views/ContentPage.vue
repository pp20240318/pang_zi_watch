<script setup>
import { ref, onMounted, computed, watch } from 'vue'
import { useRoute } from 'vue-router'
import AppHeader from '../components/AppHeader.vue'
import AppFooter from '../components/AppFooter.vue'
import { fetchPage } from '../api/public.js'

const route = useRoute()
const page = ref(null)
const loading = ref(true)
const error = ref('')

const slug = computed(() => route.params.slug)

async function loadPage() {
  loading.value = true
  error.value = ''
  try {
    page.value = await fetchPage(slug.value)
  } catch {
    error.value = '页面不存在或已下线'
    page.value = null
  } finally {
    loading.value = false
  }
}

onMounted(loadPage)
watch(slug, loadPage)
</script>

<template>
  <div class="content-page">
    <AppHeader />
    <main class="page-main">
      <div v-if="loading" class="page-state">加载中...</div>
      <div v-else-if="error" class="page-state">{{ error }}</div>
      <article v-else class="page-article">
        <h1 class="page-title">{{ page.title }}</h1>
        <div class="page-body" v-html="page.body" />
      </article>
    </main>
    <AppFooter />
  </div>
</template>

<style scoped>
.content-page {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.page-main {
  flex: 1;
  max-width: 800px;
  margin: 0 auto;
  padding: 48px 24px 64px;
  width: 100%;
}

.page-state {
  text-align: center;
  color: var(--color-muted, #666);
  padding: 80px 0;
}

.page-title {
  font-family: var(--font-serif);
  font-size: 1.75rem;
  color: var(--color-primary);
  margin: 0 0 32px;
  padding-bottom: 16px;
  border-bottom: 1px solid var(--border);
}

.page-body {
  line-height: 1.8;
  color: var(--color-text, #333);
  font-size: 0.95rem;
}

.page-body :deep(h2) {
  font-size: 1.25rem;
  color: var(--color-primary);
  margin: 32px 0 16px;
}

.page-body :deep(h3) {
  font-size: 1.05rem;
  color: var(--color-primary);
  margin: 24px 0 12px;
}

.page-body :deep(p) {
  margin: 0 0 16px;
}

.page-body :deep(ul),
.page-body :deep(ol) {
  margin: 0 0 16px;
  padding-left: 24px;
}

.page-body :deep(li) {
  margin-bottom: 8px;
}

.page-body :deep(strong) {
  color: var(--color-primary);
}
</style>
