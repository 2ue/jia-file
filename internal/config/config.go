package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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
		MaxSize          int64
		AllowedExtensions []string
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
			MaxSize          int64
			AllowedExtensions []string
		}{
			MaxSize:          10 * 1024 * 1024, // 10MB
			AllowedExtensions: []string{".txt", ".md", ".json", ".yaml", ".yml", ".go"},
		},
	}
)

// LoadConfig 加载配置文件
func LoadConfig(envPath string) (*Config, error) {
	// 如果未指定配置文件路径，使用默认路径
	if envPath == "" {
		envPath = ".env"
	}

	// 加载 .env 文件
	if err := godotenv.Load(envPath); err != nil {
		// 如果文件不存在，使用默认配置
		if os.IsNotExist(err) {
			return &defaultConfig, nil
		}
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	config := defaultConfig

	// 从环境变量加载配置
	if port := os.Getenv("SERVER_PORT"); port != "" {
		config.Server.Port = port
	}

	if level := os.Getenv("LOG_LEVEL"); level != "" {
		config.Log.Level = level
	}

	if dir := os.Getenv("LOG_DIR"); dir != "" {
		config.Log.Dir = dir
	}

	if maxSize := os.Getenv("MAX_FILE_SIZE"); maxSize != "" {
		if size, err := strconv.ParseInt(maxSize, 10, 64); err == nil {
			config.File.MaxSize = size
		}
	}

	if exts := os.Getenv("ALLOWED_EXTENSIONS"); exts != "" {
		config.File.AllowedExtensions = strings.Split(exts, ",")
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