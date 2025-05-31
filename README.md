# Jia-File

一个基于 Go 语言开发的简单文件管理系统，提供文件操作 API。

## 功能特性

- 文件列表查看
- 创建文件和目录
- 删除文件和目录
- 移动和复制文件
- 获取文件信息
- 创建文档（支持多种格式）
- 支持配置根目录限制

## 项目结构

```
.
├── api/            # API 类型定义
├── cmd/            # 主程序入口
├── internal/       # 内部包
│   ├── file/      # 文件操作
│   ├── handler/   # HTTP 处理器
│   ├── logger/    # 日志模块
│   ├── middleware/# 中间件
│   └── errors/    # 错误处理
├── scripts/       # 脚本文件
└── docs/          # 文档
```

## 快速开始

### 环境要求

- Go 1.16 或更高版本
- Air (用于热更新，可选)
- Docker (可选，用于容器化部署)

### 安装依赖

```bash
# 安装 Air (用于热更新)
go install github.com/cosmtrek/air@latest
```

### 配置说明

可以通过环境变量或 .env 文件配置以下参数：

- `PORT`: 服务器端口号（默认：8190）
- `LOG_LEVEL`: 日志级别（默认：info）
- `LOG_DIR`: 日志目录（默认：logs）
- `ROOT_PATH`: 文件操作的根目录（可选）
  - 如果设置，所有文件操作都将限制在此目录下
  - 支持相对路径和绝对路径
  - 如果未设置，则不限制文件操作范围

### 运行项目

#### 方式一：直接运行

```bash
# 编译项目
go build -o jia-file.exe ./cmd/server

# 运行服务器
./jia-file.exe
```

#### 方式二：使用 Air 热更新（推荐开发时使用）

```bash
# 在项目根目录运行
air
```

#### 方式三：使用 Docker（推荐生产环境使用）

```bash
# 构建 Docker 镜像
docker build -t jia-file .

# 运行容器
docker run -d \
  --name jia-file \
  -p 8190:8190 \
  -v /path/to/your/files:/app/files \
  -e ROOT_PATH=/app/files \
  jia-file
```

服务器将在 http://localhost:8190 启动。

### API 端点

- `GET /list?path=<path>` - 列出目录内容
- `POST /mkdir?path=<path>` - 创建目录
- `POST /touch?path=<path>` - 创建文件
- `DELETE /delete?path=<path>` - 删除文件或目录
- `POST /move?src=<src>&dst=<dst>` - 移动文件或目录
- `POST /copy?src=<src>&dst=<dst>` - 复制文件或目录
- `GET /info?path=<path>` - 获取文件信息
- `POST /document` - 创建文档

### 路径处理说明

当配置了 `ROOT_PATH` 时：
- 相对路径会自动与 `ROOT_PATH` 拼接
- 绝对路径会被验证是否在 `ROOT_PATH` 下
- 如果尝试访问 `ROOT_PATH` 外的文件，将返回错误

当未配置 `ROOT_PATH` 时：
- 直接使用传入的路径，不进行任何修改
- 支持相对路径和绝对路径

### Docker 部署说明

#### 构建镜像

```bash
docker build -t jia-file .
```

#### 运行容器

基本运行：
```bash
docker run -d --name jia-file -p 8190:8190 jia-file
```

使用数据卷和环境变量：
```bash
docker run -d \
  --name jia-file \
  -p 8190:8190 \
  -v /path/to/your/files:/app/files \
  -e ROOT_PATH=/app/files \
  -e LOG_LEVEL=info \
  jia-file
```

#### 环境变量

可以在运行容器时通过 `-e` 参数设置以下环境变量：

- `PORT`: 服务器端口号（默认：8190）
- `LOG_LEVEL`: 日志级别（默认：info）
- `LOG_DIR`: 日志目录（默认：/app/logs）
- `ROOT_PATH`: 文件操作的根目录

#### 数据卷

建议使用数据卷来持久化存储文件：

```bash
docker run -d \
  --name jia-file \
  -p 8190:8190 \
  -v /host/path/to/files:/app/files \
  -e ROOT_PATH=/app/files \
  jia-file
```

### 示例请求

```bash
# 列出目录内容
curl "http://localhost:8190/list?path=./"

# 创建目录
curl -X POST "http://localhost:8190/mkdir?path=./test"

# 创建文件
curl -X POST "http://localhost:8190/touch?path=./test/file.txt"

# 获取文件信息
curl "http://localhost:8190/info?path=./test/file.txt"

# 创建文档
curl -X POST "http://localhost:8190/document" \
  -H "Content-Type: application/json" \
  -d '{"path":"./test/doc.md","type":"md","content":"# Hello World"}'
```

## 开发说明

### 热更新配置

项目使用 Air 进行热更新，配置文件为 `.air.toml`。主要配置包括：

```toml
root = "."
tmp_dir = "tmp"

[build]
cmd = "go build -o ./tmp/main.exe ./cmd/server"
bin = "tmp/main.exe"
include_ext = ["go", "tpl", "tmpl", "html"]
exclude_dir = ["assets", "tmp", "vendor"]
delay = 1000
kill_delay = "0s"
log = "build-errors.log"
send_interrupt = false
stop_on_error = true

[log]
time = false

[color]
main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"
```

### 日志

日志文件位于 `logs` 目录下，按日期命名。包含以下级别的日志：
- INFO: 普通信息
- ERROR: 错误信息
- DEBUG: 调试信息

### 错误处理

项目使用统一的错误处理机制，所有错误都会被记录到日志中，并返回统一的错误响应格式：

```json
{
  "code": 1004,
  "message": "错误信息",
  "data": null
}
```

## 贡献指南

1. Fork 项目
2. 创建特性分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

## 许可证

MIT License 