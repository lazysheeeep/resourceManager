package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/lazysheeeep/resourceManager/rpc/internal/config"
	"github.com/lazysheeeep/resourceManager/rpc/internal/svc"
)

var configFile = flag.String("f", "etc/core.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	ctx, err := svc.NewServiceContext(c)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(ctx)
	}
}
