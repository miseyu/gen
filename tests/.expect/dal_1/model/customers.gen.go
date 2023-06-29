// Code generated by github.com/miseyu/gen. DO NOT EDIT.
// Code generated by github.com/miseyu/gen. DO NOT EDIT.
// Code generated by github.com/miseyu/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCustomer = "customers"

// Customer mapped from table <customers>
type Customer struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	BankID    int64          `gorm:"column:bank_id" json:"bank_id"`
}

// TableName Customer's table name
func (*Customer) TableName() string {
	return TableNameCustomer
}
