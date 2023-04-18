<template>
  <div class="article-page">
    <div class="head">
      <topNavigation color="#fff" scroll :displaySearch="false"></topNavigation>
    </div>
    <!-- 封面图 -->
    <div class="cover-picture" :style="{ backgroundImage: `url(${articleInfo.cover})` }">
      <div class="article-info-container">
        <div class="title">{{ articleInfo.title }}</div>
        <div class="article-info">
          <span class="info-lb">
            <SvgIcon name="camera" class="icon-small"></SvgIcon> {{
              dayjs(articleInfo.created_at).format('YYYY.MM.DD.hh.mm')
            }}
          </span>
          <span class="info-lb">
            <SvgIcon name="hot" class="icon-small"></SvgIcon>
            <span>
              {{ articleInfo.heat }}
            </span>
          </span>
          <span class="info-lb">
            <SvgIcon name="comments" class="icon-small"></SvgIcon>
            <span>
              {{ articleInfo.comments_number }}
            </span>
          </span>
          <span class="info-lb">
            <SvgIcon name="like" class="icon-small"></SvgIcon>
            <span>
              {{ articleInfo.likes_number }}
            </span>
          </span>
        </div>
      </div>
    </div>
    <!-- 文章内容 -->
    <div class="article-content">
      <div class="content">
        <!-- 其他 ql-container ql-snow -->
        <div class="ql-editor" v-html="articleInfo.content">
        </div>
      </div>
      <!-- 底部评论 -->
      <div class="comments-box">
        <div class="comments-head">
          <SvgIcon name="editor" class="icon-edit"></SvgIcon> 评论
        </div>
        <div class="comments-main">
          <commentPosting :articleID="articleID" :articleInfo="articleInfo" @updateArticleInfo="updateArticleInfo"
            :commentsID="0"></commentPosting>
        </div>
        <div class="comments-show">
          <div class="comments-show-titel"><span>Comments | </span> <span>{{ articleInfo.comments_number }} 条评论</span>
          </div>
          <div class="comments-show-info">
            <div class="comment-info-detail" v-for="commentsItem in articleInfo.comments" :key="commentsItem.id">
              <el-avatar shape="square" :size="36" :src="commentsItem.photo" />
              <div class="comment-info-content">
                <div class="content-head">
                  <div> <span class="comment-info-username">{{ commentsItem.username }}</span> <span
                      class="comment-info-other">{{ dayjs(commentsItem.created_at).format('YYYY.MM.DD.hh.mm') }}</span>
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
                <div class="comment-info-detail" v-for="lowerComments in commentsItem.lowerComments"
                  :key="lowerComments.id">
                  <el-avatar shape="square" :size="36" :src="lowerComments.photo" />
                  <div class="comment-info-content">
                    <div class="content-head">
                      <div> <span class="comment-info-username">{{ lowerComments.username }}</span> <span
                          class="comment-info-other">{{ dayjs(lowerComments.created_at).format('YYYY.MM.DD.hh.mm')
                          }}</span>
                      </div>
                      <div class="commentInfo-reply">
                        <el-tag effect="dark" round v-removeFocus @click="replyComments(lowerComments.id)">
                          回复
                        </el-tag>
                      </div>
                    </div>
                    <!-- 评论内容部分 -->
                    <div class="content-info">
                      <span v-if="lowerComments.comment_user_id != commentsItem.uid"><span class="at-user">@{{
                        lowerComments.comment_user_name
                      }} </span> : </span> {{ lowerComments.context }}
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <!-- 回复评论  dialog-->
        <el-dialog v-model="replyCommentsDialog.show" title="Shipping address">
          <commentPosting :articleID="articleID" :articleInfo="articleInfo" @updateArticleInfo="updateArticleInfo"
            :commentsID="replyCommentsDialog.commentsID"></commentPosting>
        </el-dialog>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts" >
import commentPosting from "@/components/commentPosting/commentPosting.vue"

import 'quill/dist/quill.bubble.css'
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'


import topNavigation from "@/components/topNavigation/topNavigation.vue"
import { useArticleShowProp, useInit } from "@/logic/show/article/article"
import dayjs from "dayjs"
//代码高亮
import { GetArticleCommentRes } from "@/types/show/article/article"
import { vRemoveFocus } from "@/utils/customInstruction/focus"
import { blossom } from "@/utils/effect/blossom"
import 'highlight.js/styles/agate.css'
import { onMounted, onUnmounted } from "vue"

components: {
  topNavigation
  commentPosting
}

const { articleID, articleInfo, router, route, replyCommentsDialog, global } = useArticleShowProp()

//更新评论数据
const updateArticleInfo = (commentsList: GetArticleCommentRes) => {
  articleInfo.value.comments = commentsList.comments
  articleInfo.value.comments_number = commentsList.comments_number
}

//回复二级评论
const replyComments = (commentsID: number) => {
  replyCommentsDialog.commentsID = commentsID
  replyCommentsDialog.show = !replyCommentsDialog.show
}

//樱花动效
const { startSakura, stopp } = blossom()

onMounted(async () => {
  startSakura()
  await useInit(articleID, articleInfo, route, router, global)
})

onUnmounted(() => {
  stopp()
})

</script>

<style scoped lang="scss">
// 添加相关样式
@import "@/assets/styles/views/show/article/article.scss";
</style>
