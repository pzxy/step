package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"step/middleware/grpc/proto"
	"time"

	"google.golang.org/grpc"
)

type StreamService struct{}

const PORT = "9001"

func main() {
	server := grpc.NewServer()
	s := new(StreamService)
	proto.RegisterStreamServiceServer(server, s)

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	go func() {
		for {
			select {
			case <-time.Tick(time.Second * 10):
				fmt.Printf("grpc服务器启动：%v \n", time.Now().Format("2006-01-02 15:04:05"))
			}
		}
	}()
	server.Serve(lis)
}

//1. 服务器端要实现定义的方法 这个服务器流 rpc List(StreamRequest) returns (stream StreamResponse) {};
func (s *StreamService) List(r *proto.StreamRequest, stream proto.StreamService_ListServer) error {
	for n := 0; n <= 6; n++ {
		err := stream.Send(&proto.StreamResponse{
			/**send
			通过阅读源码，可得知是 protoc 在生成时，根据定义生成了各式各样符合标准的接口方法。
			最终再统一调度内部的 SendMsg 方法，该方法涉及以下过程:
			1.消息体（对象）序列化
			2.压缩序列化后的消息体
			3.对正在传输的消息体增加 5 个字节的 header
			4.判断压缩+序列化后的消息体总字节长度是否大于预设的 maxSendMessageSize（预设值为 math.MaxInt32），若超出则提示错误
			5.写入给流的数据集
			*/
			Pt: &proto.StreamPoint{
				Name:  r.Pt.Name,
				Value: r.Pt.Value + int32(n),
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

//2. 客户端流式     rpc Record(stream StreamRequest) returns (StreamResponse) {};
func (s *StreamService) Record(stream proto.StreamService_RecordServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			/**
			SendAndClose
			在这段程序中，我们对每一个 Recv 都进行了处理，当发现 io.EOF (流关闭) 后，
			需要将最终的响应结果发送给客户端，同时关闭正在另外一侧等待的 Recv
			*/
			return stream.SendAndClose(&proto.StreamResponse{Pt: &proto.StreamPoint{Name: "gRPC Stream Server: Record", Value: 1}})
		}
		if err != nil {
			return err
		}

		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}

	return nil
}

//3 双向流     rpc Route(stream StreamRequest) returns (stream StreamResponse) {};
func (s *StreamService) Route(stream proto.StreamService_RouteServer) error {
	n := 0
	for {
		err := stream.Send(&proto.StreamResponse{
			Pt: &proto.StreamPoint{
				Name:  "gPRC Stream Client: Route",
				Value: int32(n),
			},
		})
		if err != nil {
			return err
		}

		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		n++

		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}

	return nil
}

//函数调用
func (s *StreamService) CallDemo(ctx context.Context, req *proto.CallRequest) (*proto.CallResponse, error) {
	fmt.Printf("服务器收到调用 : %v\n", req.CallString)
	rsp := &proto.CallResponse{CallString: "服务器收到调用，调用成功"}
	return rsp, nil
}
