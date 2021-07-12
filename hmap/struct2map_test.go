package hmap

import (
	"fmt"
	"testing"
	"time"
)

type TripartiteIntercomAlarm struct {
	Id             int         `gorm:"column:id;not null;primaryKey;type:bigserial;commnet:'ID'"`
	FlowNo         interface{} `json:"FlowNo" gorm:"column:flow_no;type:numeric(64,10);commnet:'事件流水号'"`
	EventTime      time.Time   `json:"EventTime" gorm:"column:event_time;type:timestamp(6) with time zone;commnet:'事件发生时间'"`
	DeviceNo       string      `json:"DeviceNo" gorm:"column:device_no;type:character varying(64);commnet:'设备编号'"`
	EventDesc      string      `json:"EventDesc" gorm:"column:event_desc;type:character varying(1024);commnet:'事件描述'"`
	TreatTime      time.Time   `json:"TreatTime" gorm:"column:treat_time;type:timestamp(6) with time zone;commnet:'处理时间'"`
	TreatResult    string      `json:"TreatResult" gorm:"column:treat_result;type:character varying(64);commnet:'处理结果'"`
	TreatFlag      int32       `json:"TreatFlag" gorm:"column:treat_flag;type:integer;commnet:'处理状态'"`
	TreatName      string      `json:"TreatName" gorm:"column:treat_name;type:character varying(64);commnet:'处理者姓名'"`
	CommCmd        int32       `json:"CommCmd" gorm:"column:comm_cmd;type:integer;commnet:'事件编号'"`
	ManagerNo      string      `json:"ManagerNo" gorm:"column:manager_no;type:character varying(64);commnet:'管理员号'"`
	IsTreatTimeout bool        `json:"IsTreatTimeout" gorm:"column:is_treat_timeout;type:boolean;commnet:'处理超时'"`
	Tripartite     string      `gorm:"column:tripartite;type:character varying(64);commnet:'数据来源的三方'"`
	CreateTime     time.Time   `gorm:"column:create_time;default:now();type:timestamp(6) with time zone;commnet:'创建时间'"`
	UpdateTime     time.Time   `gorm:"column:update_time;default:now();type:timestamp(6) with time zone;commnet:'修改时间'"`
}

func TestStruct2Map(t *testing.T) {
	m := Struct2Map(TripartiteIntercomAlarm{})
	fmt.Println(m)
}

func TestStruct2MapWithGorm(t *testing.T) {
	m, err := Struct2MapWithGorm(TripartiteIntercomAlarm{})
	fmt.Println(m, err)

}
