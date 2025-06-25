// src/http.js
import axios from 'axios';
import { useRouter } from 'vue-router';

const api = axios.create({
  baseURL: '/api',
  timeout: 5000,
});

// 页面初始化时读取本地 token 并设置到请求头
const savedToken = localStorage.getItem('token');
if (savedToken) {
  api.defaults.headers.common['Authorization'] = `Bearer ${savedToken}`;
}

// 请求拦截器 - 自动添加 token
api.interceptors.request.use(config => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers['Authorization'] = `Bearer ${token}`;
  }
  return config;
});

// 响应拦截器 - 捕获 401 错误并跳转登录页
api.interceptors.response.use(
  response => response,
  error => {
    const router = useRouter();
    if (error.response && error.response.status === 401) {
      // 清除本地 token
      localStorage.removeItem('token');
      // 删除 axios 请求头中的 authorization
      delete api.defaults.headers.common['Authorization'];
      // 跳转到登录页
      router.push('/MyLogin');
    }
    return Promise.reject(error);
  }
);

export default api;