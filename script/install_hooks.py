import os
import stat
from pathlib import Path

def install_hooks():
    # 获取项目根目录
    root_dir = Path(__file__).parent.parent
    hooks_dir = root_dir / '.git' / 'hooks'
    
    # 确保 hooks 目录存在
    hooks_dir.mkdir(exist_ok=True)
    
    # 复制 pre-commit hook
    source = root_dir / 'script' / 'hooks' / 'pre-commit'
    target = hooks_dir / 'pre-commit'
    
    # 复制文件
    with open(source, 'r') as src, open(target, 'w') as dst:
        dst.write(src.read())
    
    # 设置执行权限
    os.chmod(target, stat.S_IRUSR | stat.S_IWUSR | stat.S_IXUSR)
    
    print("Git hooks installed successfully!")

if __name__ == "__main__":
    install_hooks() 