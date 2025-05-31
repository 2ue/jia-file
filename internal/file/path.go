package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// PathProcessor 路径处理器
type PathProcessor struct {
	rootPath string
}

// NewPathProcessor 创建路径处理器
func NewPathProcessor(rootPath string) *PathProcessor {
	return &PathProcessor{
		rootPath: rootPath,
	}
}

// ProcessPath 处理路径
// 如果设置了rootPath：
//   - 对于相对路径，将其与rootPath拼接
//   - 对于绝对路径，验证是否在rootPath下
// 如果未设置rootPath：
//   - 直接返回传入的路径
func (p *PathProcessor) ProcessPath(path string) (string, error) {
	// 如果未设置rootPath，直接返回原路径
	if p.rootPath == "" {
		return path, nil
	}

	// 确保rootPath是绝对路径
	rootPath, err := filepath.Abs(p.rootPath)
	if err != nil {
		return "", fmt.Errorf("invalid root path: %v", err)
	}

	// 如果path是绝对路径
	if filepath.IsAbs(path) {
		// 检查path是否在rootPath下
		rel, err := filepath.Rel(rootPath, path)
		if err != nil {
			return "", fmt.Errorf("path is outside root directory: %v", err)
		}
		if strings.HasPrefix(rel, "..") {
			return "", fmt.Errorf("path is outside root directory")
		}
		return path, nil
	}

	// 如果path是相对路径，与rootPath拼接
	return filepath.Join(rootPath, path), nil
}

// ValidatePath 验证路径是否有效
func (p *PathProcessor) ValidatePath(path string) error {
	processedPath, err := p.ProcessPath(path)
	if err != nil {
		return err
	}

	// 检查路径是否存在
	if _, err := os.Stat(processedPath); err != nil {
		if os.IsNotExist(err) {
			return nil // 路径不存在是允许的，因为可能是新建文件/目录
		}
		return fmt.Errorf("error checking path: %v", err)
	}

	return nil
} 