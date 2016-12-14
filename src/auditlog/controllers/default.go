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
	json.Unmarshal(body, &auditLog)
	auditLog.Insert()
}

func (c *MainController) GetLog() {
	var auditLog db.AuditLog
	auditLogs := auditLog.Query()
	auditLogJson, err := json.Marshal(auditLogs)
	if err != nil {
		fmt.Println(err)
	}
	c.Ctx.Output.Body(auditLogJson)
}
