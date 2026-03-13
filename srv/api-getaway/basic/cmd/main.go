package main

import (
	_ "github.com/yuhang-jieke/orderai/srv/api-getaway/inits"
	"github.com/yuhang-jieke/orderai/srv/api-getaway/pkg"
	_ "github.com/yuhang-jieke/orderai/srv/api-getaway/pkg"
)

func main() {
	factory := pkg.NewSimpleFactory()
	httpService, _ := factory.CreateService(pkg.ServiceTypeHTTP, nil)
	httpService.Start()
}
