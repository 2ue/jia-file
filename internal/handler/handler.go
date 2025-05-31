package handler

import (
	"encoding/json"
	"jia-file/internal/file"
	"net/http"
	"path/filepath"
)

// Response 统一响应格式
type Response struct {
	Code    int         `json:"code"`    // 状态码：0表示成功，非0表示失败
	Message string      `json:"message"` // 状态描述
	Data    interface{} `json:"data"`    // 响应数据
}

// 状态码定义
const (
	CodeSuccess        = 0    // 成功
	CodeParamMissing   = 1001 // 参数缺失
	CodeMethodNotAllow = 1002 // 方法不允许
	CodePathNotExist   = 1003 // 路径不存在
	CodeOperationFail  = 1004 // 操作失败
)

// Handler HTTP处理器
type Handler struct {
	fileService file.Service
}

// NewHandler 创建HTTP处理器实例
func NewHandler(fileService file.Service) *Handler {
	return &Handler{
		fileService: fileService,
	}
}

// writeResponse 写入统一格式的响应
func (h *Handler) writeResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	json.NewEncoder(w).Encode(response)
}

// List 列出目录内容
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	if path == "" {
		h.writeResponse(w, CodeParamMissing, "Missing path parameter", nil)
		return
	}

	files, err := h.fileService.List(path)
	if err != nil {
		h.writeResponse(w, CodeOperationFail, err.Error(), nil)
		return
	}

	h.writeResponse(w, CodeSuccess, "success", files)
}

// CreateDir 创建目录
func (h *Handler) CreateDir(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.writeResponse(w, CodeMethodNotAllow, "Method not allowed", nil)
		return
	}

	path := r.URL.Query().Get("path")
	if path == "" {
		h.writeResponse(w, CodeParamMissing, "Missing path parameter", nil)
		return
	}

	if err := h.fileService.CreateDir(path); err != nil {
		h.writeResponse(w, CodeOperationFail, err.Error(), nil)
		return
	}

	h.writeResponse(w, CodeSuccess, "Directory created successfully", nil)
}

// CreateFile 创建文件
func (h *Handler) CreateFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.writeResponse(w, CodeMethodNotAllow, "Method not allowed", nil)
		return
	}

	path := r.URL.Query().Get("path")
	if path == "" {
		h.writeResponse(w, CodeParamMissing, "Missing path parameter", nil)
		return
	}

	content := r.Body
	defer content.Close()

	if err := h.fileService.CreateFile(path, nil); err != nil {
		h.writeResponse(w, CodeOperationFail, err.Error(), nil)
		return
	}

	h.writeResponse(w, CodeSuccess, "File created successfully", nil)
}

// Delete 删除文件或目录
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		h.writeResponse(w, CodeMethodNotAllow, "Method not allowed", nil)
		return
	}

	path := r.URL.Query().Get("path")
	if path == "" {
		h.writeResponse(w, CodeParamMissing, "Missing path parameter", nil)
		return
	}

	if err := h.fileService.Delete(path); err != nil {
		h.writeResponse(w, CodeOperationFail, err.Error(), nil)
		return
	}

	h.writeResponse(w, CodeSuccess, "File or directory deleted successfully", nil)
}

// Move 移动文件或目录
func (h *Handler) Move(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.writeResponse(w, CodeMethodNotAllow, "Method not allowed", nil)
		return
	}

	src := r.URL.Query().Get("src")
	dst := r.URL.Query().Get("dst")
	if src == "" || dst == "" {
		h.writeResponse(w, CodeParamMissing, "Missing src or dst parameter", nil)
		return
	}

	if err := h.fileService.Move(src, dst); err != nil {
		h.writeResponse(w, CodeOperationFail, err.Error(), nil)
		return
	}

	h.writeResponse(w, CodeSuccess, "File or directory moved successfully", nil)
}

// Copy 复制文件或目录
func (h *Handler) Copy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.writeResponse(w, CodeMethodNotAllow, "Method not allowed", nil)
		return
	}

	src := r.URL.Query().Get("src")
	dst := r.URL.Query().Get("dst")
	if src == "" || dst == "" {
		h.writeResponse(w, CodeParamMissing, "Missing src or dst parameter", nil)
		return
	}

	if err := h.fileService.Copy(src, dst); err != nil {
		h.writeResponse(w, CodeOperationFail, err.Error(), nil)
		return
	}

	h.writeResponse(w, CodeSuccess, "File or directory copied successfully", nil)
}

// GetInfo 获取文件信息
func (h *Handler) GetInfo(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	if path == "" {
		h.writeResponse(w, CodeParamMissing, "Missing path parameter", nil)
		return
	}

	info, err := h.fileService.GetInfo(path)
	if err != nil {
		println(err.Error());
		h.writeResponse(w, CodeOperationFail, err.Error(), nil)
		return
	}

	h.writeResponse(w, CodeSuccess, "success", info)
}

// CreateDocumentRequest 创建文档请求
type CreateDocumentRequest struct {
	Path    string `json:"path"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

// CreateDocument 创建文档处理函数
func (h *Handler) CreateDocument(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.writeResponse(w, CodeMethodNotAllow, "Method not allowed", nil)
		return
	}

	var req CreateDocumentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeResponse(w, CodeParamMissing, "Invalid request body", nil)
		return
	}

	// 验证请求参数
	if req.Path == "" || req.Type == "" {
		h.writeResponse(w, CodeParamMissing, "Path and type are required", nil)
		return
	}

	// 确保文件扩展名正确
	ext := filepath.Ext(req.Path)
	if ext == "" {
		req.Path = req.Path + "." + req.Type
	}

	// 创建文档
	if err := h.fileService.CreateDocument(req.Path, req.Type, req.Content); err != nil {
		h.writeResponse(w, CodeOperationFail, err.Error(), nil)
		return
	}

	h.writeResponse(w, CodeSuccess, "Document created successfully", nil)
} 