package database

import (
	"errors"
	"github.com/Capsule7446/chd-zl-client/database/zlnew"
	"github.com/Capsule7446/chd-zl-client/database/zlold"
	"gorm.io/gorm"
	"log"
)

type Idx struct {
	Idx int
}

func Tran() {
	zlnew.DB.AutoMigrate(&zlnew.ActionRecord{})
	zlnew.DB.AutoMigrate(&zlnew.CopyRecord{})
	zlnew.DB.AutoMigrate(&zlnew.Income{})
	zlnew.DB.AutoMigrate(&zlnew.Player{})
	zlnew.DB.AutoMigrate(&zlnew.RecordType{})
	zlnew.DB.AutoMigrate(&zlnew.TaskInfo{})

	TranActionRecord()
	TranCopyRecord()
	TranIncome()
	// ---

	var ps []zlold.Player
	zlold.DB.Find(&ps)
	for _, item := range ps {
		zlnew.DB.Create(&item)
	}

	var rts []zlold.RecordType
	zlnew.DB.Find(&rts)
	if len(rts) > 0 {
		zlold.DB.Find(&rts)
		for _, item := range rts {
			zlnew.DB.Create(&item)
		}
	}

	var tis []zlold.TaskInfo
	zlold.DB.Find(&tis)
	for _, item := range tis {
		zlnew.DB.Create(&item)
	}
}

func TranActionRecord() {
	var _new zlnew.ActionRecord
	if errors.Is(zlnew.DB.Last(&_new).Error, gorm.ErrRecordNotFound) {
		_new.Idx = -1
	}
	var old []zlold.ActionRecord
	if tx := zlold.DB.Find(&old, "idx > ?", _new.Idx); tx.Error != nil {
		log.Fatal("Error")
		return
	}
	for _, item := range old {
		err := zlnew.DB.Create(&item).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("创建失败：%d", item.Idx)
		}
	}

}

func TranCopyRecord() {
	var _new zlnew.CopyRecord
	if errors.Is(zlnew.DB.Last(&_new).Error, gorm.ErrRecordNotFound) {
		_new.Idx = -1
	}
	var old []zlold.CopyRecord
	if tx := zlold.DB.Find(&old, "idx > ?", _new.Idx); tx.Error != nil {
		log.Fatal("Error")
		return
	}
	for _, item := range old {
		err := zlnew.DB.Create(zlnew.GetCopyRecord(item)).Error
		if err != nil {
			log.Printf("创建失败：%d", item.Idx)
		}
	}
}

func TranIncome() {
	var _new zlnew.Income
	if errors.Is(zlnew.DB.Last(&_new).Error, gorm.ErrRecordNotFound) {
		_new.Idx = -1
	}
	var old []zlold.Income
	if tx := zlold.DB.Find(&old, "idx > ?", _new.Idx); tx.Error != nil {
		log.Fatal("Error")
		return
	}
	for _, item := range old {
		err := zlnew.DB.Create(&item).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("创建失败：%d", item.Idx)
		}
	}

}
