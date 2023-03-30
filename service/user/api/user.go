package main

import (
	"flag"
	"fmt"

	"github.com/ch3nnn/blog-admin-go/service/user/api/internal/config"
	"github.com/ch3nnn/blog-admin-go/service/user/api/internal/handler"
	"github.com/ch3nnn/blog-admin-go/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "service/user/api/etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
