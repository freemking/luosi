<template>
  <div class="product-detail">
    <a-page-header 
      :title="isEditing ? '编辑产品' : '新建产品'" 
      style="padding: 0 0 16px 0;"
    >
      <template #extra>
        <a-button @click="goBack">
          <ArrowLeftOutlined />
          返回列表
        </a-button>
      </template>
    </a-page-header>
    
    <a-card bordered="false">
      <a-form
        :model="productForm"
        layout="horizontal"
        :colon="false"
        :label-col="{ span: 6 }"
        :wrapper-col="{ span: 18 }"
      >
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item
              label="产品名称"
              name="name"
              :rules="[{ required: true, message: '请输入产品名称' }]"
            >
              <a-input v-model:value="productForm.name" placeholder="请输入产品名称" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item
              label="产品分类"
              name="category"
              :rules="[{ required: true, message: '请输入产品分类' }]"
            >
              <a-input v-model:value="productForm.category" placeholder="请输入产品分类" />
            </a-form-item>
          </a-col>
        </a-row>
        <a-form-item
          label="产品描述"
          name="description"
        >
          <a-textarea 
            v-model:value="productForm.description" 
            placeholder="请输入产品描述" 
            :rows="3"
          />
        </a-form-item>
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item
              label="规格"
              name="standard"
            >
              <a-input v-model:value="productForm.standard" placeholder="请输入规格" />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item
              label="材质"
              name="material"
            >
              <a-input v-model:value="productForm.material" placeholder="请输入材质" />
            </a-form-item>
          </a-col>
        </a-row>
        <a-form-item label="产品图片">
          <div v-for="(image, index) in productForm.images" :key="index" class="image-item">
            <a-input
              v-model:value="image.image_url"
              placeholder="图片URL"
              style="width: 85%; margin-right: 8px;"
            />
            <a-button danger @click="removeImage(index)"><DeleteOutlined /></a-button>
          </div>
          <a-button @click="addImage" type="dashed" block>
            <PlusOutlined />
            添加图片
          </a-button>
        </a-form-item>
      </a-form>
      
      <div class="action-buttons">
        <a-button @click="goBack">取消</a-button>
        <a-button type="primary" @click="handleSubmit" :loading="submitting">
          {{ isEditing ? '保存修改' : '创建产品' }}
        </a-button>
      </div>
    </a-card>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { PlusOutlined, DeleteOutlined, ArrowLeftOutlined } from '@ant-design/icons-vue'
import { useProductStore } from '../stores/auth'
import { message } from 'ant-design-vue'

const router = useRouter()
const route = useRoute()
const productStore = useProductStore()

const submitting = ref(false)
const productForm = ref({
  name: '',
  description: '',
  category: '',
  standard: '',
  material: '',
  images: []
})

const isEditing = computed(() => !!route.params.id)

const fetchProduct = async () => {
  if (isEditing.value) {
    try {
      const product = await productStore.getProduct(route.params.id)
      productForm.value.name = product.name || ''
      productForm.value.description = product.description || ''
      productForm.value.category = product.category || ''
      productForm.value.standard = product.standard || ''
      productForm.value.material = product.material || ''
      productForm.value.images = product.images ? [...product.images] : []
    } catch (err) {
      message.error('获取产品信息失败')
      goBack()
    }
  }
}

const addImage = () => {
  productForm.value.images.push({
    image_url: '',
    order: productForm.value.images.length
  })
}

const removeImage = (index) => {
  productForm.value.images.splice(index, 1)
}

const handleSubmit = async () => {
  try {
    submitting.value = true
    if (isEditing.value) {
      await productStore.updateProduct(route.params.id, productForm.value)
      message.success('产品更新成功')
    } else {
      await productStore.createProduct(productForm.value)
      message.success('产品创建成功')
    }
    goBack()
  } catch (err) {
    message.error(productStore.error || '保存产品失败')
  } finally {
    submitting.value = false
  }
}

const goBack = () => {
  router.push('/products')
}

onMounted(() => {
  fetchProduct()
})
</script>

<style scoped lang="less">
.product-detail {
  width: 100%;
}

.action-buttons {
  margin-top: 24px;
  text-align: right;
  
  .ant-btn {
    margin-left: 8px;
  }
}

.image-item {
  display: flex;
  gap: 8px;
  margin-bottom: 8px;
  align-items: center;
}

@media (max-width: 768px) {
  .image-item {
    flex-direction: column;
    align-items: stretch;

    .ant-input {
      width: 100% !important;
    }
  }
}
</style>