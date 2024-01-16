// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCompetition = "competition"

// Competition 比赛信息表
type Competition struct {
	ID         int64          `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true;comment:序号" json:"id"`             // 序号
	CreatedAt  time.Time      `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间" json:"deleted_at"`                                    // 删除时间
	Name       string         `gorm:"column:name;type:varchar(512);not null;comment:比赛名称" json:"name"`                                   // 比赛名称
	Type       string         `gorm:"column:type;type:varchar(512);not null;comment:比赛类型" json:"type"`                                   // 比赛类型
	SeasonYear int32          `gorm:"column:season_year;type:int(11);not null;comment:赛季年" json:"season_year"`                           // 赛季年
	StartTime  time.Time      `gorm:"column:start_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:开始时间" json:"start_time"` // 开始时间
	EndTime    time.Time      `gorm:"column:end_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:结束时间" json:"end_time"`     // 结束时间
}

// TableName Competition's table name
func (*Competition) TableName() string {
	return TableNameCompetition
}
