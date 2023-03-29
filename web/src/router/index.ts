import { useGlobalStore } from "@/store/main";
import { createRouter, createWebHistory } from "vue-router";

import creation from "@/router/creation";
import liveRouter from "@/router/live";
import personalRouter from "@/router/personal";
import articleShow from "@/router/show/article";
import videoShow from "@/router/show/video";
import search from "@/router/search"
import space from "@/router/space";
const routes = [
    {
        path: '',
        name: 'Home',
        meta: {
            title: '首页',
            requireAuth: false,
            keepAlive: false
        },
        component: () => import('@/views/home/home.vue')
    },
    {
        path: '/column',
        name: 'Column',
        meta: {
            title: '专栏',
            requireAuth: false,
            keepAlive: false
        },
        component: () => import('@/views/home/column.vue')
    },
    {
        path: '/live',
        name: 'Live',
        meta: {
            title: '专栏',
            requireAuth: false,
            keepAlive: false
        },
        component: () => import('@/views/home/live.vue')
    },
    {
        path: "/",
        name: "Index",
        component: () => import('@/views/Layout.vue'),
        children: [
            ...personalRouter,
            ...liveRouter,
            ...videoShow,
            ...space,
        ]
    },
    //登入
    {
        path: "/login",
        name: "Login",
        meta: {
            title: '登入',
            requireAuth: false,
            keepAlive: false
        },
        component: () => import('@/views/login/login.vue'),
    },
    ...search,
    //创作中心
    ...creation,
    //专栏展示
    ...articleShow
    //未匹配路由404
    ,
    {
        path: '/:pathMatch(.*)*',
        name: 'NotFound',
        component: () => import('@/views/404.vue'),

    }

];

const router = createRouter({
    history: createWebHistory(),
    routes,
});


router.beforeEach((to, from, next) => {
    const globalStore = useGlobalStore()
    globalStore.globalData.router.currentRouter = to.path
    // if(to.path !== "/login"){
    //     let  userInfo = useUserStore();
    //     if (userInfo.userInfoData.id == 0 ){
    //         ElMessage({
    //             message: '请先登入',
    //             type: 'error',
    //         })
    //         next({name : "Login"})
    //         return
    //     }
    // }

    next()
})

export default router;
