
## 服务器仓库

[https://github.com/Howlyao/REST_API_Server](https://github.com/Howlyao/REST_API_Server)

## 服务器配置

### fork服务器项目,clone至本地

```
# git clone repoURL
```

### go get服务器所需库函数

```
# go get github.com/boltdb/bolt    //数据库
# go get github.com/gorilla/mux    //服务器配置库 
```

### 修改部分代码文件的库路径

#### main.go
```
import(
    "log"
    "net/http"
    "github.com/Howlyao/REST_API_Server/service" //修改成克隆项目在本地的路径$GOPATH/.../.../REST_API_Server/service   

)

```

#### service/Handler.go

```
import(
    "fmt"
    "net/http"
    _ "unsafe"
    "github.com/gorilla/mux"
    "github.com/Howlyao/REST_API_Server/database" //修改成克隆项目在本地的路径$GOPATH/.../.../REST_API_Server/database   
)
```


#### 服务器项目文件说明

#####文件目录
```
REST_API_Server
│   main.go   
│
└───service
│   │   Handler.go
│   │   Routes.go
|   |   Router.go
│   
└───database
|   │   DB.go
|   │   my.db
|
└───initdb
│   │   InitialDB.go
│   │   resourses
|   |   
└───model
|   │   People.go
│   │   Film.go
│   │   Planet.go
|   |   Startship.go
│   │   Species.go
│   │   Vehicle.go
|   |   Transport.go

```
#### service

##### 服务器Route以及Handler函数定义

#### database

##### 数据库数据信息以及数据库读取函数定义

#### initdb

##### 根据获取的resources资源数据（JSON信息）,初始化，按特定格式处理插入数据库

#### model

##### Model类定义，根据Json信息，定义对应的model类。

### 服务器运行

```
# cd 项目路径
# go run main.go
```





