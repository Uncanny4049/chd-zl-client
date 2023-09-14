package zlold

import (
	"github.com/Capsule7446/chd-zl-types/types"
	"github.com/axgle/mahonia"
	"gorm.io/gorm"
	"time"
)

var decoder = mahonia.NewDecoder("gbk")

func Transform(str string) string {
	return decoder.ConvertString(str)
}

type ActionRecord types.ActionRecord

func (r *ActionRecord) AfterFind(db *gorm.DB) error {
	r.Copy = Transform(r.Copy)
	r.Action = Transform(r.Action)
	r.Item = Transform(r.Item)
	r.Role = Transform(r.Role)
	r.Date = time.Date(r.Date.Year(), r.Date.Month(),
		r.Date.Day(), r.Date.Hour(),
		r.Date.Minute(), r.Date.Second(),
		r.Date.Nanosecond(), time.Local)
	return nil
}

func (r *ActionRecord) TableName() string {
	return "ActionRecord"
}

type CopyRecord types.CopyRecord

func (r *CopyRecord) AfterFind(db *gorm.DB) error {
	r.Role = Transform(r.Role)
	r.Types = Transform(r.Types)
	return nil
}

func (r *CopyRecord) TableName() string {
	return "CopyRecord"
}

type Income types.Income

func (r *Income) AfterFind(db *gorm.DB) error {
	r.Item = Transform(r.Item)
	r.Role = Transform(r.Role)
	r.Date = r.Date.In(time.Local)
	return nil
}

func (r *Income) TableName() string {
	return "Income"
}

type Player types.Player

func (r *Player) AfterFind(db *gorm.DB) error {
	r.Name = Transform(r.Name)
	return nil
}

func (r *Player) TableName() string {
	return "Player"
}

type RecordType types.RecordType

func (r *RecordType) AfterFind(db *gorm.DB) error {
	r.Name = Transform(r.Name)
	return nil
}

func (r *RecordType) TableName() string {
	return "RecordType"
}

type TaskInfo types.TaskInfo

func (r *TaskInfo) AfterFind(db *gorm.DB) error {
	r.Role = Transform(r.Role)
	r.Date = time.Date(r.Date.Year(), r.Date.Month(),
		r.Date.Day(), r.Date.Hour(),
		r.Date.Minute(), r.Date.Second(),
		r.Date.Nanosecond(), time.Local)
	return nil
}

func (r *TaskInfo) TableName() string {
	return "TaskInfo"
}
