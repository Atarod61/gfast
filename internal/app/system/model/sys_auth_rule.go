/*
* @desc:Menumodel
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu
* @Date:   2022/3/11 14:53
 */

package model

type SysAuthRuleInfoRes struct {
	Id        uint   `orm:"id,primary"  json:"id"`        //
	Pid       uint   `orm:"pid"         json:"pid"`       // Parent ID
	Name      string `orm:"name,unique" json:"name"`      // Rule name
	Title     string `orm:"title"       json:"title"`     // Rule name
	Icon      string `orm:"icon"        json:"icon"`      // Icon
	Condition string `orm:"condition"   json:"condition"` // Condition
	Remark    string `orm:"remark"      json:"remark"`    // Remark
	MenuType  uint   `orm:"menu_type"   json:"menuType"`  // Type 0Directory 1Menu 2Button
	Weigh     int    `orm:"weigh"       json:"weigh"`     // Weight
	IsHide    uint   `orm:"is_hide" json:"isHide"`        // Display status
	IsCached  uint   `orm:"is_cached"  json:"isCached"`   // Cached or not
	IsAffix   uint   `orm:"is_affix" json:"isAffix"`      //Fixed or not
	Path      string `orm:"path"        json:"path"`      // Route address
	Redirect  string `orm:"redirect"   json:"redirect"`   // Redirect route
	Component string `orm:"component"   json:"component"` // Component path
	IsIframe  uint   `orm:"is_iframe"    json:"isIframe"` // Iframe or not
	IsLink    uint   `orm:"is_link" json:"isLink"`        // Is this an external link? 1 for yes, 0 for no
	LinkUrl   string `orm:"link_url" json:"linkUrl"`      //Link URL
}

// SysAuthRuleTreeRes menu tree structure
type SysAuthRuleTreeRes struct {
	*SysAuthRuleInfoRes
	Children []*SysAuthRuleTreeRes `json:"children"`
}

type UserMenu struct {
	Id        uint   `json:"id"`
	Pid       uint   `json:"pid"`
	Name      string `json:"name"`
	Component string `json:"component"`
	Path      string `json:"path"`
	*MenuMeta `json:"meta"`
}

type UserMenus struct {
	*UserMenu `json:""`
	Children  []*UserMenus `json:"children"`
}

type MenuMeta struct {
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	IsLink      string `json:"isLink"`
	IsHide      bool   `json:"isHide"`
	IsKeepAlive bool   `json:"isKeepAlive"`
	IsAffix     bool   `json:"isAffix"`
	IsIframe    bool   `json:"isIframe"`
}
