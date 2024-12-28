<template>
  <div class="about-container">
    <div class="about-content">
      <h1>关于我们</h1>
      <div class="content-section">
        <p>欢迎来到我们的博客！我们致力于为用户提供最新的技术资讯、教程和分享。我们希望能够通过这个平台与大家交流、学习、共同进步。</p>
        
        <p>我们的团队由一群热爱技术的人组成，我们相信通过不断的努力和创新，我们能够为更多的用户提供有价值的内容和服务。</p>
    
        <p>如果你有任何问题或者想要与我们联系，欢迎通过以下方式联系我们：</p>
        
        <div class="contact-info">
          <div class="contact-item">
            <span class="label">Email:</span>
            <a href="mailto:support@example.com">support@example.com</a>
          </div>
          <div class="contact-item">
            <span class="label">Twitter:</span>
            <a href="https://twitter.com/your_profile" target="_blank">@GoBlog</a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "About"
}
</script>

<style scoped>
.about-container {
  min-height: 100vh;
  width: 100vw;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #1a1a1a;
  padding: 0;
  position: fixed;
  top: 0;
  left: 0;
  overflow-y: auto;
}

.about-content {
  background: #222222;
  border-radius: 16px;
  padding: 3rem;
  width: 90%;
  max-width: 800px;
  position: relative;
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.1);
  margin: auto;
  transform: translateY(2rem);
}

.about-content::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #00a8ff, #00ff7f);
  border-radius: 16px 16px 0 0;
}

h1 {
  color: #ffffff;
  font-size: 2.5rem;
  margin-bottom: 2rem;
  text-align: center;
  font-weight: 700;
  letter-spacing: -1px;
  text-shadow: 0 0 10px rgba(0, 168, 255, 0.3);
}

.content-section {
  color: #888888;
  font-size: 1.0rem;
  line-height: 1.8;
}

.content-section p {
  margin-bottom: 1.5rem;
}

.contact-info {
  margin-top: 2rem;
  padding: 1.5rem;
  background: #2a2a2a;
  border-radius: 8px;
}

.contact-item {
  display: flex;
  align-items: center;
  margin-bottom: 1rem;
}

.contact-item:last-child {
  margin-bottom: 0;
}

.label {
  color: #ffffff;
  font-weight: 500;
  margin-right: 1rem;
  min-width: 80px;
}

a {
  color: #00a8ff;
  text-decoration: none;
  transition: color 0.3s ease;
}

a:hover {
  color: #00ff7f;
}

@media (max-width: 768px) {
  .about-container {
    padding-top: 70px;
  }

  .about-content {
    padding: 2rem;
    transform: translateY(1rem);
  }

  h1 {
    font-size: 2rem;
  }

  .content-section {
    font-size: 1rem;
  }

  .contact-item {
    flex-direction: column;
    align-items: flex-start;
  }

  .label {
    margin-bottom: 0.5rem;
  }
}
</style>