package message

import (
	"gofer/pkg/log"
	"time"
)

// 项目表（飞燕项目等）
type MxProject struct {
	Id            int        `gorm:"column:id;primary_key"`
	TeamId        int        `gorm:"column:team_id"`
	CloudPlatform int        `gorm:"column:cloud_platform"` // 设备平台：0-未知, 1-飞燕
	Name          string     `gorm:"column:name"`
	Projectid     string     `gorm:"column:projectid"`
	ProdKey       string     `gorm:"column:prod_key"`
	ProdSecret    string     `gorm:"column:prod_secret"`
	TestKey       string     `gorm:"column:test_key"`
	TestSecret    string     `gorm:"column:test_secret"`
	SyncStatus    int        `gorm:"column:sync_status"` // 数据同步状态：0-关闭，1-打开
	SyncKey       string     `gorm:"column:sync_key"`
	SyncSecret    string     `gorm:"column:sync_secret"`
	Mtime         time.Time `gorm:"column:mtime"`
}

func (m *MxProject) TableNames() string {
	return "mx_project"
}


func (m *MxProject)Execute()  {
	log.Info("Execute===>",m.TableNames())
}