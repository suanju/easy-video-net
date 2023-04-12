<template>
    <div class="comments-info">
        <textarea class="comments-textarea" placeholder="写下点什么..." maxlength="1000" v-model="comments.comments"></textarea>
    </div>
    <!-- 标签及发布 -->
    <div class="comments-fun">
        <div>
            <SvgIcon name="expression"
                :class="{ 'icon-small': true, 'animate__animated': true, 'animate__tada': emoji.animation }"
                @click="clickEmoji" @mouseleave="emoji.animation = false" @mouseover="emoji.animation = true">
            </SvgIcon>
        </div>

        <div> <el-button type="primary" v-removeFocus round @click="postComment(comments, articleID)">发布</el-button>
        </div>
    </div>
    <div :class="{ 'comments-emoji': true, 'animate__animated': true, 'animate__headShake': !emoji.animationBox, 'animate__flipOutX': emoji.animationBox }"
        v-show="emoji.show">
        <span class="emoji-item" v-for="emojiItem in emoji.emoji" :key="emojiItem"
            @click="comments.comments = comments.comments + emojiItem">{{ emojiItem }}</span>
    </div>
</template>
<script lang="ts" setup>
import { articlePostComment, getArticleComment } from '@/apis/contribution';
import globalScss from "@/assets/styles/global/export.module.scss";
import { ArticlePostCommentReq, CommentsInfo, GetArticleCommentReq } from '@/types/show/article/article';
import { vRemoveFocus } from "@/utils/customInstruction/focus";
import { ElButton } from 'element-plus';
import Swal from 'sweetalert2';
import { UnwrapNestedRefs, reactive } from 'vue';

const props = defineProps({
    articleID: {
        type: Number,
        required: true,
    },
    articleInfo: {
        type: Object,
        required: true,
    },
    commentsID: {
        type: Number,
        required: true,
    }
})
const emit = defineEmits(['updateArticleInfo'])



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
const postComment = async (comments: UnwrapNestedRefs<CommentsInfo>, articleID: number) => {
    try {
        const requistData = <ArticlePostCommentReq>{
            article_id: articleID,
            content: comments.comments,
            content_id: props.commentsID,
        }
        const reponse = await articlePostComment(requistData)
        console.log(reponse)

        //清空输入款
        comments.comments = ""
        const commentsList = await getArticleComment(<GetArticleCommentReq>{ articleID: articleID })
        if (!commentsList.data) {
            throw ("获取评论信息失败")
        }
        emit('updateArticleInfo', commentsList.data)
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
            title: "评论失败",
            heightAuto: false,
            confirmButtonColor: globalScss.colorButtonTheme,
            icon: "error",
        })
    }
}

</script>

<style scoped lang="scss">
@import "./static/style/commentPosting.scss"
</style>
