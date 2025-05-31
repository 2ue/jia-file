package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Config 应用配置结构
type Config struct {
	Server struct {
		Port string
	}
	Log struct {
		Level string
		Dir   string
	}
	File struct {
		RootPath string // 文件操作的根目录
	}
}

var (
	// 默认配置
	defaultConfig = Config{
		Server: struct {
			Port string
		}{
			Port: "8190",
		},
		Log: struct {
			Level string
			Dir   string
		}{
			Level: "info",
			Dir:   "logs",
		},
		File: struct {
			RootPath string
		}{
			RootPath: "", // 默认为空，表示不限制根目录
		},
	}
)

// LoadConfig 加载配置文件
func LoadConfig(envPath string) (*Config, error) {
	if envPath == "" {
		envPath = ".env"
	}
	if err := godotenv.Load(envPath); err != nil {
		if os.IsNotExist(err) {
			return &defaultConfig, nil
		}
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}
	config := defaultConfig
	if port := os.Getenv("PORT"); port != "" {
		config.Server.Port = port
	}
	if level := os.Getenv("LOG_LEVEL"); level != "" {
		config.Log.Level = level
	}
	if dir := os.Getenv("LOG_DIR"); dir != "" {
		config.Log.Dir = dir
	}
	if rootPath := os.Getenv("ROOT_PATH"); rootPath != "" {
		config.File.RootPath = rootPath
	}
	return &config, nil
}

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