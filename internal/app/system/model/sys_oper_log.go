/*
* @desc:Operation log model object
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/21 16:34
 */

package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
	"net/url"
)

// SysOperLogAdd adds operation log parameters
type SysOperLogAdd struct {
	User         *ContextUser
	Menu         *SysAuthRuleInfoRes
	Url          *url.URL
	Params       g.Map
	Method       string
	ClientIp     string
	OperatorType int
}

// SysOperLogInfoRes is the golang structure for table sys_oper_log.
type SysOperLogInfoRes struct {
	gmeta.Meta     `orm:"table:sys_oper_log"`
	OperId         uint64                   `orm:"oper_id,primary" json:"operId"`       // Log ID
	Title          string                   `orm:"title" json:"title"`                  // System module
	BusinessType   int                      `orm:"business_type" json:"businessType"`   // Operation type
	Method         string                   `orm:"method" json:"method"`                // Operation method
	RequestMethod  string                   `orm:"request_method" json:"requestMethod"` // Request method
	OperatorType   int                      `orm:"operator_type" json:"operatorType"`   // Operation type
	OperName       string                   `orm:"oper_name" json:"operName"`           // Operator
	DeptName       string                   `orm:"dept_name" json:"deptName"`           // Department name
	LinkedDeptName *LinkedSysOperLogSysDept `orm:"with:dept_id=dept_name" json:"linkedDeptName"`
	OperUrl        string                   `orm:"oper_url" json:"operUrl"`           // RequestURL
	OperIp         string                   `orm:"oper_ip" json:"operIp"`             // Host address
	OperLocation   string                   `orm:"oper_location" json:"operLocation"` // Operation location
	OperParam      string                   `orm:"oper_param" json:"operParam"`       // Request parameters
	ErrorMsg       string                   `orm:"error_msg" json:"errorMsg"`         // Error message
	OperTime       *gtime.Time              `orm:"oper_time" json:"operTime"`         // Operation time
}

type LinkedSysOperLogSysDept struct {
	gmeta.Meta `orm:"table:sys_dept"`
	DeptId     int64  `orm:"dept_id" json:"deptId"`     // Department ID
	DeptName   string `orm:"dept_name" json:"deptName"` // Department name
}

type SysOperLogListRes struct {
	OperId         uint64                   `json:"operId"`
	Title          string                   `json:"title"`
	RequestMethod  string                   `json:"requestMethod"`
	OperName       string                   `json:"operName"`
	DeptName       string                   `json:"deptName"`
	LinkedDeptName *LinkedSysOperLogSysDept `orm:"with:dept_id=dept_name" json:"linkedDeptName"`
	OperUrl        string                   `json:"operUrl"`
	OperIp         string                   `json:"operIp"`
	OperLocation   string                   `json:"operLocation"`
	OperParam      string                   `json:"operParam"`
	OperTime       *gtime.Time              `json:"operTime"`
}
