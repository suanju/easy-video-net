import { deleteFavorites, getFavoritesList } from "@/apis/personal";
import { useUserStore } from "@/store/main";
import Swal from "sweetalert2";
import { useRoute, useRouter } from "vue-router";
import globalScss from "@/assets/styles/global/export.module.scss";
import { Ref, ref } from "vue";
import { DeleteFavoritesReq, GetFavoritesListItem, GetFavoritesListRes } from "@/types/personal/collect/favorites";


export const useFavoritesProp = () => {
    const userStore = useUserStore()
    const router = useRouter()
    const route = useRoute()
    const createCollectDialogShow = ref(false)
    const favoritesList = ref(<GetFavoritesListRes>{})
    const updataInfo = ref(<GetFavoritesListItem>{})
    return {
        userStore,
        router,
        favoritesList,
        createCollectDialogShow,
        updataInfo,
        route,
    }
}


export const useDelFavorites = (id: number , favoritesList : Ref<GetFavoritesListRes>) => {
    Swal.fire({
        heightAuto: false,
        title: '确认删除该收藏夹嘛?',
        icon: 'warning',
        confirmButtonColor: globalScss.colorButtonTheme,
        showCancelButton: true,
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        reverseButtons: true
    }).then(async (result) => {
        if (result.isConfirmed) {
            try {
                await deleteFavorites(<DeleteFavoritesReq>{
                    id
                })
                Swal.fire({
                    title: "删除成功",
                    confirmButtonColor: globalScss.colorButtonTheme,
                    heightAuto: false,
                    icon: "success",
                })
                favoritesList.value = favoritesList.value.filter((item) =>{
                    return  item.id != id
                })
            } catch (err: any) {
                console.log(err)
                Swal.fire({
                    title: "删除失败",
                    heightAuto: false,
                    confirmButtonColor: globalScss.colorButtonTheme,
                    icon: "error",
                })
            }
        } else if (result.dismiss === Swal.DismissReason.cancel) {
            console.log("取消")
        }
    })
}

export const useUpdateFavorites = (item: GetFavoritesListItem ,createCollectDialogShow :Ref<boolean> ,updataInfo :Ref<GetFavoritesListItem>) => {
    updataInfo.value = item
    createCollectDialogShow.value = true
    console.log("更新", item)
}

export const useInit = async (favoritesList: Ref<GetFavoritesListRes> ,isLoading :Ref<boolean>) => {
    try {
        //获取收藏夹列表 
        const response = await getFavoritesList();
        if (!response.data) return false
        favoritesList.value = response.data
        isLoading.value = true

    } catch (err) {
        console.log(err)
        Swal.fire({
            title: "获取失败",
            heightAuto: false,
            confirmButtonColor: globalScss.colorButtonTheme,
            icon: "error",
        })
    }
}
