package main

func main() {
	//rpc.Register(rpcw.DemoService{}) //将这个结构体注册进去
	////监听端口
	//listener, err := net.Listen("tcp", ":1234")
	//if err != nil {
	//	panic(err)
	//}
	//for { //循环等待接受
	//	conn, err := listener.Accept()
	//	if err != nil {
	//		log.Printf("accept error %v", err)
	//		continue
	//	}
	//	//若是链接成功的话 goroutine 防止阻塞
	//	go jsonrpc.ServeConn(conn)
	//}

}
