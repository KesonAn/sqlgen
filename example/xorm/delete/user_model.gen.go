// Code generated by sqlgen. DO NOT EDIT!

package model

import (
	"context"
	"fmt"
	"time"

	"xorm.io/xorm"
)

// UserModel represents a user model.
type UserModel struct {
	engine xorm.EngineInterface
}

// User represents a user struct data.
type User struct {
	Id         uint64    `xorm:"pk autoincr 'id'" json:"id"`
	Name       string    `xorm:"'name'" json:"name"`
	Password   string    `xorm:"'password'" json:"password"`
	Mobile     string    `xorm:"'mobile'" json:"mobile"`
	Gender     string    `xorm:"'gender'" json:"gender"`
	Nickname   string    `xorm:"'nickname'" json:"nickname"`
	Type       int8      `xorm:"'type'" json:"type"`
	CreateTime time.Time `xorm:"'create_time'" json:"createTime"`
	UpdateTime time.Time `xorm:"'update_time'" json:"updateTime"`
}

// DeleteWhereParameter is a where parameter structure.
type DeleteWhereParameter struct {
	Id uint64
}

// DeleteByNameWhereParameter is a where parameter structure.
type DeleteByNameWhereParameter struct {
	Name string
}

// DeleteByNameAndMobileWhereParameter is a where parameter structure.
type DeleteByNameAndMobileWhereParameter struct {
	Name   string
	Mobile string
}

func (User) TableName() string {
	return "user"
}

// NewUserModel returns a new user model.
func NewUserModel(engine xorm.EngineInterface) *UserModel {
	return &UserModel{engine: engine}
}

// Insert creates  user data.
func (m *UserModel) Insert(ctx context.Context, data ...*User) error {
	if len(data) == 0 {
		return fmt.Errorf("data is empty")
	}

	var session = m.engine.Context(ctx)
	var list []User
	for _, v := range data {
		list = append(list, *v)
	}

	_, err := session.Insert(&list)
	return err
}

// Delete is generated from sql:
// delete from user where id = ?;
func (m *UserModel) Delete(ctx context.Context, where DeleteWhereParameter) error {
	var session = m.engine.Context(ctx)
	session.Where(`id = ?`, where.Id)
	_, err := session.Delete(&User{})
	return err
}

// DeleteByName is generated from sql:
// delete from user where name = ?;
func (m *UserModel) DeleteByName(ctx context.Context, where DeleteByNameWhereParameter) error {
	var session = m.engine.Context(ctx)
	session.Where(`name = ?`, where.Name)
	_, err := session.Delete(&User{})
	return err
}

// DeleteByNameAndMobile is generated from sql:
// delete from user where name = ? and mobile = ?;
func (m *UserModel) DeleteByNameAndMobile(ctx context.Context, where DeleteByNameAndMobileWhereParameter) error {
	var session = m.engine.Context(ctx)
	session.Where(`name = ? AND mobile = ?`, where.Name, where.Mobile)
	_, err := session.Delete(&User{})
	return err
}
