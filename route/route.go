package route

import (
	"encoding/json"
	"fmt"
	"goweb/db"
	code "goweb/some_important_code"
	"net/http"
	"sync"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var conns sync.Map

func Init_route(r *gin.Engine) {

	r.Static("/static", "./static")
	r.LoadHTMLGlob("./static/html/*")
}
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}
func Login_POST(c *gin.Context) {
	username := c.PostForm("username")
	passwd := c.PostForm("passwd")
	var parmcode code.Parmcode = code.OK
	if !db.Usersearch(username) {
		parmcode = code.UserNotExiest
		c.JSON(http.StatusOK, gin.H{
			"parmcode": parmcode,
			"msg":      code.Parmcode_map[parmcode],
		})
	} else {
		if db.UserPasswdMatch(username, passwd) {
			parmcode = code.OK
			userjwt, err := code.CreateJWT(username)
			if err != nil {
				parmcode = code.Fail
			}
			c.JSON(http.StatusOK, gin.H{
				"parmcode": parmcode,
				"msg":      code.Parmcode_map[parmcode],
				"jwt":      userjwt,
			})
		} else {
			parmcode = code.PasswdNotMatch
			c.JSON(http.StatusOK, gin.H{
				"parmcode": parmcode,
				"msg":      code.Parmcode_map[parmcode],
			})
		}
	}
}
func Register_POST(c *gin.Context) {
	username := c.PostForm("username")
	passwd1 := c.PostForm("passwd1")
	passwd2 := c.PostForm("passwd2")
	icode := c.PostForm("invite_code")
	var parmcode code.Parmcode = code.OK
	if passwd1 != passwd2 {
		parmcode = code.PasswdNotMatch
	}
	if len(username) < 6 {
		parmcode = code.UserNmeToShort
	}
	if len(username) > 200 {
		parmcode = code.UsernameToLong
	}
	if len(passwd1) < 6 || len(passwd2) < 6 {
		parmcode = code.PasswdToShort
	}
	if len(passwd1) > 22 || len(passwd2) > 22 {
		parmcode = code.PasswdToLong
	}
	if icode != code.Invite_code {
		parmcode = code.InvitecodeError
	}
	if db.Usersearch(username) {
		parmcode = code.UserExiest
	}
	if parmcode == code.OK {
		check := db.InsertUser(username, passwd1)
		if check {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
		} else {
			parmcode = code.Fail
			c.JSON(http.StatusOK, gin.H{
				"parmcode": parmcode,
				"msg":      code.Parmcode_map[parmcode],
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"parmcode": parmcode,
			"msg":      code.Parmcode_map[parmcode],
		})
	}
}
func CheckJWT(c *gin.Context) {
	var isLogin bool = false
	jwt := c.Request.Header.Get("Cookie")
	fmt.Println(jwt)
	username, err := code.ParseJWT(jwt)
	if err == nil {
		isLogin = true
	}
	if isLogin {
		c.Set("username", username)
	} else {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
	}
}

func Homepage(c *gin.Context) {
	cvalue, _ := c.Get("username")
	uname, _ := cvalue.(string)
	c.HTML(http.StatusOK, "homepage.html", gin.H{
		"invitecode": code.Invite_code,
		"username":   uname,
	})
}

func Msg_center(c *gin.Context) {
	cvalue, _ := c.Get("username")
	uname, _ := cvalue.(string)
	c.HTML(http.StatusOK, "msg_center.html", gin.H{
		"invitecode": code.Invite_code,
		"username":   uname,
	})
}

func Ai_chat_room(c *gin.Context) {
	c_name, _ := c.Get("username")
	uname, _ := c_name.(string)
	c.HTML(http.StatusOK, "ai_chat_room.html", gin.H{
		"invitecode": code.Invite_code,
		"username":   uname,
		"history_msg":db.Query_ai_chat_history(uname),
	})
}
//调用自己的AI回复
func ai_reply(question string) string {
	reply:="你说的是"+question+"吗？"
	return reply
}

func Ai_chat(c *gin.Context){
	var message string 
	var question string
	c_name, _ := c.Get("username")
	uname, _ := c_name.(string)
	var m map[string]interface{}
	b,_:=c.GetRawData()
	json.Unmarshal(b,&m)
	question=m["message"].(string)
	db.Insert_ai_chat(uname,uname,question,time.Now().Format("2006-01-02 15:04:05"))

	//调用自己的AI回复
	message=ai_reply(question)

	db.Insert_ai_chat(uname,"ai",message,time.Now().Format("2006-01-02 15:04:05"))
	c.JSON(http.StatusOK, gin.H{
		"username":"ai",
		"massage":message,
		"time":time.Now().Format("15:04:05"),
	})
}
func Recieve_msg(c *gin.Context) {
	cvalue, _ := c.Get("username")
	uname, _ := cvalue.(string)
	c.JSON(http.StatusOK, gin.H{
		"parmcode": code.OK,
		"name":     uname,
	})
	b, _ := c.GetRawData()
	var m db.Msg
	json.Unmarshal(b, &m)
	fmt.Println(m)
	db.Insert_msg(m)
	Send_msg(m)
}
func Send_msg(m db.Msg){
	json_msg, _ := json.Marshal(m)
	conns.Range(func(key, value interface{}) bool {
		v := value.(*websocket.Conn)
		v.WriteMessage(websocket.TextMessage, json_msg)
		return true
	})
} 
func Query_msg(c *gin.Context) {
	msgs := db.Query_msg()
	c.JSON(http.StatusOK, gin.H{
		"parmcode": code.OK,
		"msgs":     msgs,
	})
}

func Up_ws(rw http.ResponseWriter, req *http.Request) {
	upGrader := websocket.Upgrader{
		ReadBufferSize:  2048,
		WriteBufferSize: 2048,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	cnn, err := upGrader.Upgrade(rw, req, nil)
	if err != nil {
		return
	}
	l := 0
	conns.Range(func(key, value interface{}) bool {
		l++
		return true
	})
	conns.Store(l, cnn)
	for {
		_, _, err := cnn.ReadMessage()
		if err != nil {
			break
		}
	}
	defer func() {
		cnn.Close()
		conns.Delete(l)
		fmt.Println("websocket closed")
	}()
}
