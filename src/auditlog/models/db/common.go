package db

type AuditLogger struct {
	CreateTime  string `orm:"column(create_time)" json:"create_time"`
	OpId        string `orm:"column(op_id)" json:"op_id"`
	Domain      string `orm:"column(domain)" json:"domain"`
	Ip          string `orm:"column(ip)" json:"ip"`
	OpEvent     string `orm:"column(op_event)" json:"op_event"`
	Description string `orm:"column(description)" json:"description"`
}
