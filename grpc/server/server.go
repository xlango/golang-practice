package main

import (
	"log"
	"math"
	"net"

	"google.golang.org/grpc"

	pb "golang-practice/grpc/proto"
)

// 参考 /root/go/src/google.golang.org/grpc/examples/route_guide

// 定义了一个空的结构体，这是 go 语言的一个技巧
type server struct{}

// Echo 函数是 server 类的一个成员函数
// 这个 server 类必须能够实现 proto 文件中定义的所有 rpc
// 在 service.pb.go 文件中有详细的说明：
/*
   // 注意，MaxSizeServer 是 proto 中 service MaxSize 的 MaxSize + Server 拼成的！
   // MaxSizeServer is the server API for MaxSize service.
   // 他是一个 interface，只要实现了 Echo，就是这个 interface 的实现。可见，我们的 func (s *server) Echo(in *pb.Empty, stream pb.MaxSize_EchoServer) error { 实现了这个接口。注意，参数和返回值是不是和 interface 定义的一模一样？
   type MaxSizeServer interface {
       Echo(*Empty, MaxSize_EchoServer) error
   }
*/

func (s *server) Echo(in *pb.Empty, out pb.MaxSize_EchoServer) error {
	// proto 中定义 rpc Echo(Empty) returns (stream StringMessage) {};
	/*
	   in *pb.Empty 就是 Empty
	   out pb.MaxSize_EchoServer 是提供给用户的，能够调用 send 的一个 object，这个是精妙的设计提供给用户的
	   该代码中，要组织 *StringMessage 类型的返回值，使用 out.send 发送出去

	   注意，pb 是我们引用包的代号，import pb "test"

	   那么 pb.Empty 是什么呢？
	   // service.pb.go 定义的
	   type Empty struct {
	       XXX_NoUnkeyedLiteral struct{} `json:"-"`
	       XXX_unrecognized     []byte   `json:"-"`
	       XXX_sizecache        int32    `json:"-"`
	   }

	   那么 pb.MaxSize_EchoServer 是什么？
	   // service.pb.go 定义的
	   type MaxSize_EchoServer interface {
	       Send(*StringMessage) error
	       grpc.ServerStream
	   }

	   但是是否有人实现了这个接口呢？当然

	   // 在 service.pb.go 中：
	   type maxSizeEchoServer struct {
	       grpc.ServerStream
	   }

	   func (x *maxSizeEchoServer) Send(m *StringMessage) error {
	       return x.ServerStream.SendMsg(m)
	   }

	   从此，可知，pb.MaxSize_EchoServer 有 send 方法，可以将 StringMessage 发送出去。

	   那么 pb.StringMessage 是什么呢？
	   // service.pb.go 定义的
	   type StringMessage struct {
	       Ss                   []*StringSingle `protobuf:"bytes,1,rep,name=ss,proto3" json:"ss,omitempty"`
	       XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	       XXX_unrecognized     []byte          `json:"-"`
	       XXX_sizecache        int32           `json:"-"`
	   }

	   注意 Ss 和 proto 中的：
	   message StringMessage {
	       repeated StringSingle ss = 1;
	   }

	   有十分重大的关系，因为是 repeated，所以是 Ss []*StringSingle
	*/

	log.Printf("Received from client")
	var err error
	list := pb.StringMessage{}
	for i := 0; i < 5; i++ {
		feature := pb.StringSingle{
			Id:   "sssss",
			Name: "lihao",
		}
		list.Ss = append(list.Ss, &feature)
	}
	err = out.Send(&list)

	// 函数要求返回 error 类型
	return err
}

func run() error {
	sock, err := net.Listen("unix", "/var/lib/test.socket")
	if err != nil {
		return err
	}

	var options = []grpc.ServerOption{
		grpc.MaxRecvMsgSize(math.MaxInt32),
		grpc.MaxSendMsgSize(1073741824),
	}
	s := grpc.NewServer(options...)

	myServer := &server{}
	/*
	   见 service.pb.go 中
	   func RegisterMaxSizeServer(s *grpc.Server, srv MaxSizeServer) {
	       s.RegisterService(&_MaxSize_serviceDesc, srv)
	   }
	   前者是 grpc server，后者是实现了 MaxSizeServer 所有 interface 的实例，即 &server{}

	   感觉就是将 grpc server 和 handler 绑定在了一起的意思。

	   RegisterMaxSizeServer 的命名很有意思，Register(固定) + MaxSize(service MaxSize {} in proto 文件) + Server(固定)

	*/
	pb.RegisterMaxSizeServer(s, myServer)
	if err != nil {
		return err
	}

	/*
	   在 s.Serve(sock) 上监听服务

	*/
	if err := s.Serve(sock); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}

func main() {
	run()
}
