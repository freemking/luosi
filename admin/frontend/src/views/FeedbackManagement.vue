<template>
  <div class="feedback-management">
    <a-page-header title="留言管理" style="padding: 0 0 16px 0;" />
    <a-card :bordered="false">
      <a-skeleton :loading="loading" active>
        <a-table
          :data-source="feedbacks"
          :loading="loading"
          :columns="columns"
          rowKey="id"
          :scroll="{ x: 700 }"
          :pagination="{
            pageSize: 10,
            showSizeChanger: true,
            showQuickJumper: true,
            showTotal: (total) => `共 ${total} 条`,
            size: 'middle'
          }"
          :row-hover="true"
          :bordered="false"
        >
          <template #bodyCell="{ column, record }">
            <template v-if="column.key === 'action'">
              <a-button size="small" type="primary" @click="viewFeedback(record)">
                查看详情
              </a-button>
            </template>
          </template>
        </a-table>
      </a-skeleton>
    </a-card>

    <a-modal
      v-model:open="modalVisible"
      title="反馈详情"
      width="640px"
      destroyOnClose
    >
      <template #footer>
        <a-button @click="closeModal">关闭</a-button>
      </template>
      <div v-if="selectedFeedback">
        <a-descriptions :column="1" bordered :colon="false">
          <a-descriptions-item label="ID">{{ selectedFeedback.id }}</a-descriptions-item>
          <a-descriptions-item label="姓名">{{ selectedFeedback.name }}</a-descriptions-item>
          <a-descriptions-item label="邮箱">{{ selectedFeedback.email }}</a-descriptions-item>
          <a-descriptions-item label="电话">{{ selectedFeedback.phone }}</a-descriptions-item>
          <a-descriptions-item label="公司">{{ selectedFeedback.company }}</a-descriptions-item>
          <a-descriptions-item label="产品">{{ selectedFeedback.product }}</a-descriptions-item>
          <a-descriptions-item label="留言">
            <pre style="margin: 0; white-space: pre-wrap; word-break: break-word; font-family: inherit;">{{ selectedFeedback.message }}</pre>
          </a-descriptions-item>
          <a-descriptions-item label="创建时间">{{ selectedFeedback.created_at }}</a-descriptions-item>
        </a-descriptions>
      </div>
    </a-modal>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useFeedbackStore } from '../stores/auth'
import { message } from 'ant-design-vue'

const feedbackStore = useFeedbackStore()

const feedbacks = ref([])
const loading = ref(true)
const modalVisible = ref(false)
const selectedFeedback = ref(null)

const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    key: 'id',
    width: 60,
    fixed: 'left'
  },
  {
    title: '姓名',
    dataIndex: 'name',
    key: 'name',
    width: 120
  },
  {
    title: '邮箱',
    dataIndex: 'email',
    key: 'email',
    width: 180
  },
  {
    title: '产品',
    dataIndex: 'product',
    key: 'product',
    width: 120
  },
  {
    title: '创建时间',
    dataIndex: 'created_at',
    key: 'created_at',
    width: 180
  },
  {
    title: '操作',
    key: 'action',
    width: 120,
    fixed: 'right'
  }
]

const fetchFeedbacks = async () => {
  try {
    loading.value = true
    await feedbackStore.getFeedbacks()
    feedbacks.value = feedbackStore.feedbacks
  } catch (err) {
    message.error('获取反馈列表失败')
  } finally {
    loading.value = false
  }
}

const viewFeedback = async (feedback) => {
  try {
    const data = await feedbackStore.getFeedback(feedback.id)
    selectedFeedback.value = data
    modalVisible.value = true
  } catch (err) {
    message.error('获取反馈详情失败')
  }
}

const closeModal = () => {
  modalVisible.value = false
  selectedFeedback.value = null
}

onMounted(() => {
  fetchFeedbacks()
})
</script>

<style scoped lang="less">
.feedback-management {
  width: 100%;
}

@media (max-width: 768px) {
  .ant-descriptions-item-label {
    width: 80px !important;
  }
}
</style>
