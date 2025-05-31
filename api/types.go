package api

import (
	"time"
)

// Response 统一响应格式
type Response struct {
	Code    int         `json:"code"`    // 状态码：0表示成功，非0表示失败
	Message string      `json:"message"` // 状态描述
	Data    interface{} `json:"data"`    // 响应数据
}

// FileInfo 文件信息结构
type FileInfo struct {
	Name          string    `json:"name"`          // 文件名
	IsDir         bool      `json:"isDir"`         // 是否为目录
	Size          int64     `json:"size"`          // 文件大小（字节）
	SizeHuman     string    `json:"sizeHuman"`     // 人类可读的文件大小
	Path          string    `json:"path"`          // 完整路径
	Ext           string    `json:"ext"`           // 文件扩展名
	MimeType      string    `json:"mimeType"`      // MIME类型
	CreateTime    time.Time `json:"createTime"`    // 创建时间
	ModTime       time.Time `json:"modTime"`       // 修改时间
	AccessTime    time.Time `json:"accessTime"`    // 访问时间
	Mode          string    `json:"mode"`          // 文件权限
	IsHidden      bool      `json:"isHidden"`      // 是否为隐藏文件
	IsSymlink     bool      `json:"isSymlink"`     // 是否为符号链接
	SymlinkTarget string    `json:"symlinkTarget"` // 符号链接目标
}

// CreateDocumentRequest 创建文档请求
type CreateDocumentRequest struct {
	Path    string `json:"path"`    // 文档路径
	Type    string `json:"type"`    // 文档类型
	Content string `json:"content"` // 文档内容
}

// 状态码定义
const (
	CodeSuccess        = 0    // 成功
	CodeParamMissing   = 1001 // 参数缺失
	CodeMethodNotAllow = 1002 // 方法不允许
	CodePathNotExist   = 1003 // 路径不存在
	CodeOperationFail  = 1004 // 操作失败
) 