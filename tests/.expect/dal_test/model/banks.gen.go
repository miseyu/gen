// Code generated by github.com/miseyu/gen. DO NOT EDIT.
// Code generated by github.com/miseyu/gen. DO NOT EDIT.
// Code generated by github.com/miseyu/gen. DO NOT EDIT.

package model

const TableNameBank = "banks"

// Bank mapped from table <banks>
type Bank struct {
	ID      int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name    string `gorm:"column:name" json:"name"`
	Address string `gorm:"column:address" json:"address"`
	Scale   int64  `gorm:"column:scale" json:"scale"`
}

// TableName Bank's table name
func (*Bank) TableName() string {
	return TableNameBank
}
