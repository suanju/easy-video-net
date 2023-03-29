import { getReleaseInformation } from '@/apis/space';
import { GetReleaseInformationReq, GetReleaseInformationRes } from '@/types/space/space';
import Swal from 'sweetalert2';
import globalScss from "@/assets/styles/global/export.module.scss"
import { Ref, ref } from 'vue';
import { useSpaceStore } from '@/store/main';

export const useSpaceProp = () => {
    const spaceData = useSpaceStore()
    const releaseInformation = ref(<GetReleaseInformationRes>{})
    const isLoading = ref(false)
    return {
        spaceData,
        releaseInformation,
        isLoading
    }
}


export const useInit = async (id : number ,releaseInformation: Ref<GetReleaseInformationRes>, isLoading :Ref<boolean> ) => {
    try {
        const data  = await getReleaseInformation(<GetReleaseInformationReq>{
            id
        })
        if (!data.data) return false
        releaseInformation.value = data.data
        isLoading.value = true

    } catch (err) {
        console.log(err)
        Swal.fire({
            title: "获取视频专栏信息失败",
            heightAuto: false,
            confirmButtonColor: globalScss.colorButtonTheme,
            icon: "error",
        })
    }
}