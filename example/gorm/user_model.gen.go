// Code generated by sqlgen. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/shopspring/decimal"
)

// UserModel represents a user model.
type UserModel struct {
	db *gorm.DB
}

// User represents a user struct data.
type User struct {
	Id       uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name     string    `gorm:"column:name" json:"name"`
	Password string    `gorm:"column:password" json:"password"`
	Mobile   string    `gorm:"column:mobile" json:"mobile"`
	Gender   string    `gorm:"column:gender" json:"gender"`
	Nickname string    `gorm:"column:nickname" json:"nickname"`
	Type     int8      `gorm:"column:type" json:"type"`
	CreateAt time.Time `gorm:"column:create_at" json:"createAt"`
	UpdateAt time.Time `gorm:"column:update_at" json:"updateAt"`
}

// FindOneWhereParameter is a where parameter structure.
type FindOneWhereParameter struct {
	IdEqual uint64
}

// FindOneByNameWhereParameter is a where parameter structure.
type FindOneByNameWhereParameter struct {
	NameEqual string
}

// FindOneGroupByNameWhereParameter is a where parameter structure.
type FindOneGroupByNameWhereParameter struct {
	NameEqual string
}

// FindOneGroupByNameHavingNameWhereParameter is a where parameter structure.
type FindOneGroupByNameHavingNameWhereParameter struct {
	NameEqual string
}

// FindOneGroupByNameHavingNameHavingParameter is a having parameter structure.
type FindOneGroupByNameHavingNameHavingParameter struct {
	NameEqual string
}

// FindLimitWhereParameter is a where parameter structure.
type FindLimitWhereParameter struct {
	IdGT uint64
}

// FindLimitLimitParameter is a limit parameter structure.
type FindLimitLimitParameter struct {
	Count int
}

// FindLimitOffsetLimitParameter is a limit parameter structure.
type FindLimitOffsetLimitParameter struct {
	Count  int
	Offset int
}

// FindGroupLimitOffsetWhereParameter is a where parameter structure.
type FindGroupLimitOffsetWhereParameter struct {
	IdGT uint64
}

// FindGroupLimitOffsetLimitParameter is a limit parameter structure.
type FindGroupLimitOffsetLimitParameter struct {
	Count  int
	Offset int
}

// FindGroupHavingLimitOffsetWhereParameter is a where parameter structure.
type FindGroupHavingLimitOffsetWhereParameter struct {
	IdGT uint64
}

// FindGroupHavingLimitOffsetHavingParameter is a having parameter structure.
type FindGroupHavingLimitOffsetHavingParameter struct {
	IdGT uint64
}

// FindGroupHavingLimitOffsetLimitParameter is a limit parameter structure.
type FindGroupHavingLimitOffsetLimitParameter struct {
	Count  int
	Offset int
}

// FindGroupHavingOrderAscLimitOffsetWhereParameter is a where parameter structure.
type FindGroupHavingOrderAscLimitOffsetWhereParameter struct {
	IdGT uint64
}

// FindGroupHavingOrderAscLimitOffsetHavingParameter is a having parameter structure.
type FindGroupHavingOrderAscLimitOffsetHavingParameter struct {
	IdGT uint64
}

// FindGroupHavingOrderAscLimitOffsetLimitParameter is a limit parameter structure.
type FindGroupHavingOrderAscLimitOffsetLimitParameter struct {
	Count  int
	Offset int
}

// FindGroupHavingOrderDescLimitOffsetWhereParameter is a where parameter structure.
type FindGroupHavingOrderDescLimitOffsetWhereParameter struct {
	IdGT uint64
}

// FindGroupHavingOrderDescLimitOffsetHavingParameter is a having parameter structure.
type FindGroupHavingOrderDescLimitOffsetHavingParameter struct {
	IdGT uint64
}

// FindGroupHavingOrderDescLimitOffsetLimitParameter is a limit parameter structure.
type FindGroupHavingOrderDescLimitOffsetLimitParameter struct {
	Count  int
	Offset int
}

// FindOnePartWhereParameter is a where parameter structure.
type FindOnePartWhereParameter struct {
	IdGT uint64
}

// FindAllCountResult is a find all count result.
type FindAllCountResult struct {
	CountID sql.NullInt64 `gorm:"column:countID" json:"countID"`
}

// TableName returns the table name. it implemented by gorm.Tabler.
func (FindAllCountResult) TableName() string {
	return "user"
}

