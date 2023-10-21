import store from '@/store'
const serverIP = store.state.serverIP
const baseURL = `ws://${serverIP}:8082`
export function newWebSocket(intface, params) {
    var URL=`${baseURL}${intface}${params}`
    return new WebSocket(URL);
} 