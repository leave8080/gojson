package message

import (
	"gofer/pkg/log"
)

// 用户反馈表
type MxTest struct {
	Id    int    `gorm:"column:id;primary_key" json:"id"`
	Happy string `gorm:"column:happy" json:"happy"`
}

func (m *MxTest) TableNames() string {
	return "mx_test"
}

func (m *MxTest) Execute(req *MessInfo) {
	log.Info("Execute===>", m.TableNames())
	Task(req)
}
