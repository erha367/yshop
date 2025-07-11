// Code generated by https://gotool.top
package model

// ProductCate
type ProductCate struct {
	Id        int32  `gorm:"column:id;primary_key;NOT NULL"`
	Level     int32  `gorm:"column:level;default:1"`
	ParentId  int32  `gorm:"column:parent_id;default:0"`
	Name      string `gorm:"column:name"`
	CreatedAt string `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
}

// TableName 表名
func (p *ProductCate) TableName() string {
	return "product_cate"
}
