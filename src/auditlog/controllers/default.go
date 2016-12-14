package controllers

import (
	"auditlog/models/db"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

//“msg":”操作时间/操作用户/租户/ip/操作事件/内容"
func (c *MainController) WriteLog() {
	var auditLog db.AuditLog
	body := c.Ctx.Input.RequestBody
	fmt.Println("====body====", string(body))
	json.Unmarshal(body, &auditLog)
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
