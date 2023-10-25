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
        path: 'mint3DNFT',
        name: 'mint3DNFT',
        component: () => import('@/views/main/mint/mint3DNFT.vue'),
        meta: { isAuth: true }
      },
      {
        path: 'mintNFT',
        name: 'mintNFT',
        component: () => import('@/views/main/mint/mintNFT.vue'),
        meta: { isAuth: true }
      },
      {
        path: 'addmint',
        name: 'addmint',
        component: () => import('@/views/main/mint/addMintNFT'),
        meta: { isAuth: true }
      },
      {
        path: 'addmint3D',
        name: 'addmint3D',
        component: () => import('@/views/main/mint/addMint3DNFT'),
        meta: { isAuth: true }
      },
      
      {
        path: 'NFTInf',
        name: 'NFTInf',
        component: () => import('@/components/infPage/NFTInf'),
      },
      {
        path: 'NFTInf3D',
        name: 'NFTInf3D',
        component: () => import('@/components/infPage/3DNFTInf'),
      },
      {
        path: 'searchPage',
        name: 'searchPage',
        component: () => import('@/views/main/marketShop/searchPage'),
      },

    ]
  },
  {
    path: '/userhome',
    name: 'userhome',
    component: () => import('@/views/user/userHome.vue'),
    meta: { isAuth: true }
  },
  {
    path: '/mintHome',
    name: 'mintHome',
    component: () => import('@/views/main/mint/mintHome.vue'),
  },
  {
    path: '/addMintHome',
    name: 'addMintHome',
    component: () => import('@/views/main/mint/addMintHome.vue'),
  },
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
