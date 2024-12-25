package model

import (
	"time"
)

// Monitor 监控目标模型
type Monitor struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	URL       string    `gorm:"size:255;not null" json:"url"`
	Status    int8      `gorm:"default:1" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MonitorRule 监控规则模型
type MonitorRule struct {
	ID         uint64    `gorm:"primaryKey" json:"id"`
	MonitorID  uint64    `json:"monitor_id"`
	RuleType   string    `gorm:"size:50" json:"rule_type"`
	RuleConfig string    `gorm:"type:json" json:"rule_config"`
	Schedule   string    `gorm:"size:100" json:"schedule"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Notification 通知配置模型
type Notification struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Type      string    `gorm:"size:50" json:"type"`
	Config    string    `gorm:"type:json" json:"config"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
