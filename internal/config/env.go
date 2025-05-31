package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// GetEnv 获取环境变量，如果不存在则返回默认值
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetEnvInt 获取整数类型环境变量，如果不存在或转换失败则返回默认值
func GetEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

// GetEnvInt64 获取int64类型环境变量，如果不存在或转换失败则返回默认值
func GetEnvInt64(key string, defaultValue int64) int64 {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defaultValue
	}
	return intValue
}

// GetEnvBool 获取布尔类型环境变量，如果不存在或转换失败则返回默认值
func GetEnvBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return boolValue
}

// GetEnvStringSlice 获取字符串切片类型环境变量，如果不存在则返回默认值
func GetEnvStringSlice(key string, defaultValue []string) []string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return strings.Split(value, ",")
}

// LoadEnv 加载环境变量文件
func LoadEnv(filenames ...string) error {
	// 如果没有指定文件，使用默认的 .env
	if len(filenames) == 0 {
		filenames = []string{".env"}
	}

	// 尝试加载所有指定的文件
	for _, filename := range filenames {
		if err := godotenv.Load(filename); err != nil {
			// 如果文件不存在，继续尝试下一个
			if os.IsNotExist(err) {
				continue
			}
			return fmt.Errorf("error loading %s: %v", filename, err)
		}
	}

	return nil
} 