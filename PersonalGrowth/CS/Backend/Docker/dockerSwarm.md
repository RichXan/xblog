# docker swarm

## 基础概念
| 概念 | 描述 |
| --- | --- |
| Swarm | 一个集群，由多个节点组成。每个 swarm 有一个或多个 manager 节点（领导者）和 worker 节点。 |
| Node | 集群中的单个 Docker 主机。可以是 manager（管理集群状态、调度任务）或 worker（执行任务）。一个 swarm 可以有多个 manager 以实现高可用。 |
| Service | Swarm 中的核心抽象，表示一个或多个容器的分布式部署。服务定义了容器的镜像、端口、 replicas（副本数）等，支持自动扩展和滚动更新。 |
| Task | 服务的最小执行单元，通常是一个容器。Swarm 调度器将任务分配到节点上。 |
| Stack | 使用 Docker Compose 文件定义的多服务应用堆栈，可以通过 docker stack deploy 部署到 swarm。 |
| Overlay Network | Swarm 内置的网络模式，支持节点间通信，即使节点在不同主机上。 |
| Load Balancing | Swarm 自动提供服务发现和负载均衡，通过内置的路由网格（routing mesh）将流量分发到服务副本。 |
| Secrets | 安全存储敏感数据（如密码、证书），仅在运行时注入到容器中。 |
| Raft Consensus | Manager 节点使用 Raft 算法维护集群状态一致性，确保高可用（至少 3 个 manager 以容忍 1 个故障）。 |

## 基础命令
| 命令 | 描述 |
| --- | --- |
| docker swarm init | 初始化一个 swarm 集群。 |
| docker swarm join | 加入一个 swarm 集群。 |
| docker swarm leave | 离开一个 swarm 集群。 |
| docker swarm update | 更新一个 swarm 集群。 |