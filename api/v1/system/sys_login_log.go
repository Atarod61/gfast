/*
* @desc:Login Log
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu
* @Date:   2022/4/24 22:09
 */

package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

// LoginLogSearchReq query list request parameters
type LoginLogSearchReq struct {
	g.Meta        `path:"/loginLog/list" tags:"Login Log Management" method:"get" summary:"Log List"`
	LoginName     string `p:"userName"`      //Login name
	Status        string `p:"status"`        //Status
	Ipaddr        string `p:"ipaddr"`        //Login address
	SortName      string `p:"orderByColumn"` //Sort field
	SortOrder     string `p:"isAsc"`         //Sort method
	LoginLocation string `p:"loginLocation"` //Login location
	commonApi.PageReq
}

type LoginLogSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*entity.SysLoginLog `json:"list"`
}

type LoginLogDelReq struct {
	g.Meta `path:"/loginLog/delete" tags:"Login Log Management" method:"delete" summary:"Delete Logs"`
	Ids    []int `p:"ids" v:"required#ids are required"`
}

type LoginLogDelRes struct {
}

type LoginLogClearReq struct {
	g.Meta `path:"/loginLog/clear" tags:"Login Log Management" method:"delete" summary:"Clear Logs"`
}

type LoginLogClearRes struct {
}