// FindAllCountWhereWhereParameter is a where parameter structure.
type FindAllCountWhereWhereParameter struct {
	IdGT uint64
}

// FindAllCountWhereResult is a find all count where result.
type FindAllCountWhereResult struct {
	CountID sql.NullInt64 `gorm:"column:countID" json:"countID"`
}

// TableName returns the table name. it implemented by gorm.Tabler.
func (FindAllCountWhereResult) TableName() string {
	return "user"
}

// FindMaxIDResult is a find max id result.
type FindMaxIDResult struct {
	MaxID sql.NullInt64 `gorm:"column:maxID" json:"maxID"`
}

// TableName returns the table name. it implemented by gorm.Tabler.
func (FindMaxIDResult) TableName() string {
	return "user"
}

// FindMinIDResult is a find min id result.
type FindMinIDResult struct {
	MinID sql.NullInt64 `gorm:"column:minID" json:"minID"`
}

// TableName returns the table name. it implemented by gorm.Tabler.
func (FindMinIDResult) TableName() string {
	return "user"
}

// FindAvgIDResult is a find avg id result.
type FindAvgIDResult struct {
	AvgID decimal.NullDecimal `gorm:"column:avgID" json:"avgID"`
}

// TableName returns the table name. it implemented by gorm.Tabler.
func (FindAvgIDResult) TableName() string {
	return "user"
}

// UpdateWhereParameter is a where parameter structure.
type UpdateWhereParameter struct {
	IdEqual uint64
}

// UpdateOrderByIdDescWhereParameter is a where parameter structure.
type UpdateOrderByIdDescWhereParameter struct {
	IdEqual uint64
}

// UpdateOrderByIdDescLimitCountWhereParameter is a where parameter structure.
type UpdateOrderByIdDescLimitCountWhereParameter struct {
	IdEqual uint64
}

// UpdateOrderByIdDescLimitCountLimitParameter is a limit parameter structure.
type UpdateOrderByIdDescLimitCountLimitParameter struct {
	Count int
}

// DeleteOneWhereParameter is a where parameter structure.
type DeleteOneWhereParameter struct {
	IdEqual uint64
}

// DeleteOneByNameWhereParameter is a where parameter structure.
type DeleteOneByNameWhereParameter struct {
	NameEqual string
}

// DeleteOneOrderByIDAscWhereParameter is a where parameter structure.
type DeleteOneOrderByIDAscWhereParameter struct {
	NameEqual string
}

// DeleteOneOrderByIDDescWhereParameter is a where parameter structure.
type DeleteOneOrderByIDDescWhereParameter struct {
	NameEqual string
}

// DeleteOneOrderByIDDescLimitCountWhereParameter is a where parameter structure.
type DeleteOneOrderByIDDescLimitCountWhereParameter struct {
	NameEqual string
}

// DeleteOneOrderByIDDescLimitCountLimitParameter is a limit parameter structure.
type DeleteOneOrderByIDDescLimitCountLimitParameter struct {
	Count int
}

// TableName returns the table name. it implemented by gorm.Tabler.
func (User) TableName() string {
	return "user"
}

// NewUserModel returns a new user model.
func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{db: db}
}

// Create creates  user data.
func (m *UserModel) Create(ctx context.Context, data ...*User) error {
	if len(data) == 0 {
		return fmt.Errorf("data is empty")
	}

	db := m.db.WithContext(ctx)
	list := data[:]
	return db.Create(&list).Error
}

// FindOne is generated from sql:
// select * from `user` where `id` = ? limit 1;
func (m *UserModel) FindOne(ctx context.Context, where FindOneWhereParameter) (*User, error) {
	var result = new(User)
	var db = m.db.WithContext(ctx)
	db = db.Select(`*`)
	db = db.Where(`id = ?`, where.IdEqual)
	db = db.Limit(1)
	db = db.Take(result)
	return result, db.Error
}

// FindOneByName is generated from sql:
// select * from `user` where `name` = ? limit 1;
func (m *UserModel) FindOneByName(ctx context.Context, where FindOneByNameWhereParameter) (*User, error) {
	var result = new(User)
	var db = m.db.WithContext(ctx)
	db = db.Select(`*`)
	db = db.Where(`name = ?`, where.NameEqual)
	db = db.Limit(1)
	db = db.Take(result)
	return result, db.Error
}

