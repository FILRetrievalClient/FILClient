package user

import (
	"FILClient/models/db"
	"FILClient/models/user"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/jameskeane/bcrypt"
	"github.com/kataras/iris/v12"
	"net/http"
	"strconv"
	"time"
)

// 根据用户名查询用户
func UserAdminCheckLogin(username string) *user.User  {
	user := new(user.User)
	db.IsNotFound(db.GetDB().Where("username = ?",username).First(user).Error)
	return user
}

// 检查登录用户，并生成登录凭证 token
func CheckLogin(username,password string)(response user.Token,status bool,msg string)  {
	cuser := UserAdminCheckLogin(username)
	if cuser.ID == 0 {
		msg = "用户不存在"
		return
	}else {
		if ok:= bcrypt.Match(password, cuser.Password);ok{
			token := jwt.NewTokenWithClaims(jwt.SigningMethodES256,jwt.MapClaims{
				"exp":time.Now().Add(time.Hour*time.Duration(1)).Unix(),
				"iat":time.Now().Unix(),
			})
			tokenString,_ := token.SignedString([]byte("HS2JDFKhu7Y1av7b"))
			authToken := new(user.AuthToken)
			authToken.Token = tokenString
			authToken.UserId = cuser.ID
			authToken.Secret = "secret"
			authToken.Revoked = false
			authToken.ExpressIn = time.Now().Add(time.Hour* time.Duration(1)).Unix()
			authToken.CreatedAt = time.Now()
			response = authToken.AuthTokenCreate()
			status = true
			msg = "登录成功"
			return
		}else {
			msg = "用户名或密码错误"
			return
		}
	}
}


// 登录处理程序
func UserLogin(ctx iris.Context)  {
	auth := new(user.User)
	if err := ctx.ReadJSON(&auth);err != nil{
		ctx.StatusCode(iris.StatusOK)
		_,_ = ctx.JSON(user.Response{Status: false,Msg: nil,Data: "请求参数错误"})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	response,status,msg := CheckLogin(auth.Username,auth.Password)
	_,_ = ctx.JSON(user.Response{
		Status: status,
		Msg:    response,
		Data:   msg,
	})
	return
}

// 作废token
func UpdateAuthTokenByUserId(userId uint) (at *user.AuthToken) {
	db.GetDB().Model(at).Where("revoked = ?",false).
		Where("user_id = ?",userId).
		Updates(map[string]interface{}{"revoked":true})
	return
}

// 登出用户
func UserAdminLogout(userId uint)bool  {
	ot := UpdateAuthTokenByUserId(userId)
	return ot.Revoked
}

// 登出
func UserLogout (ctx iris.Context){
	aui := ctx.Values().GetString("user_id")
	id,_ := strconv.Atoi(aui)
	UserAdminLogout(uint(id))
	ctx.StatusCode(http.StatusOK)
	_,_ = ctx.JSON(user.Response{
		Status: true,
		Msg:    nil,
		Data:   "退出",
	})
}


// 获得AuthToken信息
func GetAuthTokenByToken(token string) (at *user.AuthToken) {
	at = new(user.AuthToken)
	db.GetDB().Where("token = ?",token).First(&at)
	return
}

// 验证jwt
func JwtHandler()*jwt.Middleware  {
	var mySecret = []byte("HS2JDFKhu7Y1av7b")
	return jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return mySecret,nil
		},
		SigningMethod:       jwt.SigningMethodES256,
	})
}

func AuthToken(ctx iris.Context)  {
	value := ctx.Values().Get("jwt").(*jwt.Token)
	token := GetAuthTokenByToken(value.Raw)
	if token.Revoked || token.ExpressIn < time.Now().Unix(){
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.StopExecution()
		return
	}else {
		ctx.Values().Set("user_id",token.UserId)
	}
	ctx.Next()
}