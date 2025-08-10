/*
* @desc:Returns common response parameters
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/10/27 16:30
 */

package common

import "github.com/gogf/gf/v2/frame/g"

// EmptyRes does not respond with any data
type EmptyRes struct {
	g.Meta `mime:"application/json"`
}

// ListRes returns a list of common parameters
type ListRes struct {
	CurrentPage int         `json:"currentPage"`
	Total       interface{} `json:"total"`
}
