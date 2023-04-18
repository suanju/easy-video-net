<template>
    <div class="video-page">
        <!-- 主题内容 -->
        <div class="video-content">
            <!-- 左边部分 -->
            <div class="content-left">
                <div class="vidoe-info">
                    <div class="vidoe-title">
                        {{ videoInfo?.videoInfo?.title }}
                    </div>
                    <div class="video-data">
                        <div class="video-data-list">
                            <span class="data-item">
                                <SvgIcon name="video" class="icon" color="#9499A0"></SvgIcon> {{
                                    videoInfo?.videoInfo?.heat
                                }}
                            </span>
                            <span class="data-item">
                                <SvgIcon name="barrageNumber" class="icon" color="#9499A0"></SvgIcon> {{
                                    videoInfo?.videoInfo?.barrageNumber
                                }}
                            </span>
                            <span class="data-item">
                                <SvgIcon name="time" class="icon" color="#9499A0"></SvgIcon> {{
                                    rFC3339ToTime(videoInfo?.videoInfo?.created_at)
                                }}
                            </span>
                            <span class="data-item">
                                <SvgIcon name="forbid" class="forbid-icon" color="#9499A0"></SvgIcon>未经作者授权，禁止转载
                            </span>

                            <div class="info-right">
                                <el-button :type="videoInfo?.videoInfo?.is_like ? 'primary' : 'info'" v-removeFocus
                                    class="giveLike" @click="useLikeVideo(videoInfo)">
                                    <SvgIcon name="giveLike" v-removeFocus class="icon" color="#fff" circle></SvgIcon>
                                </el-button>
                                <el-button :type="videoInfo?.videoInfo?.is_collect ? 'primary' : 'info'"
                                    @click="favoriteVideoShow = true" v-removeFocus :icon="Star" circle />
                                <el-button type="info" @click="notOpen" v-removeFocus :icon="Position" circle />
                            </div>
                        </div>
                    </div>
                </div>
                <div class="video-box">
                    <div ref="videoRef" :class="{ 'player': true, 'dplayer-comment-show': !userStore.userInfoData.token }"
                        id="dplay" />
                    <div class="video-sending">
                        <div class="live-info">
                            <span>{{ liveNumber }} 人正在看</span> , <span> 已装填 {{ videoInfo?.videoInfo?.barrageNumber }}
                                条弹幕 </span>
                        </div>
                        <div class="barrage-set" @click="videoBarrage = !videoBarrage">
                            <SvgIcon :name="videoBarrage ? 'barrageOn' : 'barrageOff'" class="barrage-icon" color="#61666D">
                            </SvgIcon>
                            <div class="barrage-input" v-show="userStore.userInfoData.token">
                                <el-input v-model="barrageInput" placeholder="发个友善的弹幕见证当下">
                                    <template #append>
                                        <div><el-button type="primary" @click="sendBarrageEvent()">发送</el-button></div>
                                    </template>
                                </el-input>
                            </div>

                            <div class="barrage-input" v-show="!userStore.userInfoData.token">
                                <el-input v-model="barrageInput" placeholder="请先登入或者注册再来嗷" readonly>
                                    <template #append>
                                        <div><el-button @click="router.push({ name: 'Login' })"
                                                type="primary">登入</el-button></div>
                                    </template>
                                </el-input>
                            </div>

                        </div>
                    </div>
                </div>
                <!-- 视频介绍 -->
                <videoIntroduce :introduce="videoInfo?.videoInfo?.introduce" :label="videoInfo?.videoInfo?.label">
                </videoIntroduce>

                <!-- 视频评论 -->
                <div class="comments-box">
                    <div class="comments-main">
                        <commentPosting :videoID="videoInfo?.videoInfo?.id" :commentsID="0"
                            @updateVideoInfo="updateVideoInfo"></commentPosting>
                    </div>

                    <div class="comments-show">
                        <div class="comments-show-titel"><span>Comments | </span> <span>{{
                            videoInfo?.videoInfo?.comments.length
                        }} 条评论</span>
                        </div>
                        <div class="comments-show-info">
                            <div class="comment-info-detail-box" v-for="commentsItem in videoInfo?.videoInfo?.comments"
                                :key="commentsItem.id">
                                <div class="comment-info-detail">
                                    <el-avatar :size="36" :src="commentsItem.photo" />
                                    <div class="comment-info-content">
                                        <div class="content-head">
                                            <div> <span class="comment-info-username">{{ commentsItem.username }}</span>
                                                <span class="comment-info-other">{{
                                                    dayjs(commentsItem.created_at).format('YYYY.MM.DD.hh.mm')
                                                }}</span>
                                            </div>
                                            <div class="commentInfo-reply">
                                                <el-button type="primary" v-removeFocus round size="small"
                                                    @click="replyComments(commentsItem.id)">回复</el-button>
                                            </div>
                                        </div>
                                        <!-- 评论内容部分 -->
                                        <div class="content-info">
                                            {{ commentsItem.context }}
                                        </div>
                                        <!-- 子评论 -->
                                        <div class="comment-info-detail-child"
                                            v-for="lowerComments in commentsItem.lowerComments" :key="lowerComments.id">
                                            <el-avatar :size="36" :src="lowerComments.photo" />
                                            <div class="comment-info-content">
                                                <div class="content-head">
                                                    <div> <span class="comment-info-username">{{
                                                        lowerComments.username
                                                    }}</span> <span class="comment-info-other">{{
    dayjs(lowerComments.created_at).format('YYYY.MM.DD.hh.mm')
}}</span>
                                                    </div>
                                                    <div class="commentInfo-reply">
                                                        <el-tag effect="dark" round v-removeFocus
                                                            @click="replyComments(lowerComments.id)">
                                                            回复
                                                        </el-tag>
                                                    </div>
                                                </div>
                                                <!-- 评论内容部分 -->
                                                <div class="content-info">
                                                    <span v-if="lowerComments.comment_user_id != commentsItem.uid"><span
                                                            class="at-user">@{{
                                                                lowerComments.comment_user_name
                                                            }} </span> : </span> {{ lowerComments.context }}
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                                <div class="comment-partition"></div>
                            </div>
                        </div>
                    </div>
                    <!-- 回复评论  dialog-->
                    <el-dialog v-model="replyCommentsDialog.show" title="Shipping address">
                        <commentPosting :videoID="videoInfo?.videoInfo?.id" @updateVideoInfo="updateVideoInfo"
                            :commentsID="replyCommentsDialog.commentsID"></commentPosting>
                    </el-dialog>
                </div>
            </div>

            <!-- 右边部分 -->
            <div class="content-right">

                <div class="user-card">
                    <div class="card-left">
                        <div class="avatar" @click="jumpSpace()"> <el-avatar :size="60"
                                :src="videoInfo?.videoInfo?.creatorInfo?.avatar" />
                        </div>
                    </div>

                    <div class="card-right">
                        <div class="username">
                            <span>{{ videoInfo?.videoInfo?.creatorInfo?.username }}</span>
                            <div class="private-letter" @click="usePersonalLetter(videoInfo?.videoInfo?.uid)">
                                <SvgIcon name="message" class="icon" color="#9499A0"></SvgIcon> 私信
                            </div>
                        </div>
                        <div class="signature">
                            <VueEllipsis3
                                :text="videoInfo?.videoInfo?.creatorInfo?.signature ? videoInfo.videoInfo.creatorInfo.signature : '这个人很勤快什么都没留下~'">
                                <template v-slot:ellipsisNode>
                                    <span>...</span>
                                </template>
                            </VueEllipsis3>
                        </div>
                        <div class="btn-list">
                            <el-button v-removeFocus type="primary" size="small" round :icon="House"
                                @click="jumpSpace()">主页</el-button>
                            <el-button class="attention" v-if="!videoInfo?.videoInfo?.creatorInfo?.is_attention"
                                v-removeFocus type="primary" size="small" round :icon="Plus"
                                @click="attention()">关注</el-button>
                            <el-button class="attention" v-if="videoInfo?.videoInfo?.creatorInfo?.is_attention"
                                v-removeFocus type="primary" size="small" round :icon="MoreFilled" color="#F1F2F3"
                                @click="attention()">已关注</el-button>
                        </div>
                    </div>
                </div>


                <div class="barrage-box">
                    <div class="barrage-top-navigation">
                        <div class="box-left"><span>弹幕列表</span>
                            <SvgIcon name="dropDown" class="icon" color="#61666D"></SvgIcon>
                        </div>
                        <span @click="barrageListShow = !barrageListShow">
                            <SvgIcon name="rightArrow" class="arrow" color="#61666D"></SvgIcon>
                        </span>
                    </div>
                    <div class="barrage-list" :style="{ height: barrageListShow ? '700px' : '0px' }">
                        <el-table :data="videoInfo.barrageList" style="width: 100%">
                            <el-table-column prop="time" label="时间" width="90">
                                <template #default="scope">
                                    <span>{{ formattingSecondTime(scope.row.time) }}</span>
                                </template>
                            </el-table-column>
                            <el-table-column prop="text" label="弹幕内容" width="260">
                                <template #default="scope">
                                    <VueEllipsis3 :text="scope.row.text">
                                        <template v-slot:ellipsisNode>
                                            <span>...</span>
                                        </template>
                                    </VueEllipsis3>
                                </template>
                            </el-table-column>
                            <el-table-column prop="sendTime" label="发送时间">
                                <template #default="scope">
                                    <span>{{ dayjs(rFC3339ToTime(scope.row.sendTime)).format('MM-DD HH:mm') }}</span>
                                </template>
                            </el-table-column>
                        </el-table>
                    </div>
                    <div class="recommend-lsit">
                        <div class="video-item" v-for="recommendItem in videoInfo?.recommendList" :key="recommendItem.id">
                            <Card :id="recommendItem.id" :title="recommendItem.title" :username="recommendItem.username"
                                :video_duration="recommendItem.video_duration"
                                v-bind:barrage-number="recommendItem.barrageNumber" :heat="recommendItem.heat"
                                :cover="recommendItem.cover" :created_at="recommendItem.created_at"></Card>
                        </div>
                    </div>
                </div>

            </div>
        </div>

        <!-- 收藏视频 -->
        <div class="dialog-list">
            <div class="dialog">
                <el-dialog v-model="favoriteVideoShow" title="收藏视频" width="24%" center :close-on-click-modal="true"
                    :append-to-body="true" :before-close="favoriteVideoShowClose" align-center>
                    <FavoriteVideo @shutDown="shutDown" @success="videoInfo.videoInfo.is_collect = true" :id="videoID">
                    </FavoriteVideo>
                </el-dialog>
            </div>
        </div>


    </div>
