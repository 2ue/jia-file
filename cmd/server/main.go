package main

import (
	"fmt"
	"jia-file/internal/file"
	"jia-file/internal/handler"
	"jia-file/internal/logger"
	"jia-file/internal/middleware"
	"log"
	"net/http"
)

func main() {
	// 初始化日志
	if err := logger.Init("logs"); err != nil {
		log.Fatal(err)
	}

	// 创建文件服务实例
	fileService := file.NewService()

	// 创建HTTP处理器实例
	h := handler.NewHandler(fileService)

	// 创建路由
	mux := http.NewServeMux()

	// 注册路由
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "Welcome to Go Web Server!")
	})

	// 文件操作路由
	mux.HandleFunc("/list", h.List)
	mux.HandleFunc("/mkdir", h.CreateDir)
	mux.HandleFunc("/touch", h.CreateFile)
	mux.HandleFunc("/delete", h.Delete)
	mux.HandleFunc("/move", h.Move)
	mux.HandleFunc("/copy", h.Copy)
	mux.HandleFunc("/info", h.GetInfo)
	mux.HandleFunc("/document", h.CreateDocument)

	// 应用中间件
	handler := middleware.LoggingMiddleware(
		middleware.RecoveryMiddleware(
			middleware.CORSMiddleware(
				middleware.PathValidationMiddleware(mux),
			),
		),
	)

	// 启动服务器
	port := ":8190"
	logger.Info("Server starting on %s...", port)
	if err := http.ListenAndServe(port, handler); err != nil {
		logger.Error("Server error: %v", err)
		log.Fatal(err)
	}
} 