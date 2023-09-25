import axios from 'axios'
// import store from '@/store'
// const IP = store.state.IP
const service = axios.create({
  // baseURL: `http://${IP}:8082`,
  baseURL: 'http://10.22.60.221:8082',
    // baseURL: 'http://10.22.60.119:8082',
    timeout: 10000,
})

// 请求拦截器
service.interceptors.request.use(
    config => {
      // 在请求头中添加授权信息
      if (localStorage.getItem("token")) {
        config.headers['Authorization'] = `Bearer ${localStorage.getItem("token")}`;
      }
      return config;
    },
    error => {
      // 请求错误处理
      return Promise.reject(error);
    }
);
  
export default service

