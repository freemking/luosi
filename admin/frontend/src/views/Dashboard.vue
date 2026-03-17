<template>
  <div class="dashboard">
    <a-page-header title="Dashboard" style="padding: 0 0 24px 0;" />
    <a-row :gutter="16">
      <a-col :xs="24" :sm="12" :md="8" :lg="6">
        <a-card :bordered="true" class="stat-card">
          <template #title>
            <a-space>
              <UserOutlined style="color: #1890ff; font-size: 24px;" />
              <span>用户总数</span>
            </a-space>
          </template>
          <div class="stat-number">{{ stats.users }}</div>
        </a-card>
      </a-col>
      <a-col :xs="24" :sm="12" :md="8" :lg="6">
        <a-card :bordered="true" class="stat-card">
          <template #title>
            <a-space>
              <AppstoreOutlined style="color: #52c41a; font-size: 24px;" />
              <span>产品总数</span>
            </a-space>
          </template>
          <div class="stat-number">{{ stats.products }}</div>
        </a-card>
      </a-col>
      <a-col :xs="24" :sm="12" :md="8" :lg="6">
        <a-card :bordered="true" class="stat-card">
          <template #title>
            <a-space>
              <MessageOutlined style="color: #faad14; font-size: 24px;" />
              <span>反馈总数</span>
            </a-space>
          </template>
          <div class="stat-number">{{ stats.feedbacks }}</div>
        </a-card>
      </a-col>
    </a-row>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { UserOutlined, AppstoreOutlined, MessageOutlined } from '@ant-design/icons-vue'
import { useUserStore, useProductStore, useFeedbackStore } from '../stores/auth'

const userStore = useUserStore()
const productStore = useProductStore()
const feedbackStore = useFeedbackStore()

const stats = ref({
  users: 0,
  products: 0,
  feedbacks: 0
})

const fetchStats = async () => {
  try {
    await Promise.all([
      userStore.getUsers(),
      productStore.getProducts(),
      feedbackStore.getFeedbacks()
    ])
    stats.value.users = userStore.users.length
    stats.value.products = productStore.products.length
    stats.value.feedbacks = feedbackStore.feedbacks.length
  } catch (err) {
    console.error('Failed to fetch stats:', err)
  }
}

onMounted(() => {
  fetchStats()
})
</script>

<style scoped lang="less">
.dashboard {
  width: 100%;
}

.stat-card {
  margin-bottom: 16px;
  border-radius: 6px;
  transition: all 0.3s;

  &:hover {
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.09);
    transform: translateY(-2px);
  }

  :deep(.ant-card-head) {
    border-bottom: none;
    padding: 16px 16px 0 16px;
  }

  :deep(.ant-card-body) {
    padding: 16px;
  }
}

.stat-number {
  font-size: 40px;
  font-weight: bold;
  text-align: center;
  color: #333333;
  padding: 16px 0 8px 0;
  line-height: 1.2;
}

@media (max-width: 768px) {
  .stat-number {
    font-size: 32px;
  }
}
</style>
