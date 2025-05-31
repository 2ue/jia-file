package file

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

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

// Service 文件服务接口
type Service interface {
	// List 列出指定目录下的文件和文件夹
	List(path string) ([]FileInfo, error)
	// CreateDir 创建目录
	CreateDir(path string) error
	// CreateFile 创建文件
	CreateFile(path string, content []byte) error
	// Delete 删除文件或目录
	Delete(path string) error
	// Move 移动文件或目录
	Move(src, dst string) error
	// Copy 复制文件或目录
	Copy(src, dst string) error
	// GetInfo 获取文件信息
	GetInfo(path string) (FileInfo, error)
	// CreateDocument 创建文档文件
	CreateDocument(path string, docType string, content string) error
}

// service 文件服务实现
type service struct{}

// NewService 创建文件服务实例
func NewService() Service {
	return &service{}
}

// formatFileSize 将文件大小转换为人类可读的格式
func formatFileSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

// detectMimeType 检测文件的MIME类型
func detectMimeType(path string, isDir bool) string {
	if isDir {
		return "inode/directory"
	}

	ext := filepath.Ext(path)
	if mimeType := mime.TypeByExtension(ext); mimeType != "" {
		return mimeType
	}

	file, err := os.Open(path)
	if err != nil {
		return "application/octet-stream"
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		return "application/octet-stream"
	}

	mimeType := http.DetectContentType(buffer)
	
	if mimeType == "application/octet-stream" {
		switch {
		case strings.HasPrefix(path, "."):
			return "text/plain"
		case strings.Contains(strings.ToLower(path), "readme"):
			return "text/plain"
		case strings.Contains(strings.ToLower(path), "license"):
			return "text/plain"
		case strings.Contains(strings.ToLower(path), "makefile"):
			return "text/plain"
		case strings.Contains(strings.ToLower(path), ".git"):
			return "application/x-git"
		}
	}

	return mimeType
}

// getFileInfo 获取单个文件的详细信息
func getFileInfo(entry os.DirEntry, path string) (FileInfo, error) {
	info, err := entry.Info()
	if err != nil {
		return FileInfo{}, err
	}

	fullPath := filepath.Join(path, entry.Name())
	ext := filepath.Ext(entry.Name())
	
	mimeType := detectMimeType(fullPath, entry.IsDir())

	isSymlink := info.Mode()&os.ModeSymlink != 0
	symlinkTarget := ""
	if isSymlink {
		if target, err := os.Readlink(fullPath); err == nil {
			symlinkTarget = target
		}
	}

	stat, err := os.Stat(fullPath)
	createTime := time.Time{}
	accessTime := time.Time{}
	if err == nil {
		createTime = stat.ModTime()
		accessTime = stat.ModTime()
	}

	return FileInfo{
		Name:          entry.Name(),
		IsDir:         entry.IsDir(),
		Size:          info.Size(),
		SizeHuman:     formatFileSize(info.Size()),
		Path:          fullPath,
		Ext:           ext,
		MimeType:      mimeType,
		CreateTime:    createTime,
		ModTime:       info.ModTime(),
		AccessTime:    accessTime,
		Mode:          info.Mode().String(),
		IsHidden:      strings.HasPrefix(entry.Name(), "."),
		IsSymlink:     isSymlink,
		SymlinkTarget: symlinkTarget,
	}, nil
}

// List 实现 Service 接口的 List 方法
func (s *service) List(path string) ([]FileInfo, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("directory does not exist: %s", path)
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %v", err)
	}

	files := make([]FileInfo, 0, len(entries))
	for _, entry := range entries {
		if fileInfo, err := getFileInfo(entry, path); err == nil {
			files = append(files, fileInfo)
		}
	}

	return files, nil
}

// CreateDir 实现 Service 接口的 CreateDir 方法
func (s *service) CreateDir(path string) error {
	return os.MkdirAll(path, 0755)
}

// CreateFile 实现 Service 接口的 CreateFile 方法
func (s *service) CreateFile(path string, content []byte) error {
	// 检查文件是否已存在
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("file already exists: %s", path)
	}

	// 确保父目录存在
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create parent directory: %v", err)
	}

	return os.WriteFile(path, content, 0644)
}

// Delete 实现 Service 接口的 Delete 方法
func (s *service) Delete(path string) error {
	// 检查文件是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("file or directory does not exist: %s", path)
	}

	return os.RemoveAll(path)
}

// Move 实现 Service 接口的 Move 方法
func (s *service) Move(src, dst string) error {
	return os.Rename(src, dst)
}

// Copy 实现 Service 接口的 Copy 方法
func (s *service) Copy(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

// GetInfo 实现 Service 接口的 GetInfo 方法
func (s *service) GetInfo(path string) (FileInfo, error) {
	info, err := os.Stat(path)
	if err != nil {
		return FileInfo{}, err
	}

	return FileInfo{
		Name:          info.Name(),
		IsDir:         info.IsDir(),
		Size:          info.Size(),
		SizeHuman:     formatFileSize(info.Size()),
		Path:          path,
		Ext:           filepath.Ext(info.Name()),
		MimeType:      detectMimeType(path, info.IsDir()),
		CreateTime:    info.ModTime(),
		ModTime:       info.ModTime(),
		AccessTime:    info.ModTime(),
		Mode:          info.Mode().String(),
		IsHidden:      strings.HasPrefix(info.Name(), "."),
		IsSymlink:     info.Mode()&os.ModeSymlink != 0,
		SymlinkTarget: "",
	}, nil
}

// CreateDocument 实现 Service 接口的 CreateDocument 方法
func (s *service) CreateDocument(path string, docType string, content string) error {
	// 确保目录存在
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("创建目录失败: %v", err)
	}

	// 创建空文件
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	return file.Close()
} 