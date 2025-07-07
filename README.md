# Go Utils

一个轻量级的 Go 语言工具包，提供日常开发所需的各种实用工具和组件。

## 🚀 特性

- **依赖注入**: 轻量级依赖注入框架
- **泛型工具**: 支持泛型的链表、Set、Map 等数据结构
- **配置管理**: 基于 YAML 的配置工具
- **对象池**: 泛型对象池
- **注册器**: 类型安全的注册处理系统
- **访问控制**: 简单的访问权限管理
- **工具函数**: 丰富的切片、Map 操作工具

## 📦 安装

```bash
go get github.com/mengri/utils
```

## 🎯 主要模块

### 1. 依赖注入 (Autowire-v2)

轻量级依赖注入框架：

```go
import "github.com/mengri/utils/autowire-v2"

type UserService struct {
    DB IDatabase `autowired:""`
}

// 自动注册
func init() {
    autowire.Auto(func() IUserService { return &UserService{} })
}

// 使用
func main() {
    autowire.Check()
    
    var userService IUserService
    autowire.Inject(&userService)
}
```

### 2. 泛型数据结构

#### 双向链表 (List)
```go
import "github.com/mengri/utils/list"

l := list.New[string]()
l.PushBack("hello")
l.PushBack("world")
l.PushFront("hi")

// 遍历
for e := l.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value)
}
```

#### 泛型 Set
```go
import "github.com/mengri/utils/utils"

set := utils.NewSet[string]()
set.Set("apple", "banana", "orange")

if set.Has("apple") {
    fmt.Println("Found apple")
}

list := set.ToList()
```

#### 泛型 Untyped Map
```go
import "github.com/mengri/utils/untyped"

m := utils.BuildUntyped[string, int]()
m.Set("age", 25)
m.Set("score", 100)

if age, ok := m.Get("age"); ok {
    fmt.Println("Age:", age)
}

keys := m.Keys()
values := m.List()
```

### 3. 对象池 (Pool)

泛型对象池：

```go
import "github.com/mengri/utils/pool"

type Buffer struct {
    data []byte
}

// 创建对象池
bufferPool := pool.New[*Buffer](func() *Buffer {
    return &Buffer{data: make([]byte, 0, 1024)}
})

// 使用对象池
buffer := bufferPool.Get()
defer bufferPool.PUT(buffer)

// 使用 buffer...
```

### 4. 配置管理 (CFTool)

基于 YAML 的配置工具：

```go
import "github.com/mengri/utils/cftool"

type DatabaseConfig struct {
    Host string `yaml:"host"`
    Port int    `yaml:"port"`
}

// 注册配置
func init() {
    cftool.Register[DatabaseConfig]("database")
}

// 从文件加载
func main() {
    cftool.ReadFile("config.yaml")
}
```

### 5. 注册器 (Register)

类型安全的注册处理系统：

```go
import "github.com/mengri/utils/register"

type User struct {
    Name string
    Age  int
}

// 注册处理器
func init() {
    register.Handle[User](func(user User) {
        fmt.Printf("User created: %s, age: %d\n", user.Name, user.Age)
    })
}

// 调用处理器
func main() {
    user := User{Name: "Alice", Age: 30}
    register.Call(user)
}
```

### 6. 访问控制 (Access)

简单的访问权限管理：

```go
import "github.com/mengri/utils/access"

// 定义访问权限
permissions := []access.Access{
    {Name: "read", CName: "读取", Desc: "读取权限"},
    {Name: "write", CName: "写入", Desc: "写入权限"},
}

// 添加权限组
access.Add("article", permissions)

// 获取权限
if perms, ok := access.Get("article"); ok {
    for _, perm := range perms {
        fmt.Printf("Permission: %s (%s)\n", perm.Name, perm.CName)
    }
}
```

### 7. 忽略工具 (Ignore)

路径忽略管理：

```go
import "github.com/mengri/utils/ignore"

// 设置忽略路径
ignore.IgnorePath("api", "GET", "/health")
ignore.IgnorePath("api", "POST", "/internal/*")

// 检查是否忽略
if ignore.IsIgnorePath("api", "GET", "/health") {
    fmt.Println("This path is ignored")
}
```

