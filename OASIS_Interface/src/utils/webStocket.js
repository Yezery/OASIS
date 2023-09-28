// import store from '@/store'
// const IP = store.state.IP
const baseURL = "ws://localhost:8082"
// const baseURL = `ws://${IP}:8082`
export function newWebSocket(intface, params) {
    var URL=`${baseURL}${intface}${params}`
    return new WebSocket(URL);
} 