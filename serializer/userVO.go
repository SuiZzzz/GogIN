package serializer

import "GoGin/dao/model"

type UserVo struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Level    byte   `json:"level"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Id       uint   `json:"id"`
}

func BuildUserVO(user *model.User) *UserVo {
	return &UserVo{
		Username: user.Username,
		Nickname: user.Nickname,
		Level:    user.Level,
		Phone:    user.Phone,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Id:       user.ID,
	}
}
