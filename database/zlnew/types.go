package zlnew

import (
	"github.com/Capsule7446/chd-zl-client/database/zlold"
	"github.com/Capsule7446/chd-zl-types/types"
	"strconv"
	"strings"
	"time"
)

type ActionRecord types.ActionRecord

func (r *ActionRecord) TableName() string {
	return "ActionRecord"
}

type CopyRecord struct {
	*types.CopyRecord
	CopyName             string    `json:"copyName,omitempty"`             //副本名称
	Status               string    `json:"status,omitempty"`               //完成状态
	StartTime            time.Time `json:"startTime,omitempty"`            // 开始时间
	EndTime              time.Time `json:"endTime,omitempty"`              // 结束时间
	TimeConsuming        float64   `json:"timeConsuming,omitempty"`        // 耗时
	Predict              int       `json:"predict,omitempty"`              //计划次数
	Actual               int       `json:"actual,omitempty"`               //完成次数
	Revival              int       `json:"revival,omitempty"`              //苏生
	NoPickEqu            int       `json:"noPickEqu,omitempty"`            //不拾取装备
	NoPickCon            int       `json:"noPickCon,omitempty"`            //不拾取消耗
	NoPickOth            int       `json:"noPickOth,omitempty"`            //不拾取其他
	NoPickTask           int       `json:"noPickTask,omitempty"`           //不拾取任务
	Blood                int       `json:"blood,omitempty"`                //血脉战魂
	PotionPlan           string    `json:"potionPlan,omitempty"`           //药水方案
	EquipmentPlan        string    `json:"equipmentPlan,omitempty"`        //装备方案
	PetFormation         string    `json:"petFormation,omitempty"`         //宠物阵型
	LimitTimeOut         int       `json:"limitTimeOut,omitempty"`         //限制-超时
	LimitRevival         int       `json:"limitRevival,omitempty"`         //限制-苏生
	ExtraCard            int       `json:"extraCard,omitempty"`            //额外翻牌
	ExtraRef             int       `json:"extraRef,omitempty"`             //额外刷新
	DeathRes             int       `json:"deathRes,omitempty"`             //死亡复活
	ExperienceWings      int       `json:"experienceWings,omitempty"`      //经验翅膀
	AggroAndMobPulling   int       `json:"aggroAndMobPulling,omitempty"`   //聚拢拉怪
	TreasureChestWaiting int       `json:"treasureChestWaiting,omitempty"` //宝箱等待
	ChangeTitle          string    `json:"changeTitle,omitempty"`          //更换称号
	BossEvasion          int       `json:"bossEvasion,omitempty"`          //Boss躲避
	DoNotAttackTheLord   int       `json:"doNotAttackTheLord,omitempty"`   //不打领主
}

func (r *CopyRecord) TableName() string {
	return "CopyRecord"
}

func GetCopyRecord(c zlold.CopyRecord) *CopyRecord {
	row := strings.Split(c.Types, ",")
	ToInt := func(s string) int {
		intNum, _ := strconv.Atoi(s)
		return intNum
	}
	ncr := CopyRecord{
		CopyRecord: &types.CopyRecord{
			Idx:   c.Idx,
			Role:  c.Role,
			Types: c.Types,
			Date:  c.Date,
		},
		CopyName:             row[0],
		Status:               row[1],
		StartTime:            ToDate(row[2], c.Date),
		EndTime:              ToDate(row[3], c.Date),
		Predict:              ToInt(row[5]),
		Actual:               ToInt(row[6]),
		Revival:              ToInt(row[7]),
		NoPickEqu:            ToInt(row[8]),
		NoPickCon:            ToInt(row[9]),
		NoPickOth:            ToInt(row[10]),
		NoPickTask:           ToInt(row[11]),
		Blood:                ToInt(row[12]),
		PotionPlan:           row[13],
		EquipmentPlan:        row[14],
		PetFormation:         row[15],
		LimitTimeOut:         ToInt(row[16]),
		LimitRevival:         ToInt(row[17]),
		ExtraCard:            ToInt(row[18]),
		ExtraRef:             ToInt(row[19]),
		DeathRes:             ToInt(row[20]),
		ExperienceWings:      ToInt(row[21]),
		AggroAndMobPulling:   ToInt(row[22]),
		TreasureChestWaiting: ToInt(row[23]),
		ChangeTitle:          row[24],
		BossEvasion:          ToInt(row[25]),
		DoNotAttackTheLord:   ToInt(row[26]),
	}
	ncr.TimeConsuming = ncr.EndTime.Sub(ncr.StartTime).Seconds()
	return &ncr
}

type Income types.Income

func (r *Income) TableName() string {
	return "Income"
}

type Player types.Player

func (r *Player) TableName() string {
	return "Player"
}

type RecordType types.RecordType

func (r *RecordType) TableName() string {
	return "RecordType"
}

type TaskInfo types.TaskInfo

func (r *TaskInfo) TableName() string {
	return "TaskInfo"
}

func ToDate(dateStr string, date time.Time) time.Time {
	ToInt := func(s string) int {
		intNum, _ := strconv.Atoi(s)
		return intNum
	}
	dateSlice := strings.Split(dateStr, ":")
	if len(dateSlice) == 3 {
		return time.Date(date.Year(), date.Month(), date.Day(), ToInt(dateSlice[0]), ToInt(dateSlice[1]), ToInt(dateSlice[2]), 0, time.Local)
	}
	return date
}
