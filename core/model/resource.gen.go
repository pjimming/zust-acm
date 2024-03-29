// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameResource = "resource"

// Resource 资源表
type Resource struct {
	ID        int64          `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true;comment:序号" json:"id"`             // 序号
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间" json:"deleted_at"`                                    // 删除时间
	Name      string         `gorm:"column:name;type:varchar(255);not null;comment:名称" json:"name"`                                     // 名称
	Code      string         `gorm:"column:code;type:varchar(255);not null;comment:编码" json:"code"`                                     // 编码
	Type      string         `gorm:"column:type;type:varchar(255);not null;comment:类型" json:"type"`                                     // 类型
	ParentID  int64          `gorm:"column:parent_id;type:bigint(20) unsigned;not null;comment:父节点id" json:"parent_id"`                 // 父节点id
	Order     int32          `gorm:"column:order;type:int(10);not null;comment:排序" json:"order"`                                        // 排序
	Icon      string         `gorm:"column:icon;type:varchar(512);not null;comment:菜单图标" json:"icon"`                                   // 菜单图标
	Component string         `gorm:"column:component;type:varchar(512);not null;comment:组件路径" json:"component"`                         // 组件路径
	Path      string         `gorm:"column:path;type:varchar(512);not null;comment:路由地址" json:"path"`                                   // 路由地址
	IsShow    bool           `gorm:"column:is_show;type:tinyint(1);not null;comment:是否显示" json:"is_show"`                               // 是否显示
	IsEnable  bool           `gorm:"column:is_enable;type:tinyint(1);not null;comment:是否启用" json:"is_enable"`                           // 是否启用
}

// TableName Resource's table name
func (*Resource) TableName() string {
	return TableNameResource
}
