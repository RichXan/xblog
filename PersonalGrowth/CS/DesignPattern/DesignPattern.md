# Go 语言设计模式

Go 语言设计模式的实例代码

## 创建型模式
- [简单工厂模式（Simple Factory）](./00SimpleFactory.md)
- [工厂方法模式（Factory Method）](./01FactoryMethod.md)
- [抽象工厂模式（Abstract Factory）](./02AbstractFactory.md)
- 创建者模式（Builder）
- 原型模式（Prototype）
- 单例模式（Singleton）

## 结构型模式
- 外观模式（Facade）
- 适配器模式（Adapter）
- 代理模式（Proxy）
- 组合模式（Composite）
- 享元模式（Flyweight）
- [装饰模式（Decorator）](./11Decorator.go)
- 桥模式（Bridge）
- [函数选项模式（FunctionOption）](./25FunctionOption.go)

## 行为型模式
- 中介者模式（Mediator）
- 观察者模式（Observer）
- 命令模式（Command）
- 迭代器模式（Iterator）
- 模板方法模式（Template Method）
- 策略模式（Strategy）
- 状态模式（State）
- 备忘录模式（Memento）
- 解释器模式（Interpreter）
- 职责链模式（Chain of Responsibility）
- 访问者模式（Visitor）

# 面向对象设计原则
- 单一职责原则（Single Responsibility Principle）
    - 一个类只负责一项职责  
- 开闭原则（Open/Closed Principle）
    - 对扩展开放，对修改关闭
- 里氏替换原则（Liskov Substitution Principle）
    - 子类可以替换父类。所有引用基类对象的地方能够透明地使用其子类的对象	
- 依赖倒置原则（Dependency Inversion Principle）
    - 高层模块不应该依赖低层模块，两者都应该依赖抽象。抽象不应该依赖于细节，细节应该依赖于抽象
- 接口隔离原则（Interface Segregation Principle）
    - 客户端不应该依赖它不需要的接口。使用多个专门的接口，而不使用单一的总接口
- 迪米特法则（Law of Demeter）
    - 一个类对另一个类的了解越少越好。一个软件实体应当尽可能少地与其他实体发生相互作用
- 合成复用原则（Composite Reuse Principle）
    - 尽量使用合成/聚合的方式，而不是继承