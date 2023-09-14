package zlnew

import (
	"github.com/Uncanny4049/chd-zl-client/module/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(c *types.ServerConfig) {
	var err error
	DB, err = gorm.Open(sqlite.Open(c.DBPath), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}
