package service

import (
	"GoGin/dao"
	"GoGin/dao/model"
	"GoGin/serializer"
	"GoGin/types"
	"GoGin/utils"
	"context"
	"encoding/hex"
	"log"
	"sync"
)

type UserService struct {
}

var UserServiceInstance *UserService

var UserServiceOnce sync.Once

func GetUserServiceInstance() *UserService {
	UserServiceOnce.Do(func() {
		UserServiceInstance = &UserService{}
	})
	return UserServiceInstance
}

// Register 注册用户
func (*UserService) Register(req *types.UserRegisterReq, ctx context.Context) *serializer.Response {

	userDao := dao.NewUserDao(ctx)

	// 判断是否已存在用户名
	isExit, err := userDao.IsExistUser(req.Username)
	if err != nil {
		log.Println("UserService.Register判断用户名是否存在发生错误")
		return &serializer.Response{
			Code:    serializer.IllegalData,
			Message: serializer.DataError,
		}
	}
	if isExit {
		return &serializer.Response{
			Code:    serializer.UsernameIsExist,
			Message: serializer.UserIsExit,
		}
	}

	// 密码加密存储
	password, err := utils.Encrypt([]byte(utils.Key), []byte(req.Password))
	if err != nil {
		log.Println("密码加密失败：", err)
		return &serializer.Response{
			Code:    serializer.IllegalData,
			Message: serializer.DataError,
		}
	}
	registerUser := model.User{
		Username: req.Username,
		Password: hex.EncodeToString(password),
		Nickname: req.Nickname,
		Level:    1,
		Phone:    req.Phone,
		Email:    req.Email,
	}

	// 用户插入
	err = userDao.InsertUser(&registerUser)

	if err != nil {
		log.Println("UserService.Register插入用户信息发生错误")
		return &serializer.Response{
			Code:    serializer.SystemError,
			Message: serializer.ServerError,
		}
	}
	return &serializer.Response{
		Code:    serializer.OK,
		Message: serializer.Succeed,
	}

}

// Login 登录用户
func (*UserService) Login(req *types.UserLoginReq, ctx context.Context) *serializer.Response {
	userDao := dao.NewUserDao(ctx)
	// 用户是否存在
	loginUser := model.User{
		Username: req.Username,
	}
	result, _ := userDao.FindByUsername(&loginUser)
	if result.ID == 0 {
		return &serializer.Response{
			Code:    serializer.UsernameNotExist,
			Message: "账号不存在",
		}
	}
	// 密码解密
	decode, _ := hex.DecodeString(result.Password)
	decrypt, err := utils.Decrypt([]byte(utils.Key), decode)
	if err != nil {
		log.Println("user_service.Login()密码解密失败：", err)
		return &serializer.Response{
			Code:    serializer.SystemError,
			Message: serializer.ServerError,
		}
	}
	if req.Password != string(decrypt) {
		return &serializer.Response{
			Code:    serializer.OK,
			Message: "密码错误，请重新登录",
		}
	}
	// 密码校验成功，签发jwt Token
	token, err := utils.GenerateToken(result.ID, result.Username)
	if err != nil {
		log.Println("user_service.Login()签发token失败:", err)
		return &serializer.Response{
			Code:    serializer.SystemError,
			Message: serializer.ServerError,
		}
	}
	return &serializer.Response{
		Code: serializer.OK,
		Data: &serializer.TokenData{
			User:  serializer.BuildUserVO(result),
			Token: token,
		},
		Message: serializer.LoginSucceed,
	}
}
