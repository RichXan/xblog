# Redis

## Redis 持久化方式

Redis 本身是基于内存的数据库，如果没有持久化机制，重启就会丢失数据。主要有两种持久化方式：

### 1. RDB（Redis Database File）快照

- 定时将内存中的数据快照保存到磁盘。
- 在 redis.conf 里通过 save 配置规则，比如：
    ```nginx
    save 900 1   # 900 秒内至少 1 次写操作，就触发快照
    save 300 10  # 300 秒内至少 10 次写操作
    save 60 10000
    ```
- 文件是二进制的 .rdb 文件，恢复速度快。
- 缺点：可能丢失最近一次快照后的数据。

### 2. AOF（Append Only File）追加日志

- 以日志形式记录每一次写命令，Redis 重启时会重放 AOF 文件来恢复数据。
- appendfsync 参数可控制刷盘策略：
    - always：每次写都刷盘，最安全，最慢。
    - everysec（默认）：每秒刷一次盘，可能丢失 1 秒数据。
    - no：依赖操作系统，性能最好但最不安全。
  - AOF 文件会随着写入增长，Redis 会做 AOF 重写（rewrite），压缩成最小指令集。

### 3. RDB+AOF 混合持久化（Redis 4.0+）

- 结合 RDB + AOF。
- AOF 重写时，先保存一份 RDB 快照，再在文件后追加 AOF 日志。
- 优点：既有 RDB 的快速恢复，又能减少 AOF 过大问题。

### Redis持久化配置方式
Redis 镜像默认使用 /data 目录保存持久化文件。你需要做两件事：
- 挂载 volume（不然容器重启数据会丢失）
- 配置 redis.conf（决定用 RDB / AOF / 混合）

**示例`docker-compose.yml`**
```yaml
version: '3.9'
services:
  redis:
    image: redis:7.2
    container_name: myredis
    restart: always
    ports:
      - "6379:6379"
    volumes:
      - ./data:/data               # 数据文件持久化目录
      - ./redis.conf:/usr/local/etc/redis/redis.conf  # 自定义配置文件
    command: ["redis-server", "/usr/local/etc/redis/redis.conf"]
```

**示例`redis.conf`（开启 AOF + RDB）**
```conf
# 开启RDB
save 900 1
save 300 10
save 60 10000
stop-writes-on-bgsave-error yes
rdbcompression yes
dbfilename dump.rdb
dir /data

# 开启AOF
appendonly yes
appendfilename "appendonly.aof"
appendfsync everysec

# 混合持久化（Redis 4.0+）
aof-use-rdb-preamble yes
```

这样：
- Redis 会定期生成 dump.rdb（RDB）
- 所有写操作也会写入 appendonly.aof（AOF）
- 两者结合，保证数据安全又能快速恢复

## 过期键的处理机制
Redis 支持给 key 设置过期时间（EXPIRE key seconds，SETEX）。但是，过期数据并不是一到时间点就立刻被删除，而是通过以下机制：

### 1. 惰性删除（Lazy Expiration）

- 客户端访问一个 key 时，Redis 会检查它是否过期，如果过期就删除。
- 这样避免了定时扫描全库的高开销，但缺点是如果从不访问，过期数据可能长期占用内存。

### 2. 定期删除（Active Expiration）

- Redis 会周期性随机抽样一些带过期时间的 key，检查是否过期，如果过期就删除
- 抽样数量和频率由 hz 参数（默认 10）控制，即每秒做 10 次过期扫描。
- 这样可以避免内存长期堆积过多过期数据。

### 3. 内存淘汰（Eviction, 当内存不足时）

- 如果开启了 maxmemory 限制，当内存满了，会根据策略淘汰键：
    - volatile-lru：从设置了过期时间的 key 中挑选最近最少使用的删除。
    - volatile-ttl：从设置了过期时间的 key 中，优先删除剩余时间最短的。
    - allkeys-lru：所有 key 中挑选最近最少使用的删除。
    - allkeys-random：随机删除 key。
    - volatile-random：从设置了过期时间的 key 中随机删除。
    - noeviction：直接返回错误，不删除。

### 过期键配置
Redis 的过期键清理机制主要由 redis.conf 参数控制：
```conf
# 每秒进行过期扫描次数（默认10次）
hz 10

# 主动过期删除参数
active-expire-effort 1   # 取值1-10，数值越大越积极
```

- 惰性删除：不需要额外配置，Redis 默认就有。
- 定期删除：由 hz 和 active-expire-effort 控制扫描频率。
- 内存淘汰策略：需要你显式配置 maxmemory-policy，比如：
```conf
# 最大内存限制（例如 512MB）
maxmemory 512mb

# 内存淘汰策略
maxmemory-policy allkeys-lru
```