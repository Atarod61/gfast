/*
* @desc:Dictionary Dataapi
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/3/18 11:59
 */

package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	commonModel "github.com/tiger1103/gfast/v3/internal/app/common/model"
	commonEntity "github.com/tiger1103/gfast/v3/internal/app/common/model/entity"
)

// GetDictReq Get dictionary information request parameters
type GetDictReq struct {
	g.Meta `path:"/dict/data/getDictData" tags:"Dictionary Management" method:"get" summary:"Get dictionary data (public)"`
	commonApi.Author
	DictType     string `p:"dictType" v:"required#Dictionary type cannot be empty"`
	DefaultValue string `p:"defaultValue"`
}

// GetDictRes complete dictionary information
type GetDictRes struct {
	g.Meta `mime:"application/json"`
	Info   *commonModel.DictTypeRes   `json:"info"`
	Values []*commonModel.DictDataRes `json:"values"`
}

// DictDataSearchReq paging request parameters
type DictDataSearchReq struct {
	g.Meta    `path:"/dict/data/list" tags:"Dictionary Management" method:"get" summary:"Dictionary Data List"`
	DictType  string `p:"dictType"`  //Dictionary type
	DictLabel string `p:"dictLabel"` //Dictionary label
	Status    string `p:"status"`    //Status
	commonApi.PageReq
}

// DictDataSearchRes dictionary data result
type DictDataSearchRes struct {
	g.Meta `mime:"application/json"`
	List   []*commonEntity.SysDictData `json:"list"`
	commonApi.ListRes
}

type DictDataReq struct {
	DictLabel string `p:"dictLabel"  v:"required#Dictionary label cannot be empty"`
	DictValue string `p:"dictValue"  v:"required#Dictionary key value cannot be empty"`
	DictType  string `p:"dictType"  v:"required#Dictionary type cannot be empty"`
	DictSort  int    `p:"dictSort"  v:"integer#Sort value must be integer"`
	CssClass  string `p:"cssClass"`
	ListClass string `p:"listClass"`
	IsDefault int    `p:"isDefault" v:"required|in:0,1#System default cannot be empty|Default value must be 0 or 1"`
	Status    int    `p:"status"    v:"required|in:0,1#Status cannot be empty|Status must be 0 or 1"`
	Remark    string `p:"remark"`
}

type DictDataAddReq struct {
	g.Meta `path:"/dict/data/add" tags:"Dictionary Management" method:"post" summary:"Add Dictionary Data"`
	*DictDataReq
}

type DictDataAddRes struct {
}

type DictDataGetReq struct {
	g.Meta   `path:"/dict/data/get" tags:"Dictionary Management" method:"get" summary:"Get Dictionary Data"`
	DictCode uint `p:"dictCode"`
}

type DictDataGetRes struct {
	g.Meta `mime:"application/json"`
	Dict   *commonEntity.SysDictData `json:"dict"`
}

type DictDataEditReq struct {
	g.Meta   `path:"/dict/data/edit" tags:"Dictionary Management" method:"put" summary:"Edit Dictionary Data"`
	DictCode int `p:"dictCode" v:"required|min:1#Primary key ID cannot be empty|Primary key ID cannot be less than 1"`
	*DictDataReq
}

type DictDataEditRes struct {
}

type DictDataDeleteReq struct {
	g.Meta `path:"/dict/data/delete" tags:"Dictionary Management" method:"delete" summary:"Delete Dictionary Data"`
	Ids    []int `p:"ids"`
}

type DictDataDeleteRes struct {
}
