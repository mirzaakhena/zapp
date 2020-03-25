import axios from 'axios'
import Vue from 'vue'
import Swal from 'sweetalert2'

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

service.interceptors.response.use((response) => {
  
  return response
}, (error) => {

  if (error.response.status === 403) {
    if (!Vue.auth.isAuthenticated()) {      
      Swal.fire("Owh!", "Your session expired. Please try to relogin ", "error")
      return Promise.resolve()  
    }    
    Swal.fire("Oops!", "You have no access to this function", "error")
    return Promise.reject(error.response)
  }

  if (error.response.status === 400) {
    Swal.fire("Oops!", error.response.data.message, "error")
    return Promise.reject(error.response)
  }

  return Promise.reject(error)
})

export default service