/*
* @desc:Dictionary type
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/3/18 11:56
 */

package model

import "github.com/gogf/gf/v2/os/gtime"

type SysDictTypeInfoRes struct {
	DictId    uint64      `orm:"dict_id,primary"  json:"dictId"`    // Dictionary primary key
	DictName  string      `orm:"dict_name"        json:"dictName"`  // Dictionary name
	DictType  string      `orm:"dict_type,unique" json:"dictType"`  // Dictionary type
	Status    uint        `orm:"status"           json:"status"`    // Status (0Normal  1Disabled)
	Remark    string      `orm:"remark"           json:"remark"`    // Remarks
	CreatedAt *gtime.Time `orm:"created_at"       json:"createdAt"` // Creation date
}
