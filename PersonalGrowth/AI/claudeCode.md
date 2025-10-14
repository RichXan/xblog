

- Claude 使用指南
    -[Claude 使用指南](https://claude.develop-on.co.kr/zh/)

- Claude 国内中转镜像
    - [88code](https://88code.org/)
    - [claudebuddy](https://claudebuddy.fun/)
    - [aicodemirror](https://www.aicodemirror.com/)

- Claude MPC
    - [Claude Code Template](https://www.aitmpl.com/)
    - Sequential Thinking: `claude mcp add thinking -s user -- npx -y @modelcontextprotocol/server-sequential-thinking`
    - Backend Architect: `npx claude-code-templates@latest --agent=development-team/backend-architect --yes`
    - Golang Pro: `npx claude-code-templates@latest --agent=programming-languages/golang-professional --yes`



- 常用命令
| 命令	| 说明	|
| ------- | ------- |
| claude	| 启动交互式会话	|
| claude -p "生成一个登录页面"	| 一次性查询	|
| claude --continue	| 恢复上次对话	|
| /clear	| 清空会话	|
| /cost	| 查看 token 使用	|
| /memory	| 编辑项目记忆文件（CLAUDE.md）	|
| /review	| 请求代码审查	|
| /config	| 查看或修改配置	|
| claude update	| 更新 Claude Code 到最新版本	|

export ANTHROPIC_BASE_URL="https://www.88code.org/api"
export ANTHROPIC_AUTH_TOKEN="88_b00d0c860a46a60f3b4a4963ff2ad8c2ba4497b268198d947aecddff989ef41b"