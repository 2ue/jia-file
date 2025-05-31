# Jia-File API 文档

## 基本信息

- 基础URL: `http://localhost:8190`
- 所有响应格式均为JSON
- 所有路径参数必须是绝对路径，不支持相对路径

## 响应格式

所有API响应都遵循以下格式：

```json
{
    "code": 0,       // 状态码，0表示成功
    "message": "",   // 状态描述
    "data": null     // 响应数据
}
```

## 状态码

- 0: 成功
- 400: 请求参数错误
- 401: 未授权
- 403: 禁止访问
- 404: 资源不存在
- 500: 服务器内部错误

## API 端点

### 1. 列出目录内容

- **URL**: `/api/files/list`
- **方法**: `GET`
- **参数**:
  - `path`: 要列出内容的目录的绝对路径
- **响应**:
```json
{
    "code": 0,
    "message": "success",
    "data": [
        {
            "name": "example.txt",
            "isDir": false,
            "size": 1024,
            "sizeHuman": "1 KB",
            "path": "/absolute/path/to/example.txt",
            "ext": ".txt",
            "mimeType": "text/plain",
            "createTime": "2024-01-01T00:00:00Z",
            "modTime": "2024-01-01T00:00:00Z",
            "accessTime": "2024-01-01T00:00:00Z",
            "mode": "-rw-r--r--",
            "isHidden": false,
            "isSymlink": false,
            "symlinkTarget": ""
        }
    ]
}
```

### 2. 创建目录

- **URL**: `/api/files/mkdir`
- **方法**: `POST`
- **参数**:
  - `path`: 要创建的目录的绝对路径
- **响应**:
```json
{
    "code": 0,
    "message": "success",
    "data": null
}
```

### 3. 创建文件

- **URL**: `/api/files/create`
- **方法**: `POST`
- **参数**:
  - `path`: 要创建的文件的绝对路径
  - `content`: 文件内容
- **响应**:
```json
{
    "code": 0,
    "message": "success",
    "data": null
}
```

### 4. 删除文件或目录

- **URL**: `/api/files/delete`
- **方法**: `DELETE`
- **参数**:
  - `path`: 要删除的文件或目录的绝对路径
- **响应**:
```json
{
    "code": 0,
    "message": "success",
    "data": null
}
```

### 5. 移动文件或目录

- **URL**: `/api/files/move`
- **方法**: `POST`
- **参数**:
  - `src`: 源文件或目录的绝对路径
  - `dst`: 目标位置的绝对路径
- **响应**:
```json
{
    "code": 0,
    "message": "success",
    "data": null
}
```

### 6. 复制文件或目录

- **URL**: `/api/files/copy`
- **方法**: `POST`
- **参数**:
  - `src`: 源文件或目录的绝对路径
  - `dst`: 目标位置的绝对路径
- **响应**:
```json
{
    "code": 0,
    "message": "success",
    "data": null
}
```

### 7. 获取文件信息

- **URL**: `/api/files/info`
- **方法**: `GET`
- **参数**:
  - `path`: 要获取信息的文件或目录的绝对路径
- **响应**:
```json
{
    "code": 0,
    "message": "success",
    "data": {
        "name": "example.txt",
        "isDir": false,
        "size": 1024,
        "sizeHuman": "1 KB",
        "path": "/absolute/path/to/example.txt",
        "ext": ".txt",
        "mimeType": "text/plain",
        "createTime": "2024-01-01T00:00:00Z",
        "modTime": "2024-01-01T00:00:00Z",
        "accessTime": "2024-01-01T00:00:00Z",
        "mode": "-rw-r--r--",
        "isHidden": false,
        "isSymlink": false,
        "symlinkTarget": ""
    }
}
```

### 8. 创建文档

- **URL**: `/api/files/document`
- **方法**: `POST`
- **参数**:
  - `path`: 要创建的文档的绝对路径
  - `type`: 文档类型（如：txt, md, json等）
  - `content`: 文档内容
- **响应**:
```json
{
    "code": 0,
    "message": "success",
    "data": null
}
```

## 错误处理

当发生错误时，API会返回相应的错误码和错误信息：

```json
{
    "code": 400,
    "message": "Invalid path parameter: path must be absolute",
    "data": null
}
```

## 安全说明

1. 所有路径参数必须是绝对路径，不支持相对路径
2. 路径中不允许包含 `..` 或 `.` 等特殊字符
3. 所有文件操作都会进行权限检查
4. 建议在生产环境中启用HTTPS 