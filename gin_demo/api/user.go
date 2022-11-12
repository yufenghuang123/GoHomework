package api
//编写 apis
//到现在我们就需要开始写我们的逻辑了，写之前我们先思考一下大概需要些什么逻辑。
//
//登录：
//
//1.传入用户名。
//2.验证是否有该用户，没有则直接退出。
//3.验证密码是否正确。
//4.正确则返回我们的 token 或者是 Set Cookie。
//注册：
//
//1.传入用户名和密码
//2.验证用户名是否重复，若重复也直接退出。
//3.注册成功。
import (
	"fmt"
	"gin-demo/api/middleware"
	"gin-demo/dao"
	"gin-demo/model"
	"gin-demo/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)


// 新增以下代码
func getUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	utils.RespSuccess(c, username.(string))
}
func register(c *gin.Context)  {
	if err := c.ShouldBind(&model.User{});err !=nil{
		utils.RespSuccess(c,"verification failed")
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	//验证用户名是否重复
	flag :=dao.SelectUser(username)
	fmt.Println(flag)
	if flag{
		//以json格式返回信息
		utils.RespFail(c,"user already exists")
		return
	}
	dao.AddUser(username,password)

	utils.RespSuccess(c,"add user successful")
}

func login(c *gin.Context)  {
	if err :=c.ShouldBind(&model.User{});err !=nil{
		utils.RespFail(c,"verification")
		return
	}
	//传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag :=dao.SelectUser(username)
	if !flag{
		utils.RespFail(c,"user doesn't exists")
		return
	}
	selectpassword:=dao.SelectUserPasswordUsername(username)
	if selectpassword != password{
		utils.RespFail(c,"wrong password")
		return
	}
	//正确则登录成功
	//创建一个我们自己的声明
	claim := model.MyClaims{
		Username: username, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // 过期时间
			Issuer:    "Yxh",                                // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	//c.SetCookie("gin_demo_cookie", "test", 3600, "/", "localhost", false, true)
	tokenString, _ := token.SignedString(middleware.Secret)
	utils.RespSuccess(c,tokenString)

}
