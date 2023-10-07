package dao

import (
	"GoGin/dao/model"
	"context"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewSession(ctx)}
}

// IsExistUser 判断用户名是否存在
func (u *UserDao) IsExistUser(username string) (isExist bool, err error) {
	var result model.User
	u.DB.Table("user").Select("id").Where("username = ?", username).First(&result)
	// 用户已存在
	if result.ID != 0 {
		return true, nil
	}
	return false, nil
}

// InsertUser 插入用户
func (u *UserDao) InsertUser(user *model.User) error {
	tx := u.DB.Begin()
	err := tx.Table("user").Select("username", "password", "nickname").Create(user).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return err
}

// FindByUsername 校验密码
func (u *UserDao) FindByUsername(user *model.User) (result *model.User, err error) {
	u.DB.Model(&model.User{}).
		Where("username = ?", user.Username).First(&result)
	return result, nil
}

// FindById 根据用户id查找用户
func (u *UserDao) FindById(id uint) (result *model.User) {
	u.Session(&gorm.Session{QueryFields: true}).Table("user").
		Where("id = ?", id).First(&result)
	return result
}
