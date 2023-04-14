package main

import (
	"flag"
	"fmt"
	commonConfig "github.com/ch3nnn/blog-admin-go/common/config"
	"github.com/ch3nnn/blog-admin-go/service/user/rpc/internal/svc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"

	"github.com/ch3nnn/blog-admin-go/service/user/rpc/internal/config"
	"github.com/ch3nnn/blog-admin-go/service/user/rpc/internal/server"
	"github.com/ch3nnn/blog-admin-go/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "service/user/rpc/etc/consul.yaml", "the config file")

func main() {
	flag.Parse()

	var consulConfig commonConfig.ConsulConfig
	conf.MustLoad(*configFile, &consulConfig)

	// 读取 consul 配置中心
	var c config.Config
	commonConfig.LoadYAMLConf(commonConfig.ConsulConfig.NewClient(consulConfig), consulConfig.Consul.Key, &c)

	ctx := svc.NewServiceContext(c)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// 服务注册 consul
	_ = consul.RegisterService(c.ListenOn, c.Consul)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
