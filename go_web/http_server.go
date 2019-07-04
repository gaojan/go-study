package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

/*
http server包执行流程
1、创建Listen Socket,监听指定的端口，等待客户端请求到来。
2、Listen Socket接受客户端的请求，得到Client Socket，接下来通过Client Socket与客户端通信。
3、处理客户端的请求，首先从Client Socket读取HTTP请求的协议头，如果POST方法，还可能要读取客户端
提交的数据，然后交给相应的handler处理请求，handler处理完毕准备好客户端需要的数据，通过Client Socket写给客户端。

三个问题： 1、如何监听端口。 2、如何接收客户端请求。 3、如何分配handler
1、ListenAndServe监听端口。2、接收用户请求并创建Conn  srv.Serve(I net.Listener)  3、处理连接 *Conn.serve{}
*/

// http包建立web服务器
func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                //解析参数，默认是不会解析的
	fmt.Println("Form:", r.Form) // 这些信息是输出到服务器端端打印信息，比如记录日志什么的
	fmt.Println("path:", r.URL.Scheme)
	fmt.Println("url_long:", r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello astaxie!") // 这个写入到w也就是response的是输出到客户端的信息
}

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello world!")
}

//　Web服务内部有支持高并发的特性，这跟tornado很相似
func main() {
	// http://localhost:9090/?url_long=111&url_long=222&url_long=998
	http.HandleFunc("/", sayHelloName) // 设置访问的路由
	http.HandleFunc("/hello", sayHelloWorld)
	err := http.ListenAndServe("127.0.0.1:9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
