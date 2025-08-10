package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/research/controller"
)

func (router *Router) BindResearchController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/Research", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.Research,
		)
	})
}
