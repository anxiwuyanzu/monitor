# Monitor 个人监控系统

一个基于 Go 语言开发的个人网站监控系统，支持自定义监控规则和多种通知方式。

## 技术栈

### 后端
- Go 1.21+
- Gin Web 框架
- GORM ORM 框架
- MySQL 数据库
- Cron 定时任务

### 前端
- Vue 3
- Element Plus
- Axios

## 项目结构

```
.
├── cmd/                    # 主程序入口
│   └── server/            # 服务器入口
├── config/                # 配置文件
├── internal/              # 内部包
│   ├── handler/          # HTTP 处理器
│   ├── middleware/       # 中间件
│   ├── model/           # 数据模型
│   ├── repository/      # 数据仓库
│   └── service/         # 业务逻辑
├── pkg/                  # 公共包
│   ├── database/        # 数据库相关
│   ├── logger/          # 日志相关
│   └── utils/           # 工具函数
└── web/                 # 前端代码
```

## 功能特性

- [ ] 监控目标管理（增删改查）
- [ ] 监控规则配置
- [ ] 定时任务调度
- [ ] 多种通知方式（钉钉、PushPlus等）
- [ ] 监控日志记录
- [ ] 数据可视化

## 快速开始

### 1. 克隆项目
```bash
git clone <repository-url>
cd monitor
```

### 2. 配置项目
1. 复制配置文件模板：
```bash
cp config/config.yaml.example config/config.yaml
```

2. 修改配置文件 `config/config.yaml`：
   - 数据库配置：
     - 设置数据库连接信息（host、port、username、password）
     - 确保数据库已创建（默认数据库名：monitor）
   - 日志配置：
     - 设置日志级别和存储位置
   - 通知配置（可选）：
     - 钉钉机器人：设置 webhook 和 secret
     - PushPlus：设置 token
   - 监控配置：
     - 可根据需要调整监控间隔和超时设置

### 3. 运行后端服务
```bash
go run cmd/server/main.go
```

### 4. 运行前端服务（待实现）
```bash
cd web
npm install
npm run dev
```

## 配置说明

### 服务器配置
```yaml
server:
  port: 8080  # 服务器端口
  mode: debug # 运行模式：debug/release
```

### 数据库配置
```yaml
database:
  driver: mysql
  host: localhost     # 数据库主机地址
  port: 3306         # 数据库端口
  username: root     # 数据库用户名
  password: ******   # 数据库密码
  dbname: monitor    # 数据库名称
```

### 通知配置
支持多种通知方式：

1. 钉钉机器人：
```yaml
dingtalk:
  enabled: true
  webhook: https://oapi.dingtalk.com/robot/send?access_token=xxx
  secret: your-secret
```

2. PushPlus：
```yaml
pushplus:
  enabled: true
  token: your-token
```

## 数据库设计

包含三个主要表：
- monitors：监控目标表
- monitor_rules：监控规则表
- notifications：通知配置表

## 贡献指南

1. Fork 本仓库
2. 创建您的特性分支：`git checkout -b feature/AmazingFeature`
3. 提交您的更改：`git commit -m 'feat: Add some AmazingFeature'`
4. 推送到分支：`git push origin feature/AmazingFeature`
5. 提交 Pull Request

## 许可证

[MIT License](LICENSE) 