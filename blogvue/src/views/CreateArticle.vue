<template>
  <DefaultLayout>
    <div class="article-container">
      <div class="article-box">
        <h2>写文章</h2>
        <div class="form-group">
          <input 
            v-model="article.title" 
            type="text" 
            placeholder="文章标题"
            class="title-input"
          >
        </div>
        <div class="form-group">
          <textarea 
            v-model="article.content" 
            placeholder="文章内容..."
            class="content-input"
          ></textarea>
        </div>
        <div class="button-group">
          <button @click="handleSubmit" class="submit-btn">发布文章</button>
          <button @click="$router.push('/articles')" class="cancel-btn">取消</button>
        </div>
      </div>
    </div>
  </DefaultLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { articleApi } from '../api/article'
import DefaultLayout from '../layouts/DefaultLayout.vue'
import { ElMessage } from 'element-plus'

const router = useRouter()
const article = ref({
  title: '',
  content: ''
})

const handleSubmit = async () => {
  // 检查登录状态
  const token = localStorage.getItem('token')
  const user = JSON.parse(localStorage.getItem('user') || '{}')
  console.log('提交时的用户信息:', { token, user })

  if (!token || !user.id) {
    ElMessage.error('用户未登录或登录已过期')
    router.push('/login')
    return
  }

  if (!article.value.title.trim()) {
    ElMessage.warning('标题不能为空')
    return
  }
  if (!article.value.content.trim()) {
    ElMessage.warning('内容不能为空')
    return
  }

  try {
    ElMessage.info('正在发布文章...')
    console.log('发布文章:', {
      title: article.value.title,
      content: article.value.content,
      user: user
    })
    await articleApi.create(article.value)
    ElMessage.success('发布成功')
    router.push('/articles')
  } catch (error) {
    console.error('发布文章失败:', {
      error,
      response: error.response,
      data: error.response?.data,
      status: error.response?.status
    })
    if (error.response?.status === 401) {
      ElMessage.error('登录已过期，请重新登录')
      router.push('/login')
    } else {
      ElMessage.error(error.response?.data || error.message || '发布失败')
    }
  }
}

onMounted(() => {
  const token = localStorage.getItem('token')
  const user = JSON.parse(localStorage.getItem('user') || '{}')
  console.log('组件加载时的用户信息:', { token, user })
  if (!token) {
    ElMessage.warning('请先登录')
    router.push('/login')
    return
  }

  // 检查用户信息
  if (!user.id) {
    console.error('用户信息不完整:', user)
    localStorage.removeItem('token')
    localStorage.removeItem('user')
    ElMessage.error('用户信息不完整，请重新登录')
    router.push('/login')
    return
  }
})
</script>

<style scoped>
.article-container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: flex-start;
  padding: 6rem 2rem 2rem 2rem;
  background: #1a1a1a;
  position: absolute;
  width: 100%;
  top: 0;
  left: 0;
}

.article-box {
  background: #222222;
  padding: 2.5rem;
  border-radius: 16px;
  width: 100%;
  max-width: 1200px;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.2);
}

h2 {
  color: #ffffff;
  font-size: 1.8rem;
  margin-bottom: 2rem;
}

.form-group {
  margin-bottom: 1.5rem;
}

.title-input {
  width: 100%;
  padding: 1rem;
  font-size: 1.2rem;
  background: #1a1a1a;
  border: 1px solid #333;
  border-radius: 8px;
  color: #fff;
}

.content-input {
  width: 100%;
  min-height: 400px;
  padding: 1rem;
  font-size: 1.1rem;
  background: #1a1a1a;
  border: 1px solid #333;
  border-radius: 8px;
  color: #fff;
  resize: vertical;
  line-height: 1.6;
}

.button-group {
  display: flex;
  gap: 1rem;
  margin-top: 2rem;
}

.submit-btn, .cancel-btn {
  padding: 0.8rem 1.5rem;
  border-radius: 8px;
  border: none;
  cursor: pointer;
  font-size: 1rem;
  min-width: 120px;
}

.submit-btn {
  background: #00ff7f;
  color: #000;
}

.cancel-btn {
  background: #666;
  color: #fff;
}

.submit-btn:hover, .cancel-btn:hover {
  opacity: 0.9;
  transform: translateY(-2px);
}

@media (max-width: 768px) {
  .article-container {
    padding: 70px 1rem 1rem 1rem;
  }

  .article-box {
    padding: 1.5rem;
  }

  .content-input {
    min-height: 300px;
  }
}
</style> 