/*
* @desc:xxxxFunctional Description
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/11/3 10:04
 */

package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

type PersonalInfoReq struct {
	g.Meta `path:"/personal/getPersonalInfo" tags:"User Management" method:"get" summary:"Logged-in User Information"`
	commonApi.Author
}

type PersonalInfoRes struct {
	g.Meta   `mime:"application/json"`
	User     *entity.SysUser `json:"user"`
	Roles    []string        `json:"roles"`
	DeptName string          `json:"deptName"`
}

// SetPersonalReq adds and modifies user public request fields
type SetPersonalReq struct {
	Nickname  string `p:"nickname" v:"required#Nickname cannot be empty"`
	Mobile    string `p:"mobile" v:"required|phone#Mobile number cannot be empty|Invalid phone number format"`
	Remark    string `p:"remark"`
	Sex       int    `p:"sex"`
	UserEmail string `p:"userEmail" v:"required|email#Email cannot be empty|Invalid email format"`
	Describe  string `p:"describe"`
	Avatar    string `p:"avatar"`
}

// PersonalEditReq Modify personal information
type PersonalEditReq struct {
	g.Meta `path:"/personal/edit" tags:"User Management" method:"put" summary:"Update Profile"`
	*SetPersonalReq
	commonApi.Author
}

type PersonalEditRes struct {
	commonApi.EmptyRes
	UserInfo *model.LoginUserRes `json:"userInfo"`
	Token    string              `json:"token"`
}

type PersonalResetPwdReq struct {
	g.Meta   `path:"/personal/resetPwd" tags:"User Management" method:"put" summary:"Reset Personal Password"`
	Password string `p:"password" v:"required|password#Password cannot be empty|Password must start with a letter，only contain letters、numbers and underscores，and be 6~18 characters long"`
	commonApi.Author
}

type PersonalResetPwdRes struct {
}
