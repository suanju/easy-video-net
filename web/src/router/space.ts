export default [{
    path: 'space/:id',
    name: 'Space',
    meta: {
        title: '用户信息',
        requireAuth: false,
        keepAlive: false
    },
    component: () => import('@/views/space/Layout.vue'),
    children: [
        {
            path: 'individual',
            name: 'SpaceIndividual',
            meta: {
                title: '个人空间',
                requireAuth: false,
                keepAlive: false
            },
            component: () => import('@/views/space/space.vue')
        },
        {
            path: 'myAttention',
            name: 'MyAttention',
            meta: {
                title: '我的关注',
                requireAuth: false,
                keepAlive: false
            },
            component: () => import('@/views/space/myAttention.vue')
        },
        {
            path: 'myVermicelli',
            name: 'MyVermicelli',
            meta: {
                title: '我的粉丝',
                requireAuth: false,
                keepAlive: false
            },
            component: () => import('@/views/space/myVermicelli.vue')
        }
    ]
}]