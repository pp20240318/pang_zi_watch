<template>
  <el-menu :default-active="active" router :collapse="sidebar.collapsed">
    <el-menu-item v-if="can('dashboard:view')" index="/" route="/">
      <el-icon><Odometer /></el-icon>
      <span>仪表盘</span>
    </el-menu-item>
    <el-menu-item v-if="canAny(['banner:read','banner:write'])" index="/banners" route="/banners">
      <el-icon><Picture /></el-icon>
      <span>轮播管理</span>
    </el-menu-item>
    <el-menu-item v-if="canAny(['brand:read','brand:write'])" index="/brands" route="/brands">
      <el-icon><Collection /></el-icon>
      <span>品牌管理</span>
    </el-menu-item>
    <el-menu-item v-if="canAny(['product:read','product:write'])" index="/products" route="/products">
      <el-icon><Goods /></el-icon>
      <span>商品管理</span>
    </el-menu-item>
    <el-menu-item v-if="canAny(['order:read','order:write'])" index="/orders" route="/orders">
      <el-icon><List /></el-icon>
      <span>订单管理</span>
    </el-menu-item>
    <el-sub-menu index="system">
      <template #title>
        <el-icon><Setting /></el-icon>
        <span>系统</span>
      </template>
      <el-menu-item v-if="can('admin:manage')" index="/admins" route="/admins">管理员</el-menu-item>
      <el-menu-item v-if="can('admin:roles')" index="/roles" route="/roles">角色权限</el-menu-item>
      <el-menu-item index="/change-password" route="/change-password">修改密码</el-menu-item>
    </el-sub-menu>
  </el-menu>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { useAuthStore } from '../store/auth'
import { useSidebarStore } from '../store/sidebar'

const route = useRoute()
const auth = useAuthStore()
const sidebar = useSidebarStore()
const active = computed(() => route.path)

function can(p: string) {
  return auth.hasPermission(p)
}
function canAny(perms: string[]) {
  return auth.hasAnyPermission(perms)
}
</script>
