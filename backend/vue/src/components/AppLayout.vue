<template>
  <div class="layout">
    <Topbar :title="title" />
    <section class="body" :class="{ 'sidebar-collapsed': sidebarCollapsed }">
      <aside class="sider" :class="{ collapsed: sidebarCollapsed }">
        <SidebarMenu />
      </aside>
      <main class="content">
        <slot />
      </main>
    </section>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useSidebarStore } from '../store/sidebar'
import Topbar from './Topbar.vue'
import SidebarMenu from './SidebarMenu.vue'

withDefaults(defineProps<{ title?: string }>(), { title: '胖子腕表管理后台' })
const sidebarCollapsed = computed(() => useSidebarStore().collapsed)
</script>

<style scoped>
.layout { min-height: 100vh; display: flex; flex-direction: column; background: #f5f5f5; }
.body { flex: 1; display: grid; grid-template-columns: 220px 1fr; min-height: 0; }
.body.sidebar-collapsed { grid-template-columns: 64px 1fr; }
.sider { background: #fff; border-right: 1px solid #e5e6eb; overflow: hidden; }
.content { padding: 16px; overflow-y: auto; }
</style>
