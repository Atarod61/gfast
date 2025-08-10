/*
* @desc:User model object
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu
* @Date:   2022/3/7 11:47
 */

package model

import (
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

// LoginUserRes login return
type LoginUserRes struct {
	Id           uint64 `orm:"id,primary"       json:"id"`           //
	UserName     string `orm:"user_name,unique" json:"userName"`     // Username
	UserNickname string `orm:"user_nickname"    json:"userNickname"` // User nickname
	UserPassword string `orm:"user_password"    json:"userPassword"` // Login password; encrypted with cmf_password
	UserSalt     string `orm:"user_salt"        json:"userSalt"`     // Encrypted salt
	UserStatus   uint   `orm:"user_status"      json:"userStatus"`   // User status; 0: Disabled, 1: Normal, 2: Unverified
	IsAdmin      int    `orm:"is_admin"         json:"isAdmin"`      // Backend administrator status; 1: Yes; 0: No
	Avatar       string `orm:"avatar" json:"avatar"`                 //Avatar
	DeptId       uint64 `orm:"dept_id"       json:"deptId"`          //Department ID
}

// SysUserRoleDeptRes User data with department, role, and position information
type SysUserRoleDeptRes struct {
	*entity.SysUser
	Dept     *entity.SysDept       `json:"dept"`
	RoleInfo []*SysUserRoleInfoRes `json:"roleInfo"`
	Post     []*SysUserPostInfoRes `json:"post"`
}

type SysUserRoleInfoRes struct {
	RoleId uint   `json:"roleId"`
	Name   string `json:"name"`
}

type SysUserPostInfoRes struct {
	PostId   int64  `json:"postId"`
	PostName string `json:"postName"`
}

type SysUserSimpleRes struct {
	gmeta.Meta   `orm:"table:sys_user"`
	Id           uint64 `orm:"id"       json:"id"`                   //
	Avatar       string `orm:"avatar" json:"avatar"`                 // Avatar
	Sex          int    `orm:"sex" json:"sex"`                       // Gender
	UserName     string `orm:"user_name" json:"userName"`            // Username
	UserNickname string `orm:"user_nickname"    json:"userNickname"` // User nickname
}