// FindOneGroupByName is generated from sql:
// select * from `user` where `name` = ? group by name limit 1;
func (m *UserModel) FindOneGroupByName(ctx context.Context, where FindOneGroupByNameWhereParameter) (*User, error) {
	var result = new(User)
	var db = m.db.WithContext(ctx)
	db = db.Select(`*`)
	db = db.Where(`name = ?`, where.NameEqual)
	db = db.Group(`name`)
	db = db.Limit(1)
	db = db.Take(result)
	return result, db.Error
}

// FindOneGroupByNameHavingName is generated from sql:
// select * from `user` where `name` = ? group by name having name = ? limit 1;
func (m *UserModel) FindOneGroupByNameHavingName(ctx context.Context, where FindOneGroupByNameHavingNameWhereParameter, having FindOneGroupByNameHavingNameHavingParameter) (*User, error) {
	var result = new(User)
	var db = m.db.WithContext(ctx)
	db = db.Select(`*`)
	db = db.Where(`name = ?`, where.NameEqual)
	db = db.Group(`name`)
	db = db.Having(`name = ?`, having.NameEqual)
	db = db.Limit(1)
	db = db.Take(result)
	return result, db.Error
}

// FindAll is generated from sql:
// select * from `user`;
func (m *UserModel) FindAll(ctx context.Context) ([]*User, error) {
	var result []*User
	var db = m.db.WithContext(ctx)
	db = db.Select(`*`)
	db = db.Find(&result)
	return result, db.Error
}

// FindLimit is generated from sql:
// select * from `user` where id > ? limit ?;
func (m *UserModel) FindLimit(ctx context.Context, where FindLimitWhereParameter, limit FindLimitLimitParameter) ([]*User, error) {
	var result []*User
	var db = m.db.WithContext(ctx)
	db = db.Select(`*`)
	db = db.Where(`id > ?`, where.IdGT)
	db = db.Limit(limit.Count)
	db = db.Find(&result)
	return result, db.Error
}

// FindLimitOffset is generated from sql:
// select * from `user` limit ?, ?;
func (m *UserModel) FindLimitOffset(ctx context.Context, limit FindLimitOffsetLimitParameter) ([]*User, error) {
	var result []*User
	var db = m.db.WithContext(ctx)
	db = db.Select(`*`)
	db = db.Offset(limit.Offset).Limit(limit.Count)
	db = db.Find(&result)
	return result, db.Error
}

// FindGroupLimitOffset is generated from sql:
// select * from `user` where id > ? group by name limit ?, ?;
func (m *UserModel) FindGroupLimitOffset(ctx context.Context, where FindGroupLimitOffsetWhereParameter, limit FindGroupLimitOffsetLimitParameter) ([]*User, error) {
	var result []*User
	var db = m.db.WithContext(ctx)
	db = db.Select(`*`)
	db = db.Where(`id > ?`, where.IdGT)
	db = db.Group(`name`)
	db = db.Offset(limit.Offset).Limit(limit.Count)
	db = db.Find(&result)
	return result, db.Error
}

// FindGroupHavingLimitOffset is generated from sql:
// select * from `user` where id > ? group by name having id > ? limit ?, ?;
func (m *UserModel) FindGroupHavingLimitOffset(ctx context.Context, where FindGroupHavingLimitOffsetWhereParameter, having FindGroupHavingLimitOffsetHavingParameter, limit FindGroupHavingLimitOffsetLimitParameter) ([]*User, error) {
	var result []*User
	var db = m.db.WithContext(ctx)
	db = db.Select(`*`)
	db = db.Where(`id > ?`, where.IdGT)
	db = db.Group(`name`)
	db = db.Having(`id > ?`, having.IdGT)
	db = db.Offset(limit.Offset).Limit(limit.Count)
	db = db.Find(&result)
	return result, db.Error
}

// FindGroupHavingOrderAscLimitOffset is generated from sql:
// select * from `user` where id > ? group by name having id > ? order by id limit ?, ?;
func (m *UserModel) FindGroupHavingOrderAscLimitOffset(ctx context.Context, where FindGroupHavingOrderAscLimitOffsetWhereParameter, having FindGroupHavingOrderAscLimitOffsetHavingParameter, limit FindGroupHavingOrderAscLimitOffsetLimitParameter) ([]*User, error) {
	var result []*User
	var db = m.db.WithContext(ctx)
	db = db.Select(`*`)
	db = db.Where(`id > ?`, where.IdGT)
	db = db.Group(`name`)
	db = db.Having(`id > ?`, having.IdGT)
	db = db.Order(`id`)
	db = db.Offset(limit.Offset).Limit(limit.Count)
	db = db.Find(&result)
	return result, db.Error
}

