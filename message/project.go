package message

import (
	"gofer/pkg/log"
	"time"
)

// 项目表（飞燕项目等）
type MxProject struct {
	Id            int       `gorm:"column:id;primary_key" json:"id" validate:"required"`
	TeamId        int       `gorm:"column:team_id" json:"team_id" validate:"required"`
	CloudPlatform int       `gorm:"column:cloud_platform" json:"cloud_platform" validate:"required"` // 设备平台：0-未知, 1-飞燕
	Name          string    `gorm:"column:name" json:"name"`
	Projectid     string    `gorm:"column:projectid" json:"projectid"`
	ProdKey       string    `gorm:"column:prod_key" json:"prod_key"`
	ProdSecret    string    `gorm:"column:prod_secret" json:"prod_secret"`
	TestKey       string    `gorm:"column:test_key" json:"test_key"`
	TestSecret    string    `gorm:"column:test_secret" json:"test_secret"`
	SyncStatus    int       `gorm:"column:sync_status" json:"sync_status"` // 数据同步状态：0-关闭，1-打开
	SyncKey       string    `gorm:"column:sync_key" json:"sync_key"`
	SyncSecret    string    `gorm:"column:sync_secret" json:"sync_secret"`
	Mtime         time.Time `gorm:"column:mtime" json:"mtime"`
}

func (m *MxProject) TableNames() string {
	return "mx_project"
}

func (m *MxProject) Execute(req *MessInfo) {
	log.Info("Execute===>", m.TableNames())
}
