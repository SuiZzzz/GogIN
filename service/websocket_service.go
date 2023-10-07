package service

import (
	"GoGin/dao"
	"GoGin/dao/model"
	"GoGin/serializer"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type Client struct {
}

var ClientInstance *Client
var once sync.Once

func NewClient() *Client {
	once.Do(func() {
		ClientInstance = &Client{}
	})
	return ClientInstance
}

func (*Client) HandleMessage(conn *websocket.Conn, ctx context.Context) *serializer.ClientResp {
	if conn == nil {
		log.Println("websocket_service err: websocket conn is nil")
		return &serializer.ClientResp{
			Message: nil,
			Code:    500,
			Error:   "websocket_service err: websocket conn is nil",
		}
	}
	_, bytes, err := conn.ReadMessage()
	if err != nil {
		log.Println("websocket_service err:", err)
		return &serializer.ClientResp{
			Message: nil,
			Code:    500,
			Error:   "websocket_service err" + err.Error(),
		}
	}

	req := &serializer.ClientReq{}
	err = json.Unmarshal(bytes, req)
	if err != nil {
		log.Println("websocket_service err:", err)
		return &serializer.ClientResp{
			Message: nil,
			Code:    500,
			Error:   "websocket_service err" + err.Error(),
		}
	}
	switch req.Type {
	case serializer.Notification:
		return getNotification(ctx, req)
	case serializer.Group:
		return getGroup(ctx, req)
	default:
		log.Println("websocket_service err: unknown type")
		return &serializer.ClientResp{
			Message: nil,
			Code:    500,
			Error:   "websocket_service err: unknown type",
		}
	}

}

func getNotification(ctx context.Context, req *serializer.ClientReq) *serializer.ClientResp {
	client := dao.NewClientDao(ctx)
	messages := client.FindNotificationsBatches(req.UserId)
	user := dao.NewUserDao(ctx)
	var message []string
	var auditType []string
	for _, n := range *messages {
		text := ""
		// 审核状态
		audit := model.Audit[n.Audit]
		auditType = append(auditType, audit)
		if n.AuditorId == 0 || n.AuditorId != 0 && n.AuditorId == req.UserId {
			// 当前为审核人员
			username := user.FindById(n.SenderId).Nickname
			text = fmt.Sprintf("审核状态：%s\n用户 %s 向您发送审核申请，审核语句为：\n%s", audit, username, n.SQL)
		} else if req.UserId != n.AuditorId && req.UserId != 0 {
			// 当前为普通用户
			auditorName := user.FindById(n.AuditorId).Nickname
			text = fmt.Sprintf("审核状态：%s\n您向 %s 发送审核申请，审核语句为：\n%s", audit, auditorName, n.SQL)
		}
		message = append(message, text)
	}
	resp := serializer.ClientResp{Message: message, Audit: auditType}
	return &resp
}

func getGroup(ctx context.Context, req *serializer.ClientReq) *serializer.ClientResp {
	client := dao.NewClientDao(ctx)
	users := client.FindUser()
	var message []string
	var nickname []string
	for _, user := range *users {
		level := ""
		text := ""
		switch user.Level {
		case 1:
			level = "用户"
		case 2:
			level = "管理员"
		case 3:
			level = "超级管理员"
		}
		text = fmt.Sprintf("id：%d\n用户名：%s\n用户级别：%s", user.ID, user.Nickname, level)
		nickname = append(nickname, user.Nickname)
		message = append(message, text)
	}
	return &serializer.ClientResp{Message: message, Nickname: nickname}
}
