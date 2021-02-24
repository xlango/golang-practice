package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "golang-practice/grpc/proto"

	"google.golang.org/grpc"
)

func main() {
	// 通过 grpc.Dial 获得一条连接
	conn, err := grpc.Dial("unix:///var/lib/test.socket", grpc.WithInsecure())
	// 如果要增加 Recv 可以接受的一个消息的数据量，必须增加 grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(100000000))
	//conn, err := grpc.Dial("unix:///var/lib/test.socket", grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(100000000)))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	/*
	   在 service.pb.go 中

	   // 接口 interface
	   type MaxSizeClient interface {
	       Echo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MaxSize_EchoClient, error)
	   }

	   type maxSizeClient struct {
	       cc *grpc.ClientConn
	   }
	   // 传入一个连接，返回一个 MaxSizeClient 的实例，这个实例实现了 MaxSizeClient 接口 Echo，实际上是 maxSizeClient 的实例
	   func NewMaxSizeClient(cc *grpc.ClientConn) MaxSizeClient {
	       return &maxSizeClient{cc}
	   }

	   注意名字，NewMaxSizeClient = New + MaxSize（service MaxSize {} in proto 文件）+ Client
	*/
	client := pb.NewMaxSizeClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Second)
	defer cancel()

	/*
	   在 service.pb.go 中，参数是 1 context，2 Empty，返回值是 MaxSize_EchoClient, error
	   func (c *maxSizeClient) Echo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MaxSize_EchoClient, error) {
	       stream, err := c.cc.NewStream(ctx, &_MaxSize_serviceDesc.Streams[0], "/test.MaxSize/Echo", opts...)
	       if err != nil {
	           return nil, err
	       }
	       x := &maxSizeEchoClient{stream}
	       if err := x.ClientStream.SendMsg(in); err != nil {
	           return nil, err
	       }
	       if err := x.ClientStream.CloseSend(); err != nil {
	           return nil, err
	       }
	       return x, nil
	   }

	   // MaxSize_EchoClient 是一个 interface
	   // 必须实现 Recv 方法
	   type MaxSize_EchoClient interface {
	       Recv() (*StringMessage, error)
	       grpc.ClientStream
	   }

	   type maxSizeEchoClient struct {
	       grpc.ClientStream
	   }

	   func (x *maxSizeEchoClient) Recv() (*StringMessage, error) {
	       m := new(StringMessage)
	       if err := x.ClientStream.RecvMsg(m); err != nil {
	           return nil, err
	       }
	       return m, nil
	   }

	*/
	//stream 是实现 MaxSize_EchoClient 的实例
	stream, err := client.Echo(ctx, &pb.Empty{})

	for {
		// stream 有一个最重要的方法，就是 Recv()，Recv 的返回值就是 *pb.StringMessage，这里面包含了多个 Ss []*StringSingle
		data, err := stream.Recv()
		if err != nil {
			fmt.Printf("error %v", err)
			return
		}
		fmt.Printf("%v", data)
	}

}
