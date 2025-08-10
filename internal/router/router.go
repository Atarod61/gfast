/*
* @desc:Route Binding
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu
* @Date:   2022/2/18 16:23
 */

package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	commonRouter "github.com/tiger1103/gfast/v3/internal/app/common/router"
	commonService "github.com/tiger1103/gfast/v3/internal/app/common/service"
	researchRouter "github.com/tiger1103/gfast/v3/internal/app/research/router"
	systemRouter "github.com/tiger1103/gfast/v3/internal/app/system/router"
	"github.com/tiger1103/gfast/v3/library/libRouter"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/api/v1", func(group *ghttp.RouterGroup) {
		//Cross-origin processing. For security reasons, comment this line in the production environment
		group.Middleware(commonService.Middleware().MiddlewareCORS)
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		// Bind backend routes
		systemRouter.R.BindController(ctx, group)

		researchRouter.R.BindController(ctx, group)
		// Bind public routes
		commonRouter.R.BindController(ctx, group)
		//Automatically bind defined modules
		if err := libRouter.RouterAutoBind(ctx, router, group); err != nil {
			panic(err)
		}
	})
}
