/*
* @desc:Dictionary type
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu
* @Date:   2022/4/14 21:30
 */

package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	commonModel "github.com/tiger1103/gfast/v3/internal/app/common/model"
	commonEntity "github.com/tiger1103/gfast/v3/internal/app/common/model/entity"
)

type DictTypeSearchReq struct {
	g.Meta   `path:"/dict/type/list" tags:"Dictionary Management" method:"get" summary:"Dictionary Type List"`
	DictName string `p:"dictName"` //Dictionary name
	DictType string `p:"dictType"` //Dictionary type
	Status   string `p:"status"`   //Dictionary status
	commonApi.PageReq
}

type DictTypeSearchRes struct {
	g.Meta       `mime:"application/json"`
	DictTypeList []*commonModel.SysDictTypeInfoRes `json:"dictTypeList"`
	commonApi.ListRes
}

type DictTypeAddReq struct {
	g.Meta   `path:"/dict/type/add" tags:"Dictionary Management" method:"post" summary:"Add Dictionary Type"`
	DictName string `p:"dictName"  v:"required#Dictionary name cannot be empty"`
	DictType string `p:"dictType"  v:"required#Dictionary type cannot be empty"`
	Status   uint   `p:"status"  v:"required|in:0,1#Status cannot be empty|Status must be 0 or 1"`
	Remark   string `p:"remark"`
}

type DictTypeAddRes struct {
}

type DictTypeGetReq struct {
	g.Meta `path:"/dict/type/get" tags:"Dictionary Management" method:"get" summary:"Get Dictionary Type"`
	DictId uint `p:"dictId" v:"required#Type ID cannot be empty"`
}

type DictTypeGetRes struct {
	g.Meta   `mime:"application/json"`
	DictType *commonEntity.SysDictType `json:"dictType"`
}

type DictTypeEditReq struct {
	g.Meta   `path:"/dict/type/edit" tags:"Dictionary Management" method:"put" summary:"Edit Dictionary Type"`
	DictId   int64  `p:"dictId" v:"required|min:1#Primary key ID cannot be empty|Primary key ID must be greater than 0"`
	DictName string `p:"dictName"  v:"required#Dictionary name cannot be empty"`
	DictType string `p:"dictType"  v:"required#Dictionary type cannot be empty"`
	Status   uint   `p:"status"  v:"required|in:0,1#Status cannot be empty|Status must be 0 or 1"`
	Remark   string `p:"remark"`
}

type DictTypeEditRes struct {
}

type DictTypeDeleteReq struct {
	g.Meta  `path:"/dict/type/delete" tags:"Dictionary Management" method:"delete" summary:"Delete Dictionary Type"`
	DictIds []int `p:"dictIds" v:"required#Dictionary type ID cannot be empty"`
}

type DictTypeDeleteRes struct {
}

type DictTypeAllReq struct {
	g.Meta `path:"/dict/type/optionSelect" tags:"Dictionary Management" method:"get" summary:"Get Dictionary Select Options"`
}

type DictTYpeAllRes struct {
	g.Meta   `mime:"application/json"`
	DictType []*commonEntity.SysDictType `json:"dictType"`
}
