/*
* @desc:Department Management Parameters
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/4/6 15:07
 */

package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

type DeptSearchReq struct {
	g.Meta   `path:"/dept/list" tags:"Department Management" method:"get" summary:"Department List"`
	DeptName string `p:"deptName"`
	Status   string `p:"status"`
}

type DeptSearchRes struct {
	g.Meta   `mime:"application/json"`
	DeptList []*entity.SysDept `json:"deptList"`
}

type DeptAddReq struct {
	g.Meta   `path:"/dept/add" tags:"Department Management" method:"post" summary:"Add Department"`
	ParentID int    `p:"parentId"  v:"required#Parent cannot be empty"`
	DeptName string `p:"deptName"  v:"required#Department name cannot be empty"`
	OrderNum int    `p:"orderNum"  v:"required#Sort order cannot be empty"`
	Leader   string `p:"leader"`
	Phone    string `p:"phone"`
	Email    string `p:"email"  v:"email#Invalid email format"`
	Status   uint   `p:"status"  v:"required#Status is required"`
}

type DeptAddRes struct {
}

type DeptEditReq struct {
	g.Meta   `path:"/dept/edit" tags:"Department Management" method:"put" summary:"Edit Department"`
	DeptId   int    `p:"deptId" v:"required#deptId cannot be empty"`
	ParentID int    `p:"parentId"  v:"required#Parent cannot be empty"`
	DeptName string `p:"deptName"  v:"required#Department name cannot be empty"`
	OrderNum int    `p:"orderNum"  v:"required#Sort order cannot be empty"`
	Leader   string `p:"leader"`
	Phone    string `p:"phone"`
	Email    string `p:"email"  v:"email#Invalid email format"`
	Status   uint   `p:"status"  v:"required#Status is required"`
}

type DeptEditRes struct {
}

type DeptDeleteReq struct {
	g.Meta `path:"/dept/delete" tags:"Department Management" method:"delete" summary:"Delete Department"`
	Id     uint64 `p:"id" v:"required#id cannot be empty"`
}

type DeptDeleteRes struct {
}

type DeptTreeSelectReq struct {
	g.Meta `path:"/dept/treeSelect" tags:"Department Management" method:"get" summary:"Get Department Tree Menu"`
}

type DeptTreeSelectRes struct {
	g.Meta `mime:"application/json"`
	Deps   []*model.SysDeptTreeRes `json:"deps"`
}
