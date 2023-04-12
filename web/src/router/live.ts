export default [{
    path: "/live/room/:id",
    name: "liveRoom",
    component: () => import('@/views/live/room.vue'),
}]