# 单例模式（Singleton）
单例模式确保一个类只有一个实例，并提供一个全局访问点。这种模式属于创建型模式，它提供了一种创建对象的最佳方式。

## 主要解决的问题
- 如何确保一个类只有一个实例
- 如何提供全局访问点
- 如何控制共享资源
- 如何延迟初始化

## 应用实例
1. 数据库连接池
2. 线程池管理
3. 系统配置管理
4. 日志记录器

## 使用场景
1. 资源管理
   - 连接池管理
   - 线程池管理
   - 缓存管理
2. 系统配置
   - 全局配置
   - 环境变量
   - 系统参数
3. 工具类
   - 日志工具
   - 工厂管理器
   - 注册表管理
4. 设备访问
   - 打印机管理
   - 驱动程序
   - 硬件访问

## 优缺点
### 优点
1. 内存节约
   - 只有一个实例
   - 避免资源浪费
2. 访问便捷
   - 全局访问点
   - 统一管理
3. 延迟加载
   - 按需初始化
   - 提高性能

### 缺点
1. 扩展困难
   - 职责过重
   - 耦合度高
2. 并发问题
   - 线程安全
   - 死锁风险
3. 测试困难
   - 全局状态
   - 单元测试

## 代码实现

```golang
package designpattern

import "sync"

/*
===========================
    1. 通用单例实现
===========================
*/
// Singleton 是单例模式接口
type Singleton interface {
    foo()
}

// singleton 是单例模式实现
type singleton struct{}

func (s *singleton) foo() {}

var (
    instance *singleton
    once     sync.Once
)

// GetInstance 用于获取单例模式对象
func GetInstance() Singleton {
    once.Do(func() {
        instance = &singleton{}
    })
    return instance
}

/*
===========================
    2. 饿汉式实现
===========================
*/

var hungryInstance = &singleton{}

// GetHungryInstance 饿汉式单例
func GetHungryInstance() Singleton {
    return hungryInstance
}

/* 
===========================
    3. 懒汉式实现（双重检查）
===========================
*/

var (
    lazyInstance *singleton
    mu          sync.Mutex
)

// GetLazyInstance 懒汉式单例
func GetLazyInstance() Singleton {
    if lazyInstance == nil {
        mu.Lock()
        defer mu.Unlock()
        if lazyInstance == nil {
            lazyInstance = &singleton{}
        }
    }
    return lazyInstance
}
```

## 使用示例

```golang
func main() {
    // 1. 使用 sync.Once 的实现
    instance1 := GetInstance()
    instance2 := GetInstance()
    // instance1 == instance2 总是成立
    
    // 2. 饿汉式
    hungry1 := GetHungryInstance()
    hungry2 := GetHungryInstance()
    // hungry1 == hungry2 总是成立
    
    // 3. 懒汉式
    lazy1 := GetLazyInstance()
    lazy2 := GetLazyInstance()
    // lazy1 == lazy2 总是成立
}
```

## 类图
```mermaid
classDiagram
    %% 接口定义
    class Singleton {
        <<interface>>
        +foo()
    }

    %% 单例实现
    class singleton {
        -instance *singleton
        -once sync.Once
        +foo()
    }

    %% 获取实例的方法
    class GetInstance {
        <<function>>
        +GetInstance() Singleton
    }

    class GetHungryInstance {
        <<function>>
        +GetHungryInstance() Singleton
    }

    class GetLazyInstance {
        <<function>>
        +GetLazyInstance() Singleton
    }

    %% 实现关系
    Singleton <|.. singleton
    GetInstance ..> singleton : creates
    GetHungryInstance ..> singleton : creates
    GetLazyInstance ..> singleton : creates
```

## 说明
1. 单例模式的主要角色：
   - Singleton（单例）：包含一个实例和访问方法
   - Instance（实例）：唯一的实例对象
   - Client（客户端）：使用单例的代码
2. 实现要点：
   - 构造函数私有化
   - 线程安全保证
   - 延迟加载实现
3. 设计考虑：
   - 是否需要线程安全
   - 是否需要延迟加载
   - 是否需要序列化
4. 相关模式：
   - 工厂模式：创建对象
   - 建造者模式：复杂对象
   - 原型模式：对象克隆

## Go语言中的单例模式应用

### 1. 数据库连接池

```golang
type DBConnection struct {
    // 数据库连接配置
}

var (
    dbInstance *DBConnection
    dbOnce     sync.Once
)

func GetDBConnection() *DBConnection {
    dbOnce.Do(func() {
        dbInstance = &DBConnection{
            // 初始化数据库连接
        }
    })
    return dbInstance
}
```

### 2. 配置管理器

```golang
type Config struct {
    settings map[string]interface{}
    mu       sync.RWMutex
}

var (
    configInstance *Config
    configOnce     sync.Once
)

func GetConfig() *Config {
    configOnce.Do(func() {
        configInstance = &Config{
            settings: make(map[string]interface{}),
        }
        // 加载配置文件
    })
    return configInstance
}

func (c *Config) Get(key string) interface{} {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.settings[key]
}

func (c *Config) Set(key string, value interface{}) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.settings[key] = value
}
```

### 3. 日志管理器

```golang
type Logger struct {
    file *os.File
    mu   sync.Mutex
}

var (
    loggerInstance *Logger
    loggerOnce     sync.Once
)

func GetLogger() *Logger {
    loggerOnce.Do(func() {
        loggerInstance = &Logger{
            // 初始化日志文件
        }
    })
    return loggerInstance
}

func (l *Logger) Log(message string) {
    l.mu.Lock()
    defer l.mu.Unlock()
    // 写入日志
}
```

### 4. 应用状态管理

```golang
type AppState struct {
    state map[string]interface{}
    mu    sync.RWMutex
}

var (
    stateInstance *AppState
    stateOnce     sync.Once
)

func GetAppState() *AppState {
    stateOnce.Do(func() {
        stateInstance = &AppState{
            state: make(map[string]interface{}),
        }
    })
    return stateInstance
}

func (a *AppState) GetState(key string) interface{} {
    a.mu.RLock()
    defer a.mu.RUnlock()
    return a.state[key]
}

func (a *AppState) SetState(key string, value interface{}) {
    a.mu.Lock()
    defer a.mu.Unlock()
    a.state[key] = value
}
```

## 最佳实践

1. **使用 sync.Once**
   - 保证线程安全
   - 避免双重检查锁定的复杂性
   - 确保初始化代码只执行一次

2. **并发安全**
   - 对共享资源的访问使用互斥锁
   - 读多写少的场景使用读写锁
   - 避免竞态条件

3. **延迟初始化**
   - 只在首次使用时初始化
   - 避免启动时的资源浪费
   - 处理初始化失败的情况

4. **资源清理**
   - 实现 Close 或 Shutdown 方法
   - 正确处理资源释放
   - 考虑优雅关闭的场景

5. **测试友好**
   - 提供重置单例的方法（仅用于测试）
   - 使用接口而不是具体类型
   - 方便进行单元测试

```golang
// 测试辅助函数示例
func ResetForTest() {
    instance = nil
    once = sync.Once{}
}
```

## 注意事项

1. 避免过度使用单例模式
2. 考虑使用依赖注入作为替代方案
3. 确保单例的状态是线程安全的
4. 注意初始化顺序和依赖关系
5. 在测试中要特别注意单例的重置
