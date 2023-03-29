export default [{
    path: "/search/:text",
    name: "Search",
    meta: {
        title: '搜索',
        requireAuth: false,
        keepAlive: false
    },
    component: () => import('@/views/search/Layout.vue'),
    children: [
        {
            path: 'video',
            name: 'VideoSearch',
            meta: {
                title: '视频搜索',
                requireAuth: false,
                keepAlive: false
            },
            component: () => import('@/views/search/video.vue')
        },
        {
            path: 'user',
            name: 'UserSearch',
            meta: {
                title: '用户搜索',
                requireAuth: false,
                keepAlive: false
            },
            component: () => import('@/views/search/user.vue')
        },
    ]
}];