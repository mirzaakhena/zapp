import axios from 'axios'
import Vue from 'vue'

const service = axios.create({
  baseURL: process.env.NODE_ENV === 'development' ? 'http://localhost:8081' : '',
  timeout: 5000
})

service.interceptors.request.use((config) => {
  config.headers['Authorization'] = 'Bearer ' + Vue.auth.getToken()
  return config
}, (error) => {
  return Promise.reject(error)
})

export default service