// FindGroupHavingOrderDescLimitOffset is generated from sql:
// select * from `user` where id > ? group by name having id > ? order by id desc limit ?, ?;
func (m *UserModel) FindGroupHavingOrderDescLimitOffset(ctx context.Context, where FindGroupHavingOrderDescLimitOffsetWhereParameter, having FindGroupHavingOrderDescLimitOffsetHavingParameter, limit FindGroupHavingOrderDescLimitOffsetLimitParameter) ([]*User, error) {
	var result []*User
	var db = m.db.WithContext(ctx)
	db = db.Select(`*`)
	db = db.Where(`id > ?`, where.IdGT)
	db = db.Group(`name`)
	db = db.Having(`id > ?`, having.IdGT)
	db = db.Order(`id desc`)
	db = db.Offset(limit.Offset).Limit(limit.Count)
	db = db.Find(&result)
	return result, db.Error
}

// FindOnePart is generated from sql:
// select `name`, `password`, `mobile` from `user` where id > ? limit 1;
func (m *UserModel) FindOnePart(ctx context.Context, where FindOnePartWhereParameter) (*User, error) {
	var result = new(User)
	var db = m.db.WithContext(ctx)
	db = db.Select(`name, password, mobile`)
	db = db.Where(`id > ?`, where.IdGT)
	db = db.Limit(1)
	db = db.Take(result)
	return result, db.Error
}

// FindAllCount is generated from sql:
// select count(id) AS countID from `user`;
func (m *UserModel) FindAllCount(ctx context.Context) (*FindAllCountResult, error) {
	var result = new(FindAllCountResult)
	var db = m.db.WithContext(ctx)
	db = db.Select(`count(id) AS countID`)
	db = db.Limit(1)
	db = db.Take(result)
	return result, db.Error
}

// FindAllCountWhere is generated from sql:
// select count(id) AS countID from `user` where id > ?;
func (m *UserModel) FindAllCountWhere(ctx context.Context, where FindAllCountWhereWhereParameter) (*FindAllCountWhereResult, error) {
	var result = new(FindAllCountWhereResult)
	var db = m.db.WithContext(ctx)
	db = db.Select(`count(id) AS countID`)
	db = db.Where(`id > ?`, where.IdGT)
	db = db.Limit(1)
	db = db.Take(result)
	return result, db.Error
}

// FindMaxID is generated from sql:
// select max(id) AS maxID from `user`;
func (m *UserModel) FindMaxID(ctx context.Context) (*FindMaxIDResult, error) {
	var result = new(FindMaxIDResult)
	var db = m.db.WithContext(ctx)
	db = db.Select(`max(id) AS maxID`)
	db = db.Limit(1)
	db = db.Take(result)
	return result, db.Error
}

// FindMinID is generated from sql:
// select min(id) AS minID from `user`;
func (m *UserModel) FindMinID(ctx context.Context) (*FindMinIDResult, error) {
	var result = new(FindMinIDResult)
	var db = m.db.WithContext(ctx)
	db = db.Select(`min(id) AS minID`)
	db = db.Limit(1)
	db = db.Take(result)
	return result, db.Error
}

// FindAvgID is generated from sql:
// select avg(id) AS avgID from `user`;
func (m *UserModel) FindAvgID(ctx context.Context) (*FindAvgIDResult, error) {
	var result = new(FindAvgIDResult)
	var db = m.db.WithContext(ctx)
	db = db.Select(`avg(id) AS avgID`)
	db = db.Limit(1)
	db = db.Take(result)
	return result, db.Error
}

// Update is generated from sql:
// update `user` set `name` = ?, `password` = ?, `mobile` = ?, `gender` = ?, `nickname` = ?, `type` = ?, `create_at` = ?, `update_at` = ? where `id` = ?;
func (m *UserModel) Update(ctx context.Context, data *User, where UpdateWhereParameter) error {
	var db = m.db.WithContext(ctx)
	db = db.Model(&User{})
	db = db.Where(`id = ?`, where.IdEqual)
	db = db.Updates(map[string]interface{}{
		"name":      data.Name,
		"password":  data.Password,
		"mobile":    data.Mobile,
		"gender":    data.Gender,
		"nickname":  data.Nickname,
		"type":      data.Type,
		"create_at": data.CreateAt,
		"update_at": data.UpdateAt,
	})
	return db.Error
}

