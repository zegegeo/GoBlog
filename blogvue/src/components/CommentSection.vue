<template>
  <div class="comments-section">
    <h3>评论区</h3>
    <!-- 发表评论 -->
    <div class="comment-form">
      <textarea 
        v-model="newComment" 
        placeholder="写下你的评论..."
      ></textarea>
      <button @click="submitComment">发表评论</button>
    </div>
    <!-- 评论列表 -->
    <div class="comment-list">
      <div v-for="comment in comments" :key="comment.id" class="comment-item">
        <div class="comment-header">
          <span class="author">{{ comment.author.username }}</span>
          <span class="time">{{ comment.created_at }}</span>
        </div>
        <div class="comment-content">{{ comment.content }}</div>
        <div v-if="isCommentOwner(comment)" class="comment-actions">
          <button @click="deleteComment(comment.id)" class="delete-btn">删除</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { commentApi } from '../api/comment'

const props = defineProps({
  articleId: {
    type: Number,
    required: true
  }
})

const comments = ref([])
const newComment = ref('')
const currentUser = JSON.parse(localStorage.getItem('user') || '{}')

const loadComments = async () => {
  try {
    const response = await commentApi.getList(props.articleId)
    comments.value = response.data.comments
  } catch (error) {
    console.error('加载评论失败:', error)
  }
}

const submitComment = async () => {
  if (!newComment.value.trim()) return
  try {
    await commentApi.create(props.articleId, newComment.value)
    newComment.value = ''
    await loadComments()
  } catch (error) {
    console.error('发表评论失败:', error)
    alert('发表评论失败')
  }
}

const deleteComment = async (commentId) => {
  if (!confirm('确定要删除这条评论吗？')) return
  try {
    await commentApi.delete(commentId)
    await loadComments()
  } catch (error) {
    console.error('删除评论失败:', error)
    alert('删除评论失败')
  }
}

const isCommentOwner = (comment) => {
  return comment.user_id === currentUser.id
}

onMounted(loadComments)
</script>

<style scoped>
.comments-section {
  margin-top: 2rem;
  padding-top: 2rem;
  border-top: 1px solid #333;
}

.comment-form {
  margin-bottom: 2rem;
}

.comment-form textarea {
  width: 100%;
  min-height: 100px;
  padding: 1rem;
  margin-bottom: 1rem;
  background: #1a1a1a;
  border: 1px solid #333;
  border-radius: 8px;
  color: #fff;
  resize: vertical;
}

.comment-item {
  background: #1a1a1a;
  padding: 1rem;
  border-radius: 8px;
  margin-bottom: 1rem;
}

.comment-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 0.5rem;
}

.author {
  color: #00a8ff;
}

.time {
  color: #666;
}

.comment-content {
  color: #fff;
  line-height: 1.6;
}

.comment-actions {
  margin-top: 0.5rem;
  text-align: right;
}

.delete-btn {
  background: #ff4757;
  color: #fff;
  border: none;
  padding: 0.3rem 0.8rem;
  border-radius: 4px;
  cursor: pointer;
}

button {
  background: #00ff7f;
  color: #000;
  border: none;
  padding: 0.8rem 1.5rem;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.3s ease;
}

button:hover {
  opacity: 0.9;
  transform: translateY(-1px);
}
</style> 