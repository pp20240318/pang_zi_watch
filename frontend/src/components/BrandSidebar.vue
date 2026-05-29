<script setup>
defineProps({
  brands: { type: Array, required: true },
  activeId: { type: String, required: true },
})

const emit = defineEmits(['select'])
</script>

<template>
  <aside class="brand-sidebar">
    <h3 class="sidebar-title">品牌系列</h3>
    <ul class="brand-list">
      <li v-for="brand in brands" :key="brand.id">
        <button
          type="button"
          class="brand-item"
          :class="{ active: brand.id === activeId, 'brand-item--contact': brand.contact }"
          @click="emit('select', brand.id)"
        >
          <span class="brand-name">{{ brand.name }}</span>
          <span v-if="!brand.contact" class="brand-count">{{ brand.count }}</span>
        </button>
      </li>
    </ul>

    <div class="sidebar-promo">
      <img
        src="https://images.unsplash.com/photo-1612817159947-7b7d7f8b5e0e?w=400&q=80&auto=format&fit=crop"
        alt="精品腕表"
        class="promo-img"
      />
    </div>
  </aside>
</template>

<style scoped>
.brand-sidebar {
  background: #fff;
  border: 1px solid var(--border);
  border-radius: 8px;
  padding: 20px 0;
  height: fit-content;
  position: sticky;
  top: 96px;
}

.sidebar-title {
  font-family: var(--font-serif);
  font-size: 1.1rem;
  color: var(--color-primary);
  margin: 0 20px 16px;
  padding-bottom: 12px;
  border-bottom: 1px solid var(--border);
}

.brand-list {
  list-style: none;
  margin: 0;
  padding: 0;
}

.brand-item {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 20px;
  border: none;
  background: transparent;
  cursor: pointer;
  font-size: 0.9rem;
  color: var(--color-text);
  text-align: left;
  transition: background 0.2s, color 0.2s;
  border-left: 3px solid transparent;
}

.brand-item:hover {
  background: var(--color-bg-soft);
  color: var(--color-accent);
}

.brand-item.active {
  background: rgba(201, 162, 39, 0.08);
  color: var(--color-accent);
  border-left-color: var(--color-accent);
  font-weight: 600;
}

.brand-count {
  font-size: 0.75rem;
  color: var(--color-muted);
  background: var(--color-bg-soft);
  padding: 2px 8px;
  border-radius: 10px;
}

.brand-item.active .brand-count {
  background: rgba(201, 162, 39, 0.2);
  color: var(--color-accent);
}

.sidebar-promo {
  margin: 20px 16px 0;
  border-radius: 6px;
  overflow: hidden;
  position: relative;
}

.promo-img {
  width: 100%;
  height: 120px;
  object-fit: cover;
  display: block;
}

.brand-item--contact .brand-name {
  font-size: 0.82rem;
  line-height: 1.4;
}

.brand-item--contact {
  margin-top: 8px;
  border-top: 1px solid var(--border);
  padding-top: 16px;
}

.brand-item--contact.active {
  border-left-color: var(--color-accent);
}
</style>
