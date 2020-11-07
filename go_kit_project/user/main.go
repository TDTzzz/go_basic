package main

import (
	"context"
	"flag"
	"fmt"
	"go_basic/go_kit_project/user/dao"
	"go_basic/go_kit_project/user/endpoint"
	"go_basic/go_kit_project/user/redis"
	"go_basic/go_kit_project/user/service"
	"go_basic/go_kit_project/user/transport"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {

	var (
		//服务地址和服务名
		servicePort = flag.Int("service.port", 10086, "service port")
	)

	flag.Parse()

	ctx := context.Background()
	errChan := make(chan error)

	err := dao.InitMysql("127.0.0.1", "3306", "root", "123456", "test_db")

	if err != nil {
		log.Fatal(err)
	}
	err = redis.InitRedis("127.0.0.1", "6379", "")
	if err != nil {
		log.Fatal(err)
	}

	userService := service.MakeUserServiceImpl(&dao.UserDAOImpl{})

	userEndpoints := &endpoint.UserEndpoints{
		endpoint.MakeRegisterEndpoint(userService),
		endpoint.MakeLoginEndpoint(userService),
	}
	r := transport.MakeHttpHandler(ctx, userEndpoints)

	go func() {
		errChan <- http.ListenAndServe(":"+strconv.Itoa(*servicePort), r)
	}()
	go func() {
		// 监控系统信号，等待 ctrl + c 系统信号通知服务关闭
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	
	error := <-errChan
	log.Println(error)
}
