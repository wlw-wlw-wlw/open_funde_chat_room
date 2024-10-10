package code

import (
	"errors"
	"math/rand"
	"sync"
	"time"
	"github.com/golang-jwt/jwt/v4"
)

var Invite_code string = "wlw"
var wnum sync.WaitGroup

var Login_sign_kay=[]byte("wlw_wlw_wlw_wlw")
const TokenExpireDuration = time.Hour
type Custom_info struct{
	Username string
	jwt.RegisteredClaims
}

type Parmcode int

const (
	//总
	OK   = Parmcode(1)
	Fail = Parmcode(0)
	//前端部分错误
	PasswdNotMatch = Parmcode(101)
	UserExiest     = Parmcode(102)
	UserNmeToShort = Parmcode(103)
	PasswdToShort = Parmcode(104)
	UsernameToLong   = Parmcode(105)
	PasswdToLong   = Parmcode(106)
	//邀请码部分错误
	InvitecodeError = Parmcode(201)
	//数据库部分错误
	UserNotExiest = Parmcode(301)
	PaswwdError   = Parmcode(302)
)

var Parmcode_map map[Parmcode]string

func init() {
	Parmcode_map = make(map[Parmcode]string)
	Parmcode_map[OK] = "操作成功"
	Parmcode_map[Fail] = "未知原因的失败"

	Parmcode_map[PasswdNotMatch] = "密码不匹配"
	Parmcode_map[UserExiest] = "用户已存在"
	Parmcode_map[UserNmeToShort] = "用户名太短"
	Parmcode_map[PasswdToShort] = "密码太短"
	Parmcode_map[UsernameToLong] = "用户名太长"
	Parmcode_map[PasswdToShort] = "密码太长"

	Parmcode_map[InvitecodeError] = "邀请码错误"

	Parmcode_map[UserNotExiest] = "用户不存在"
	Parmcode_map[PaswwdError] = "密码错误"

	go set_in_code()
	wnum.Add(1)
}
func set_in_code() {
	defer wnum.Done()
	var r int
	by := make([]byte, 7)
	for {
		for i := 0; i < 7; i++ {
			r = rand.Int()
			r = r % 36
			if r < 10 {
				r += 48
			} else {
				r += 87
			}
			by[i] = byte(r)
		}
		Invite_code = string(by)
		time.Sleep(5 * time.Minute)
	}
}

func CreateJWT(username string)(string,error){
	claims:=Custom_info{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "wlw", // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(Login_sign_kay)
}

func ParseJWT(JWTstring string) (string,error){
	var c =new(Custom_info)
	token,err:=jwt.ParseWithClaims(JWTstring,c,func(tocken *jwt.Token)(i interface{},e error){
		return Login_sign_kay,nil
	} )
	if err!=nil{
		return "",err
	}
	if claims, ok := token.Claims.(*Custom_info); ok && token.Valid { // 校验token
		return claims.Username, nil
	}
	return "", errors.New("invalid token")
}