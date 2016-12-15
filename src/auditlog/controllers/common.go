package controllers

import (
	"github.com/astaxie/beego/context"
)

type AuditLogger struct {
	CreateTime  string `orm:"column(create_time)" json:"create_time"`
	OpId        string `orm:"column(op_id)" json:"op_id"`
	Domain      string `orm:"column(domain)" json:"domain"`
	Ip          string `orm:"column(ip)" json:"ip"`
	OpEvent     string `orm:"column(op_event)" json:"op_event"`
	Description string `orm:"column(description)" json:"description"`
}

func SetOutput(ctx *context.Context, status int, body interface{}) {
	ctx.Output.Status = status
	if body == nil {
		body = ""
	}
	switch body.(type) {
	case string:
		ctx.Output.Body([]byte(body.(string)))
	case error:
		ctx.Output.Body([]byte(body.(error).Error()))
	case []byte:
		ctx.Output.Body(body.([]byte))
	}
}
