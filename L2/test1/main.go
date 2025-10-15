package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// 注册路由
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		time.Sleep(2 * time.Second) // 模拟耗时
		fmt.Fprintf(w, "Hello, %s (耗时: %v)\n", r.RemoteAddr, time.Since(start))
	})

	// 启动 HTTP 服务器
	go func() {
		fmt.Println("服务器启动在 :8080")
		http.ListenAndServe(":8080", nil)
	}()

	// 稍等一下让服务器启动
	time.Sleep(500 * time.Millisecond)

	// 并发发起多个请求
	for i := 0; i < 3; i++ {
		go func(id int) {
			resp, err := http.Get("http://localhost:8080/hello")
			if err != nil {
				fmt.Println("请求出错：", err)
				return
			}
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("客户端 goroutine %d 收到响应: %s", id, body)
		}(i)
	}

	// 防止主协程提前退出
	time.Sleep(5 * time.Second)
}


// type Handler interface {
//     ServeHTTP(ResponseWriter, *Request)
// }

// type ServeMux struct {}
// func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
//     // ...
// }

// var h Handler = &ServeMux{}
// h.ServeHTTP(w, r) // 动态分派调用
