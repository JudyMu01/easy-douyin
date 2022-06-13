# easy-douyin
* [📦 Install ](#install) -- Install relevant dependencies and the project
* [🔧 Usage ](#usage) -- Commands to run the server
* [📦 Directory structure ](#directory) -- What does each file do
## Install


### Install ffmpeg
```markdown
Browse and install ffmpeg according to your OS for generation of video covers.
https://ffbinaries.com/downloads
```

### Install dousheng apk
```markdown
Install dousheng android apk to test the program on your mobile phone.
https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7#
```

## Usage

```shell
go build && ./easy-douyin
```


## Directory

    .
    ├── controller                          # API functions and response structure
    │   ├── comment.go
    │   ├── common.go
    │   ├── favorite.go
    │   ├── feed.go
    │   ├── publish.go
    │   ├── relation.go
    │   └── user.go
    ├── public                              # public static resources on server
    │   ├── covers                          # store static pictures of video cover
    │   │   └──*.png
    │   └── videos                          # store static videos
    │       └──*.mp4
    ├── repository                          # init, models and CRUD of database
    │   ├── comment.go
    │   ├── db_init.go
    │   ├── follow.go
    │   ├── like.go
    │   ├── user.go
    │   └── video.go
    ├── service                             # realisation of functions in controller
    │   ├── comment.go
    │   ├── follow.go
    │   ├── like.go
    │   ├── user.go
    │   └── video.go
    ├── test                                # test files
    │   └── ...
    ├── util
    │   └── MD5.go                          # encryption function
    ├── middleware
    │   ├── logger.go                       # provide error log
    │   ├── auth.go                         # authority function handler
    │   └── jwt.go                          # generate and parse jwt
    ├── .gitattributes
    ├── .gitignore
    ├── go.mod
    ├── go.sum
    ├── main.go                             # start of execution
    └── router.go                           # path configuration
  


## Some explain
* User login data is stored in memory and is valid during a single run
* After the video is uploaded, it is saved to the local public directory for access  http://127.0.0.1:8080/public/videos/video_name 