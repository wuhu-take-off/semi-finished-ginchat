package models

import (
	"ginchat/utils"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string    `gorm:"column:name" json:"name"`
	Password      string    `gorm:"column:password" json:"password"`
	Phone         string    `gorm:"column:phone" json:"phone" valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string    `gorm:"column:email" json:"email" valid:"email"`
	Identity      string    `gorm:"column:identity" json:"identity"`
	ClientIp      string    `gorm:"column:client_ip" json:"clientIp"`
	ClientPort    string    `gorm:"column:client_port" json:"clientPort"`
	LoginTime     time.Time `gorm:"column:login_time" json:"loginTime"`
	HeartbeatTime time.Time `gorm:"column:heartbeat_time" json:"heartbeatTime"`
	LogOutTime    time.Time `gorm:"column:login_out_time" json:"logOutTime"`
	IsLogout      bool      `gorm:"column:is_logout" json:"isLogout"`
	DeviceInfo    string    `gorm:"column:device_info" json:"deviceInfo"`
}

// 查看UserBasic类型在数据库中是否存在,如果不存在依照UserBasic类型在数据库中建立对应表格
func InitUserBasic() {
	err := utils.GetDB().AutoMigrate(UserBasic{})
	if err != nil {
		panic(err)
	}
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	db := utils.GetDB()

	data := make([]*UserBasic, 10)
	db.Find(&data)
	return data

}

func CreateUser(user *UserBasic) *gorm.DB {
	return utils.GetDB().Create(user)
}

func DeleteUser(user *UserBasic) *gorm.DB {
	return utils.GetDB().Delete(user)
}

func UpdateUser(user *UserBasic) *gorm.DB {
	return utils.GetDB().Model(&user).Updates(UserBasic{
		Name:     user.Name,
		Password: user.Password,
		Phone:    user.Phone,
		Email:    user.Email,
	})
}

func FindUserByName(name string) *gorm.DB {
	user := UserBasic{}
	return utils.GetDB().Where("name = ?", name).First(&user)
}
func FindUserByPhone(phone string) *gorm.DB {
	user := UserBasic{}
	return utils.GetDB().Where("phone = ?", phone).First(&user)
}
func FindUserByEmail(email string) *gorm.DB {
	user := UserBasic{}
	return utils.GetDB().Where("email = ?", email).First(&user)
}

func FindUserByNameAndPwd(name, password string) UserBasic {
	user := UserBasic{}
	utils.GetDB().Where("name = ? and password = ?", name, utils.Md5Encode(password)).First(&user)

	//token加密
	encode := utils.Md5Encode(strconv.FormatInt(time.Now().Unix(), 10))
	utils.GetDB().Model(&user).Where("id = ?", user.ID).Update("identity", encode)
	return user
}
