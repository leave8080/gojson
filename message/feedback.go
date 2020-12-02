package message

import (
	"gofer/pkg/log"
	"time"
)

// 用户反馈表
type MxFeedback struct {
	Id        int        `gorm:"column:id;primary_key" json:"id"`
	UserId    int        `gorm:"column:user_id" json:"user_id"`
	ProductId int        `gorm:"column:product_id" json:"product_id"`
	Phone     string     `gorm:"column:phone" json:"phone"`
	Content   string     `gorm:"column:content" json:"content"`
	Type      int        `gorm:"column:type" json:"type"`     // 反馈类型：0-其他问题，1-设备问题，2-配网问题，3-反馈故障，4-功能建议
	Status    int        `gorm:"column:status" json:"status"`   // 反馈状态：0-未解决，1-已解决
	Platform  int        `gorm:"column:platform" json:"platform"` // 设备平台：0-未知，1-Android，2-iOS
	Version   string     `gorm:"column:version" json:"version"`
	Ctime     time.Time `gorm:"column:ctime" json:"ctime"`
	Mtime     time.Time `gorm:"column:mtime" json:"mtime"`
}

func (m *MxFeedback) TableNames() string {
	return "mx_feedback"
}

func (m *MxFeedback)Execute()  {
	log.Info("Execute===>",m.TableNames())
}