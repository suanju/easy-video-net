export default [{
    path: "/articleShow",
    name: "ArticleShowPage",

    children: [
        {
            path: "article",
            name: "ArticleShow",
            component: () => import('@/views/show/article/article.vue'),
        }
    ]
}]
