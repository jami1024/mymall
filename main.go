package main

import (
	serviceAccount "mymall/service/account/account"
	serviceOrder "mymall/service/account/order"
	"mymall/service/server"

	"github.com/gin-gonic/gin"
)


func main() {
	g := gin.New()
	//登陆、注册、退出
	g.POST("/login", server.Login)
	g.POST("/register", server.Register)
	g.POST("/logout",server.Logout)

	account := g.Group("/account")
	order := account.Group("/order")
	order.POST("/list", serviceOrder.List)
	account.POST("/profile", serviceAccount.Profile)
	g.Run()
}