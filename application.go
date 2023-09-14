package main

import (
	"encoding/json"
	"github.com/Capsule7446/chd-zl-client/database"
	"github.com/Capsule7446/chd-zl-client/database/zlnew"
	"github.com/Capsule7446/chd-zl-client/database/zlold"
	"github.com/Capsule7446/chd-zl-client/module/types"
)

func main() {
	config := types.ReadConfig()
	zlold.Init(config)
	zlnew.Init(config)
	database.Tran()
	var crs []zlnew.CopyRecord
	zlnew.DB.Find(&crs)
	marshal, _ := json.Marshal(crs)
	println(string(marshal))
	//database.Tran()
	//var cls []zlnew.CopyRecord
	//zlnew.DB.Find(&cls)
	//for _, cl := range cls {
	//	var idxs []zlnew.ActionRecord
	//	zlnew.DB.Where(&zlnew.ActionRecord{
	//		Role: cl.Role,
	//	}).Where("date BETWEEN ? AND ?", cl.StartTime, cl.EndTime).Find(&idxs)
	//	if len(idxs) > 0 {
	//		log.Println(cl.CopyName)
	//		for _, ar := range idxs {
	//			log.Println(fmt.Sprintf("%s %s %s", ar.Copy, ar.Action, ar.Item))
	//		}
	//	}
	//	log.Println("------------------")
	//}
	//database.Tran()
}
