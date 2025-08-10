/*
* @desc:Routing Processing
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/11/16 11:09
 */

package libRouter

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gregex"
	"reflect"
)

// RouterAutoBindBefore  collects controllers that need to be bound without verifying the user's login status and automatically binds them
// Routing method naming conventions must be：BeforeBindXXXController
func RouterAutoBindBefore(ctx context.Context, R interface{}, group *ghttp.RouterGroup) (err error) {
	return bind(ctx, R, group, "before")
}

// RouterAutoBind collects controllers that need to be bound and automatically binds them
// The routing method naming rule must be：BindXXXController
func RouterAutoBind(ctx context.Context, R interface{}, group *ghttp.RouterGroup) (err error) {
	return bind(ctx, R, group)
}

func bind(ctx context.Context, R interface{}, group *ghttp.RouterGroup, option ...string) (err error) {
	var rule string
	if len(option) > 0 && option[0] == "before" {
		rule = `^BeforeBind(.+)Controller$`
	} else {
		rule = `^Bind(.+)Controller$`
	}
	//TypeOf will return the type of the target data, such as int/float/struct/pointer, etc
	typ := reflect.TypeOf(R)
	//ValueOf returns the value of the target data
	val := reflect.ValueOf(R)
	if val.Elem().Kind() != reflect.Struct {
		err = gerror.New("expect struct but a " + val.Elem().Kind().String())
		return
	}
	for i := 0; i < typ.NumMethod(); i++ {
		if match := gregex.IsMatchString(rule, typ.Method(i).Name); match {
			//Call binding method
			val.Method(i).Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(group)})
		}
	}
	return
}
