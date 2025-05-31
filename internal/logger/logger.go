package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	// InfoLogger 信息日志记录器
	InfoLogger *log.Logger
	// ErrorLogger 错误日志记录器
	ErrorLogger *log.Logger
	// DebugLogger 调试日志记录器
	DebugLogger *log.Logger
)

// Init 初始化日志记录器
func Init(logDir string) error {
	// 确保日志目录存在
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return fmt.Errorf("创建日志目录失败: %v", err)
	}

	// 创建日志文件
	logFile := filepath.Join(logDir, fmt.Sprintf("app-%s.log", time.Now().Format("2006-01-02")))
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("打开日志文件失败: %v", err)
	}

	// 初始化日志记录器
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	DebugLogger = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}

// Info 记录信息日志
func Info(format string, v ...interface{}) {
	if InfoLogger != nil {
		InfoLogger.Printf(format, v...)
	}
}

// Error 记录错误日志
func Error(format string, v ...interface{}) {
	if ErrorLogger != nil {
		ErrorLogger.Printf(format, v...)
	}
}

// Debug 记录调试日志
func Debug(format string, v ...interface{}) {
	if DebugLogger != nil {
		DebugLogger.Printf(format, v...)
	}
} 