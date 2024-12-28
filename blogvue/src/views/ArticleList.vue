<template>
  <DefaultLayout>
    <div class="article-container">
      <div class="article-box">
        <div class="header">
          <h2>文章列表</h2>
          <router-link to="/articles/create" class="create-btn">
            <span>写文章</span>
          </router-link>
        </div>
        <div class="article-list">
          <div v-for="article in articles" :key="article.id" class="article-item">
            <div class="article-header">
              <div class="article-info">
                <h3>{{ article.title }}</h3>
                <p class="article-meta">
                  <span class="author">作者: {{ article.author?.username || '未知作者' }}</span>
                  <span class="time">发布于: {{ article.created_at }}</span>
                  <span class="views">阅读: {{ article.view_count }}</span>
                </p>
              </div>
            </div>
            <div class="article-actions">
              <router-link :to="`/articles/${article.id}`" class="action-btn view">
                查看
              </router-link>
              <template v-if="isAuthor(article)">
                <router-link :to="`/articles/${article.id}/edit`" class="action-btn edit">
                  编辑
                </router-link>
                <button @click="deleteArticle(article.id)" class="action-btn delete">
                  删除
                </button>
              </template>
            </div>
          </div>
        </div>
      </div>
    </div>
  </DefaultLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { articleApi } from '../api/article'
import DefaultLayout from '../layouts/DefaultLayout.vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'

const router = useRouter()

const articles = ref([])
const total = ref(0)
const currentPage = ref(1)
const pageSize = ref(10)
const currentUser = JSON.parse(localStorage.getItem('user') || '{}')

const isAuthor = (article) => {
  return article.user_id === currentUser.id
}

const loadArticles = async () => {
  try {
    console.log('开始获取文章列表')
    const response = await articleApi.getList(currentPage.value, pageSize.value)
    console.log('文章列表响应:', response.data)
    articles.value = response.data.articles
    total.value = response.data.total
  } catch (error) {
    console.error('加载文章失败:', error)
    if (error.response?.status === 401) {
      console.log('token失效，跳转到登录页')
      router.push('/login')
      return
    }
    alert('加载文章失败: ' + error.message)
  }
}

const deleteArticle = async (id) => {
  if (!confirm('确定要删除这篇文章吗？')) return
  try {
    await articleApi.delete(id)
    await loadArticles()
    ElMessage.success('删除成功')
  } catch (error) {
    console.error('删除文章失败:', error)
    ElMessage.error('删除文章失败: ' + (error.response?.data || error.message))
  }
}

onMounted(() => {
  const token = localStorage.getItem('token')
  console.log('当前token:', token)
  if (!token) {
    console.log('未找到token，跳转到登录页')
    router.push('/login')
    return
  }
  loadArticles()
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

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

.header h2 {
  color: #ffffff;
  font-size: 1.8rem;
  font-weight: 600;
}

.create-btn {
  background: #00ff7f;
  color: #000;
  padding: 0.8rem 1.5rem;
  border-radius: 8px;
  text-decoration: none;
  font-weight: 500;
  transition: all 0.3s ease;
}

.create-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 5px 15px rgba(0, 255, 127, 0.3);
}

.article-item {
  background: #1a1a1a;
  padding: 2rem;
  border-radius: 8px;
  margin-bottom: 1rem;
  transition: all 0.3s ease;
}

.article-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
}

.article-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}

.article-info {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.article-item h3 {
  color: #ffffff;
  font-size: 1.4rem;
  margin-bottom: 0.5rem;
  text-align: left;
}

.article-meta {
  display: flex;
  gap: 1rem;
  color: #666;
  font-size: 0.9rem;
}

.author {
  color: #00a8ff;
}

.time, .views {
  color: #666;
}

.article-actions {
  display: flex;
  gap: 1rem;
  margin-top: 1.5rem;
}

.action-btn {
  padding: 0.8rem 1.5rem;
  border-radius: 6px;
  border: none;
  cursor: pointer;
  font-size: 1rem;
  text-decoration: none;
  transition: all 0.3s ease;
  min-width: 100px;
  text-align: center;
}

.view {
  background: #00a8ff;
  color: #fff;
}

.edit {
  background: #ffd700;
  color: #000;
}

.delete {
  background: #ff4757;
  color: #fff;
}

.action-btn:hover {
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

  .article-item h3 {
    font-size: 1.2rem;
  }

  .action-btn {
    padding: 0.6rem 1rem;
    min-width: 80px;
    font-size: 0.9rem;
  }

  .article-item {
    padding: 1.5rem;
  }
}
</style>
