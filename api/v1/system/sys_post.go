/*
* @desc:Position-related parameters
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu
* @Date:   2022/4/7 23:09
 */

package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

type PostSearchReq struct {
	g.Meta   `path:"/post/list" tags:"Post Management" method:"get" summary:"Post List"`
	PostCode string `p:"postCode"` //Position code
	PostName string `p:"postName"` //Position name
	Status   string `p:"status"`   //Status
	commonApi.PageReq
}

type PostSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	PostList []*entity.SysPost `json:"postList"`
}

type PostAddReq struct {
	g.Meta   `path:"/post/add" tags:"Post Management" method:"post" summary:"Add Post"`
	PostCode string `p:"postCode" v:"required#Post code cannot be empty"`
	PostName string `p:"postName" v:"required#Post name cannot be empty"`
	PostSort int    `p:"postSort" v:"required#Post sorting cannot be empty"`
	Status   uint   `p:"status" v:"required#Status cannot be empty"`
	Remark   string `p:"remark"`
}

type PostAddRes struct {
}

type PostEditReq struct {
	g.Meta   `path:"/post/edit" tags:"Status cannot be empty" method:"put" summary:"Edit Post"`
	PostId   int64  `p:"postId" v:"required#ID is required"`
	PostCode string `p:"postCode" v:"required#Post code cannot be empty"`
	PostName string `p:"postName" v:"required#Post name cannot be empty"`
	PostSort int    `p:"postSort" v:"required#Post sorting cannot be empty"`
	Status   uint   `p:"status" v:"required#Status cannot be empty"`
	Remark   string `p:"remark"`
}

type PostEditRes struct {
}

type PostDeleteReq struct {
	g.Meta `path:"/post/delete" tags:"Post Management" method:"delete" summary:"Delete Post"`
	Ids    []int `p:"ids"`
}

type PostDeleteRes struct {
}

