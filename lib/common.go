package lib

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"time"
)

// JSON 数据 code 返回值
const (
	GeneralErrorCode = 10003
	SuccessCode = 0
	ServerErrorCode = 500
	ServerNotFount = 404
	UnAuthorized = 401
	)
type NewStringMap map[string]interface{}

func (this *NewStringMap) Get(key string) interface{}{
	return (*this)[key]
}
func (this *NewStringMap) Set(key string,value interface{}) {
	(*this)[key] = value
}

// MD5 加密
func EncryptStr(salt,str string) (password string) {
	pattern := []byte(str + salt)
	has := md5.New()
	code := has.Sum(pattern)
	return hex.EncodeToString(code)
}

// 检验密码
func CheckPassword(salt,s_password,c_password string) bool {
	curr_password := EncryptStr(salt,c_password)
	if curr_password == s_password {
		return true
	}
	return false
}

// 生成MD5 盐
func GenerateSalt(length int) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 检查 Token 是否有效
func CheckToken(secret,en_token string) (interface{},bool) {
	logs.Info("received token: ",en_token)
	token,err := jwt.Parse(en_token,func(token *jwt.Token) (interface{},error){
		return []byte(secret),nil
	})
	if err != nil {
		logs.Error(err.Error())
		return nil,false
	}
	if !token.Valid {
		logs.Error("Token expired")
		return nil,false
	}
	claim,ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil ,false
	}
	username := claim["username"].(string)
	logs.Info("当前token 代表用户为: ",username)
	return username,true
}

