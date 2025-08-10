package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

type UserMenusReq struct {
	g.Meta `path:"/user/getUserMenus" tags:"User Management" method:"get" summary:"Get user menus"`
	commonApi.Author
}

type UserMenusRes struct {
	g.Meta      `mime:"application/json"`
	MenuList    []*model.UserMenus `json:"menuList"`
	Permissions []string           `json:"permissions"`
}

// UserSearchReq user search request parameters
type UserSearchReq struct {
	g.Meta   `path:"/user/list" tags:"User Management" method:"get" summary:"User list"`
	DeptId   string `p:"deptId"` //Departmentid
	Mobile   string `p:"mobile"`
	Status   string `p:"status"`
	KeyWords string `p:"keyWords"`
	commonApi.PageReq
	commonApi.Author
}

type UserSearchRes struct {
	g.Meta   `mime:"application/json"`
	UserList []*model.SysUserRoleDeptRes `json:"userList"`
	commonApi.ListRes
}

type UserGetParamsReq struct {
	g.Meta `path:"/user/params" tags:"User Management" method:"get" summary:"Get user maintenance parameters"`
}

type UserGetParamsRes struct {
	g.Meta   `mime:"application/json"`
	RoleList []*entity.SysRole `json:"roleList"`
	Posts    []*entity.SysPost `json:"posts"`
}

// SetUserReq adds and modifies common user request fields
type SetUserReq struct {
	DeptId   uint64  `p:"deptId" v:"required#Department cannot be empty"` //Department
	Email    string  `p:"email" v:"email#Invalid email format"`       //Email address
	NickName string  `p:"nickName" v:"required#Nickname cannot be empty"`
	Mobile   string  `p:"mobile" v:"required|phone#Mobile cannot be empty|Invalid mobile format"`
	PostIds  []int64 `p:"postIds"`
	Remark   string  `p:"remark"`
	RoleIds  []int64 `p:"roleIds"`
	Sex      int     `p:"sex"`
	Status   uint    `p:"status"`
	IsAdmin  int     `p:"isAdmin"` // Whether the background administrator 1 Yes  0   No
}

// UserAddReq adds user parameters
type UserAddReq struct {
	g.Meta `path:"/user/add" tags:"User Management" method:"post" summary:"Add user"`
	*SetUserReq
	UserName string `p:"userName" v:"required#Username cannot be empty"`
	Password string `p:"password" v:"required|password#Password cannot be empty|Password must start with a letter，only contain letters、numbers and underscores，and must be 6~18 characters long"`
	UserSalt string
}

type UserAddRes struct {
}

// UserEditReq Modify user parameters
type UserEditReq struct {
	g.Meta `path:"/user/edit" tags:"User Management" method:"put" summary:"Edit user"`
	*SetUserReq
	UserId int64 `p:"userId" v:"required#User ID is required"`
}

type UserEditRes struct {
}

type UserGetEditReq struct {
	g.Meta `path:"/user/getEdit" tags:"User Management" method:"get" summary:"Get user details"`
	Id     uint64 `p:"id"`
}

type UserGetEditRes struct {
	g.Meta         `mime:"application/json"`
	User           *entity.SysUser `json:"user"`
	CheckedRoleIds []uint          `json:"checkedRoleIds"`
	CheckedPosts   []int64         `json:"checkedPosts"`
}

// UserResetPwdReq resets user password status parameters
type UserResetPwdReq struct {
	g.Meta   `path:"/user/resetPwd" tags:"User Management" method:"put" summary:"Reset user password"`
	Id       uint64 `p:"userId" v:"required#User ID is required"`
	Password string `p:"password" v:"required|password#Password required|Password must start with a letter，can only contain letters、numbers and underscores，and must be 6~18 characters long"`
}

type UserResetPwdRes struct {
}

// UserStatusReq sets user status parameters
type UserStatusReq struct {
	g.Meta     `path:"/user/setStatus" tags:"User Management" method:"put" summary:"Set user status"`
	Id         uint64 `p:"userId" v:"required#User ID is required"`
	UserStatus uint   `p:"status" v:"required#User status is required"`
}

type UserStatusRes struct {
}

type UserDeleteReq struct {
	g.Meta `path:"/user/delete" tags:"User Management" method:"delete" summary:"Delete user"`
	Ids    []int `p:"ids"  v:"required#ids are required"`
}

type UserDeleteRes struct {
}

type UserGetByIdsReq struct {
	g.Meta `path:"/user/getUsers" tags:"User Management" method:"get" summary:"Get multiple users"`
	commonApi.Author
	Ids []int `p:"ids" v:"required#ids are required"`
}

type UserGetByIdsRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.SysUserSimpleRes `json:"list"`
}
