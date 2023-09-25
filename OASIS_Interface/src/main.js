import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import Web3 from 'web3'
Vue.config.productionTip = false
import 'animate.css';
import './plugins/element.js'
import Identicon from "identicon.js"
import axios from 'axios'
Vue.prototype.axios = axios
Vue.prototype.Web3 = Web3
Vue.prototype.Identicon = Identicon
const eventBus = new Vue();
export default eventBus;
new Vue({
  Web3,
  router,
  store,
  render: h => h(App)
}).$mount('#app')
