/*
* @desc:Operation Log
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/12/21 14:37
 */

package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

// SysOperLogSearchReq pagination request parameters
type SysOperLogSearchReq struct {
	g.Meta        `path:"/operLog/list" tags:"Operation Log" method:"get" summary:"Operation Log List"`
	Title         string `p:"title"`         //System module
	RequestMethod string `p:"requestMethod"` //Request method
	OperName      string `p:"operName"`      //Operator
	commonApi.PageReq
	commonApi.Author
}

// SysOperLogSearchRes returns a list of results
type SysOperLogSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*model.SysOperLogListRes `json:"list"`
}

// SysOperLogGetReq gets a data request
type SysOperLogGetReq struct {
	g.Meta `path:"/operLog/get" tags:"Operation Log" method:"get" summary:"Get operation log details"`
	commonApi.Author
	OperId uint64 `p:"operId" v:"required#Primary key is required"` //Get by primary key
}

// SysOperLogGetRes retrieves a data result
type SysOperLogGetRes struct {
	g.Meta `mime:"application/json"`
	*model.SysOperLogInfoRes
}

// SysOperLogDeleteReq deletes data request
type SysOperLogDeleteReq struct {
	g.Meta `path:"/operLog/delete" tags:"Operation Log" method:"delete" summary:"Delete operation log"`
	commonApi.Author
	OperIds []uint64 `p:"operIds" v:"required#Primary key is required"` //Delete by primary key
}

// SysOperLogDeleteRes deletes data and returns a result
type SysOperLogDeleteRes struct {
	commonApi.EmptyRes
}

type SysOperLogClearReq struct {
	g.Meta `path:"/operLog/clear" tags:"Operation Log" method:"delete" summary:"Clear logs"`
	commonApi.Author
}

type SysOperLogClearRes struct {
	commonApi.EmptyRes
}