</template>

<script lang="ts" setup>
import globalScss from "@/assets/styles/global/export.module.scss";
import commentPosting from "@/components/commentPosting/videoCommentPosting.vue";
import FavoriteVideo from "@/components/favoriteVideo/favoriteVideo.vue";
import topNavigation from "@/components/topNavigation/topNavigation.vue";
import videoIntroduce from "@/components/videoIntroduce/videoIntroduce.vue";
import Card from "@/components/videoPageVideoCard/videoPageVideoCard.vue";
import useAttention from "@/hooks/useAttention";
import usePersonalLetter from "@/hooks/usePersonalLetter";
import { useInit, useLikeVideo, useSendBarrage, useVideoProp, useWebSocket } from "@/logic/show/video/video";
import { GetVideoCommentRes } from '@/types/show/video/video';
import { formattingSecondTime, rFC3339ToTime } from "@/utils/conversion/timeConversion";
import { vRemoveFocus } from "@/utils/customInstruction/focus";
import { House, MoreFilled, Plus, Position, Star } from '@element-plus/icons-vue';
import dayjs from 'dayjs';
import DPlayer from "dplayer";
import Swal from "sweetalert2";
import { onMounted, onUnmounted, ref, watch } from "vue";
import { VueEllipsis3 } from 'vue-ellipsis-3';


