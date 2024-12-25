# Monitor Service

一个基于 Go 语言开发的监控服务系统，支持 HTTP 接口监控和自定义告警通知。

## 项目结构

```
monitor/
├── api                     # API 定义和文档
│   └── v1                 # API 版本控制
│       └── openapi.yaml   # API 规范文档
├── build                  # 构建相关
│   ├── package           # 打包脚本
│   ├── ci                # CI 配置文件
│   └── docker           # Docker 相关文件
│       ├── Dockerfile
│       └── docker-compose.yml
├── cmd                    # 主程序入口
│   └── server           # 服务器程序
│       └── main.go      # 主程序入口点
├── configs               # 配置文件
│   ├── config.yaml      # 基础配置
│   ├── config.dev.yaml  # 开发环境配置
│   └── config.prod.yaml # 生产环境配置
├── deployments          # 部署配置
│   ├── kubernetes      # K8s 配置
│   └── terraform       # 基础设施配置
├── docs                 # 项目文档
│   ├── architecture.md # 架构文档
│   └── api.md         # API 文档
├── internal            # 内部代码
│   ├── app            # 应用核心
│   │   └── server.go  # 服务器设置
│   ├── config         # 配置管理
│   │   └── config.go
│   ├── domain         # 领域模型
│   │   ├── entity    # 实体定义
│   │   ├── valueobject # 值对象
│   │   └── repository # 仓储接口
│   ├── infrastructure # 基础设施
│   │   ├── database  # 数据库
│   │   └── cache     # 缓存
│   ├── interfaces    # 接口层
│   │   ├── api      # API 处理
│   │   └── job      # 后台任务
│   ├── usecases     # 业务用例
│   │   ├── monitor
│   │   └── notification
│   └── pkg          # 内部工具包
│       ├── validator
│       └── errors
├── pkg              # 外部工具包
│   ├── httpclient
│   └── utils
├── scripts         # 脚本文件
│   ├── migrations # 数据库迁移
│   └── seed       # 测试数据
├── test           # 测试代码
│   ├── integration # 集成测试
│   ├── e2e        # 端到端测试
│   └── mocks      # 测试模拟
├── web            # Web 资源
│   ├── static    # 静态文件
│   └── templates # 模板文件
├── .gitignore
├── .golangci.yml  # 代码检查配置
├── Makefile       # 项目管理
├── go.mod         # 依赖管理
└── README.md      # 项目说明
```

## 功能特性

- HTTP 接口监控
  - 支持 GET/POST/PUT/DELETE 等方法
  - 自定义请求头和请求体
  - 灵活的监控规则配置
  - 可配置重试策略

- 告警通知
  - 钉钉机器人通知
  - PushPlus 通知
  - 自定义通知模板

- 监控管理
  - 监控任务的增删改查
  - 监控日志查询
  - 监控统计分析

## 快速开始

### 环境要求

- Go 1.21+
- MySQL 8.0+
- Redis 6.0+ (可选)

### 安装

1. 克隆项目
```bash
git clone https://github.com/yourusername/monitor.git
cd monitor
```

2. 安装依赖
```bash
go mod download
```

3. 配置数据库
```bash
# 复制配置文件
cp configs/config.example.yaml configs/config.yaml
# 修改配置文件中的数据库连接信息
```

4. 运行服务
```bash
make run
```

### 配置说明

配置文件位于 `configs/config.yaml`，主要包含以下配置项：

- 服务器配置
  - 监听地址和端口
  - 运行模式

- 数据库配置
  - 连接信息
  - 连接池设置

- 日志配置
  - 日志级别
  - 日志文件

- 监控配置
  - 默认监控间隔
  - 超时设置
  - 重试策略

- 通知配置
  - 钉钉机器人配置
  - PushPlus 配置

## API 文档

详细的 API 文档请参考 `api/v1/openapi.yaml` 或访问 [API 文档](docs/api.md)。

## 开发指南

### 代码规范

项目使用 [golangci-lint](https://golangci-lint.run/) 进行代码检查，配置文件为 `.golangci.yml`。

### 提交规范

提交信息格式：
```
<type>(<scope>): <description>

[optional body]

[optional footer]
```

类型（type）：
- feat: 新功能
- fix: 修复
- docs: 文档
- style: 格式
- refactor: 重构
- test: 测试
- chore: 构建/工具

### 测试

运行测试：
```bash
make test        # 运行单元测试
make test-cover  # 运行测试并生成覆盖率报告
make test-e2e    # 运行端到端测试
```

## 部署

### Docker 部署

```bash
# 构建镜像
make docker-build

# 运行容器
make docker-run
```

### Kubernetes 部署

部署配置文件位于 `deployments/kubernetes/`，包含：
- Deployment
- Service
- ConfigMap
- Secret

## 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交改动 (`git commit -m 'feat: Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件 