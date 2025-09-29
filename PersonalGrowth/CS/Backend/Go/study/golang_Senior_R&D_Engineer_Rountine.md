# Golang 高级研发工程师学习路线图
## 阶段 1：基础打牢（初级 → 中级）

目标：掌握 Go 语言核心特性，能独立开发小型项目。

- 语言核心
- Go 基础语法、内存模型、slice/map 底层实现
    - 并发编程：goroutine、channel、select
    - 错误处理模式（errors、fmt.Errorf、xerrors、wrap error）
    - 标准库：context、net/http、io、sync、time

- 工程实践
    - Go Modules 依赖管理
    - 单元测试、基准测试、mock
    - 常用工具：go fmt、golangci-lint、go vet、go doc

- 练手项目
    - CLI 工具（如一个 log 分析器）
    - REST API 小服务（用户注册登录）

## 阶段 2：进阶能力（中级 → 高级）

目标：能开发中大型 Web 服务，熟悉数据库与常见框架。
- Web 开发
    - Gin/Echo/Fiber 框架
    - Middleware 机制（认证、限流、日志）

- 数据库 & 缓存
    - ORM：GORM、Ent
    - Redis 缓存：分布式锁、缓存穿透/击穿/雪崩
    - SQL 优化、事务机制

- 微服务通信
    - gRPC & Protobuf
    - 拦截器、中间件、负载均衡

- 消息队列
    - Kafka、RabbitMQ、NATS
    - 消息重试、幂等性

- 运维工具
    - Docker 容器化
    - CI/CD 基础（GitHub Actions/GitLab CI）

- 练手项目
    - 博客系统（REST API + 数据库 + Redis 缓存）
    - 微服务小项目（用户服务 + 订单服务 + gRPC 通信）

## 阶段 3：高级研发（高级 → Tech Lead）

目标：具备架构思维，能处理高并发、分布式与性能优化问题。

- 架构设计
    - 分层架构（domain/service/repo）
    - DDD（领域驱动设计）基础
    - API 设计规范（RESTful / gRPC / GraphQL）

- 高并发与分布式
    - 分布式锁（Redis/Etcd/Zookeeper）
    - 分布式事务（Saga, TCC, Outbox Pattern）
    - 一致性哈希、CAP 理论、Raft 原理

- 性能优化
    - pprof 调优（CPU、内存、阻塞、goroutine 泄露）
    - 零拷贝 I/O、连接池优化
    - 高性能日志库（zap, zerolog）

- 可观测性
    - Prometheus + Grafana 监控
    - OpenTelemetry / Jaeger 链路追踪
    - ELK / Loki 日志平台

- 安全
    - OAuth2 / OIDC / JWT
    - TLS/mTLS 双向认证

- 练手项目
    - 高并发 IM 聊天服务
    - 电商下单系统（包含订单、库存、支付微服务 + 分布式事务）

## 阶段 4：架构师方向（Tech Lead → 架构师）

目标：能带领团队设计复杂系统，结合云原生和企业级实践。

- 云原生
    - Kubernetes (K8s)：服务编排、Helm、Operator
    - Service Mesh（Istio, Linkerd）
    - Serverless（Knative, OpenFaaS, AWS Lambda）

- 系统设计
    - Event-Driven 架构（CQRS、Event Sourcing）
    - API Gateway 设计（Kong, Envoy, Nginx + 自研）
    - 高可用系统（限流、熔断、降级、隔离舱）

- DevOps & 工程化
    - CI/CD 流程：灰度发布、蓝绿部署、金丝雀发布
    - Monorepo / Polyrepo 管理方式
    - 代码规范 & Review 机制

- 前沿探索
    - AI + Go：高并发 AI API Gateway
    - 边缘计算 & IoT：Go 在嵌入式和边缘节点上的应用
    - 大数据处理：Go + Kafka/Flink/ClickHouse

## 总结学习路径

1. 基础 → 熟练 Go 语法 + 并发模型
2. 进阶 → 掌握 Web 开发、数据库、微服务通信
3. 高级 → 深入分布式、性能优化、系统设计
4. 架构 → 云原生、DevOps、复杂系统架构能力


## Golang 学习路线表格
阶段	学习目标	核心知识点	项目练习
阶段 1：基础	掌握 Go 语言核心语法与并发模型	- Go 基础语法、内存模型
- slice/map 底层实现
- goroutine、channel、select
- context、errors 包
- 标准库（net/http、io、sync）
- 单元测试、基准测试、mock	- CLI 工具（如日志分析器）
- 简单 REST API 服务（用户注册/登录）
阶段 2：进阶	能开发中小型 Web 服务，熟悉数据库与常用框架	- Web 框架：Gin/Echo/Fiber
- 中间件机制（认证、限流、日志）
- 数据库：GORM、Ent、SQL 优化、事务
- Redis：缓存、分布式锁、缓存穿透/击穿
- gRPC & Protobuf
- 消息队列（Kafka/RabbitMQ/NATS）
- Docker 基础、CI/CD 入门	- 博客系统（REST API + DB + Redis）
- 微服务小项目（用户服务 + 订单服务 + gRPC 通信）
阶段 3：高级	能设计高并发分布式系统，进行性能优化	- 分层架构（domain/service/repo）
- DDD 基础、API 设计规范
- 分布式锁、事务（Saga、TCC、Outbox）
- CAP 理论、Raft/Paxos
- pprof 调优（CPU/内存/goroutine）
- 高性能日志库（zap/zerolog）
- Prometheus + Grafana
- OpenTelemetry/Jaeger
- OAuth2/OIDC/JWT
- TLS/mTLS 安全	- 高并发 IM 聊天服务
- 电商下单系统（含订单、库存、支付微服务 + 分布式事务）
阶段 4：架构师方向	能设计和落地复杂系统，带团队解决业务问题	- Kubernetes (K8s)：服务编排、Helm、Operator
- Service Mesh (Istio, Linkerd)
- Serverless (Knative, OpenFaaS, AWS Lambda)
- Event-Driven 架构（CQRS、Event Sourcing）
- API Gateway 设计（Kong/Envoy/Nginx + 自研）
- 系统高可用：限流、熔断、降级、隔离舱
- CI/CD：灰度发布、蓝绿部署、金丝雀发布
- 工程化：Monorepo/Polyrepo、代码规范	- 云原生微服务平台
- 高可用 API Gateway
- IoT/边缘计算服务
- Go + AI API Gateway（支持高并发 AI 调用）