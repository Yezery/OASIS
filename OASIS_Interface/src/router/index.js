import Vue from 'vue'
import VueRouter from 'vue-router'
import store from '@/store'
Vue.use(VueRouter)
// 解决重复点击路由时出错误
const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push (location) {
    return originalPush.call(this, location).catch(err => err)
}
const routes = [
  {
    path: '/',
    redirect: '/home/MarketShop'
  },
  {
    path: '/home',
    name: 'home',
    component: () => import('@/views/layout.vue'),
    meta: { isAuth: true },
    children: [
      {
        path: 'MarketShop',
        name: 'MarketShop',
        component: () => import('@/views/main/marketShop/nftMarket.vue'),
      },
      {
        path: 'MarketShop3D',
        name: 'MarketShop3D',
        component: () => import('@/views/main/marketShop/nftMarket3D.vue'),
      },
      {
        path: 'ImitNFT',
        name: 'ImitNFT',
        component: () => import('@/views/main/imit/imitNFT.vue'),
        meta: { isAuth: true }
      },
      {
        path: 'Imit3DNFT',
        name: 'Imit3DNFT',
        component: () => import('@/views/main/imit/imit3DNFT.vue'),
        meta: { isAuth: true }
      },
      {
        path: 'addImit',
        name: 'addImit',
        component: () => import('@/views/main/imit/addImitNFT'),
        meta: { isAuth: true }
      },
      {
        path: 'addImit3D',
        name: 'addImit3D',
        component: () => import('@/views/main/imit/addImit3DNFT'),
        meta: { isAuth: true }
      },
      
      {
        path: 'NFTInf',
        name: 'NFTInf',
        component: () => import('@/components/infPage/NFTInf'),
        meta: { isAuth: true }
      },
      {
        path: 'NFTInf3D',
        name: 'NFTInf3D',
        component: () => import('@/components/infPage/3DNFTInf'),
        meta: { isAuth: true }
      },

    ]
  },
  {
    path: '/userhome',
    name: 'userhome',
    component: () => import('@/views/user/userHome.vue'),
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
        next(); 
      }
      router.push("/"); 
    }
  } else {
    next();
  }
})
export default router
