package db

import (
	"github.com/astaxie/beego/orm"
)

type AuditLog struct {
	Id             int    `orm:"auto"`
	CreateTime     string `json:"create_time"`
	OpId           string `json:"op_id"`
	Domain         string `json:"domain"`
	Ip             string `json:"ip"`
	OpEvent        string `json:"op_event"`
	OpEventContent string `json:"op_event _content"`
}

func (auditLog *AuditLog) Insert() {
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Insert(auditLog)
	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
}

func (auditLog *AuditLog) Query() *[]AuditLog {
	var auditLogs []AuditLog
	o := orm.NewOrm()
	o.QueryTable(auditLog).All(&auditLogs)
	return &auditLogs
}