components: {
    commentPosting
    topNavigation
    Card
    videoIntroduce
    VueEllipsis3
    FavoriteVideo
}

var dp = ref(<DPlayer>{})  //播放器配置对象
var socket: WebSocket //socket
const { videoRef, route, router, userStore, videoID, videoInfo, barrageInput, barrageListShow, liveNumber, replyCommentsDialog, videoBarrage, global } = useVideoProp()

const sendBarrageEvent = () => {
    useSendBarrage(barrageInput, dp.value, userStore, videoInfo, socket)
}

//回复二级评论
const replyComments = (commentsID: number) => {
    replyCommentsDialog.commentsID = commentsID
    replyCommentsDialog.show = !replyCommentsDialog.show
}

//更新评论数据
const updateVideoInfo = (commentsList: GetVideoCommentRes) => {
    videoInfo.videoInfo.comments = commentsList.comments
    videoInfo.videoInfo.comments_number = commentsList.comments_number
}

const attention = async () => {
    if (await useAttention(videoInfo.videoInfo.creatorInfo.id)) {
        videoInfo.videoInfo.creatorInfo.is_attention = !videoInfo.videoInfo.creatorInfo.is_attention
    }
}

//跳转用户空间
const jumpSpace = () => {
    router.push({ name: "SpaceIndividual", params: { "id": videoInfo.videoInfo.creatorInfo.id } })
}


//收藏视频
const favoriteVideoShow = ref(false)
//关闭createCollectDialogShow
const shutDown = () => {
    favoriteVideoShow.value = false
}

const favoriteVideoShowClose = (done: () => void) => {
    done()
}




const watchPath = watch(() => route.path, async () => {
    console.log(route.path)
    if (!route.path.includes('/videoShow/video')) {
        return false
    }
    console.log(videoID)
    dp.value = await useInit(videoRef, route, router, videoID, videoInfo, global) as DPlayer
    if (userStore.userInfoData.token) {
        let socketLer = useWebSocket(userStore, videoInfo, router, liveNumber)
        socketLer ? socket = socketLer : ""
    }
}, { immediate: true, deep: true })


const watchDanmaku = watch(videoBarrage, (newVal, oldVal) => {
    if (videoBarrage.value) {
        dp.value.danmaku.show()
    } else {
        dp.value.danmaku.hide()
    }

})


const notOpen = () => {
    Swal.fire({
        title: "敬请期待",
        heightAuto: false,
        confirmButtonColor: globalScss.colorButtonTheme,
        icon: "info",
    })
}

onMounted(async () => {

})

onUnmounted(() => {
    //销毁监听和socket
    watchPath()
    watchDanmaku()
    socket?.close()
})

</script>

<style lang="scss" scoped>
@import "@/assets/styles/views/show/video/video.scss";
</style>
