import { Router, useRouter } from 'vue-router';
import Swal from 'sweetalert2';
import globalScss from "@/assets/styles/global/export.module.scss"
import { Ref, ref } from 'vue';
import { deleteRecordByID, getRecordList } from '@/apis/personal';
import { useRoute } from 'vue-router';
import { DeleteRecordByIDReq, GetRecordListItem, GetRecordListReq, GetRecordListRes } from '@/types/personal/record/record'; import { useGlobalStore } from '@/store/main';
import { PageInfo } from '@/types/idnex';

export const useRecordProp = () => {
    const loading = useGlobalStore().globalData.loading
    const route = useRoute()
    const router = useRouter()
    const recordList = ref(<GetRecordListRes>[])
    const pageInfo = ref(<PageInfo>{
            page: 1,
            size: 8,
    })
    //是否首次加载
    const isLoading = ref(true)
    //是否正在加载更多
    const isLoadMore = ref(false)
    //是否全部加载完成
    const isTheEnd = ref(false)
    return {
        recordList,
        isLoading,
        route,
        router,
        pageInfo,
        loading,
        isLoadMore,
        isTheEnd
    }
}

export const useJump = (item: GetRecordListItem, router: Router) => {
    if (item.type == "视频") {
        router.push({ name: "VideoShow", query: { videoID: item.to_id } })
    } else if (item.type == "专栏") {
        router.push({ name: "ArticleShow", query: { articleID: item.to_id } })
    } else {
        router.push({ name: "liveRoom", query: { roomID: item.to_id } })
    }
}

export const useDelRecord = async (recordList: Ref<GetRecordListRes>, id: number, loading: any) => {
    try {
        loading.loading = true
        await deleteRecordByID(<DeleteRecordByIDReq>{
            id
        })
        loading.loading = false
        recordList.value = recordList.value.filter((item) => {
            if (item.id == id) item.is_delete = true
            return item
        })
        setTimeout(() => {
            recordList.value = recordList.value.filter((item) => {
                return item.id != id
            })
        }, 400)

    } catch (err) {
        loading.loading = false
        Swal.fire({
            title: "删除失败",
            heightAuto: false,
            confirmButtonColor: globalScss.colorButtonTheme,
            icon: "error",
        })
    }
}


export const useLoadData = async (recordList: Ref<GetRecordListRes>, isLoading: Ref<boolean>, page: Ref<PageInfo>, isTheEnd: Ref<boolean>,) => {
    try {
        const data = await getRecordList(<GetRecordListReq>{
            page_info: page.value
        }
        )
        page.value.page++
        if (!data.data) return false
        if (data.data.length == 0) isTheEnd.value = true
        data.data = data.data.filter((item) => {
            item.is_delete = false
            return item
        })
        recordList.value = [...recordList.value,...data.data] 
        console.log(recordList)
        isLoading.value = false
    } catch (err) {
        console.log(err)
        Swal.fire({
            title: "获取历史记录失败",
            heightAuto: false,
            confirmButtonColor: globalScss.colorButtonTheme,
            icon: "error",
        })
    }
}