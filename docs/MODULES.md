# Jia-File 模块文档

## 目录结构

```
jia-file/
├── api/            # API 类型定义和常量
├── cmd/            # 主程序入口
├── internal/       # 内部实现
│   ├── file/      # 文件操作
│   ├── handler/   # HTTP 处理器
│   ├── logger/    # 日志模块
│   ├── middleware/# 中间件
│   └── errors/    # 错误处理
├── scripts/       # 工具脚本
└── docs/          # 文档
```

## 模块说明

### api 模块

包含 API 相关的类型定义和常量。

- `types.go`: 定义 API 请求和响应的数据结构
- 所有路径相关的参数都必须是绝对路径

### cmd 模块

主程序入口点。

- `server/main.go`: 服务器启动代码
- 负责初始化各个模块并启动 HTTP 服务器

### internal 模块

核心功能实现。

#### file 模块

文件操作相关功能。

- 所有文件操作都使用绝对路径
- 提供以下功能：
  - 列出目录内容
  - 创建目录
  - 创建文件
  - 删除文件或目录
  - 移动文件或目录
  - 复制文件或目录
  - 获取文件信息
  - 创建文档

#### handler 模块

HTTP 请求处理。

- 参数验证
- 请求处理
- 响应格式化
- 错误处理

#### logger 模块

日志功能。

- 多级别日志支持
- 日志文件管理
- 日志格式化

#### middleware 模块

HTTP 中间件。

- 请求日志记录
- 错误恢复
- CORS 支持
- 路径验证（确保使用绝对路径）

#### errors 模块

错误处理。

- 统一错误类型
- 错误码定义
- 错误信息格式化

### scripts 模块

工具脚本。

- 开发环境设置
- 部署脚本
- 测试脚本

### docs 模块

项目文档。

- `API.md`: API 文档
- `CHANGELOG.md`: 更新日志
- `MODULES.md`: 模块文档

## 模块依赖

```
cmd/server
  ├── internal/file
  ├── internal/handler
  ├── internal/logger
  ├── internal/middleware
  └── internal/errors
```

## 开发指南

### 添加新功能

1. 在 `api` 模块中定义新的类型和常量
2. 在 `internal/file` 中实现核心功能
3. 在 `internal/handler` 中添加 HTTP 处理逻辑
4. 在 `internal/middleware` 中添加必要的中间件
5. 更新相关文档

### 修改现有功能

1. 确保所有路径参数都使用绝对路径
2. 在 `internal/middleware` 中添加路径验证
3. 更新相关文档和测试

### 添加新中间件

1. 在 `internal/middleware` 中创建新的中间件函数
2. 在 `main.go` 中注册中间件
3. 更新相关文档

## 安全注意事项

1. 所有文件操作必须使用绝对路径
2. 路径中不允许包含 `..` 或 `.` 等特殊字符
3. 所有文件操作都需要进行权限检查
4. 建议在生产环境中启用 HTTPS 