# Common Utils

一个功能丰富的 Go 语言工具包，提供企业级应用开发所需的各种通用组件和工具。

## 🚀 特性

- **依赖注入**: 提供两个版本的依赖注入框架
- **缓存系统**: 支持 Redis 的高性能缓存
- **数据库存储**: 基于 GORM 的数据库操作封装
- **权限管理**: 完整的权限控制系统
- **HTTP 服务器**: 基于 Gin 的 Web 服务器
- **插件系统**: 灵活的插件架构
- **工具函数**: 丰富的辅助工具函数

## 📦 安装

```bash
go get github.com/mengri/utils
```

## 🎯 主要模块

### 1. 依赖注入 (Autowire)

提供两个版本的依赖注入框架：

#### Autowire V1
```go
import "github.com/mengri/utils/autowire"

type UserService struct {
    DB *gorm.DB `autowired:""`
}

func main() {
    autowire.Autowired(&UserService{})
}
```

#### Autowire V2
```go
import "github.com/mengri/utils/autowire-v2"

// 自动注册
func init() {
    autowire.Auto(func() IUserService { return &UserService{} })
}
```

### 2. 缓存系统 (Cache)

支持 Redis 的 KV 缓存系统：

```go
import "github.com/mengri/utils/cache"

type User struct {
    ID   int64  `json:"id"`
    Name string `json:"name"`
}

// 创建缓存
cache := cache.CreateKvCache[User, int64](
    redisClient, 
    time.Minute*10,
    func(id int64) string { return fmt.Sprintf("user:%d", id) },
)

// 使用缓存
user, err := cache.Get(ctx, 123)
err = cache.Set(ctx, 123, &user)
```

### 3. 数据库存储 (Store)

基于 GORM 的数据库操作封装：

```go
import "github.com/mengri/utils/store"

type User struct {
    ID   int64  `gorm:"primaryKey"`
    Name string `gorm:"not null"`
}

func (u *User) TableName() string { return "users" }
func (u *User) IdValue() int64 { return u.ID }

// 使用存储
store := store.NewBaseStore[User](db)
user, err := store.Get(ctx, 123)
err = store.Save(ctx, &user)
```

### 4. 权限管理 (Permit)

完整的权限控制系统：

```go
import "github.com/mengri/utils/permit"

// 添加权限
err := permit.Add(ctx, "read", "article", "article:123")

// 检查权限
allowed, err := permit.Check(ctx, "article", []string{"article:123"}, []string{"read"})

// 获取权限
targets, err := permit.Granted(ctx, "read", "article")
```

### 5. HTTP 服务器 (Server)

基于 Gin 的 Web 服务器：

```go
import "github.com/mengri/utils/server"

server := server.NewServer()
server.GET("/users", func(c *gin.Context) {
    // 处理逻辑
})
```

### 6. 插件系统 (Plugins)

灵活的插件架构：

```go
import "github.com/mengri/utils/plugins"

type MyPlugin struct{}

func (p *MyPlugin) Initialize() error {
    // 初始化逻辑
    return nil
}

// 注册插件
plugins.Register("my-plugin", &MyPlugin{})
```

### 7. 泛型链表 (List)

泛型双向链表实现：

```go
import "github.com/mengri/utils/list"

l := list.New[int]()
l.PushBack(1)
l.PushBack(2)
l.PushFront(0)

// 遍历
for e := l.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value)
}
```

### 8. 自动化工具 (Auto)

提供自动化处理工具：

```go
import "github.com/mengri/utils/auto"

type Config struct {
    Database Label `aolabel:"database"`
    Redis    Label `aolabel:"redis"`
}

// 自动标签处理
labels := auto.CreateLabels(&config)
```

### 9. 环境配置 (Env)

环境变量和配置管理：

```go
import "github.com/mengri/utils/env"

// 获取环境变量
dbURL := env.GetEnv("DATABASE_URL", "default_url")
isDebug := env.IsDebug()
```

### 10. 编码工具 (Encode)

提供 YAML 等编码工具：

```go
import "github.com/mengri/utils/encode"

// YAML 编码/解码
data, err := encode.YamlMarshal(config)
err = encode.YamlUnmarshal(data, &config)
```

## 🛠️ 工具函数

### 切片操作
```go
import "github.com/mengri/utils/utils"

// 切片转换
ids := utils.SliceToSlice(users, func(u *User) int64 { return u.ID })

// 切片转 Map
userMap := utils.SliceToMap(users, func(u *User) int64 { return u.ID })
```

### 其他工具
```go
// MD5 计算
hash := utils.MD5(data)

// Set 操作
set := utils.NewSet[string]()
set.Add("item1")
set.Add("item2")
```

## 📋 依赖

- **Go**: 1.22.2+
- **Gin**: Web 框架
- **GORM**: ORM 框架
- **Redis**: 缓存支持
- **UUID**: 唯一 ID 生成

## 🔧 配置

### 数据库配置
```go
import "github.com/mengri/utils/store/store_mysql"

config := store_mysql.Config{
    Host:     "localhost",
    Port:     3306,
    Database: "mydb",
    Username: "user",
    Password: "password",
}
```

### Redis 配置
```go
import "github.com/mengri/utils/cache/cache_redis"

redisClient := cache_redis.NewRedisClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "",
    DB:       0,
})
```

## 🧪 测试

运行测试：
```bash
go test ./...
```

运行特定模块测试：
```bash
go test ./autowire/
go test ./cache/
go test ./store/
```

## 📖 示例

### 完整的 Web 应用示例
```go
package main

import (
    "context"
    "github.com/mengri/utils/autowire-v2"
    "github.com/mengri/utils/server"
    "github.com/mengri/utils/store"
    "github.com/gin-gonic/gin"
)

type User struct {
    ID   int64  `json:"id" gorm:"primaryKey"`
    Name string `json:"name" gorm:"not null"`
}

func (u *User) TableName() string { return "users" }
func (u *User) IdValue() int64 { return u.ID }

type UserService struct {
    store store.IBaseStore[User] `autowired:""`
}

func (s *UserService) GetUser(ctx context.Context, id int64) (*User, error) {
    return s.store.Get(ctx, id)
}

func init() {
    autowire.Auto(func() *UserService { return &UserService{} })
}

func main() {
    // 初始化依赖
    autowire.Check()
    
    // 创建服务器
    srv := server.NewServer()
    
    srv.GET("/users/:id", func(c *gin.Context) {
        // 处理用户请求
        c.JSON(200, gin.H{"message": "success"})
    })
    
    srv.Run(":8080")
}
```

## 📚 API 文档

详细的 API 文档可以通过 `godoc` 生成：
```bash
godoc -http=:6060
```

## 🤝 贡献

欢迎贡献代码！请确保：

1. 添加适当的测试
2. 遵循 Go 代码规范
3. 添加必要的文档注释
4. 提交前运行 `go fmt` 和 `go vet`

## 📄 许可证

本项目采用 MIT 许可证。

## 🔗 相关链接

- [Go 官方文档](https://golang.org/doc/)
- [Gin 框架文档](https://gin-gonic.com/docs/)
- [GORM 文档](https://gorm.io/docs/)
- [Redis Go 客户端](https://redis.uptrace.dev/)

## 🆕 版本信息

可以通过以下方式查看版本信息：
```go
import "github.com/mengri/utils/utils"

fmt.Println(string(utils.VersionsInfo()))
```

## 📞 支持

如有问题或建议，请提交 Issue 或联系开发团队。

