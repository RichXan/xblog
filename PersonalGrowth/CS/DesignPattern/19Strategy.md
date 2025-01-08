# 策略模式（Strategy）
策略模式定义了一系列算法，并将每个算法封装起来，使它们可以相互替换，且算法的变化不会影响使用算法的客户。策略模式属于行为型模式。

## 主要解决的问题
- 如何在运行时选择算法
- 如何封装算法族
- 如何消除大量的条件判断
- 如何使算法的变化独立于使用算法的客户

## 应用实例
1. Java的Comparator接口
2. Spring的Resource访问策略
3. 支付系统的支付方式
4. 日志框架的日志级别

## 使用场景
1. 算法变体
   - 排序算法
   - 压缩算法
   - 缓存策略
2. 业务规则
   - 折扣计算
   - 佣金计算
   - 税费计算
3. 系统行为
   - 认证方式
   - 日志记录
   - 资源访问
4. 配置选项
   - 导出格式
   - 通知方式
   - 渲染模式

## 优缺点
### 优点
1. 算法可替换
   - 运行时切换
   - 独立演化
2. 避免条件语句
   - 消除if-else
   - 提高可维护性
3. 扩展性好
   - 易于增加策略
   - 符合开闭原则

### 缺点
1. 策略数量增加
   - 类数量增多
   - 维护成本上升
2. 客户了解策略
   - 需要了解所有策略
   - 选择成本增加
3. 对象创建
   - 每个策略都是对象
   - 内存占用增加

## 代码实现

```golang
package designpattern

// Strategy 定义策略接口
type Strategy interface {
    Execute(data interface{}) interface{}
}

// ConcreteStrategyA 具体策略A
type ConcreteStrategyA struct{}

func (s *ConcreteStrategyA) Execute(data interface{}) interface{} {
    return "Executing strategy A with " + fmt.Sprint(data)
}

// ConcreteStrategyB 具体策略B
type ConcreteStrategyB struct{}

func (s *ConcreteStrategyB) Execute(data interface{}) interface{} {
    return "Executing strategy B with " + fmt.Sprint(data)
}

// Context 上下文
type Context struct {
    strategy Strategy
}

func NewContext(strategy Strategy) *Context {
    return &Context{strategy: strategy}
}

func (c *Context) SetStrategy(strategy Strategy) {
    c.strategy = strategy
}

func (c *Context) ExecuteStrategy(data interface{}) interface{} {
    return c.strategy.Execute(data)
}

// 实际应用示例：支付系统
type PaymentStrategy interface {
    Pay(amount float64) string
}

// CreditCardPayment 信用卡支付
type CreditCardPayment struct {
    cardNumber string
    cvv        string
}

func NewCreditCardPayment(cardNumber, cvv string) *CreditCardPayment {
    return &CreditCardPayment{
        cardNumber: cardNumber,
        cvv:        cvv,
    }
}

func (p *CreditCardPayment) Pay(amount float64) string {
    return fmt.Sprintf("Paid %.2f using Credit Card %s", amount, p.cardNumber)
}

// PayPalPayment PayPal支付
type PayPalPayment struct {
    email string
}

func NewPayPalPayment(email string) *PayPalPayment {
    return &PayPalPayment{email: email}
}

func (p *PayPalPayment) Pay(amount float64) string {
    return fmt.Sprintf("Paid %.2f using PayPal account %s", amount, p.email)
}

// AlipayPayment 支付宝支付
type AlipayPayment struct {
    id string
}

func NewAlipayPayment(id string) *AlipayPayment {
    return &AlipayPayment{id: id}
}

func (p *AlipayPayment) Pay(amount float64) string {
    return fmt.Sprintf("Paid %.2f using Alipay account %s", amount, p.id)
}

// PaymentContext 支付上下文
type PaymentContext struct {
    strategy PaymentStrategy
}

func NewPaymentContext(strategy PaymentStrategy) *PaymentContext {
    return &PaymentContext{strategy: strategy}
}

func (c *PaymentContext) SetStrategy(strategy PaymentStrategy) {
    c.strategy = strategy
}

func (c *PaymentContext) ProcessPayment(amount float64) string {
    return c.strategy.Pay(amount)
}
```

## 使用示例

```golang
func main() {
    // 基本示例
    context := NewContext(&ConcreteStrategyA{})
    result := context.ExecuteStrategy("data")
    
    context.SetStrategy(&ConcreteStrategyB{})
    result = context.ExecuteStrategy("data")
    
    // 支付示例
    creditCard := NewCreditCardPayment("1234-5678-9012-3456", "123")
    paypal := NewPayPalPayment("user@example.com")
    alipay := NewAlipayPayment("user123")
    
    payment := NewPaymentContext(creditCard)
    fmt.Println(payment.ProcessPayment(100.50))
    
    payment.SetStrategy(paypal)
    fmt.Println(payment.ProcessPayment(50.75))
    
    payment.SetStrategy(alipay)
    fmt.Println(payment.ProcessPayment(200.00))
}
```

## 类图
```mermaid
classDiagram
    %% 基本策略模式
    class Strategy {
        <<interface>>
        +Execute(data interface{}) interface{}
    }

    class ConcreteStrategyA {
        +Execute(data interface{}) interface{}
    }

    class ConcreteStrategyB {
        +Execute(data interface{}) interface{}
    }

    class Context {
        -strategy Strategy
        +SetStrategy(Strategy)
        +ExecuteStrategy(interface{}) interface{}
    }

    %% 支付系统示例
    class PaymentStrategy {
        <<interface>>
        +Pay(amount float64) string
    }

    class CreditCardPayment {
        -cardNumber string
        -cvv string
        +Pay(amount float64) string
    }

    class PayPalPayment {
        -email string
        +Pay(amount float64) string
    }

    class AlipayPayment {
        -id string
        +Pay(amount float64) string
    }

    class PaymentContext {
        -strategy PaymentStrategy
        +SetStrategy(PaymentStrategy)
        +ProcessPayment(float64) string
    }

    %% 关系
    Strategy <|.. ConcreteStrategyA
    Strategy <|.. ConcreteStrategyB
    Context o-- Strategy
    PaymentStrategy <|.. CreditCardPayment
    PaymentStrategy <|.. PayPalPayment
    PaymentStrategy <|.. AlipayPayment
    PaymentContext o-- PaymentStrategy
```

## 说明
1. 策略模式的主要角色：
   - Strategy（策略）：定义所有支持的算法的公共接口
   - ConcreteStrategy（具体策略）：实现了Strategy接口的具体算法
   - Context（上下文）：持有一个Strategy的引用
2. 实现要点：
   - 策略的选择机制
   - 策略的参数传递
   - 策略的创建方式
3. 设计考虑：
   - 是否需要策略的默认实现
   - 是否需要策略的组合
   - 是否需要策略的缓存
4. 相关模式：
   - 工厂模式：创建策略对象
   - 状态模式：不同的行为方式
   - 命令模式：不同的命令执行
