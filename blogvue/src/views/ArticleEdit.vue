<template>
  <DefaultLayout>
    <div class="article-container">
      <div class="article-box">
        <div class="form-header">
          <h2>{{ isEdit ? '编辑文章' : '写文章' }}</h2>
        </div>
        <form @submit.prevent="handleSubmit">
          <div class="form-group">
            <input 
              v-model="form.title" 
              type="text" 
              placeholder="请输入文章标题"
              required
            >
          </div>
          <div class="form-group">
            <textarea 
              v-model="form.content" 
              placeholder="请输入文章内容"
              rows="10"
              required
            ></textarea>
          </div>
          <button type="submit" class="submit-btn">
            {{ isEdit ? '保存修改' : '发布文章' }}
          </button>
        </form>
      </div>
    </div>
  </DefaultLayout>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { articleApi } from '../api/article'
import DefaultLayout from '../layouts/DefaultLayout.vue'

const route = useRoute()
const router = useRouter()
const isEdit = computed(() => route.params.id)

const form = ref({
  title: '',
  content: ''
})

onMounted(async () => {
  if (isEdit.value) {
    try {
      // 暂时使用模拟数据
      const mockArticle = {
        id: route.params.id,
        title: '如何使用 Go-Zero 构建微服务',
        content: '本文将介绍如何使用 Go-Zero 框架构建现代化的微服务应用。Go-Zero 是一个功能强大的微服务框架，提供了完整的工程化体系。\n\n' +
                '1. Go-Zero 的特点\n' +
                '- 完整的微服务架构\n' +
                '- 内置服务发现\n' +
                '- 内置负载均衡\n' +
                '- 内置熔断器\n\n' +
                '2. 如何开始\n' +
                '首先需要安装 Go-Zero 工具链...'
      }
      form.value = {
        title: mockArticle.title,
        content: mockArticle.content
      }
      // 后续替换为实际 API 调用
      // const response = await articleApi.getDetail(route.params.id)
      // form.value = {
      //   title: response.data.title,
      //   content: response.data.content
      // }
    } catch (error) {
      console.error('加载文章失败:', error)
      alert('加载文章失败')
      router.push('/articles')
    }
  }
})

const handleSubmit = async () => {
  try {
    console.log('Submitting article:', form.value)
    if (isEdit.value) {
      await articleApi.update(route.params.id, form.value)
      alert('更新成功')
    } else {
      const response = await articleApi.create(form.value)
      console.log('Create article response:', response)
      alert('发布成功')
    }
    router.push('/articles')
  } catch (error) {
    console.error('保存文章失败:', {
      error,
      response: error.response,
      data: error.response?.data,
      status: error.response?.status,
      headers: error.response?.headers
    })
    alert('保存文章失败: ' + (error.response?.data?.message || error.message))
  }
}
</script>

<style scoped>
.article-container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: flex-start;
  padding: 80px 2rem 2rem 2rem;
  background: #1a1a1a;
  position: fixed;
  width: 100%;
  top: 0;
  left: 0;
  overflow: auto;
}

.article-box {
  background: #222222;
  padding: 2.5rem;
  border-radius: 16px;
  width: 100%;
  max-width: 1200px;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.2);
}

.form-header {
  margin-bottom: 2rem;
}

.form-header h2 {
  color: #ffffff;
  font-size: 1.8rem;
  font-weight: 600;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-group input,
.form-group textarea {
  width: 100%;
  padding: 1rem;
  border: 1px solid #444;
  border-radius: 8px;
  background: #1a1a1a;
  color: #fff;
  font-size: 1rem;
  transition: all 0.3s ease;
}

.form-group textarea {
  resize: vertical;
  min-height: 500px;
  line-height: 1.6;
  font-size: 1.1rem;
}

.form-group input:focus,
.form-group textarea:focus {
  border-color: #00ff7f;
  outline: none;
}

.submit-btn {
  width: 100%;
  padding: 1rem;
  background: #00ff7f;
  color: #000;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
}

.submit-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 5px 15px rgba(0, 255, 127, 0.3);
}

.form-group input {
  font-size: 1.3rem;
  padding: 1.2rem;
}

/* 适配移动设备 */
@media (max-width: 768px) {
  .article-container {
    padding: 70px 1rem 1rem 1rem;
  }

  .article-box {
    padding: 1.5rem;
  }

  .form-group textarea {
    min-height: 300px;
  }
}
</style>
