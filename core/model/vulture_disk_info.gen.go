// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameVultureDiskInfo = "vulture_disk_info"

// VultureDiskInfo 硬盘信息
type VultureDiskInfo struct {
	ID              int64      `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement:true;comment:主键" json:"id"`                 // 主键
	Created         time.Time  `gorm:"column:created;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created"` // 创建时间
	Updated         time.Time  `gorm:"column:updated;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:修改时间" json:"updated"` // 修改时间
	Deleted         *time.Time `gorm:"column:deleted;type:timestamp;comment:删除时间(软删除标识)" json:"deleted"`                             // 删除时间(软删除标识)
	Creator         *string    `gorm:"column:creator;type:varchar(255);comment:修改人" json:"creator"`                                  // 修改人
	VultureDiskID   *int64     `gorm:"column:vulture_disk_id;type:bigint(20);comment:硬盘id" json:"vulture_disk_id"`                   // 硬盘id
	RefID           int64      `gorm:"column:ref_id;type:bigint(20);not null;comment:集群id" json:"ref_id"`                            // 集群id
	Host            *string    `gorm:"column:host;type:varchar(255);comment:主机ip" json:"host"`                                       // 主机ip
	RaidModel       *string    `gorm:"column:raid_model;type:varchar(255);comment:raid型号" json:"raid_model"`                         // raid型号
	Location        *string    `gorm:"column:location;type:varchar(255);comment:位置" json:"location"`                                 // 位置
	Slot            *string    `gorm:"column:slot;type:varchar(255);comment:插槽编号" json:"slot"`                                       // 插槽编号
	Manufacturer    *string    `gorm:"column:manufacturer;type:varchar(255);comment:厂商" json:"manufacturer"`                         // 厂商
	Model           *string    `gorm:"column:model;type:varchar(255);comment:型号" json:"model"`                                       // 型号
	Wwn             *string    `gorm:"column:wwn;type:varchar(255);comment:WWN" json:"wwn"`                                          // WWN
	SerialNumber    *string    `gorm:"column:serial_number;type:varchar(255);comment:序列号" json:"serial_number"`                      // 序列号
	BusType         *string    `gorm:"column:bus_type;type:varchar(255);comment:总线类型" json:"bus_type"`                               // 总线类型
	MediaType       *string    `gorm:"column:media_type;type:varchar(255);comment:媒体类型" json:"media_type"`                           // 媒体类型
	Size            *int64     `gorm:"column:size;type:bigint(20);comment:容量" json:"size"`                                           // 容量
	ErrorCount      *int32     `gorm:"column:error_count;type:int(11);comment:错误数" json:"error_count"`                               // 错误数
	PartNumber      *string    `gorm:"column:part_number;type:varchar(255);comment:部件号" json:"part_number"`                          // 部件号
	FirmwareVersion *string    `gorm:"column:firmware_version;type:varchar(255);comment:固件版本" json:"firmware_version"`               // 固件版本
	TransferSpeed   *string    `gorm:"column:transfer_speed;type:varchar(255);comment:传输速率" json:"transfer_speed"`                   // 传输速率
	FirmwareState   *string    `gorm:"column:firmware_state;type:varchar(255);comment:固件状态" json:"firmware_state"`                   // 固件状态
	ForeignState    *string    `gorm:"column:foreign_state;type:varchar(255);comment:掉线状态" json:"foreign_state"`                     // 掉线状态
	InquiryData     *string    `gorm:"column:inquiry_data;type:varchar(255);comment:诊断数据" json:"inquiry_data"`                       // 诊断数据
	ControllerID    *string    `gorm:"column:controller_id;type:varchar(255);comment:RAID控制器ID" json:"controller_id"`                // RAID控制器ID
	Light           *string    `gorm:"column:light;type:varchar(255);comment:硬盘指示灯状态 on/off" json:"light"`                           // 硬盘指示灯状态 on/off
	DiskLife        int32      `gorm:"column:disk_life;type:int(11);not null;comment:磁盘寿命" json:"disk_life"`                         // 磁盘寿命
	Path            string     `gorm:"column:path;type:varchar(255);not null;comment:路径" json:"path"`                                // 路径
	MountPoint      string     `gorm:"column:mount_point;type:varchar(255);not null;comment:挂载点" json:"mount_point"`                 // 挂载点
}

// TableName VultureDiskInfo's table name
func (*VultureDiskInfo) TableName() string {
	return TableNameVultureDiskInfo
}