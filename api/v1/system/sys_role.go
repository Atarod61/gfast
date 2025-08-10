/*
* @desc:Roleapi
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/3/30 9:16
 */

package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

type RoleListReq struct {
	g.Meta   `path:"/role/list" tags:"Role Management" method:"get" summary:"Role List"`
	RoleName string `p:"roleName"`   //Parameter name
	Status   string `p:"roleStatus"` //Status
	commonApi.PageReq
}

type RoleListRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*entity.SysRole `json:"list"`
}

type RoleGetParamsReq struct {
	g.Meta `path:"/role/getParams" tags:"Role Management" method:"get" summary:"Role Edit Parameters"`
}

type RoleGetParamsRes struct {
	g.Meta `mime:"application/json"`
	Menu   []*model.SysAuthRuleInfoRes `json:"menu"`
}

type RoleAddReq struct {
	g.Meta    `path:"/role/add" tags:"Role Management" method:"post" summary:"Add Role"`
	Name      string `p:"name" v:"required#Role name cannot be empty"`
	Status    uint   `p:"status"    `
	ListOrder uint   `p:"listOrder" `
	Remark    string `p:"remark"    `
	MenuIds   []uint `p:"menuIds"`
}

type RoleAddRes struct {
}

type RoleGetReq struct {
	g.Meta `path:"/role/get" tags:"Role Management" method:"get" summary:"Get role information"`
	Id     uint `p:"id" v:"required#Role ID cannot be empty"`
}

type RoleGetRes struct {
	g.Meta  `mime:"application/json"`
	Role    *entity.SysRole `json:"role"`
	MenuIds []int           `json:"menuIds"`
}

type RoleEditReq struct {
	g.Meta    `path:"/role/edit" tags:"Role Management" method:"put" summary:"Edit role"`
	Id        int64  `p:"id" v:"required#Role ID is required"`
	Name      string `p:"name" v:"required#Role name cannot be empty"`
	Status    uint   `p:"status"    `
	ListOrder uint   `p:"listOrder" `
	Remark    string `p:"remark"    `
	MenuIds   []uint `p:"menuIds"`
}

type RoleEditRes struct {
}

type RoleDeleteReq struct {
	g.Meta `path:"/role/delete" tags:"Role Management" method:"delete" summary:"Delete role"`
	Ids    []int64 `p:"ids" v:"required#Role ID cannot be empty"`
}

type RoleDeleteRes struct {
}