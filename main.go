package main

import (
	"goweb/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	route.Init_route(r)
	r.GET("/login", route.Login)
	r.GET("/register", route.Register)
	r.POST("/login", route.Login_POST)
	r.POST("/register", route.Register_POST)
	r.GET("/homepage", route.CheckJWT, route.Homepage)

	r.GET("/msg_center", route.CheckJWT, route.Msg_center)
	r.GET("/ai_chat_room", route.CheckJWT, route.Ai_chat_room)
	r.POST("/ai_chat", route.CheckJWT, route.Ai_chat)
	r.POST("/send_msg", route.CheckJWT, route.Recieve_msg)
	r.GET("/query_msg", route.CheckJWT, route.Query_msg)
	r.GET("/ws_receive_msg", route.CheckJWT,func(c *gin.Context){
		route.Up_ws(c.Writer,c.Request)
	})
	r.Run(":9999")
}
 