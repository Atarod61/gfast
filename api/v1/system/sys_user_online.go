/*
* @desc:Online User
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2023/1/10 16:57
 */

package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

// SysUserOnlineSearchReq list search parameters
type SysUserOnlineSearchReq struct {
	g.Meta   `path:"/online/list" tags:"Online User Management" method:"get" summary:"List"`
	Username string `p:"userName"`
	Ip       string `p:"ipaddr"`
	commonApi.PageReq
	commonApi.Author
}

// SysUserOnlineSearchRes list results
type SysUserOnlineSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*entity.SysUserOnline `json:"list"`
}

type SysUserOnlineForceLogoutReq struct {
	g.Meta `path:"/online/forceLogout" tags:"Online User Management" method:"delete" summary:"Force User Logout"`
	commonApi.Author
	Ids []int `p:"ids" v:"required#ids cannot be empty"`
}

type SysUserOnlineForceLogoutRes struct {
	commonApi.EmptyRes
}