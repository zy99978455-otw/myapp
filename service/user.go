package service

import (
	"errors"
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
