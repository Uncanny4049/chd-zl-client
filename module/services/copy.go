package services

import (
	"fmt"
	"time"

	"github.com/Uncanny4049/chd-zl-client/database/zlnew"
	"github.com/Uncanny4049/chd-zl-types/types"
)

type CopyRecord struct {
	zlnew.CopyRecord
	Data  [][]string     `json:"data"` // 数据
	Type4 map[string]int `json:"type4"`
	Card  [][]string     `json:"card"` // 翻牌
}

func GetCopyByDate(role string, date time.Time) []CopyRecord {
	start := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	end := time.Date(date.Year(), date.Month(), date.Day()+1, 0, 0, 0, 0, date.Location())
	var ncrs []zlnew.CopyRecord
	zlnew.DB.Debug().Where(zlnew.CopyRecord{
		CopyRecord: types.CopyRecord{
			Role: role,
		},
	}).Where("date BETWEEN ? AND ?", start, end).Find(&ncrs)
	result := make([]CopyRecord, 0)
	for _, item := range ncrs {
		result = append(result, GetCopyInfo(item))
	}
	return result
}

func GetCopyInfo(copy zlnew.CopyRecord) CopyRecord {
	var ars []zlnew.ActionRecord
	zlnew.DB.Where(&zlnew.ActionRecord{
		Role: copy.Role,
	}).Where("date BETWEEN ? AND ?", copy.StartTime, copy.EndTime).Find(&ars)
	newCopy := CopyRecord{
		CopyRecord: copy,
		Data:       make([][]string, 5),
		Type4:      make(map[string]int),
		Card:       make([][]string, 0),
	}
	for _, item := range ars {
		switch item.Action {
		case "使用苏生复活":
			newCopy.Data[0] = append(newCopy.Data[0], item.Copy) //使用苏生复活
		case "掉线重连":
			newCopy.Data[1] = append(newCopy.Data[1], item.Copy) //掉线
		case "死亡回城":
			newCopy.Data[2] = append(newCopy.Data[2], item.Copy) // 死亡
		case "特殊执行":
			newCopy.Data[3] = append(newCopy.Data[3], item.Item) // 特殊执行
		case "激活图鉴":
			newCopy.Data[4] = append(newCopy.Data[4], fmt.Sprintf("【%s】\t %s", item.Copy, item.Item)) //使用苏生复活
		case "获得道具":
			if newCopy.Type4[item.Item] == 0 {
				newCopy.Type4[item.Item] = 1
			} else {
				newCopy.Type4[item.Item] += 1
			}
		case "免费翻牌":
			newCopy.Card = append(newCopy.Card, []string{item.Item})
		case "刷新翻牌":
			newCopy.Card[len(newCopy.Card)-1] = append(newCopy.Card[len(newCopy.Card)-1], item.Item)
		default:
		}
	}

	return newCopy
}

func GetAllRole() []string {
	var records []zlnew.CopyRecord
	zlnew.DB.Select("DISTINCT Role").Find(&records)
	name := make([]string, 0)
	for _, item := range records {
		name = append(name, item.Role)
	}
	return name
}
