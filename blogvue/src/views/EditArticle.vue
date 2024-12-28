<template>
  <DefaultLayout>
    <div class="article-container">
      <div class="article-box">
        <h2>编辑文章</h2>
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
          <button @click="handleSubmit" class="submit-btn">保存修改</button>
          <button @click="$router.push(`/articles/${route.params.id}`)" class="cancel-btn">取消</button>
        </div>
      </div>
    </div>
  </DefaultLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { articleApi } from '../api/article'
import DefaultLayout from '../layouts/DefaultLayout.vue'
import { ElMessage } from 'element-plus'

const route = useRoute()
const router = useRouter()
const article = ref({
  title: '',
  content: ''
})

const loadArticle = async () => {
  try {
    const id = parseInt(route.params.id)
    if (!id) {
      console.error('Invalid article ID:', route.params.id)
      router.push('/articles')
      return
    }

    console.log('Loading article with ID:', id)
    const response = await articleApi.getDetail(id)
    if (response.data && response.data.article) {
      article.value = response.data.article
    } else {
      throw new Error('Invalid article data')
    }
  } catch (error) {
    console.error('加载文章失败:', error)
    if (error.response?.status === 401) {
      router.push('/login')
    } else {
      alert('加载文章失败: ' + (error.response?.data || error.message))
      router.push('/articles')
    }
  }
}

const handleSubmit = async () => {
  try {
    const id = parseInt(route.params.id)
    if (isNaN(id) || id <= 0) {
      throw new Error('无效的文章ID')
    }

    if (!article.value.title.trim() || !article.value.content.trim()) {
      ElMessage.warning('标题和内容不能为空')
      return
    }

    ElMessage.info('正在更新文章...')
    await articleApi.update(id, {
      title: article.value.title,
      content: article.value.content
    })
    ElMessage.success('更新成功')
    router.push(`/articles/${id}`)
  } catch (error) {
    console.error('更新文章失败:', error)
    ElMessage.error(error.response?.data || error.message || '更新失败')
  }
}

onMounted(() => {
  const token = localStorage.getItem('token')
  if (!token) {
    router.push('/login')
    return
  }
  loadArticle()
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