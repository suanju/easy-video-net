<div align="center">

<img src="./.github/logo.png" height="150px" width="150px">

# easy-video-net

*Golang + Vue编写的视频弹幕网，支持视频，专栏 ，直播 ，im等功能*

__本项目仅处于实验性阶段，依然存在许多问题，但后续会持续维护__



</div>

## 预览
> 体验地址 [easy-video](http://easy-video.top/)

![image-20230419151645612](https://user-images.githubusercontent.com/64412088/233002215-359b2337-6224-4318-811c-b2195f3cef4a.png)
![image-20230419151937238](https://user-images.githubusercontent.com/64412088/233002263-ff599b43-00c7-4d9a-8caf-2797500b1787.png)
![image-20230419151819626](https://user-images.githubusercontent.com/64412088/233002291-0ff90253-5e13-4240-9d89-43fff9e455b5.png)
![image-20230419151958183](https://user-images.githubusercontent.com/64412088/233002317-6bb54307-b696-48a7-9f73-4a24bdd65261.png)
![image-20230419152335308](https://user-images.githubusercontent.com/64412088/233002344-96b837f1-8174-4d21-9bb7-5d1ea4fad625.png)
![image-20230419151844222](https://user-images.githubusercontent.com/64412088/233002384-374e5375-dad6-4516-9a45-2466ad63d1bb.png)

## 简介

`easy-video-net` 是基于Golang + Vue开发的前后端分离项目
- Server端采用 Golang + Gin + Gorm
- Web端采用 Vue3 + Typescript + Element-Plus

### 主要模块

1. 视频上传分享支持转码及弹幕功能
2. 稿件投稿使用富文本编辑器进行简单发布
3. 一个简单直播功能 需要使用[livego](https://github.com/gwuhaolin/livego)搭建直播服务
4. 简单的消息通知 及其im功能
5. 个人相关及其相关发布信息的CRUD

### 特点

- 上传支持 不用存储、 不同质量、 实现本地及阿里云oss存储、 支持分片上传、 断点续传、 oss直传
- 视频本地存储使用ffmpeg进行视频转码 , oss使用阿里云智能媒体转码
- 简单实现直播功能并且采用protobuf进行通信

### 项目结构

```
easy-vide-net
├─service 服务端代码
│  ├─assets 静态资源
│  ├─config 配置文件
│  ├─consts 常量定义
│  ├─controllers 控制器
│  ├─global 全局使用
│  ├─interaction 结构体定义
│  │  ├─receive   请求  
│  │  └─response   响应
│  ├─logic  逻辑处理
│  ├─middlewares 中间件
│  ├─models  模型定义
│  ├─proto   proto文件
│  ├─router   路由定义
│  ├─runtime 运行时文件如日志
│  └─utils 工具类
│              
└─web
    │  .env  配置文件
    └─src
        ├─apis 接口定义
        ├─assets 静态资源
        ├─components 组件
        ├─hooks 
        ├─logic 逻辑代码
        ├─proto proto文件
        ├─router 路由定义
        ├─store 状态管理
        ├─types 类型定义
        ├─utils 工具类
        └─views 视图文件
```


## 构建 || 运行

### 依赖
- Server端依赖
    1. golang v1.18
    2. mysql  v8.0
    3. reids  v3.0
    4. ffmpeg  v4.2

- Web端依赖
    1. node v16.16

### Server端构建 || 运行
1. 配置 [./service/config/config.ini](./service/config/config.ini)
2. `cd service && go mod tidy`
3. `go build`
4. `./easy-video-net.exe`


### Web端构建 || 运行
- `cd web && npm i`
- `npm run dev`

## License
查看 [./LICENSE]()