import axios from 'axios'
import { getApiUrl } from '../config/api'

// 创建axios实例
const instance = axios.create({
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 请求拦截器
instance.interceptors.request.use(
  config => {
    console.log('发送请求:', {
      url: config.url,
      method: config.method,
      data: config.data,
      headers: config.headers
    })
    return config
  },
  error => {
    console.error('请求错误:', error)
    return Promise.reject(error)
  }
)

// 响应拦截器
instance.interceptors.response.use(
  response => {
    console.log('收到响应:', response)
    return response
  },
  error => {
    console.error('响应错误:', {
      status: error.response?.status,
      data: error.response?.data,
      message: error.message,
      config: error.config
    })
    return Promise.reject(error)
  }
)

export const userApi = {
  // 用户登录
  login(username, password) {
    console.log('Login request data:', { 
      username: username.trim(), 
      password: password.trim() 
    })
    return instance.post(getApiUrl('/user/login'), { 
      username: username.trim(), 
      password: password.trim() 
    })
    .then(response => {
      const { token, user } = response.data
      return {
        data: {
          token,
          user: {
            id: user.id,
            username: user.username,
            createdAt: user.createdAt
          }
        }
      }
    })
  },

  // 用户注册
  register(data) {
    if (!data.username || !data.password) {
      return Promise.reject(new Error('用户名和密码不能为空'))
    }
    console.log('Register request data:', data)
    return instance.post(getApiUrl('/user/register'), {
      username: data.username.trim(),
      password: data.password.trim()
    })
  }
} 