package service

import (
	"errors"
	"myapp/common"
	"myapp/model"
	"myapp/util"
	"myapp/vo"
)

// ValidateUser 用户数据合法性校验
func ValidateUser(user *model.User) error {
	if len(user.Name) == 0 {
		return errors.New("用户名不能为空")
	}

	if len(user.Nickname) == 0 {
		return errors.New("昵称不能为空")
	}

	if len(user.Password) == 0 {
		return errors.New("密码不能为空")
	}

	return nil
}

// RegisterUser 用户注册
func RegisterUser(user *model.User) error {
	// 对于用户密码进行MD5加密
	user.Password = util.MD5(user.Password)

	// 调用首次创建方法
	return common.DB.FirstOrCreate(&model.User{}, user).Error
}

// ValidateLogin 用户数据合法性校验
func ValidateLogin(user *model.User) error {
	if len(user.Name) == 0 {
		return errors.New("用户名不能为空")
	}

	if len(user.Password) == 0 {
		return errors.New("密码不能为空")
	}
	return nil
}

func Login(user *model.User) (bool, error) {
	user.Password = util.MD5(user.Password)
	var users []model.User
	err := common.DB.Find(&users, &user).Error

	if err != nil || len(users) == 0 {
		return false, err
	}
	return true, nil
}

func GetUserByName(username string) (*vo.UserVO, error) {
	queryUser := model.User{Name: username}

	var users []model.User
	err := common.DB.Find(&users, queryUser).Error

	if err != nil || len(users) == 0 {
		return nil, err
	}

	user := users[0]

	userVO := vo.UserVO{ID: user.ID, Name: user.Name, Nickname: user.Nickname}
	return &userVO, nil
}

func DeleteById(id int) error {
	var user model.User

	result := common.DB.First(&user, id)

	if result.Error != nil {
		return result.Error
	}

	result = common.DB.Delete(&user)
	return result.Error
}
