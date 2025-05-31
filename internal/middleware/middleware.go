package middleware

import (
	"encoding/json"
	"jia-file/api"
	"jia-file/internal/logger"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

// LoggingMiddleware 日志中间件
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// 调用下一个处理器
		next.ServeHTTP(w, r)

		// 记录请求日志
		logger.Info("%s %s %s %v",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}

// RecoveryMiddleware 错误恢复中间件
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("Panic recovered: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// CORSMiddleware CORS中间件
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// MethodMiddleware 方法限制中间件
func MethodMiddleware(methods ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, method := range methods {
				if r.Method == method {
					next.ServeHTTP(w, r)
					return
				}
			}
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		})
	}
}

// PathValidationMiddleware 路径验证中间件
func PathValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 检查路径参数
		path := r.URL.Query().Get("path")
		if path != "" && !isValidPath(path) {
			w.Header().Set("Content-Type", "application/json")
			response := api.Response{
				Code:    api.CodeParamMissing,
				Message: "Path must be an absolute path",
				Data:    nil,
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		// 检查源路径和目标路径
		src := r.URL.Query().Get("src")
		dst := r.URL.Query().Get("dst")
		if (src != "" && !isValidPath(src)) || (dst != "" && !isValidPath(dst)) {
			w.Header().Set("Content-Type", "application/json")
			response := api.Response{
				Code:    api.CodeParamMissing,
				Message: "Source and destination paths must be absolute paths",
				Data:    nil,
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// isValidPath 验证路径是否为绝对路径
func isValidPath(path string) bool {
	// 检查是否为绝对路径
	if !filepath.IsAbs(path) {
		return false
	}

	// 检查是否包含 .. 或 .
	if strings.Contains(path, "..") || strings.Contains(path, "./") {
		return false
	}

	return true
} 