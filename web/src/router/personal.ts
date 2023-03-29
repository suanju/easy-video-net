export default [{
  path: 'personal',
  name: 'Personal',
  meta: {
    title: '用户信息',
    requireAuth: false,
    keepAlive: false
  },
  component: () => import('@/views/personal/Layout.vue'),
  children: [
    {
      path: '',
      name: 'UserShow',
      meta: {
        title: '首页',
        requireAuth: false,
        keepAlive: false
      },
      component: () => import('@/views/personal/userInfo/userShow.vue')
    },
    {
      path: 'userinfo',
      name: 'UserInfo',
      meta: {
        title: '用户信息',
        requireAuth: false,
        keepAlive: false
      },
      component: () => import('@/views/personal/userInfo/userInfo.vue')
    },{
      path: 'picturesetting',
      name: 'PictureSetting',
      meta: {
        title: '用户信息',
        requireAuth: false,
        keepAlive: false
      },
      component: () => import('@/views/personal/userInfo/pictureSetting.vue')
    } ,{
      path: 'safety',
      name: 'Safety',
      meta: {
        title: '安全',
        requireAuth: false,
        keepAlive: false
      },
      component: () => import('@/views/personal/userInfo/safety.vue')
    },{
      path: 'livesetup',
      name: 'LiveSetUp',
      meta: {
        title: '直播设置',
        requireAuth: false,
        keepAlive: false
      },
      component: () => import('@/views/personal/live/setUp.vue')
    },
    {
      path: 'favorites',
      name: 'Favorites',
      meta: {
        title: '我的收藏',
        requireAuth: false,
        keepAlive: false
      },
      component: () => import('@/views/personal/collect/favorites.vue')
    },
    {
      path: 'collectList/:id',
      name: 'CollectList',
      meta: {
        title: '收藏夹',
        requireAuth: false,
        keepAlive: false
      },
      component: () => import('@/views/personal/collect/collectList.vue')
    },{
      path: 'record',
      name: 'Record',
      meta: {
        title: '历史记录',
        requireAuth: false,
        keepAlive: false
      },
      component: () => import('@/views/personal/record/record.vue')
    },
  ]
}] 