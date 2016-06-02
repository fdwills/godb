// auto generated file
// DO NOT EDIT!

package model

import (
	"time"
)

type UserExtra struct {
	Key2     int32
	Key1     int32
	UpdateAt time.Time `sql:"DEFAULT:current_timestamp"`
	Uid      int32     `gorm:"primary_key"`
	CreateAt time.Time `sql:"DEFAULT:current_timestamp"`
}

func (UserExtra) TableName() string {
	return "user_extra"
}
