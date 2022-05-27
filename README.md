# easy-douyin

## 抖音项目服务端简单示例

具体功能内容参考飞书说明文档

工程无其他依赖，直接编译运行即可

```shell
go build && ./simple-demo
```

### 功能说明

已实现接口

* 登录/douyin/user/login/
* 注册/douyin/user/register/
* 用户信息/douyin/user/

未实现接口 douyin/
* 视频流接口/feed/
* 投稿接口/publish/action/
* 发布列表/publish/list/
* 赞操作/favorite/action/
* 点赞列表/favorite/list/
* 评论操作/comment/action/
* 评论列表/comment/list/"
* 关注操作/relation/action/
* 关注列表/relation/follow/list/
* 粉丝列表/relation/follower/list/

### 说明
* 用户登录数据保存在内存中，单次运行过程中有效
* 视频上传后会保存到本地 public 目录中，访问时用 127.0.0.1:8080/static/video_name 即可

### 测试数据

测试数据写在 demo_data.go 中，用于列表接口的 mock 测试