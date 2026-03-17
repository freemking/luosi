<template>
  <a-layout>
    <a-layout-sider 
      v-model="collapsed"
      :trigger="null"
      breakpoint="md"
      collapsed-width="0"
      :width="collapsed ? 0 : 208"
    >
      <div class="logo">
        <template v-if="!collapsed">
          <span class="logo-text">Admin Panel</span>
        </template>
      </div>
      <a-menu
        :selected-keys="selectedKeys"
        mode="inline"
        theme="dark"
        :inline-indent="collapsed ? 0 : 16"
      >
        <a-menu-item key="dashboard">
          <router-link to="/">
            <template #icon><DashboardOutlined /></template>
            <span>Dashboard</span>
          </router-link>
        </a-menu-item>
        <a-menu-item key="users" v-if="isSuperAdmin">
          <router-link to="/users">
            <template #icon><UserOutlined /></template>
            <span>用户管理</span>
          </router-link>
        </a-menu-item>
        <a-menu-item key="products">
          <router-link to="/products">
            <template #icon><AppstoreOutlined /></template>
            <span>产品管理</span>
          </router-link>
        </a-menu-item>
        <a-menu-item key="feedbacks">
          <router-link to="/feedbacks">
            <template #icon><MessageOutlined /></template>
            <span>反馈管理</span>
          </router-link>
        </a-menu-item>
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-header class="layout-header">
        <div class="header-left">
          <a-button 
            type="text" 
            @click="toggle"
            :icon="collapsed ? MenuUnfoldOutlined : MenuFoldOutlined"
          />
        </div>
        <div class="header-right">
          <a-space size="middle">
            <a-avatar :style="{ backgroundColor: '#1890ff' }">
              <template #icon><UserOutlined /></template>
            </a-avatar>
            <span class="username">{{ user?.username }}</span>
            <a-divider type="vertical" />
            <a-button @click="handleLogout" type="primary" danger size="small">
              退出登录
            </a-button>
          </a-space>
        </div>
      </a-header>
      <a-content class="layout-content">
        <router-view />
      </a-content>
    </a-layout>
  </a-layout>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { 
  DashboardOutlined, 
  UserOutlined, 
  AppstoreOutlined, 
  MessageOutlined,
  MenuUnfoldOutlined,
  MenuFoldOutlined
} from '@ant-design/icons-vue'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const collapsed = ref(false)
const selectedKeys = ref([
  window.location.pathname === '/' || window.location.pathname === '' 
    ? 'dashboard' 
    : window.location.pathname.substring(1)
])

const user = computed(() => authStore.user)
const isSuperAdmin = computed(() => authStore.isSuperAdmin)

const toggle = () => {
  collapsed.value = !collapsed.value
}

const handleLogout = () => {
  authStore.logout()
  router.push('/login')
}

const handleResize = () => {
  if (window.innerWidth < 768) {
    collapsed.value = true
  }
}

onMounted(() => {
  handleResize()
  window.addEventListener('resize', handleResize)
})

watch(() => window.location.pathname, () => {
  selectedKeys.value = [
    window.location.pathname === '/' || window.location.pathname === '' 
      ? 'dashboard' 
      : window.location.pathname.substring(1)
  ]
})
</script>

<style scoped lang="less">
.logo {
  height: 64px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #002140;
  padding: 0 16px;

  .logo-text {
    color: #fff;
    font-size: 16px;
    font-weight: 600;
    letter-spacing: 0.5px;
  }
}

.layout-header {
  background: #fff;
  padding: 0 16px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
  height: 64px;
  line-height: 64px;
}

.header-left {
  display: flex;
  align-items: center;
}

.header-right {
  display: flex;
  align-items: center;

  .username {
    color: #333333;
    font-size: 14px;
  }
}

.layout-content {
  margin: 24px 16px;
  padding: 24px;
  background: #fff;
  min-height: calc(100vh - 64px - 48px);
  border-radius: 6px;
  background-color: #f5f5f5;
}

@media (max-width: 768px) {
  .layout-content {
    margin: 16px 8px;
    padding: 16px;
  }

  .header-right {
    .username {
      display: none;
    }
  }
}
</style>
