/*
	平台主函数
*/
package main

import (
	"fmt"
	_ "incentiveblog/config"
	"incentiveblog/router"
	_ "net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("welcome to incentiveblog")
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	//静态文件处理
	e.Static("/", "boke")     // /index.html -- > boke/index.html
	e.Static("/boke", "boke") // /boke/login.html -> boke/login.html

	//加一个测试路由
	e.GET("/ping", router.Pong)
	e.POST("/register", router.Register)            //注册
	e.POST("/login", router.Login)                  //登陆
	e.POST("/checktoken", router.CheckLogin)        //检测登陆
	e.POST("/content", router.UploadContent)        //上传文件
	e.GET("/content/:hash", router.DownloadContent) //下载文件
	e.POST("/detail", router.GetDetail)             //查看用户积分明细
	e.POST("/blog", router.PublishBlog)             //发表博客
	e.GET("/blog/:userid", router.GetBlogs)         //查看博客
	e.POST("/concern", router.UserConcern)          //用户关注
	e.GET("/concern/:userid", router.CountConcern)  //粉丝统计

	e.Logger.Fatal(e.Start(":8086"))
}
