// auto generated file
// DO NOT EDIT!

package model

import (
	"database/sql"
	"time"
)

type User struct {
	Uid      int32     `gorm:"primary_key"`
	UpdateAt time.Time `sql:"DEFAULT:current_timestamp"`
	CreateAt time.Time `sql:"DEFAULT:current_timestamp"`
	Phone    sql.NullString
	Password sql.NullString
	Name     sql.NullString
}

func (User) TableName() string {
	return "user"
}
