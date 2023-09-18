import axios from 'axios'
// import store from '@/store'
// const IP = store.state.IP
const service = axios.create({
    // baseURL: `http://${IP}:8082`,
    baseURL: 'http://10.22.60.54:8082',
    timeout: 10000 
})
export default service

