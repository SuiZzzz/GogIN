package model

import (
	"gorm.io/gorm"
	"time"
)

var Audit = map[byte]string{
	0: "未审核",
	1: "已通过",
	2: "未通过",
}

type Notification struct {
	gorm.Model
	SenderId   uint      `json:"sender_id" gorm:"default:0;comment:'发送者id'"`
	AuditorId  uint      `json:"auditor_id" gorm:"default:0;comment:'接收者id，系统默认0'"`
	ReceiverId uint      `json:"receiver_id" gorm:"default:0;comment:'审核人id'"`
	Type       byte      `json:"type" gorm:"default:0;comment:'消息类型，默认0'"`
	Audit      byte      `json:"audit" gorm:"default:0;comment:'是否审核，0未审核，1已通过，2未通过'"`
	SQL        string    `json:"sql" gorm:"type:varchar(100);default:'';comment:'sql语句"`
	SQLResult  bool      `json:"sql_result" gorm:"default:false;comment:'SQL查询结果，默认false'"`
	Text       string    `json:"text" gorm:"type:varchar(60);default:'';comment:'消息"`
	Error      string    `json:"error" gorm:"type:varchar(60);default:'';comment:'错误"`
	AuditedAt  time.Time `json:"audited_at" gorm:"default:'2000-01-01 00:00:00';comment:'审核时间"`
}
