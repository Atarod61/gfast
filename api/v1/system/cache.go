/*
* @desc:Cache Processing
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2023/2/1 18:12
 */

package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
)

type CacheRemoveReq struct {
	g.Meta `path:"/cache/remove" tags:"Cache Management" method:"delete" summary:"Clear Cache"`
	commonApi.Author
}

type CacheRemoveRes struct {
	commonApi.EmptyRes
}
