package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// IgnoreConfig 忽略配置
type IgnoreConfig struct {
	Paths      []string `json:"paths"`      // 忽略的路径
	Extensions []string `json:"extensions"` // 忽略的文件扩展名
	Patterns   []string `json:"patterns"`   // 忽略的文件模式
}

// LoadIgnoreConfig 加载忽略配置
func LoadIgnoreConfig(configPath string) (*IgnoreConfig, error) {
	// 如果未指定配置文件路径，使用默认路径
	if configPath == "" {
		configPath = filepath.Join("internal", "config", "ignore.json")
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	// 解析配置
	var config IgnoreConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
} 