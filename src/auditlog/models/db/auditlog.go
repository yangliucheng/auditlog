package db

import (
	"github.com/astaxie/beego/orm"
)

type AuditLog struct {
	Id             int    `orm:"auto"`
	CreateTime     string `orm:"column(create_time)" json:"create_time"`
	OpId           string `json:"op_id"`
	Domain         string `json:"domain"`
	Ip             string `json:"ip"`
	OpEvent        string `json:"op_event"`
	OpEventContent string `json:"op_event_content"`
}

func (auditLog *AuditLog) Insert() error {
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Insert(auditLog)
	if err != nil {
		err = o.Rollback()
		return err
	} else {
		err = o.Commit()
		return err
	}
	return nil
}

func (auditLog *AuditLog) Query(count int) *[]AuditLog {
	var auditLogs []AuditLog
	o := orm.NewOrm()
	o.QueryTable(auditLog).Exclude("op_event","").OrderBy("-create_time").Limit(count).All(&auditLogs)
	return &auditLogs
}
