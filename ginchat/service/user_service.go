package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

// GetUserList
// @Summary 用户列表
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	list := models.GetUserList()
	c.JSONP(200, gin.H{
		"message": list,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "二次密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")

	if models.FindUserByName(user.Name).Error != gorm.ErrRecordNotFound {
		c.JSON(-1, gin.H{
			"message": "用户名已注册",
		})
		return
	}

	if password != repassword {
		c.JSONP(-1, gin.H{
			"message": "两次密码不一致!",
		})
		return
	}

	user.Password = utils.Md5Encode(password)
	models.CreateUser(&user)
	c.JSONP(200, gin.H{
		"message": "新增用户成功",
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "ID"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	atoi, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(atoi)
	models.DeleteUser(&user)
	c.JSONP(200, gin.H{
		"message": "删除用户成功",
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "用户名"
// @param password formData string false "密码"
// @param rePassword formData string false "二次密码"
// @param phone formData string false "电话号码"
// @param email formData string false "邮箱"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}

	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		c.JSONP(200, gin.H{
			"message": "参数不匹配",
		})
		return
	}

	if user.Password != c.PostForm("rePassword") {
		c.JSONP(-1, gin.H{
			"message": "两次密码不一致",
		})
		return
	}

	models.UpdateUser(&user)
	c.JSONP(200, gin.H{
		"message": "修改用户成功",
	})
}

// FindUserByNameAndPsw
// @Summary 验证账号和密码
// @Tags 用户模块
// @param name formData string false "用户名"
// @param password formData string false "密码"
// @Success 200 {string} json{"code","message","data}
// @Router /user/findUserByNameAndPsw [post]
func FindUserByNameAndPsw(c *gin.Context) {

	name := c.PostForm("name")
	password := c.PostForm("password")
	user := models.FindUserByNameAndPwd(name, password)
	if user.Name == "" {
		c.JSONP(-1, gin.H{
			"code":    0,
			"message": "用户名或密码错误",
		})
		return
	}
	c.JSONP(200, gin.H{
		"code":    1, //1成功,0失败
		"message": "登录成功",
		"data":    user,
	})
}

// 防止跨域站点伪造请求
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMessage(c *gin.Context) {
	upgrade, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	fmt.Println(upgrade)
	if err != nil {
		panic(err)
	}
	defer func(ws *websocket.Conn) {
		err2 := ws.Close()
		if err2 != nil {
			panic(err2)
		}
	}(upgrade)
	MsgHandler(upgrade, c)
}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	message, err := utils.Subscribe(c, utils.PublishKey)
	fmt.Println("MsgHandler", message.Payload)
	if err != nil {
		fmt.Println(err)
	}

	format := time.Now().Format("2006-01-02 15:04:05")
	res := fmt.Sprintf("[ws][%s]:%s", format, message)
	err = ws.WriteMessage(1, []byte(res))
	if err != nil {
		fmt.Println(err)
	}
}
