# uv使用基础教程
1. uv创建虚拟环境
```bash
uv venv
source .venv/bin/activate
```
2. 退出虚拟环境
```bash
deactivate
```

3. 临时添加镜像源
```bash
export UV_INDEX_URL=https://pypi.tuna.tsinghua.edu.cn/simple/
```

3. 安装依赖
```bash
uv pip install -r requirements.txt
```

4. 使用uv
```bash
uv run python --version
```

