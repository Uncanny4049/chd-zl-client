package services

import (
	"fmt"
	"time"

	"github.com/Uncanny4049/chd-zl-client/database/zlnew"
	"github.com/Uncanny4049/chd-zl-types/types"
)

type CopyRecord struct {
	*zlnew.CopyRecord
	Type1 []string       `json:"type1"`
	Type2 []string       `json:"type2"`
	Type3 []string       `json:"type3"`
	Type4 map[string]int `json:"type4"`
}

func GetCopyByDate(role string, date time.Time) []CopyRecord {
	start := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	end := time.Date(date.Year(), date.Month(), date.Day()+1, 0, 0, 0, 0, date.Location())
	var ncrs []zlnew.CopyRecord
	zlnew.DB.Where(zlnew.CopyRecord{
		CopyRecord: &types.CopyRecord{
			Role: role,
		},
	}).Where("date BETWEEN ? AND ?", start, end).Find(&ncrs)
	result := make([]CopyRecord, 0)
	for _, item := range ncrs {
		result = append(result, GetCopyInfo(&item))
	}
	return result
}

func GetCopyInfo(copy *zlnew.CopyRecord) CopyRecord {
	var ars []zlnew.ActionRecord
	zlnew.DB.Where(&zlnew.ActionRecord{
		Role: copy.Role,
	}).Where("date BETWEEN ? AND ?", copy.StartTime, copy.EndTime).Find(&ars)
	newCopy := CopyRecord{
		CopyRecord: copy,
		Type1:      make([]string, 0),
		Type2:      make([]string, 0),
		Type3:      make([]string, 0),
		Type4:      make(map[string]int),
	}
	for _, item := range ars {
		switch item.Action {
		case "使用苏生复活":
			newCopy.Type1 = append(newCopy.Type1, item.Copy) //使用苏生复活
		case "掉线重连":
			newCopy.Type2 = append(newCopy.Type2, item.Copy) //使用苏生复活
		case "激活图鉴":
			newCopy.Type3 = append(newCopy.Type3, fmt.Sprintf("【%s】\t %s", item.Copy, item.Item)) //使用苏生复活
		case "获得道具":
			if newCopy.Type4[item.Item] == 0 {
				newCopy.Type4[item.Item] = 1
			} else {
				newCopy.Type4[item.Item] += 1
			}
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
