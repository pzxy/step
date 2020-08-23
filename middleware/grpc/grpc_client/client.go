package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"step/go_grpc/proto"
	"time"
)

const PORT = "9001"

func main() {
	//有必要的话加上ip：PORT
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err:%v", err)
	}
	defer conn.Close()
	client := proto.NewStreamServiceClient(conn)
	fmt.Printf("grpc客户端启动：%v \n", time.Now().Format("2006-01-02 15:04:05"))

	err = printLists(client, &proto.StreamRequest{Pt: &proto.StreamPoint{Name: "gRPC Stream Client: List", Value: 2020}})
	if err != nil {
		log.Fatalf("printLists.err: %v", err)
	}

	err = printRecord(client, &proto.StreamRequest{Pt: &proto.StreamPoint{Name: "gRPC Stream Client: Record", Value: 2020}})
	if err != nil {
		log.Fatalf("printRecord.err: %v", err)
	}

	err = printRoute(client, &proto.StreamRequest{Pt: &proto.StreamPoint{Name: "gRPC Stream Client: Route", Value: 2020}})
	if err != nil {
		log.Fatalf("printRoute.err: %v", err)
	}

	err = printCallDemo(client, &proto.CallRequest{CallString: "客户端向服务端函数调用"})
	if err != nil {
		log.Fatalf("printRoute.err: %v", err)
	}
}

//1. 接受来自服务器的流   rpc List(StreamRequest) returns (stream StreamResponse) {};
func printLists(client proto.StreamServiceClient, r *proto.StreamRequest) error {
	stream, err := client.List(context.Background(), r)
	if err != nil {
		return err
	}
	for {
		resp, err := stream.Recv()
		/** Recv
		RecvMsg 会从流中读取完整的 gRPC 消息体，另外通过阅读源码可得知：
		（1）RecvMsg 是阻塞等待的
		（2）RecvMsg 当流成功/结束（调用了 Close）时，会返回 io.EOF
		（3）RecvMsg 当流出现任何错误时，流会被中止，错误信息会包含 RPC 错误码。而在 RecvMsg 中可能出现如下错误：
		1. io.EOF
		2. io.ErrUnexpectedEOF
		3. transport.ConnectionError
		4. google.golang.org/grpc/codes
		同时需要注意，默认的 MaxReceiveMessageSize 值为 1024  1024  4，建议不要超出
		*/
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)
	}

	return nil
}

//2. 服务器流    rpc Record(stream StreamRequest) returns (StreamResponse) {};
func printRecord(client proto.StreamServiceClient, r *proto.StreamRequest) error {
	stream, err := client.Record(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n < 6; n++ {
		err := stream.Send(r)
		if err != nil {
			return err
		}
	}

	resp, err := stream.CloseAndRecv()
	/**
	stream.CloseAndRecv 和 stream.SendAndClose 是配套使用的流方法，相信聪明的你已经秒懂它的作用了
	*/
	if err != nil {
		return err
	}

	log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)

	return nil
}

//3 双向流     rpc Route(stream StreamRequest) returns (stream StreamResponse) {};
func printRoute(client proto.StreamServiceClient, r *proto.StreamRequest) error {
	stream, err := client.Route(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n <= 6; n++ {
		err = stream.Send(r)
		if err != nil {
			return err
		}

		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)
	}

	stream.CloseSend()

	return nil
}

//4 函数调用
func printCallDemo(client proto.StreamServiceClient, r *proto.CallRequest) error {
	rsp, err := client.CallDemo(context.Background(), r)
	fmt.Printf("调用服务器函数后返回:%v", rsp.CallString)
	if err != nil {
		return err
	}
	return nil
}
