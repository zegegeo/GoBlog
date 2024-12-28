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
    const token = localStorage.getItem('token')
    console.log('发送请求:', {
      url: config.url,
      method: config.method,
      headers: config.headers,
      token: token
    })
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`
    }
    return config
  },
  error => {
    console.error('Request error:', error)
    return Promise.reject(error)
  }
)

// 添加响应拦截器
instance.interceptors.response.use(
  response => {
    return response
  },
  error => {
    console.error('响应错误:', {
      status: error.response?.status,
      data: error.response?.data,
      message: error.message
    })
    if (error.response?.status === 401) {
      localStorage.removeItem('token')
      localStorage.removeItem('user')
      router.push('/login')
    }
    return Promise.reject(error)
  }
)

export const articleApi = {
  // 获取文章列表
  getList(page = 1, pageSize = 10) {
    console.log('获取文章列表:', { page, pageSize })
    return instance.get(getApiUrl('/articles'), {
      params: { page, pageSize }
    }).then(response => {
      console.log('Article list response:', response.data)
      return response
    })
  },
  
  // 获取文章详情
  getDetail(id) {
    if (!id) {
      return Promise.reject(new Error('文章ID不能为空'))
    }
    const url = getApiUrl(`/articles/${id}`)
    console.log('发送文章详情请求:', {
      url: url,
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('token')}`
      },
      token: localStorage.getItem('token')
    })
    return instance.get(url)
      .then(response => {
        console.log('文章详情响应:', response.data)
        // 确保作者信息存在
        if (response.data.article && !response.data.article.author) {
          response.data.article.author = {
            username: '未知作者'
          }
        }
        return response
      }).catch(error => {
        console.error('获取文章详情失败:', {
          status: error.response?.status,
          data: error.response?.data,
          url: error.config?.url,
          method: error.config?.method,
          message: error.message
        })
        throw error
      })
  },
  
  // 创建文章
  create(data) {
    const token = localStorage.getItem('token')
    const user = JSON.parse(localStorage.getItem('user') || '{}')
    
    // 检查用户信息
    if (!user.id) {
      console.error('用户信息:', user)
      return Promise.reject(new Error('用户信息缺失，请重新登录'))
    }
    
    // 添加用户ID到请求数据
    const requestData = {
      title: data.title.trim(),
      content: data.content.trim(),
      userId: parseInt(user.id)
    }
    
    console.log('Creating article:', requestData)
    return instance.post(getApiUrl('/articles'), requestData)
      .then(response => {
        console.log('Article created:', response.data)
        return response
      })
      .catch(error => {
        console.error('Create article failed:', {
          error,
          requestData,
          user,
          response: error.response?.data
        })
        throw error
      })
  },
  
  // 更新文章
  update(id, data) {
    if (!id) {
      return Promise.reject(new Error('文章ID不能为空'))
    }
    
    const user = JSON.parse(localStorage.getItem('user') || '{}')
    if (!user.id) {
      return Promise.reject(new Error('用户未登录'))
    }
    
    const requestData = {
      title: data.title.trim(),
      content: data.content.trim(),
      user_id: parseInt(user.id)
    }
    
    console.log('Updating article:', {
      id,
      data: requestData
    })
    
    return instance.put(getApiUrl(`/articles/${id}`), requestData)
      .then(response => {
        console.log('Article updated:', response)
        return response
      })
      .catch(error => {
        console.error('Update article failed:', {
          error,
          id,
          data: requestData,
          response: error.response?.data
        })
        throw error
      })
  },
  
  // 删除文章
  delete(id) {
    return instance.delete(getApiUrl(`/articles/${id}`))
  }
} 