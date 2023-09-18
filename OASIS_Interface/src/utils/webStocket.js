// import store from '@/store'
// const IP = store.state.IP
// const baseURL = "ws://10.22.61.214:8082"
const baseURL = "ws://10.22.60.54:8082"
// const baseURL = `ws://${IP}:8082`
export function newWebSocket(intface, params) {
    var URL=`${baseURL}${intface}${params}`
    return new WebSocket(URL);
} 