// UpdateOrderByIdDesc is generated from sql:
// update `user` set `name` = ?, `password` = ?, `mobile` = ?, `gender` = ?, `nickname` = ?, `type` = ?, `create_at` = ?, `update_at` = ? where `id` = ? order by id desc;
func (m *UserModel) UpdateOrderByIdDesc(ctx context.Context, data *User, where UpdateOrderByIdDescWhereParameter) error {
	var db = m.db.WithContext(ctx)
	db = db.Model(&User{})
	db = db.Where(`id = ?`, where.IdEqual)
	db = db.Order(`id desc`)
	db = db.Updates(map[string]interface{}{
		"name":      data.Name,
		"password":  data.Password,
		"mobile":    data.Mobile,
		"gender":    data.Gender,
		"nickname":  data.Nickname,
		"type":      data.Type,
		"create_at": data.CreateAt,
		"update_at": data.UpdateAt,
	})
	return db.Error
}

// UpdateOrderByIdDescLimitCount is generated from sql:
// update `user` set `name` = ?, `password` = ?, `mobile` = ?, `gender` = ?, `nickname` = ?, `type` = ?, `create_at` = ?, `update_at` = ? where `id` = ? order by id desc limit ?;
func (m *UserModel) UpdateOrderByIdDescLimitCount(ctx context.Context, data *User, where UpdateOrderByIdDescLimitCountWhereParameter, limit UpdateOrderByIdDescLimitCountLimitParameter) error {
	var db = m.db.WithContext(ctx)
	db = db.Model(&User{})
	db = db.Where(`id = ?`, where.IdEqual)
	db = db.Order(`id desc`)
	db = db.Limit(limit.Count)
	db = db.Updates(map[string]interface{}{
		"name":      data.Name,
		"password":  data.Password,
		"mobile":    data.Mobile,
		"gender":    data.Gender,
		"nickname":  data.Nickname,
		"type":      data.Type,
		"create_at": data.CreateAt,
		"update_at": data.UpdateAt,
	})
	return db.Error
}

// DeleteOne is generated from sql:
// delete from `user` where `id` = ?;
func (m *UserModel) DeleteOne(ctx context.Context, where DeleteOneWhereParameter) error {
	var db = m.db.WithContext(ctx)
	db = db.Where(`id = ?`, where.IdEqual)
	db = db.Delete(&User{})
	return db.Error
}

// DeleteOneByName is generated from sql:
// delete from `user` where `name` = ?;
func (m *UserModel) DeleteOneByName(ctx context.Context, where DeleteOneByNameWhereParameter) error {
	var db = m.db.WithContext(ctx)
	db = db.Where(`name = ?`, where.NameEqual)
	db = db.Delete(&User{})
	return db.Error
}

// DeleteOneOrderByIDAsc is generated from sql:
// delete from `user` where `name` = ? order by id;
func (m *UserModel) DeleteOneOrderByIDAsc(ctx context.Context, where DeleteOneOrderByIDAscWhereParameter) error {
	var db = m.db.WithContext(ctx)
	db = db.Where(`name = ?`, where.NameEqual)
	db = db.Order(`id`)
	db = db.Delete(&User{})
	return db.Error
}

// DeleteOneOrderByIDDesc is generated from sql:
// delete from `user` where `name` = ? order by id desc;
func (m *UserModel) DeleteOneOrderByIDDesc(ctx context.Context, where DeleteOneOrderByIDDescWhereParameter) error {
	var db = m.db.WithContext(ctx)
	db = db.Where(`name = ?`, where.NameEqual)
	db = db.Order(`id desc`)
	db = db.Delete(&User{})
	return db.Error
}

// DeleteOneOrderByIDDescLimitCount is generated from sql:
// delete from `user` where `name` = ? order by id desc limit ?;
func (m *UserModel) DeleteOneOrderByIDDescLimitCount(ctx context.Context, where DeleteOneOrderByIDDescLimitCountWhereParameter, limit DeleteOneOrderByIDDescLimitCountLimitParameter) error {
	var db = m.db.WithContext(ctx)
	db = db.Where(`name = ?`, where.NameEqual)
	db = db.Order(`id desc`)
	db = db.Limit(limit.Count)
	db = db.Delete(&User{})
	return db.Error
}
