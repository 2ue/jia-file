package main

import (
	"fmt"
	"jia-file/internal/file"
	"jia-file/internal/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	// 获取当前工作目录作为基础路径
	basePath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// 创建文件服务实例
	fileService := file.NewService(basePath)

	// 创建HTTP处理器实例
	h := handler.NewHandler(fileService)

	// 注册路由
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "Welcome to Go Web Server!")
	})

	// 文件操作路由
	http.HandleFunc("/list", h.List)
	http.HandleFunc("/mkdir", h.CreateDir)
	http.HandleFunc("/touch", h.CreateFile)
	http.HandleFunc("/delete", h.Delete)
	http.HandleFunc("/move", h.Move)
	http.HandleFunc("/copy", h.Copy)
	http.HandleFunc("/info", h.GetInfo)
	http.HandleFunc("/document", h.CreateDocument)

	// 启动服务器
	port := ":8190"
	fmt.Printf("Server starting on %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
} 