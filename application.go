package main

import (
	"encoding/json"
	"github.com/Capsule7446/chd-zl-client/database/zlnew"
	"github.com/Capsule7446/chd-zl-client/database/zlold"
	"github.com/Capsule7446/chd-zl-client/module/services"
	"github.com/Capsule7446/chd-zl-client/module/types"
	"log"
)

func main() {
	config := types.ReadConfig()
	zlold.Init(config)
	zlnew.Init(config)
	//services.Tran()
	var ncrs []zlnew.CopyRecord
	zlnew.DB.Find(&ncrs)
	for _, item := range ncrs {
		info := services.GetCopyInfo(&item)
		if len(info.Type1) == 0 && len(info.Type2) == 0 && len(info.Type3) == 0 && len(info.Type4) == 0 {
			continue
		}
		marshal, _ := json.Marshal(info)
		log.Println(string(marshal))
	}
}
