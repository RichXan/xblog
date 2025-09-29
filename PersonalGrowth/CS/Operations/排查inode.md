# 排查inode占用问题
1. 找出占用 inode 最多的目录
```
sudo find / -xdev -printf '%h\n' | sort | uniq -c | sort -rn | head -20
```

2. 如何查看 inode 使用情况
```
df -i
```

3. baikal php
PHP 默认将 session 存在文件里，每次访问生成一个文件

如果没有定期清理，文件会无限堆积

Docker overlay2 会保留镜像层的文件，每个小文件占用一个 inode

最终导致根分区 inode 被耗尽 → 系统无法创建临时文件 → bash 报错

2️⃣ docker-compose.yml 示例
```
version: '3.8'

services:
  baikal:
    image: ckulka/baikal:nginx
    container_name: baikal
    ports:
      - "7800:80"
    volumes:
      # 挂载 Baikal 数据库和 session 目录到宿主机
      - /mnt/data/baikal-sessions:/var/lib/php/sessions
      - /mnt/data/baikal-data:/var/www/html/Specific
      - /mnt/data/baikal-config:/var/www/html/html/config
      # 挂载自定义 php.ini 配置
      - ./php_custom.ini:/usr/local/etc/php/conf.d/custom.ini
    environment:
      - TZ=Asia/Shanghai
```

3️⃣ 自定义 PHP 配置文件 php_custom.ini

在 docker-compose.yml 里挂载到容器 /usr/local/etc/php/conf.d/custom.ini，内容示例：

```
; PHP session 配置
session.save_path = "/var/lib/php/sessions"
session.gc_maxlifetime = 3600        ; 1小时
session.gc_probability = 1
session.gc_divisor = 100
```

含义：

session.gc_maxlifetime：Session 文件存活时间，单位秒

session.gc_probability / session.gc_divisor：垃圾回收概率 = probability/divisor = 1/100，每 100 次请求触发一次清理