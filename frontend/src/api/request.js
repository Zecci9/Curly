import axios from 'axios'

const request = axios.create({
  // 开发环境建议保持 /api/v1，让 Vite proxy 转发到 Go。
  // 生产环境也推荐由 Nginx/Caddy 将同域 /api 反代到 Go。
  baseURL: import.meta.env.VITE_API_BASE_URL || '/api/v1',
  timeout: 8000,
})

request.interceptors.request.use((config) => {
  const token = localStorage.getItem('curly.token')
  if (token) config.headers.Authorization = `Bearer ${token}`
  return config
})

request.interceptors.response.use(
  (res) => res.data,
  (err) => {
    const message =
      err.response?.data?.message ||
      (err.code === 'ERR_NETWORK'
        ? '无法连接 Curly 后端，请确认 Go HTTPS 服务已启动。'
        : err.message) ||
      '请求失败'

    return Promise.reject(new Error(message))
  },
)

export default request
