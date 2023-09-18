import Vue from 'vue'
import VueRouter from 'vue-router'
import store from '@/store'
Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: '/home/MarketShop'
  },
  {
    path: '/home',
    name: 'home',
    component: () => import('@/views/market/index.vue'),
    meta: { isAuth: true },
    // 
    children: [
      {
        path: 'ImitNFT',
        name: 'ImitNFT',
        component: () => import('@/views/market/imitNFT.vue'),
        meta: { isAuth: true }
      },
      {
        path: 'addImit',
        name: 'addImit',
        component: () => import('@/views/market/addMint.vue'),
        meta: { isAuth: true }
      },
      {
        path: 'MarketShop',
        name: 'MarketShop',
        component: () => import('@/views/market/NFTMarket.vue'),
      },
      {
        path: 'NFTInf',
        name: 'NFTInf',
        component: () => import('@/views/market/NFTInf.vue'),
        meta: { isAuth: true }
      },

    ]
  },
  {
    path: '/userhome',
    name: 'userhome',
    component: () => import('@/views/user/UserHome.vue'),
    meta: { isAuth: true }
  }
]

const router = new VueRouter({
  routes
})

//全局前置路由守卫————初始化的时候被调用、每次路由切换之前被调用
router.beforeEach((to, from, next) => {
  if (to.meta.isAuth) {
    const marketNFTInf = store.state.marketNFTInf;
    if (store.state.isconnect) {
      next();
    } else {
      if (!marketNFTInf) {
        router.push("/"); // 只在需要重定向时导航
      } else {
        next(); // 允许访问路由或已经在目标路由上，无需重定向
      }
      next();
    }
  } else {
    next();
  }
})
export default router
