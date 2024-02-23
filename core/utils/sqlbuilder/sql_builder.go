package sqlbuilder

import (
	"context"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// SQLBuilder 构建SQL语句
type SQLBuilder struct {
	session *gorm.DB
}

// NewSQLBuilder 新建SQLBuilder
func NewSQLBuilder(ctx context.Context, session *gorm.DB) *SQLBuilder {
	return &SQLBuilder{
		session: session.WithContext(ctx),
	}
}

// AndStringLikeIn 添加string类型的Like查询条件
func (sb *SQLBuilder) AndStringLikeIn(field string, value ...string) *SQLBuilder {
	if len(value) <= 0 {
		return sb
	}
	for _, item := range value {
		sb.session.Where(fmt.Sprintf("%s like ?", field), fmt.Sprintf("%%%s%%", item))
	}
	return sb
}

// AndStringInLike 添加string类型的Like查询条件
// select * from auth where name like
func (sb *SQLBuilder) AndStringInLike(value string, field ...string) *SQLBuilder {
	if len(field) <= 0 {
		return sb
	}
	if len(value) <= 0 {
		return sb
	}
	for i, item := range field {
		if i == 0 {
			sb.session.Where(fmt.Sprintf("%s like ?", item), fmt.Sprintf("%%%s%%", value))
			continue
		}
		sb.session.Or(fmt.Sprintf("%s like ?", item), fmt.Sprintf("%%%s%%", value))
	}
	return sb
}

// AndStringLike 添加string类型的Like查询条件
func (sb *SQLBuilder) AndStringLike(field, value string) *SQLBuilder {
	if len(value) <= 0 {
		return sb
	}
	sb.session.Where(fmt.Sprintf("%s like ?", field), fmt.Sprintf("%%%s%%", value))
	return sb
}

// AndStringLeftLike 添加string类型的左Like查询条件
func (sb *SQLBuilder) AndStringLeftLike(field, value string) *SQLBuilder {
	if len(value) <= 0 {
		return sb
	}
	sb.session.Where(fmt.Sprintf("%s like ?", field), fmt.Sprintf("%%%s", value))
	return sb
}

// AndStringRightLike 添加string类型的右Like查询条件
func (sb *SQLBuilder) AndStringRightLike(field, value string) *SQLBuilder {
	if len(value) <= 0 {
		return sb
	}
	sb.session.Where(fmt.Sprintf("%s like ?", field), fmt.Sprintf("%s%%", value))
	return sb
}

// AndStringEQ 添加string类型的相等查询条件
func (sb *SQLBuilder) AndStringEQ(field, value string) *SQLBuilder {
	if len(value) <= 0 {
		return sb
	}
	sb.session.Where(fmt.Sprintf("%s = ?", field), value)
	return sb
}

// AndStringNotEQ 添加string类型的相等查询条件
func (sb *SQLBuilder) AndStringNotEQ(field, value string) *SQLBuilder {
	if len(value) <= 0 {
		return sb
	}
	sb.session.Where(fmt.Sprintf("%s != ?", field), value)
	return sb
}

// AndStringRegex 添加string类型的正则查询条件
func (sb *SQLBuilder) AndStringRegex(field, value string) *SQLBuilder {
	if len(value) <= 0 {
		return sb
	}
	sb.session.Where(fmt.Sprintf("%s REGEXP ?", field), value)
	return sb
}

// AndStringIn 添加string类型的相等查询条件
func (sb *SQLBuilder) AndStringIn(field string, value []string) *SQLBuilder {
	if len(value) <= 0 {
		return sb
	}
	sb.session.Where(fmt.Sprintf("%s in ?", field), value)
	return sb
}

// AndIntEQ 添加Int类型的相等查询条件
func (sb *SQLBuilder) AndIntEQ(field string, value int) *SQLBuilder {
	if value <= 0 {
		return sb
	}
	sb.session.Where(fmt.Sprintf("%s = ?", field), value)
	return sb
}

// AndIntGT 添加Int类型的大于查询条件
func (sb *SQLBuilder) AndIntGT(field string, value int) *SQLBuilder {
	sb.session.Where(fmt.Sprintf("%s > ?", field), value)
	return sb
}

// AndIntGTE 添加Int类型的大于等于查询条件
func (sb *SQLBuilder) AndIntGTE(field string, value int) *SQLBuilder {
	sb.session.Where(fmt.Sprintf("%s >= ?", field), value)
	return sb
}

// AndIntLT 添加Int类型的小于查询条件
func (sb *SQLBuilder) AndIntLT(field string, value int) *SQLBuilder {
	sb.session.Where(fmt.Sprintf("%s < ?", field), value)
	return sb
}

// AndIntLTE 添加Int类型的小于等于查询条件
func (sb *SQLBuilder) AndIntLTE(field string, value int) *SQLBuilder {
	sb.session.Where(fmt.Sprintf("%s <= ?", field), value)
	return sb
}

// AndIntIn 添加Int类型的相等查询条件
func (sb *SQLBuilder) AndIntIn(field string, value []int64) *SQLBuilder {
	if len(value) <= 0 {
		return sb
	}
	sb.session.Where(fmt.Sprintf("%s in (?)", field), value)
	return sb
}

// OrStringLikeIn 添加string类型的Like查询条件
func (sb *SQLBuilder) OrStringLikeIn(field string, value ...string) *SQLBuilder {
	if len(value) <= 0 {
		return sb
	}
	for _, item := range value {
		sb.session.Or(fmt.Sprintf("%s like ?", field), fmt.Sprintf("%%%s%%", item))
	}
	return sb
}

// OrStringLike 添加string类型的Like查询条件
func (sb *SQLBuilder) OrStringLike(field, value string) *SQLBuilder {
	if len(value) <= 0 {
		return sb
	}
	sb.session.Or(fmt.Sprintf("%s like ?", field), fmt.Sprintf("%%%s%%", value))
	return sb
}

// OrStringLeftLike 添加string类型的左Like查询条件
func (sb *SQLBuilder) OrStringLeftLike(field, value string) *SQLBuilder {
	if len(value) <= 0 {
		return sb
	}
	sb.session.Or(fmt.Sprintf("%s like ?", field), fmt.Sprintf("%%%s", value))
	return sb
}

// OrStringRightLike 添加string类型的右Like查询条件
func (sb *SQLBuilder) OrStringRightLike(field, value string) *SQLBuilder {
	if len(value) <= 0 {
		return sb
	}
	sb.session.Or(fmt.Sprintf("%s like ?", field), fmt.Sprintf("%s%%", value))
	return sb
}

// OrStringEQ 添加string类型的相等查询条件
func (sb *SQLBuilder) OrStringEQ(field, value string) *SQLBuilder {
	if len(value) <= 0 {
		return sb
	}
	sb.session.Or(fmt.Sprintf("%s = ?", field), value)
	return sb
}

// OrStringIn 添加string类型的相等查询条件
func (sb *SQLBuilder) OrStringIn(field string, value []string) *SQLBuilder {
	if len(value) <= 0 {
		return sb
	}
	sb.session.Or(fmt.Sprintf("%s in (?)", field), value)
	return sb
}

// OrIntEQ 添加Int类型的相等查询条件
func (sb *SQLBuilder) OrIntEQ(field string, value int) *SQLBuilder {
	sb.session.Or(fmt.Sprintf("%s = ?", field), value)
	return sb
}

// OrIntGT 添加Int类型的大于查询条件
func (sb *SQLBuilder) OrIntGT(field string, value int) *SQLBuilder {
	sb.session.Or(fmt.Sprintf("%s > ?", field), value)
	return sb
}

// OrIntGTE 添加Int类型的大于等于查询条件
func (sb *SQLBuilder) OrIntGTE(field string, value int) *SQLBuilder {
	sb.session.Or(fmt.Sprintf("%s >= ?", field), value)
	return sb
}

// OrIntLT 添加Int类型的小于查询条件
func (sb *SQLBuilder) OrIntLT(field string, value int) *SQLBuilder {
	sb.session.Or(fmt.Sprintf("%s < ?", field), value)
	return sb
}

// OrIntLTE 添加Int类型的小于等于查询条件
func (sb *SQLBuilder) OrIntLTE(field string, value int) *SQLBuilder {
	sb.session.Or(fmt.Sprintf("%s <= ?", field), value)
	return sb
}

// OrIntIn 添加Int类型的相等查询条件
func (sb *SQLBuilder) OrIntIn(field string, value []int) *SQLBuilder {
	if len(value) <= 0 {
		return sb
	}
	sb.session.Or(fmt.Sprintf("%s in (?)", field), value)
	return sb
}

// OrderDesc 添加倒序
func (sb *SQLBuilder) OrderDesc(name ...string) *SQLBuilder {
	sb.session.Order(fmt.Sprintf("%s desc", strings.Join(name, ",")))
	return sb
}

// OrderAsc 添加正序
func (sb *SQLBuilder) OrderAsc(name ...string) *SQLBuilder {
	sb.session.Order(fmt.Sprintf("%s asc", strings.Join(name, ",")))
	return sb
}

// Limit 添加分页
func (sb *SQLBuilder) Limit(limit int) *SQLBuilder {
	sb.session.Limit(limit)
	return sb
}

// Offset 分页偏移
func (sb *SQLBuilder) Offset(offset int) *SQLBuilder {
	sb.session.Offset(offset)
	return sb
}

// AndTimeGT 添加Int类型的大于查询条件
func (sb *SQLBuilder) AndTimeGT(field string, value string) *SQLBuilder {
	if len(value) <= 0 {
		return sb
	}
	sb.session.Where(fmt.Sprintf("%s > ?", field), value)
	return sb
}

// AndTimeLT 添加Int类型的小于查询条件
func (sb *SQLBuilder) AndTimeLT(field string, value string) *SQLBuilder {
	if len(value) <= 0 {
		return sb
	}
	sb.session.Where(fmt.Sprintf("%s < ?", field), value)
	return sb
}

// ToSession 转换session
func (sb *SQLBuilder) ToSession() *gorm.DB {
	return sb.session
}
