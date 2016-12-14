package controllers

import (
	"auditlog/models/db"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type MainController struct {
	beego.Controller
}

//“msg":”操作时间/操作用户/租户/ip/操作事件/内容"
func (c *MainController) WriteLog() {
	var auditLog db.AuditLog
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &auditLog)
	createTime64, _ := strconv.ParseInt(auditLog.CreateTime, 10, 64)
	timeUnix := time.Unix(createTime64, 0)
	layout := "2006-01-02 15:04:05"
	auditLog.CreateTime = timeUnix.Format(layout)
	err := auditLog.Insert()
	if err != nil {
		SetOutput(c.Ctx, 304, []byte(err.Error()))
	}
	SetOutput(c.Ctx, 201, []byte("插入数据看成功"))
}

func (c *MainController) GetLog() {

	auditLoggers := make([]AuditLogger, 0)
	var auditLog db.AuditLog
	var operationDes db.OperationDes
	operationDesmap := operationDes.Query()
	auditLogs := auditLog.Query()
	for _, a := range *auditLogs {
		auditLogger := new(AuditLogger)
		auditLogger.CreateTime = a.CreateTime
		auditLogger.OpId = a.OpId
		auditLogger.Domain = a.Domain
		auditLogger.Ip = a.Ip
		auditLogger.OpEvent = a.OpEvent
		auditLogger.Description = operationDesmap[a.OpEvent]
		auditLoggers = append(auditLoggers, *auditLogger)
	}
	auditLoggersJson, err := json.Marshal(auditLoggers)
	if err != nil {
		SetOutput(c.Ctx, 304, []byte(err.Error()))
	}
	SetOutput(c.Ctx, 200, auditLoggersJson)
}
