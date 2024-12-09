# docker配置vpn网络代理

## 原因：

有时因为网络原因，比如公司 NAT，或其它啥的，需要使用代理。Docker 的代理配置，略显复杂，因为有三种场景。但基本原理都是一致的，都是利用 Linux 的 http_proxy 等环境变量。

## 配置步骤：

在执行docker pull时，是由守护进程dockerd来执行。因此，代理需要配在dockerd的环境中。而这个环境，则是受systemd所管控，因此实际是systemd的配置。

`sudo vim /etc/systemd/system/docker.service.d/http-proxy.conf`

在这个proxy.conf文件（可以是任意*.conf的形式）中，添加以下内容：
```
[Service]
Environment="HTTP_PROXY=http://127.0.0.1:7890"
Environment="HTTPS_PROXY=http://127.0.0.1:7890"
```

然后，执行`sudo systemctl daemon-reload`，使配置生效。

`sudo systemctl restart docker`，重启docker服务。


# docker配置镜像源

docker默认使用的是docker hub的镜像源，如果需要使用国内的镜像源，可以配置国内的镜像源。

`sudo vim /etc/docker/daemon.json`

```
{
  "registry-mirrors": [
    "https://registry.docker-cn.com",
    "http://mirrors.ustc.edu.cn"
  ]
}
```

然后，执行`sudo systemctl daemon-reload`，使配置生效。

`sudo systemctl restart docker`，重启docker服务。

# 主机配置dns域名解析服务器

`sudo vim /etc/resolv.conf`

```
nameserver 8.8.8.8
nameserver 8.8.4.4
```

重启dns服务

`sudo systemctl restart systemd-resolved`

# 主机配置网关

`sudo vim /etc/network/interfaces`

```
auto eth0
iface eth0 inet static
    address 192.168.1.100
    netmask 255.255.255.0
    gateway 192.168.1.1
```
