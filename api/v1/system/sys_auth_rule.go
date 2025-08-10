/*
* @desc:Menuapi
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/3/18 10:27
 */

package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

type RuleSearchReq struct {
	g.Meta `path:"/menu/list" tags:"Menu Management" method:"get" summary:"Menu List"`
	commonApi.Author
	Title     string `p:"menuName" `
	Component string `p:"component"`
}

type RuleListRes struct {
	g.Meta `mime:"application/json"`
	Rules  []*model.SysAuthRuleTreeRes `json:"rules"`
}

type RuleAddReq struct {
	g.Meta `path:"/menu/add" tags:"Menu Management" method:"post" summary:"Add Menu"`
	commonApi.Author
	MenuType  uint   `p:"menuType"  v:"min:0|max:2#Menu type minimum value is:min|Menu type maximum value is:max"`
	Pid       uint   `p:"parentId"  v:"min:0"`
	Name      string `p:"name" v:"required#Please enter the rule name"`
	Title     string `p:"menuName" v:"required|length:1,100#Please enter a title|Title length must be between:min and:max characters"`
	Icon      string `p:"icon"`
	Weigh     int    `p:"menuSort" `
	Condition string `p:"condition" `
	Remark    string `p:"remark" `
	IsHide    uint   `p:"isHide"`
	Path      string `p:"path"`
	Redirect  string `p:"redirect"`
	Roles     []uint `p:"roles"`
	Component string `p:"component" v:"required-if:menuType,1#Component path cannot be empty"`
	IsLink    uint   `p:"isLink"`
	IsIframe  uint   `p:"isIframe"`
	IsCached  uint   `p:"isKeepAlive"`
	IsAffix   uint   `p:"isAffix"`
	LinkUrl   string `p:"linkUrl"`
}

type RuleAddRes struct {
}

type RuleGetParamsReq struct {
	g.Meta `path:"/menu/getParams" tags:"Menu Management" method:"get" summary:"Get parameters for adding、editing menu"`
	commonApi.Author
}

type RuleGetParamsRes struct {
	g.Meta `mime:"application/json"`
	Roles  []*entity.SysRole           `json:"roles"`
	Menus  []*model.SysAuthRuleInfoRes `json:"menus"`
}

type RuleInfoReq struct {
	g.Meta `path:"/menu/get" tags:"Menu Management" method:"get" summary:"Get menu details"`
	commonApi.Author
	Id uint `p:"id" v:"required#Menu ID is required"`
}

type RuleInfoRes struct {
	g.Meta  `mime:"application/json"`
	Rule    *entity.SysAuthRule `json:"rule"`
	RoleIds []uint              `json:"roleIds"`
}

type RuleUpdateReq struct {
	g.Meta `path:"/menu/update" tags:"Menu Management" method:"put" summary:"Update menu"`
	commonApi.Author
	Id        uint   `p:"id" v:"required#id is required"`
	MenuType  uint   `p:"menuType"  v:"min:0|max:2#Menu type minimum value is:min|Menu type maximum value is:max"`
	Pid       uint   `p:"parentId"  v:"min:0"`
	Name      string `p:"name" v:"required#Please enter the rule name"`
	Title     string `p:"menuName" v:"required|length:1,100#Please enter a title|Title length must be between:min and:max characters"`
	Icon      string `p:"icon"`
	Weigh     int    `p:"menuSort" `
	Condition string `p:"condition" `
	Remark    string `p:"remark" `
	IsHide    uint   `p:"isHide"`
	Path      string `p:"path"`
	Redirect  string `p:"redirect"`
	Roles     []uint `p:"roles"`
	Component string `p:"component" v:"required-if:menuType,1#Component path cannot be empty"`
	IsLink    uint   `p:"isLink"`
	IsIframe  uint   `p:"isIframe"`
	IsCached  uint   `p:"isKeepAlive"`
	IsAffix   uint   `p:"isAffix"`
	LinkUrl   string `p:"linkUrl"`
}

type RuleUpdateRes struct {
}

type RuleDeleteReq struct {
	g.Meta `path:"/menu/delete" tags:"Menu Management" method:"delete" summary:"Delete menu"`
	commonApi.Author
	Ids []int `p:"ids" v:"required#Menu id are required"`
}

type RuleDeleteRes struct {
}
