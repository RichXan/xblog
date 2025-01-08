# Go 语言设计模式

Go 语言设计模式的实例代码

## 创建型模式
- [简单工厂模式（Simple Factory）](/PersonalGrowth/CS/DesignPattern/00SimpleFactory.md)
- [工厂方法模式（Factory Method）](/PersonalGrowth/CS/DesignPattern/01FactoryMethod.md)
- [抽象工厂模式（Abstract Factory）](/PersonalGrowth/CS/DesignPattern/02AbstractFactory.md)
- [创建者模式（Builder）](/PersonalGrowth/CS/DesignPattern/03Builder.md)
- [原型模式（Prototype）](/PersonalGrowth/CS/DesignPattern/04Prototype.md)
- [单例模式（Singleton）](/PersonalGrowth/CS/DesignPattern/05Singleton.md)

## 结构型模式
- [外观模式（Facade）](/PersonalGrowth/CS/DesignPattern/06Facade.md)
- [适配器模式（Adapter）](/PersonalGrowth/CS/DesignPattern/07Adapter.md)
- [代理模式（Proxy）](/PersonalGrowth/CS/DesignPattern/08Proxy.md)
- [享元模式（Flyweight）](/PersonalGrowth/CS/DesignPattern/09Flyweight.md)
- [组合模式（Composite）](/PersonalGrowth/CS/DesignPattern/10Composite.md)
- [装饰模式（Decorator）](/PersonalGrowth/CS/DesignPattern/11Decorator.md)
- [桥模式（Bridge）](/PersonalGrowth/CS/DesignPattern/12Bridge.md)
- [过滤器模式（Filter、Criteria Pattern）](/PersonalGrowth/CS/DesignPattern/13Filter.md)
- [函数选项模式（FunctionOption）](/PersonalGrowth/CS/DesignPattern/25FunctionOption.md)

## 行为型模式
- [中介者模式（Mediator）](/PersonalGrowth/CS/DesignPattern/14Mediator.md)
- [观察者模式（Observer）](/PersonalGrowth/CS/DesignPattern/15Observer.md)
- [命令模式（Command）](/PersonalGrowth/CS/DesignPattern/16Command.md)
- [迭代器模式（Iterator）](/PersonalGrowth/CS/DesignPattern/17Iterator.md)
- [模板方法模式（Template Method）](/PersonalGrowth/CS/DesignPattern/18TemplateMethod.md)
- [策略模式（Strategy）](/PersonalGrowth/CS/DesignPattern/19Strategy.md)
- [状态模式（State）](/PersonalGrowth/CS/DesignPattern/20State.md)
- [备忘录模式（Memento）](/PersonalGrowth/CS/DesignPattern/21Memento.md)
- [解释器模式（Interpreter）](/PersonalGrowth/CS/DesignPattern/22Interpreter.md)
- [职责链模式（Chain of Responsibility）](/PersonalGrowth/CS/DesignPattern/23ChainOfResponsibility.md)
- [访问者模式（Visitor）](/PersonalGrowth/CS/DesignPattern/24Visitor.md)
- 空对象模式（Null Object Pattern）


## 设计原则
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

开发中的使用：
- 单例模型（初始化时getConfig, getLogger）
- 外观模型（handler->service->dao）
- 适配器模式（一个接口中支持新旧版本的实现）


> 参考资料：
- [菜鸟教程](https://www.runoob.com/design-pattern)
- [senghoo/golang-design-pattern](https://github.com/senghoo/golang-design-pattern)
- [youlookwhat/DesignPattern](https://github.com/youlookwhat/DesignPattern)
