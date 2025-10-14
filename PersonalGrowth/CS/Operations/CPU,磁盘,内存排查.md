# CPU
- `uptime`: 查看系统整体负载情况。
  - load average：过去 1 分钟 / 5 分钟 / 15 分钟的平均负载。
  - 一般认为 load ≈ CPU核心数 时负载较高。
    - 例如 4 核 CPU，如果 load > 4，说明有点吃紧。

- `top / htop`: 查看 CPU 使用率。按 1 可以查看各个CPU状态
  - top 界面按 1 → 会显示每个核心的使用情况
  - htop 更直观，会用柱状图显示所有核心
    - s → 用户态 CPU 使用率
    - sy → 内核态 CPU 使用率
    - id → 空闲 CPU 百分比
    - wa → I/O 等待，如果很高说明 CPU 在等磁盘。

- `lscpu`: 查看物理 CPU、核心数、逻辑 CPU（超线程）
  - Socket(s) → 物理 CPU 个数
  - Core(s) per socket → 每个 CPU 的核心数
  - Thread(s) per core → 每个核心的线程数（超线程）
  - CPU(s) → 总逻辑 CPU 数 = Socket × Core × Thread

# 磁盘
- `df -h` : 查看硬盘使用情况，如果出现95%以上的情况，则需要扩容磁盘空间。
- `df -i` : 查看inode使用情况。


# 内存
- `free -h`: 查看内存使用情况。


# 交换内存
- `free -h`: 查看内存使用情况。
- `vmstat 1 5`: 
  - si（swap in）和 so（swap out）数值大 → 系统频繁在换页
  - 看 si 和 so 是否大于 0。若频繁换页，需要立刻调整。