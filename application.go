package main

import (
	"encoding/json"
	"time"

	"github.com/Uncanny4049/chd-zl-client/database/zlnew"
	"github.com/Uncanny4049/chd-zl-client/database/zlold"
	"github.com/Uncanny4049/chd-zl-client/module/services"
	"github.com/Uncanny4049/chd-zl-client/module/types"
)

func main() {
	config := types.ReadConfig()
	zlold.Init(config)
	zlnew.Init(config)
	//services.Tran()

	//data := services.GetCopyByDate("\uE812\uE812我饿了\uE812\uE812", time.Date(2023, 9, 5, 0, 0, 0, 0, time.Local))
	data := services.GetAllRole()
	d := make(map[string][]services.CopyRecord)
	for _, role := range data {
		d[role] = services.GetCopyByDate(role, time.Date(2023, 9, 5, 0, 0, 0, 0, time.Local))
	}
	marshal, _ := json.Marshal(d)
	println(string(marshal))
}
