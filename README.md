# IssuePilot

**基于 Go 的 GitHub Issue 智能分析与研发协作 Agent 平台**

IssuePilot 是一个面向软件研发团队的后端项目，旨在整合 GitHub Issue 数据，通过 AI 能力辅助开发者理解、分析和管理仓库中的问题。

## 项目目标

帮助开发者更高效地处理 GitHub Issue，实现：

- GitHub 仓库绑定与管理
- GitHub Issue 同步、查询与筛选
- Issue 内容摘要与智能分类
- Issue 优先级评估与处理建议
- 相似 Issue 检索
- 基于自然语言的 Issue 查询与分析 Agent

## 技术栈

| 技术 | 用途 |
|---|---|
| Go | 后端开发语言 |
| Gin | HTTP Web 框架 |
| GORM | ORM 数据库操作 |
| PostgreSQL | 关系型数据库 |
| pgvector | 向量检索 |
| Redis | 缓存与异步任务 |
| go-github | GitHub API 集成 |
| JWT | 用户身份认证 |
| Apifox / Postman | API 接口测试 |

## 项目架构规划

```text
issuepilot/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── config/
│   ├── router/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   ├── model/
│   └── middleware/
├── configs/
├── docs/
├── scripts/
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## 当前开发进度

### 已完成

- [x] Go 项目初始化
- [x] Gin 框架集成
- [x] HTTP 服务启动
- [x] 健康检查接口
- [x] Git 本地版本管理
- [x] GitHub 远程仓库创建与代码推送

### 开发计划

- [ ] 项目配置管理
- [ ] 路由与业务模块分层
- [ ] 统一 API 响应格式
- [ ] PostgreSQL 数据库集成
- [ ] 用户注册、登录与 JWT 认证
- [ ] GitHub 仓库绑定
- [ ] GitHub Issue 同步与管理
- [ ] AI 智能分析
- [ ] 相似 Issue 检索
- [ ] Agent 自然语言查询
- [ ] 自动化测试与工程化部署

## 本地运行

### 环境要求

- Go 1.27.1 或兼容版本
- Git

### 克隆项目

```bash
git clone https://github.com/zzz-363/IssuePilot.git
cd IssuePilot
```

### 安装依赖

```bash
go mod tidy
```

### 启动服务

```bash
go run ./cmd/api
```

服务默认监听：

```text
http://localhost:8080
```

### 健康检查

请求：

```http
GET http://localhost:8080/api/v1/health
```

预期响应：

```json
{
  "code": 200,
  "message": "ok",
  "data": {
    "service": "issuepilot",
    "status": "healthy"
  }
}
```

## 项目原则

- 初期采用模块化单体架构，不急于拆分微服务。
- AI 能力优先通过云端 API 接入，不要求本地运行大模型。
- Agent 初期只提供查询和分析能力，不自动修改 GitHub Issue。
- 通过接口测试验证每个阶段的功能。

## License

待定。
