package db

import (
	"github.com/astaxie/beego/orm"
)

type OperationDes struct {
	OpEvent     string `orm:"pk"`
	OpType      string
	Description string
}

func (operationDes *OperationDes) Query() map[string]string {
	var operationDess []OperationDes
	operationDesmap := make(map[string]string, 0)
	o := orm.NewOrm()
	o.QueryTable(operationDes).Exclude("op_type","query").All(&operationDess)
	for _, v := range operationDess {
		operationDesmap[v.OpEvent] = v.Description
	}
	return operationDesmap
}
