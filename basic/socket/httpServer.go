package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func main() {
	testServer1()
}
func testServer1() {
	// 这里没有指定server，所以是用http包中默认的server来处理客户端请求
	// 所以是用http包中默认的server来处理客户端请求	并没有构造Server对象，
	// 而是使用http提供的缺省Server对象，通过http.Handle指定了匹配"/foo"的请求的Handler，
	// 通过HandleFunc方法来指定匹配“/bar”的请求的处理函数。
	// 最后通过http.ListenAndServe方法启动缺省的Server对象，并在8080端口监听客户端请求。

	//Handler是一个接口类型，包含了ServerHTTP方法，该方法对客户端的请求（Request对象）进行处理，
	// 并通过ResponseWriter将响应信息传送回客户端。
	http.Handle("/foo", nil)

	//	我们调用了http.Handle方法和HandleFunc方法的时候其实是将"/foo"，
	//	”/bar"与相应的Handler的对应关系存到http.DefaultServeMux对象中。
	//	所以当Server对象的Handler属性为nil时，
	//	Server对象便可以通过http.DefaultServeMux对象来实现请求的处理逻辑。

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	//Server对象包含了多个属性，其中较为常用的两个属性是Addr和Handler属性。Addr属性即是Server对象监听的地址（端口号）
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func testServer2() {
	// 自定义server
	//通过创建Server对象对一些属性进行了指定。

	//func ListenAndServe(addr string, handler Handler) error {
	//	server := &Server{Addr: addr, Handler: handler}
	//	return server.ListenAndServe()
	//}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
	//Handler是一个接口类型，包含了ServerHTTP方法，该方法对客户端的请求（Request对象）进行处理，
	// 并通过ResponseWriter将响应信息传送回客户端。
	// 可以Handler是Server对象处理请求的逻辑的关键所在。
	// 但是上面ListenAndServe方法的Handler参数为nil，这时那么Server对象将使用http.DefaultServeMux来处理请求。

	//http.DefaultServeMux是一个ServeMux对象，
	//ServeMux对象是一个HTTP请求的分发器，将匹配了相应URL的请求分发给相应的处理函数进行处理。其定义如下：

	//type ServeMux struct {
	//	mu    sync.RWMutex
	//	m     map[string]muxEntry
	//	hosts bool // whether any patterns contain hostnames
	//}
	//
	//type muxEntry struct {
	//	h       Handler
	//	pattern string
	//}

	//ServeMux正是通过map结构（map[string]muxEntry）来匹配请求和Handler的，
	//起到了路由的作用。并且ServeMux本身也是一个Handler，
	//因为它实现了Handler接口的ServeHTTP方法。

}
