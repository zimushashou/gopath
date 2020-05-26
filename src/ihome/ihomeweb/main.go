package main

import (
        "github.com/julienschmidt/httprouter"
        "github.com/micro/go-log"
        "github.com/micro/go-web"
        "ihome/ihomeweb/handler"
        _ "ihome/ihomeweb/models"
        "net/http"
)

func main() {
	// create new web service
        service := web.NewService(
                web.Name("go.micro.web.ihomeweb"),
                web.Version("latest"),
                web.Address(":8999"),
        )

	// initialise service
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }
        router := httprouter.New()
        router.NotFound = http.FileServer(http.Dir("html"))             // must router.NotFound
        //获取地区信息
        router.GET("/api/v1.0/areas",handler.GetArea)
        //下面两个目前并不实现服务
        //获取session
        router.GET("/api/v1.0/session",handler.GetSession)
        //获取index
        router.GET("/api/v1.0/house/index",handler.GetIndex)
        // register html handler
	service.Handle("/", router)
	// run service
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
