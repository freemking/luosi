<template>
  <div class="product-management">
    <a-page-header title="产品管理" style="padding: 0 0 16px 0;" />
    <div class="toolbar">
      <a-button type="primary" @click="showAddModal">
        <PlusOutlined />
        新建产品
      </a-button>
    </div>
    <a-card bordered="false">
      <a-skeleton :loading="loading" active>
        <a-table
          :data-source="products"
          :loading="loading"
          :columns="columns"
          rowKey="id"
          :scroll="{ x: 800 }"
          :pagination="{
            pageSize: 10,
            showSizeChanger: true,
            showQuickJumper: true,
            showTotal: (total) => `共 ${total} 条`,
            size: 'middle'
          }"
          :row-hover="true"
          bordered="false"
        >
          <template #bodyCell="{ column, record }">
            <template v-if="column.key === 'images'">
              <div v-if="record.images && record.images.length > 0" class="image-list">
                <img 
                  v-for="img in record.images" 
                  :key="img.id" 
                  :src="img.image_url" 
                  class="image-thumb"
                  title="查看图片"
                />
              </div>
              <span v-else style="color: #666666">暂无图片</span>
            </template>
            <template v-if="column.key === 'action'">
              <a-space size="small">
                <a-button size="small" @click="handleEdit(record)">编辑</a-button>
                <a-button size="small" danger @click="showDeleteModal(record)">删除</a-button>
              </a-space>
            </template>
          </template>
        </a-table>
      </a-skeleton>
    </a-card>

    <a-modal
      v-model:open="deleteModalVisible"
      title="确认删除"
      @ok="handleDelete"
      :confirmLoading="deleting"
      ok-text="确认删除"
      cancel-text="取消"
    >
      <a-alert
        message="警告"
        description="确定要删除此产品吗？删除后无法恢复。"
        type="warning"
        show-icon
      />
    </a-modal>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { PlusOutlined, DeleteOutlined } from '@ant-design/icons-vue'
import { useProductStore } from '../stores/auth'
import { message } from 'ant-design-vue'

const router = useRouter()
const productStore = useProductStore()
const loading = ref(true)
const deleteModalVisible = ref(false)
const deleting = ref(false)
const currentId = ref(null)

const productForm = ref({
  name: '',
  description: '',
  category: '',
  standard: '',
  material: '',
  images: []
})

const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    key: 'id',
    width: 60,
    fixed: 'left'
  },
  {
    title: '产品名称',
    dataIndex: 'name',
    key: 'name',
    width: 150
  },
  {
    title: '分类',
    dataIndex: 'category',
    key: 'category',
    width: 120
  },
  {
    title: '材质',
    dataIndex: 'material',
    key: 'material',
    width: 120
  },
  {
    title: '图片',
    key: 'images',
    width: 200
  },
  {
    title: '操作',
    key: 'action',
    width: 120,
    fixed: 'right'
  }
]

const products = computed(() => productStore.products)

const fetchProducts = async () => {
  try {
    loading.value = true
    await productStore.getProducts()
  } catch (err) {
    message.error('获取产品列表失败')
  } finally {
    loading.value = false
  }
}

const showAddModal = () => {
  router.push('/products/create')
}

const handleEdit = (record) => {
  router.push(`/products/${record.id}`)
}

const showDeleteModal = (record) => {
  currentId.value = record.id
  deleteModalVisible.value = true
}

const handleDelete = async () => {
  try {
    deleting.value = true
    await productStore.deleteProduct(currentId.value)
    message.success('产品删除成功')
    deleteModalVisible.value = false
    fetchProducts()
  } catch (err) {
    message.error(productStore.error || '删除产品失败')
  } finally {
    deleting.value = false
  }
}

onMounted(() => {
  fetchProducts()
})
</script>

<style scoped lang="less">
.product-management {
  width: 100%;
}

.toolbar {
  margin-bottom: 16px;
  text-align: right;
}

.image-list {
  display: flex;
  gap: 4px;
  flex-wrap: wrap;
}

.image-thumb {
  width: 40px;
  height: 40px;
  object-fit: cover;
  border-radius: 4px;
  border: 1px solid #d9d9d9;
  transition: all 0.3s;
  cursor: pointer;

  &:hover {
    border-color: #1890ff;
    transform: scale(1.1);
  }
}

.image-item {
  display: flex;
  gap: 8px;
  margin-bottom: 8px;
  align-items: center;
}

@media (max-width: 768px) {
  .toolbar {
    text-align: left;
  }

  .image-item {
    flex-direction: column;
    align-items: stretch;

    .ant-input {
      width: 100% !important;
    }
  }
}
</style>
