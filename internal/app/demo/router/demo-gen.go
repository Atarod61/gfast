package router


import (
    "context"
    "github.com/gogf/gf/v2/net/ghttp"
    "github.com/tiger1103/gfast/v3/internal/app/demo/controller"
)
func (router *Router) BindDemoGenController(ctx context.Context, group *ghttp.RouterGroup) {
    group.Group("/demoGen", func(group *ghttp.RouterGroup) {
        group.Bind(
            controller.DemoGen,
        )
    })
}