# 拉取仓库代码

若有submodule，需要使用--recursive参数拉取代码，并且提前在submodule的仓库中添加sshkey

```
# 创建目录
sudo mkdir deploy-family-calendar-server
# 更改目录权限给dev
sudo chown dev:dev /mnt/data/docker/deploy-family-calendar-server/
# 查看目录权限是不是dev用户的
ls -ld ./deploy-family-calendar-server/
# 使用dev用户拉取仓库
sudo -Hu dev git clone git@sgit.oazon.com-calendarloop:calendar-loop/deploy-family-calendar-server.git
```

## 测试SSH Key是否生效
如果执行
```
ssh -T git@sgit.oazon.com
```
如果返回
Welcome to GitLab, @你的用户名! → SSH Key 已经识别。

如果返回 Permission denied (publickey) → SSH Key 还没生效（可能没上传对）。

## 如果还是拉取失败的话，在66网段的机器使用ipv4拉去代码
```
GIT_SSH_COMMAND="ssh -4" git clone git@sgit.oazon.com:calendar-loop/deploy-family-calendar-server.git
```

# 拉取submodule的代码
1. 在git仓库中添加sshkey
2. 在主仓库中执行指令拉取submodule
```
git submodule update --recursive
```

# 如果deploy构建镜像失败时
尝试手动拉取镜像：
```bash
docker pull node:20-alpine
```
如果手动拉取成功，重新运行 ./deploy

如果失败，观察详细错误信息