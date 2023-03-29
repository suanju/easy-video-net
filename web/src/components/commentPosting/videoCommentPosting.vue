<template>
    <div class="comments-info">
        <textarea class="comments-textarea" placeholder="写下点什么..." maxlength="1000"
            v-model="comments.comments"></textarea>
    </div>
    <!-- 标签及发布 -->
    <div class="comments-fun">
        <div>
            <SvgIcon name="expression"
                :class="{ 'icon-small': true, 'animate__animated': true, 'animate__tada': emoji.animation }"
                @click="clickEmoji" @mouseleave="emoji.animation = false" @mouseover="emoji.animation = true">
            </SvgIcon>
        </div>

        <div> <el-button type="primary" v-removeFocus round @click="postComment(comments, videoID)">发布</el-button>
        </div>
    </div>
    <div :class="{ 'comments-emoji': true, 'animate__animated': true, 'animate__headShake': !emoji.animationBox, 'animate__flipOutX': emoji.animationBox }"
        v-show="emoji.show">
        <span class="emoji-item" v-for="emojiItem in emoji.emoji" :key="emojiItem"
            @click="comments.comments = comments.comments + emojiItem">{{ emojiItem }}</span>
    </div>
</template>
<script lang="ts" setup>import { getVideoComment, videoPostComment } from '@/apis/contribution';
import {  CommentsInfo } from '@/types/show/article/article';
import { ElButton } from 'element-plus';
import Swal from 'sweetalert2';
import { vRemoveFocus } from "@/utils/customInstruction/focus"
import globalScss from "@/assets/styles/global/export.module.scss"
import {  UnwrapNestedRefs, reactive } from 'vue';
import { GetVideoCommentReq, VideoPostCommentReq } from '@/types/show/video/video';

const props = defineProps({
    videoID: {
        type: Number,
        required: true,
        default: 1
    },
    commentsID: {
        type: Number,
        required: true,
    }
})
const emit = defineEmits(['updateVideoInfo'])



const emoji = reactive({
    show: false,
    animation: false,
    animationBox: true,
    emoji: [
        "😀", "😄", "😁", "😆", "😅", "🤣", "😂", "🙂", "🙃", "😉", "😊", "😇", "🥰", "😍", "🤩", "😚", "🤗", "🤨", "😐", "😑", "😶", "🤐", "😏", "😒", "😮‍💨", "🤥", "😌", "😪", "🤤", "😷", "🤒", "🤕", "🥵", "😵",
        "😕", "🙁", "☹️", "😳", "😞", "😭", "🥱", "😩", "😰", "😲", "😯", "😠", "😩", "😧", "😯", "🥺"
    ]
})

const comments = reactive(<CommentsInfo>{
    comments: "",
})


//点击表情触发
const clickEmoji = () => {
    emoji.animationBox = !emoji.animationBox
    if (emoji.show) {
        //收回面板的动画效果
        setTimeout(() => {
            emoji.show = !emoji.show
        }, 800);
    } else {
        emoji.show = !emoji.show
    }
}

//回复评论
const postComment = async (comments: UnwrapNestedRefs<CommentsInfo>, videoID: number) => {
    try {
        const requistData = <VideoPostCommentReq>{
            video_id: videoID,
            content: comments.comments,
            content_id: props.commentsID,
        }
        const reponse = await videoPostComment(requistData)
        console.log(reponse)

        //清空输入款
        comments.comments = ""
        const commentsList = await getVideoComment(<GetVideoCommentReq>{ video_id: videoID })

        if (!commentsList.data) {
            throw ("获取评论信息失败")
        }

        emit('updateVideoInfo', commentsList.data)

        const Toast = Swal.mixin({
            toast: true,
            position: 'top',
            showConfirmButton: false,
            timer: 3000,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({
            icon: 'success',
            title: '评论成功'
        })

    } catch (err) {
        Swal.fire({
            title: err as string,
            heightAuto: false,
            confirmButtonColor: globalScss.colorButtonTheme,
            icon: "error",
        })
    }
}

</script>

<style scoped lang="scss">
@import "./static/style/videoCommentPosting.scss"
</style>
