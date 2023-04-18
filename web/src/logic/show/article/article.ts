
import { getArticleContributionByID } from "@/apis/contribution";
import globalScss from "@/assets/styles/global/export.module.scss";
import { useGlobalStore } from "@/store/main";
import { GetArticleContributionByIDReq, GetArticleContributionByIDRes } from "@/types/show/article/article";
import Swal from 'sweetalert2';
import { Ref, reactive, ref } from "vue";
import { RouteLocationNormalizedLoaded, Router, useRoute, useRouter } from 'vue-router';

export const useArticleShowProp = () => {
    const articleID = ref(0)
    const global = useGlobalStore()
    const articleInfo = ref(<GetArticleContributionByIDRes>{})
    const router = useRouter()
    const route = useRoute()

    //回复二级评论
    const replyCommentsDialog = reactive({
        show: false,
        commentsID: 0,
    })

    return {
        articleID,
        articleInfo,
        router,
        route,
        replyCommentsDialog,
        global
    }
}
export const useInit = async (articleID: Ref<number>, articleInfo: Ref<GetArticleContributionByIDRes>, route: RouteLocationNormalizedLoaded, router: Router, global: any) => {
    try {
        if (!route.params.id) {
            router.back()
            Swal.fire({
                title: "缺少文章ID",
                heightAuto: false,
                confirmButtonColor: globalScss.colorButtonTheme,
                icon: "error",
            })
            router.back()
            return
        }
        articleID.value = Number(route.params.id)
        global.globalData.loading.loading = true
        window.onresize = function () {
            const canvasSnow = document.getElementById('canvas_sakura') as HTMLEmbedElement;
            if (!canvasSnow) return false
            canvasSnow.width = String(window.innerWidth);
            canvasSnow.height = String(window.innerHeight);
        }
        //获取文章内容
        const response = await getArticleContributionByID(<GetArticleContributionByIDReq>{
            articleID: articleID.value
        })

        if (!response.data) throw "文章内容为空"
        articleInfo.value = response.data
        articleInfo.value.comments = response.data?.comments

        global.globalData.loading.loading = false

    } catch (err) {
        global.globalData.loading.loading = false
        console.log(err)
    }
}
