/*
* @desc:System parameter configuration
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu
* @Date:   2022/4/18 21:11
 */

package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	commonEntity "github.com/tiger1103/gfast/v3/internal/app/common/model/entity"
)

type ConfigSearchReq struct {
	g.Meta     `path:"/config/list" tags:"System Configuration Management" method:"get" summary:"System Configuration List"`
	ConfigName string `p:"configName"` //Parameter name
	ConfigKey  string `p:"configKey"`  //Parameter key name
	ConfigType string `p:"configType"` //Status
	commonApi.PageReq
}

type ConfigSearchRes struct {
	g.Meta `mime:"application/json"`
	List   []*commonEntity.SysConfig `json:"list"`
	commonApi.ListRes
}

type ConfigReq struct {
	ConfigName  string `p:"configName"  v:"required#Parameter name cannot be empty"`
	ConfigKey   string `p:"configKey"  v:"required#Parameter key cannot be empty"`
	ConfigValue string `p:"configValue"  v:"required#Parameter value cannot be empty"`
	ConfigType  int    `p:"configType"    v:"required|in:0,1#System built-in flag is required|System built-in type must be 0 or 1"`
	Remark      string `p:"remark"`
}

type ConfigAddReq struct {
	g.Meta `path:"/config/add" tags:"System Configuration Management" method:"post" summary:"Add System Configuration"`
	*ConfigReq
}

type ConfigAddRes struct {
}

type ConfigGetReq struct {
	g.Meta `path:"/config/get" tags:"System Configuration Management" method:"get" summary:"Get System Configuration"`
	Id     int `p:"id"`
}

type ConfigGetRes struct {
	g.Meta `mime:"application/json"`
	Data   *commonEntity.SysConfig `json:"data"`
}

type ConfigEditReq struct {
	g.Meta   `path:"/config/edit" tags:"System Configuration Management" method:"put" summary:"Edit System Configuration"`
	ConfigId int64 `p:"configId" v:"required|min:1#Primary key ID cannot be empty|Invalid primary key ID parameter"`
	*ConfigReq
}

type ConfigEditRes struct {
}

type ConfigDeleteReq struct {
	g.Meta `path:"/config/delete" tags:"System Configuration Management" method:"delete" summary:"Delete System Configuration"`
	Ids    []int `p:"ids"`
}

type ConfigDeleteRes struct {
}
