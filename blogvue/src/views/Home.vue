<template>
  <div class="home-container">
    <div class="welcome-section">
      <div class="content">
        <h1>GoBlog</h1>
        <h2>分享你的故事</h2>
        <div class="auth-buttons">
          <button class="btn login-btn" @click="goToLogin">
            <span>立即登录</span>
            <div class="btn-bg"></div>
          </button>
          <button class="btn register-btn" @click="goToRegister">
            <span>免费注册</span>
            <div class="btn-bg"></div>
          </button>
        </div>
        <p class="welcome-text">欢迎使用博客系统，开启你的创作之旅</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router';

const router = useRouter();

// 跳转到登录页面
const goToLogin = () => {
  router.push('/login');
};

// 跳转到注册页面
const goToRegister = () => {
  router.push('/register');
};
</script>

<style scoped>
/* 添加全局样式重置 */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.home-container {
  min-height: 100vh;
  width: 100vw;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #1a1a1a;
  margin: 0;
  padding: 0;
  /* 防止水平滚动条 */
  overflow-x: hidden;
  position: fixed;
  top: 0;
  left: 0;
}

.welcome-section {
  background: #222222;
  width: 100vw;
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  padding: 0;
  margin: 0;
}

.welcome-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #00a8ff, #00ff7f);
}

.content {
  text-align: center;
}

h1 {
  color: #ffffff;
  margin-bottom: 1rem;
  font-size: 4.5rem;
  font-weight: 700;
  letter-spacing: -1px;
  text-shadow: 0 0 10px rgba(0, 168, 255, 0.3);
}

h2 {
  color: #888888;
  font-size: 2rem;
  font-weight: 400;
  margin-bottom: 4rem;
}

.auth-buttons {
  display: flex;
  justify-content: center;
  gap: 3rem;
  margin: 3rem 0;
}

.btn {
  padding: 1.2rem 5rem;
  border: none;
  border-radius: 8px;
  font-size: 1.3rem;
  cursor: pointer;
  min-width: 250px;
  position: relative;
  overflow: hidden;
  transition: all 0.3s ease;
  background: transparent;
}

.btn span {
  position: relative;
  z-index: 1;
  font-weight: 500;
}

.btn-bg {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  transition: transform 0.5s ease;
}

.btn:hover {
  transform: translateY(-3px);
  box-shadow: 0 5px 20px rgba(0, 168, 255, 0.2);
}

.login-btn {
  border: 2px solid #00a8ff;
  color: #00a8ff;
}

.login-btn:hover {
  color: #222222;
}

.login-btn .btn-bg {
  background: #00a8ff;
  transform: translateX(-100%);
}

.login-btn:hover .btn-bg {
  transform: translateX(0);
}

.register-btn {
  border: 2px solid #00ff7f;
  color: #00ff7f;
}

.register-btn:hover {
  color: #222222;
}

.register-btn .btn-bg {
  background: #00ff7f;
  transform: translateX(-100%);
}

.register-btn:hover .btn-bg {
  transform: translateX(0);
}

.welcome-text {
  color: #666666;
  font-size: 1.2rem;
  margin-top: 2rem;
  line-height: 1.6;
}

@media (max-width: 768px) {
  .welcome-section {
    padding: 2rem 1rem;
  }

  h1 {
    font-size: 3rem;
  }

  h2 {
    font-size: 1.5rem;
    margin-bottom: 3rem;
  }

  .auth-buttons {
    flex-direction: column;
    align-items: center;
    gap: 1.5rem;
  }

  .btn {
    width: 100%;
    max-width: 300px;
    padding: 1rem 2rem;
  }
}
</style>