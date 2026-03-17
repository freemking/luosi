<template>
  <div class="login-page">
    <div class="login-background"></div>
    <a-row 
      justify="center" 
      align="middle" 
      style="min-height: 100vh; position: relative; z-index: 1; padding: 16px;"
    >
      <a-col 
        :xs="24" 
        :sm="18" 
        :md="14" 
        :lg="10" 
        :xl="8"
      >
        <div class="login-header">
          <h1>管理后台</h1>
          <p>Admin Management System</p>
        </div>
        <a-card 
          :bordered="false" 
          class="login-card"
          size="default"
        >
          <a-form
            ref="loginFormRef"
            :model="loginForm"
            @submit.prevent="handleLogin"
            :colon="false"
            :label-align="'left'"
          >
            <a-form-item
              label="用户名"
              name="username"
              :rules="[{ required: true, message: '请输入用户名' }]"
              class="form-item-spacing"
            >
              <a-input 
                v-model:value="loginForm.username"
                placeholder="请输入用户名"
                size="large"
                style="width: 100%"
              />
            </a-form-item>

            <a-form-item
              label="密&nbsp&nbsp&nbsp&nbsp码"
              name="password"
              :rules="[{ required: true, message: '请输入密码' }]"
              class="form-item-spacing"
            >
              <a-input-password 
                v-model:value="loginForm.password"
                placeholder="请输入密码"
                size="large"
                @keyup.enter="handleLogin"
                style="width: 100%"
              />
            </a-form-item>

            <a-form-item
              label="验证码"
              name="captcha"
              :rules="[{ required: true, message: '请输入验证码', whitespace: true }]"
              class="form-item-spacing"
            >
              <a-row :gutter="8" align="middle" style="width: 100%">
                <a-col :span="14">
                  <a-input 
                    v-model:value="loginForm.captcha"
                    placeholder="请输入验证码"
                    size="large"
                    style="width: 100%"
                  />
                </a-col>
                <a-col :span="10">
                  <img 
                    :src="captchaImage" 
                    @click="refreshCaptcha" 
                    alt="验证码" 
                    class="captcha-img"
                  />
                </a-col>
              </a-row>
            </a-form-item>

            <div v-if="error" class="error-message">
              {{ error }}
            </div>

            <a-form-item class="form-item-spacing submit-btn">
              <a-button 
                type="primary" 
                html-type="submit" 
                :loading="loading" 
                :disabled="loading"
                block
                size="large"
              >
                {{ loading ? '登录中...' : '登录' }}
              </a-button>
            </a-form-item>
          </a-form>
        </a-card>
      </a-col>
    </a-row>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import { message } from 'ant-design-vue'

const router = useRouter()
const authStore = useAuthStore()

const loginFormRef = ref(null)
const loginForm = ref({
  username: 'admin',
  password: '1qaz@WSX',
  captcha: ''
})

const captchaId = ref('')
const captchaImage = ref('')
const loading = ref(false)
const error = ref('')

const refreshCaptcha = async () => {
  try {
    const data = await authStore.getCaptcha()
    captchaId.value = data.captcha_id
    captchaImage.value = data.captcha
    loginForm.value.captcha = ''
  } catch (err) {
    message.error('获取验证码失败')
    console.error('Failed to refresh captcha:', err)
  }
}

const handleLogin = async () => {
  if (!loginFormRef.value) {
    return
  }
  try {
    await loginFormRef.value.validate()
  } catch (err) {
    // 检查是否是因为outOfDate导致的验证失败
    if (err.outOfDate) {
      try {
        // 重新验证
        await loginFormRef.value.validate()
      } catch (secondErr) {
        return
      }
    } else {
      return
    }
  }
  error.value = ''
  try {
    loading.value = true
    await authStore.login(
      loginForm.value.username,
      loginForm.value.password,
      loginForm.value.captcha.trim(),
      captchaId.value
    )
    message.success('登录成功')
    router.push('/')
  } catch (err) {
    error.value = authStore.error || '登录失败'
    message.error(error.value)
    refreshCaptcha()
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  refreshCaptcha()
})
</script>

<style scoped lang="less">
.login-page {
  width: 100vw;
  height: 100vh;
  position: fixed;
  top: 0;
  left: 0;
  overflow: hidden;
  background-color: #f5f5f5;
}

.login-background {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, #1890ff 0%, #096dd9 100%);
  z-index: 0;
}

.login-header {
  text-align: center;
  color: #fff;
  margin-bottom: 24px;
  padding: 0 8px;

  h1 {
    font-size: 32px;
    margin: 0 0 8px 0;
    font-weight: 600;
    text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    line-height: 1.25;
  }

  p {
    font-size: 14px;
    margin: 0;
    opacity: 0.9;
    letter-spacing: 2px;
    color: rgba(255, 255, 255, 0.85);
  }
}

.login-card {
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
  border-radius: 8px;
  padding: 8px;
  max-width: 460px;
  margin: 0 auto;

  :deep(.ant-card-body) {
    padding: 24px;
  }
}

.form-item-spacing {
  margin-bottom: 16px;
}

.captcha-img {
  width: 100%;
  height: 40px;
  object-fit: cover;
  cursor: pointer;
  border-radius: 4px;
  border: 1px solid #d9d9d9;
  transition: border-color 0.3s;

  &:hover {
    border-color: #1890ff;
  }
}

.error-message {
  color: #ff4d4f;
  margin-bottom: 16px;
  text-align: center;
  font-size: 14px;
  line-height: 1.5;
}

.submit-btn {
  margin-bottom: 0;
}

@media (max-width: 768px) {
  .login-header {
    margin-bottom: 16px;
    
    h1 {
      font-size: 28px;
    }
  }

  .login-card {
    :deep(.ant-card-body) {
      padding: 16px;
    }
  }

  :deep(.ant-form-item-label) {
    padding-bottom: 4px;
  }
}
</style>
