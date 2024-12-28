<template>
  <DefaultLayout>
    <div class="article-container">
      <div class="article-box">
        <div v-if="article.id" class="article-header">
          <h1>{{ article.title }}</h1>
          <div class="article-meta">
            <span class="author">作者: {{ article.author?.username }}</span>
            <span class="separator">|</span>
            <span class="time">发布于: {{ article.created_at }}</span>
            <span class="separator">|</span>
            <span class="views">阅读: {{ article.view_count }}</span>
          </div>
        </div>
        <div v-if="article.id" class="article-content">
          {{ article.content }}
        </div>
        <div v-if="article.id && isAuthor" class="article-actions">
          <router-link :to="`/articles/${article.id}/edit`" class="edit-btn">
            编辑文章
          </router-link>
          <button @click="handleDelete" class="delete-btn">
            删除文章
          </button>
        </div>
        <div v-if="!article.id" class="loading">
          加载中...
        </div>
        <CommentSection v-if="article.id" :article-id="Number(route.params.id)" />
      </div>
    </div>
  </DefaultLayout>
</template>

<script setup>
import { ref, onMounted, defineAsyncComponent, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { articleApi } from '../api/article'

const DefaultLayout = defineAsyncComponent(() => import('../layouts/DefaultLayout.vue'))
const CommentSection = defineAsyncComponent(() => import('../components/CommentSection.vue'))

const route = useRoute()
const router = useRouter()
const article = ref({
  id: 0,
  title: '',
  content: '',
  author: {
    username: ''
  },
  created_at: '',
  view_count: 0
})

const currentUser = JSON.parse(localStorage.getItem('user') || '{}')
const isAuthor = computed(() => article.value.user_id === currentUser.id)

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
    console.log('Article detail response:', response)
    if (response.data && response.data.article) {
      article.value = response.data.article
    } else {
      throw new Error('Invalid article data')
    }
  } catch (error) {
    console.error('加载文章失败:', {
      status: error.response?.status,
      data: error.response?.data,
      message: error.message,
      stack: error.stack
    })
    if (error.response?.status === 401 || error.response?.data?.message === '未登录或登录已过期') {
      localStorage.removeItem('token')
      localStorage.removeItem('user')
      router.push('/login')
    } else if (error.response?.status === 404) {
      alert('文章不存在')
      router.push('/articles')
    } else {
      alert('加载文章失败: ' + (error.response?.data?.message || error.response?.data || error.message))
      router.push('/articles')
    }
  }
}

const handleDelete = async () => {
  if (!confirm('确定要删除这篇文章吗？')) return
  try {
    await articleApi.delete(route.params.id)
    alert('删除成功')
    router.push('/articles')
  } catch (error) {
    console.error('删除文章失败:', error)
    alert('删除文章失败')
  }
}

onMounted(() => {
  const token = localStorage.getItem('token')
  if (!token) {
    router.push('/login')
    return
  }

  const id = parseInt(route.params.id)
  if (!id) {
    console.error('Invalid article ID on mount:', route.params.id)
    router.push('/articles')
    return
  }

  console.log('Component mounted, loading article with ID:', id)
  loadArticle()
})
</script>

<style scoped>
.article-container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: flex-start;
  padding: 80px 2rem 2rem 2rem;
  background: #1a1a1a;
  position: absolute;
  width: 100%;
  top: 0;
  left: 0;
  min-height: 100%;
}

.article-box {
  background: #222222;
  padding: 2.5rem;
  border-radius: 16px;
  width: 100%;
  max-width: 1200px;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.2);
  margin-bottom: 2rem;
}

.article-header {
  margin-bottom: 2rem;
  border-bottom: 1px solid #333;
  padding-bottom: 1.5rem;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.article-header h1 {
  color: #ffffff;
  font-size: 2rem;
  font-weight: 600;
  margin-bottom: 1rem;
  width: 100%;
  text-align: center;
}

.article-meta {
  color: #666;
  font-size: 1rem;
  display: flex;
  gap: 1rem;
  align-items: center;
  margin: 0 auto;
  padding: 0.5rem 0;
}

.author, .time, .views {
  display: inline-flex;
  align-items: center;
  color: #888;
  white-space: nowrap;
}

.separator {
  color: #444;
  padding: 0 0.5rem;
  opacity: 0.5;
}

.article-content {
  color: #fff;
  line-height: 1.8;
  font-size: 1.1rem;
  margin: 2rem 0;
  text-align: justify;
  white-space: pre-wrap;
  min-height: 300px;
  padding: 2rem;
  background: #1a1a1a;
  border-radius: 8px;
}

.article-actions {
  display: flex;
  gap: 1rem;
  margin-top: 2rem;
  padding-top: 2rem;
  border-top: 1px solid #333;
  justify-content: flex-end;
}

.edit-btn,
.delete-btn {
  padding: 0.8rem 1.5rem;
  border-radius: 8px;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.3s ease;
  min-width: 120px;
  text-align: center;
}

.edit-btn {
  background: #ffd700;
  color: #000;
  text-decoration: none;
}

.delete-btn {
  background: #ff4757;
  color: #fff;
  border: none;
}

.edit-btn:hover,
.delete-btn:hover {
  transform: translateY(-2px);
  opacity: 0.9;
}

@media (max-width: 768px) {
  .article-container {
    padding: 70px 1rem 1rem 1rem;
  }

  .article-box {
    padding: 1.5rem;
  }

  .article-header h1 {
    font-size: 1.6rem;
  }

  .article-content {
    font-size: 1rem;
  }

  .edit-btn,
  .delete-btn {
    padding: 0.6rem 1rem;
    min-width: 100px;
    font-size: 0.9rem;
  }
}

.loading {
  text-align: center;
  padding: 2rem;
  color: #666;
  font-size: 1.2rem;
}
</style>