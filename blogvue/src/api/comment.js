import axios from 'axios'
import { getApiUrl } from '../config/api'

// 创建axios实例
const instance = axios.create({
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 添加请求拦截器
instance.interceptors.request.use(
  config => {
    // 从localStorage获取token
    const token = localStorage.getItem('token')
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`
    }
    return config
  },
  error => {
    return Promise.reject(error)
  }
)

export const commentApi = {
  // 获取评论列表
  getList(articleId, page = 1, pageSize = 10) {
    return instance.get(getApiUrl(`/articles/${articleId}/comments`), {
      params: { page, pageSize }
    })
  },
  
  // 创建评论
  create(articleId, content) {
    return instance.post(getApiUrl(`/articles/${articleId}/comments`), {
      content
    })
  },
  
  // 删除评论
  delete(commentId) {
    return instance.delete(getApiUrl(`/comments/${commentId}`))
  }
} 