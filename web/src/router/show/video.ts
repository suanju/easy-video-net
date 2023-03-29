export default [{
    path: "/videoShow",
    name: "VideoShowPage",
    children: [
        {
            path: "video",
            name: "VideoShow",
            component: () => import('@/views/show/video/video.vue'),
        }
    ]
}]
