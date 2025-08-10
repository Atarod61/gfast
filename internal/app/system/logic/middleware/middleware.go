/*
* @desc:Middleware
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/23 15:05
 */

package middleware

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	commonService "github.com/tiger1103/gfast/v3/internal/app/common/service"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libResponse"
)

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}

type sMiddleware struct{}

// Ctx custom context object
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	ctx := r.GetCtx()
	// Initialize login user information
	data, err := service.GfToken().ParseToken(r)
	if err != nil {
		// Execute the next request logic
		r.Middleware.Next()
	}
	if data != nil {
		context := new(model.Context)
		err = gconv.Struct(data.Data, &context.User)
		if err != nil {
			g.Log().Error(ctx, err)
			// Execute the next request logic
			r.Middleware.Next()
		}
		service.Context().Init(r, context)
	}
	// Execute the next request logic
	r.Middleware.Next()
}

// Auth permission judgment processing middleware
func (s *sMiddleware) Auth(r *ghttp.Request) {
	ctx := r.GetCtx()
	//Get the login user id
	adminId := service.Context().GetUserId(ctx)
	accessParams := r.Get("accessParams").Strings()
	accessParamsStr := ""
	if len(accessParams) > 0 && accessParams[0] != "undefined" {
		accessParamsStr = "?" + gstr.Join(accessParams, "&")
	}
	url := gstr.TrimLeft(r.Request.URL.Path, "/") + accessParamsStr
	/*if r.Method != "GET" && adminId != 1 && url!="api/v1/system/login" {
		libResponse.FailJson(true, r, "Sorry! This is a demo system, data cannot be deleted or modified!")
	}*/
	//Get the user ID without authentication
	tagSuperAdmin := false
	service.SysUser().NotCheckAuthAdminIds(ctx).Iterator(func(v interface{}) bool {
		if gconv.Uint64(v) == adminId {
			tagSuperAdmin = true
			return false
		}
		return true
	})
	if tagSuperAdmin {
		r.Middleware.Next()
		//Don't execute further
		return
	}
	//Get the menu ID corresponding to the address
	menuList, err := service.SysAuthRule().GetMenuList(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
		libResponse.FailJson(true, r, "data request failed")
	}
	var menu *model.SysAuthRuleInfoRes
	for _, m := range menuList {
		ms := gstr.SubStr(m.Name, 0, gstr.Pos(m.Name, "?"))
		if m.Name == url || ms == url {
			menu = m
			break
		}
	}
	//Only validate rules that exist in the database
	if menu != nil {
		//Do not check permissions if the interface is accessible without logging in
		excludePaths := g.Cfg().MustGet(ctx, "gfToken.excludePaths").Strings()
		for _, p := range excludePaths {
			if gstr.Equal(menu.Name, gstr.TrimLeft(p, "/")) {
				r.Middleware.Next()
				return
			}
		}
		//Skip if a condition doesn't require verification
		if gstr.Equal(menu.Condition, "nocheck") {
			r.Middleware.Next()
			return
		}
		menuId := menu.Id
		//Menu not stored in the database, no permissions verified
		if menuId != 0 {
			//Check permissions and perform operations
			enforcer, err := commonService.CasbinEnforcer(ctx)
			if err != nil {
				g.Log().Error(ctx, err)
				libResponse.FailJson(true, r, "Failed to get permission")
			}
			hasAccess := false
			hasAccess, err = enforcer.Enforce(fmt.Sprintf("%s%d", service.SysUser().GetCasBinUserPrefix(), adminId), gconv.String(menuId), "All")
			if err != nil {
				g.Log().Error(ctx, err)
				libResponse.FailJson(true, r, "permission determination failed")
			}
			if !hasAccess {
				libResponse.FailJson(true, r, "no access permission")
			}
		}
	} else if menu == nil && accessParamsStr != "" {
		libResponse.FailJson(true, r, "no access permission")
	}
	r.Middleware.Next()
}
