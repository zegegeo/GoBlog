<template>
  <nav class="navbar">
    <div class="nav-brand">
      <router-link to="/" class="brand">GoBlog</router-link>
    </div>
    <!-- 未登录状态的菜单 -->
    <div v-if="!isLoggedIn" class="nav-menu">
      <router-link to="/login" class="nav-item">登录</router-link>
      <router-link to="/register" class="nav-item">注册</router-link>
      <router-link to="/about" class="nav-item">关于</router-link>
    </div>
    <!-- 登录状态的菜单 -->
    <div v-else class="nav-menu">
      <router-link to="/articles" class="nav-item">文章列表</router-link>
      <router-link to="/articles/create" class="nav-item">写文章</router-link>
      <a @click="handleLogout" class="nav-item logout">退出</a>
    </div>
  </nav>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'

const router = useRouter()
const isLoggedIn = ref(false)

const checkLoginStatus = () => {
  const token = localStorage.getItem('token')
  isLoggedIn.value = !!token
}

const handleLogout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  isLoggedIn.value = false
  ElMessage.success('已退出登录')
  router.push('/login')
}

// 监听登录状态变化
onMounted(() => {
  checkLoginStatus()
  // 添加事件监听器以检测登录状态变化
  window.addEventListener('storage', checkLoginStatus)
})
</script>

<style scoped>
.navbar {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 2rem;
  background: #222222;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);
  z-index: 1000;
}

.nav-brand {
  font-size: 1.5rem;
  font-weight: bold;
}

.brand {
  color: #00ff7f;
  text-decoration: none;
  font-weight: 700;
  font-size: 1.8rem;
}

.nav-menu {
  display: flex;
  gap: 0.8rem;
  align-items: center;
}

.nav-item {
  color: #ffffff;
  text-decoration: none;
  padding: 0.5rem 0.8rem;
  border-radius: 6px;
  transition: all 0.3s ease;
  font-size: 1.1rem;
  font-weight: 500;
}

.nav-item:hover {
  background: rgba(255, 255, 255, 0.1);
  transform: translateY(-1px);
}

.logout {
  cursor: pointer;
  color: #ff4757;
  font-weight: 600;
}

.logout:hover {
  background: rgba(255, 71, 87, 0.1);
}

@media (max-width: 768px) {
  .navbar {
    padding: 1rem;
  }

  .nav-menu {
    gap: 1rem;
  }

  .nav-item {
    padding: 0.5rem 1rem;
    font-size: 1rem;
  }
}
</style>