### 8. 工具函数

#### 切片操作
```go
import "github.com/mengri/utils/utils"

users := []User{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
}

// 切片转换
names := utils.SliceToSlice(users, func(u User) string { return u.Name })

// 切片转 Map
userMap := utils.SliceToMap(users, func(u User) string { return u.Name })

// 切片转 Map 数组
ageGroups := utils.SliceToMapArray(users, func(u User) int { return u.Age / 10 })
```

#### MD5 计算
```go
import "github.com/mengri/utils/utils"

data := []byte("hello world")
hash := utils.MD5(data)
fmt.Println(hash)
```

#### 集合操作
```go
import "github.com/mengri/utils/utils"

a := []string{"apple", "banana", "orange"}
b := []string{"banana", "grape", "orange"}

// 求交集
intersection := utils.Intersection(a, b)
fmt.Println(intersection) // [banana orange]
```

### 9. 环境配置 (Env)

环境变量和配置管理：

```go
import "github.com/mengri/utils/env"

// 获取环境变量
dbURL := env.GetEnv("DATABASE_URL", "localhost:3306")
isDebug := env.IsDebug()

// 获取工作目录
workDir := env.GetWorkDir()
```

### 10. 编码工具 (Encode)

YAML 编码解码：

```go
import "github.com/mengri/utils/encode"

type Config struct {
    Name string `yaml:"name"`
    Port int    `yaml:"port"`
}

config := Config{Name: "myapp", Port: 8080}

// 编码
data, err := encode.YamlMarshal(config)
if err != nil {
    panic(err)
}

// 解码
var newConfig Config
err = encode.YamlUnmarshal(data, &newConfig)
if err != nil {
    panic(err)
}
```

### 11. 版本信息

```go
import "github.com/mengri/utils/version"

// 获取版本信息
fmt.Println(string(version.VersionsInfo()))
```

## 🛠️ 完整示例

```go
package main

import (
    "fmt"
    "github.com/mengri/utils/autowire-v2"
    "github.com/mengri/utils/cftool"
    "github.com/mengri/utils/register"
    "github.com/mengri/utils/utils"
)

type Config struct {
    AppName string `yaml:"app_name"`
    Port    int    `yaml:"port"`
}

type Logger interface {
    Log(message string)
}

type ConsoleLogger struct{}

func (l *ConsoleLogger) Log(message string) {
    fmt.Printf("[LOG] %s\n", message)
}

type App struct {
    Logger Logger `autowired:""`
    Config Config `autowired:""`
}

func init() {
    // 注册依赖
    autowire.Auto(func() Logger { return &ConsoleLogger{} })
    autowire.Auto(func() *App { return &App{} })
    
    // 注册配置
    cftool.Register[Config]("app")
    
    // 注册处理器
    register.Handle[string](func(msg string) {
        fmt.Println("Received message:", msg)
    })
}

func main() {
    // 加载配置
    cftool.InitFor("app", []byte(`
app_name: "My App"
port: 8080
`))
    
    // 检查依赖
    autowire.Check()
    
    // 获取应用实例
    var app *App
    autowire.Inject(&app)
    
    app.Logger.Log("Application started")
    
    // 使用工具函数
    numbers := []int{1, 2, 3, 4, 5}
    doubled := utils.SliceToSlice(numbers, func(n int) int { return n * 2 })
    fmt.Println("Doubled:", doubled)
    
    // 使用注册器
    register.Call("Hello, World!")
}
```

## 📋 依赖

- **Go**: 1.22.2+
- **YAML**: gopkg.in/yaml.v3

## 🧪 测试

运行测试：
```bash
go test ./...
```

运行特定模块测试：
```bash
go test ./autowire-v2/
go test ./list/
go test ./register/
```

## 📚 API 文档

生成文档：
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
- [YAML 规范](https://yaml.org/spec/)

## 🆕 版本信息

```go
import "github.com/mengri/utils/version"

fmt.Println(string(version.VersionsInfo()))
```

---

*这是一个轻量级的 Go 工具包，专注于提供实用的工具函数和组件，无外部重型依赖。*

