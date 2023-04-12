export default [{
    path: "/articleShow",
    name: "ArticleShowPage",

    children: [
        {
            path: "article/:id",
            name: "ArticleShow",
            component: () => import('@/views/show/article/article.vue'),
        }
    ]
}]
