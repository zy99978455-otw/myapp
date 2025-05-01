package service

import (
	"myapp/common"
	"myapp/model"
)

func Migrate() error {
	return common.DB.AutoMigrate(&model.User{}, &model.Comment{})
}
