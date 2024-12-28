// 定义可用的API基础URL
export const API_BASE_URLS = [
  'http://localhost:8888',
  'http://192.168.31.23:8888'
];

// 获取可用的API地址
export const getApiUrl = (path) => {
  console.log('API Request URL:', path)
  const baseUrl = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8888'
  const fullUrl = `${baseUrl}/api/v1${path}`
  console.log('API Request URL:', fullUrl)
  return fullUrl
}; 