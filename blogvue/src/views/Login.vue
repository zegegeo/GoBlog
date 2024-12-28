<template>
  <div class="login-container">
    <div class="login-box">
      <div class="form-header">
        <h2>欢迎回来</h2>
      </div>
      <form @submit.prevent="handleLogin">
        <div class="form-group">
          <input 
            v-model="form.username" 
            type="text" 
            placeholder="请输入用户名"
            required
          >
        </div>
        <div class="form-group">
          <input 
            v-model="form.password" 
            type="password" 
            placeholder="请输入密码"
            required
          >
        </div>
        <button type="submit" class="login-btn">
          <span>登录</span>
        </button>
      </form>
      <div class="form-footer">
        <p>还没有账号？ <router-link to="/register">立即注册</router-link></p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { userApi } from '../api/user'
import { ElMessage } from 'element-plus'

const router = useRouter()
const form = ref({
  username: '',
  password: ''
})

const handleLogin = async () => {
  try {
    ElMessage.info('正在登录...')
    const response = await userApi.login(form.value.username, form.value.password)
    
    console.log('登录响应:', response.data)
    
    // 保存token和用户信息
    const userData = {
      id: response.data.user.id,
      username: response.data.user.username,
      createdAt: response.data.user.created_at
    }
    
    localStorage.setItem('token', response.data.token)
    localStorage.setItem('user', JSON.stringify(userData))
    
    console.log('保存的用户信息:', {
      token: response.data.token,
      user: userData
    })
    
    ElMessage.success('登录成功')
    router.push('/articles')
  } catch (error) {
    console.error('Login error:', error)
    ElMessage.error('登录失败: ' + (error.response?.data || error.message))
  }
}
</script>

<style scoped>
.login-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #1a1a1a;
  padding: 20px;
  position: fixed;
  width: 100%;
  top: 0;
  left: 0;
  overflow: hidden;
}

.login-box {
  background: #222222;
  padding: 2.5rem;
  border-radius: 16px;
  width: 100%;
  max-width: 500px;
  min-height: 500px;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.2);
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.form-header {
  text-align: center;
  margin-bottom: 2.5rem;
}

.form-header h2 {
  color: #ffffff;
  font-size: 1.8rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
}

.form-header p {
  color: #888;
  font-size: 1rem;
}

.form-group {
  margin-bottom: 1.5rem;
}

label {
  display: block;
  color: #ffffff;
  margin-bottom: 0.5rem;
  font-size: 0.95rem;
}

input {
  width: 100%;
  padding: 0.8rem 1rem;
  border: 1px solid #444;
  border-radius: 8px;
  background: #1a1a1a;
  color: #fff;
  font-size: 1rem;
  transition: all 0.3s ease;
}

input:focus {
  border-color: #00a8ff;
  outline: none;
  box-shadow: 0 0 0 3px rgba(0, 168, 255, 0.1);
}

.login-btn {
  width: 100%;
  padding: 1rem;
  background: linear-gradient(135deg, #00a8ff, #00a1ff);
  border: none;
  border-radius: 8px;
  color: #fff;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  margin-top: 1rem;
}

.login-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 5px 15px rgba(0, 168, 255, 0.3);
}

.login-btn:active {
  transform: translateY(0);
}

.form-footer {
  text-align: center;
  margin-top: 2rem;
  color: #666;
  font-size: 0.9rem;
}

.form-footer a {
  color: #00a8ff;
  text-decoration: none;
  font-weight: 500;
  transition: color 0.3s;
}

.form-footer a:hover {
  color: #0097e6;
  text-decoration: underline;
}
</style>