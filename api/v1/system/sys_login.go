/*
* @desc:Login
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu
* @Date:   2022/4/27 21:51
 */

package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

type UserLoginReq struct {
	g.Meta     `path:"/login" tags:"Authentication" method:"post" summary:"User Login"`
	Username   string `p:"username" v:"required#Username cannot be empty"`
	Password   string `p:"password" v:"required#Password cannot be empty"`
	VerifyCode string `p:"verifyCode" v:"required#Verification code cannot be empty"`
	VerifyKey  string `p:"verifyKey"`
}

type UserLoginRes struct {
	g.Meta      `mime:"application/json"`
	UserInfo    *model.LoginUserRes `json:"userInfo"`
	Token       string              `json:"token"`
	MenuList    []*model.UserMenus  `json:"menuList"`
	Permissions []string            `json:"permissions"`
}

type UserLoginOutReq struct {
	g.Meta `path:"/logout" tags:"Authentication" method:"get" summary:"User Logout"`
	commonApi.Author
}

type UserLoginOutRes struct {
}
