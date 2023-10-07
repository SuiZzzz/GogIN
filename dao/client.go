package dao

import (
	"GoGin/dao/model"
	"context"
	"gorm.io/gorm"
)

type ClientDao struct {
	*gorm.DB
}

func NewClientDao(ctx context.Context) *ClientDao {
	return &ClientDao{NewSession(ctx)}
}

// FindNotificationsBatches 根据用户id查找所有通知
func (c *ClientDao) FindNotificationsBatches(send uint) *[]model.Notification {
	var messages []model.Notification
	c.Session(&gorm.Session{QueryFields: true}).Table("notification").
		Where("sender_id = ?", send).Or("receiver_id = ?", send).
		Order("created_at desc").Find(&messages)
	return &messages
}

// FindUser 查找所有用户
func (c *ClientDao) FindUser() *[]model.User {
	var users []model.User
	c.Session(&gorm.Session{QueryFields: true}).Table("user").Find(&users)
	return &users
}
