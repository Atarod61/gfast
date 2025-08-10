/*
* @desc:Public interface related
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/3/30 9:28
 */

package common

import "github.com/tiger1103/gfast/v3/internal/app/common/model"

// PageReq public request parameters
type PageReq struct {
	model.PageReq
}

type Author struct {
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}"`
}
