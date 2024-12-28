<template>
  <div class="register-container">
    <div class="register-box">
      <h2>注册</h2>
      <form @submit.prevent="handleRegister">
        <div class="form-group">
          <input 
            v-model="form.username" 
            type="text" 
            placeholder="用户名"
            required
          >
        </div>
        <div class="form-group">
          <input 
            v-model="form.password" 
            type="password" 
            placeholder="密码"
            required
          >
        </div>
        <div class="form-group">
          <input 
            v-model="form.confirmPassword" 
            type="password" 
            placeholder="确认密码"
            required
          >
        </div>
        <button type="submit">注册</button>
        <div class="login-link">
          已有账号？<router-link to="/login">去登录</router-link>
        </div>
      </form>
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
  password: '',
  confirmPassword: ''
})

const handleRegister = async () => {
  if (form.value.password !== form.value.confirmPassword) {
    ElMessage.warning('两次输入的密码不一致')
    return
  }

  try {
    ElMessage.info('正在注册...')
    const data = {
      username: form.value.username.trim(),
      password: form.value.password.trim()
    }
    await userApi.register(data)
    ElMessage.success('注册成功')
    router.push('/login')
  } catch (error) {
    console.error('Register error:', error)
    ElMessage.error('注册失败: ' + (error.response?.data || error.message))
  }
}
</script>

<style scoped>
.register-container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #1a1a1a;
  padding: 20px;
  position: fixed;
  width: 100%;
  top: 0;
  left: 0;
  overflow: hidden;
}

.register-box {
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

h2 {
  color: #ffffff;
  text-align: center;
  margin-bottom: 2rem;
  font-size: 1.8rem;
  font-weight: 600;
}

.form-group {
  margin-bottom: 1.5rem;
}

input {
  width: 100%;
  padding: 0.8rem 1rem;
  border: 1px solid #333;
  border-radius: 8px;
  background: #1a1a1a;
  color: #fff;
  font-size: 1rem;
  transition: all 0.3s ease;
}

input:focus {
  border-color: #00ff7f;
  outline: none;
  box-shadow: 0 0 0 3px rgba(0, 255, 127, 0.1);
}

button {
  width: 100%;
  padding: 1rem;
  background: #00ff7f;
  color: #000;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  margin-top: 1rem;
  transition: all 0.3s ease;
}

button:hover {
  opacity: 0.9;
  transform: translateY(-2px);
  box-shadow: 0 5px 15px rgba(0, 255, 127, 0.3);
}

.login-link {
  text-align: center;
  margin-top: 2rem;
  color: #666;
  font-size: 0.9rem;
}

.login-link a {
  color: #00ff7f;
  text-decoration: none;
  font-weight: 500;
  transition: color 0.3s;
}

.login-link a:hover {
  text-decoration: underline;
}

form {
  width: 100%;
  max-width: 320px;
  margin: 0 auto;
}
</style>