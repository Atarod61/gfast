/*
* @desc:Dictionary Data Management
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/28 9:22
 */

package sysDictData

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/common/consts"
	"github.com/tiger1103/gfast/v3/internal/app/common/dao"
	"github.com/tiger1103/gfast/v3/internal/app/common/model"
	"github.com/tiger1103/gfast/v3/internal/app/common/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/common/service"
	systemConsts "github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterSysDictData(New())
}

func New() *sSysDictData {
	return &sSysDictData{}
}

type sSysDictData struct {
}

// GetDictWithDataByType Get options by dictionary key type
func (s *sSysDictData) GetDictWithDataByType(ctx context.Context, dictType, defaultValue string) (dict *system.GetDictRes,
	err error) {
	cache := service.Cache()
	cacheKey := consts.CacheSysDict + "_" + dictType
	//Get from cache
	iDict := cache.GetOrSetFuncLock(ctx, cacheKey, func(ctx context.Context) (value interface{}, err error) {
		err = g.Try(ctx, func(ctx context.Context) {
			//Get from database
			dict = &system.GetDictRes{}
			//Get type data
			err = dao.SysDictType.Ctx(ctx).Where(dao.SysDictType.Columns().DictType, dictType).
				Where(dao.SysDictType.Columns().Status, 1).Fields(model.DictTypeRes{}).Scan(&dict.Info)
			liberr.ErrIsNil(ctx, err, "Failed to obtain dictionary type")
			if dict.Info == nil {
				return
			}
			err = dao.SysDictData.Ctx(ctx).Fields(model.DictDataRes{}).
				Where(dao.SysDictData.Columns().DictType, dictType).
				Where(dao.SysDictData.Columns().Status, 1).
				Order(dao.SysDictData.Columns().DictSort + " asc," +
					dao.SysDictData.Columns().DictCode + " asc").
				Scan(&dict.Values)
			liberr.ErrIsNil(ctx, err, "Failed to obtain dictionary data")
		})
		value = dict
		return
	}, 0, consts.CacheSysDictTag)
	if !iDict.IsEmpty() {
		err = gconv.Struct(iDict, &dict)
		if err != nil {
			return
		}
	}
	//Set the given default value
	for _, v := range dict.Values {
		if defaultValue != "" {
			if gstr.Equal(defaultValue, v.DictValue) {
				v.IsDefault = 1
			} else {
				v.IsDefault = 0
			}
		}
	}
	return
}

// List Get dictionary data
func (s *sSysDictData) List(ctx context.Context, req *system.DictDataSearchReq) (res *system.DictDataSearchRes, err error) {
	res = new(system.DictDataSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysDictData.Ctx(ctx)
		if req != nil {
			if req.DictLabel != "" {
				m = m.Where(dao.SysDictData.Columns().DictLabel+" like ?", "%"+req.DictLabel+"%")
			}
			if req.Status != "" {
				m = m.Where(dao.SysDictData.Columns().Status+" = ", gconv.Int(req.Status))
			}
			if req.DictType != "" {
				m = m.Where(dao.SysDictData.Columns().DictType+" = ?", req.DictType)
			}
			res.Total, err = m.Count()
			liberr.ErrIsNil(ctx, err, "Failed to obtain dictionary data")
			if req.PageNum == 0 {
				req.PageNum = 1
			}
			res.CurrentPage = req.PageNum
		}
		if req.PageSize == 0 {
			req.PageSize = systemConsts.PageSize
		}
		err = m.Page(req.PageNum, req.PageSize).Order(dao.SysDictData.Columns().DictSort + " asc," +
			dao.SysDictData.Columns().DictCode + " asc").Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "Failed to obtain dictionary data")
	})
	return
}

func (s *sSysDictData) Add(ctx context.Context, req *system.DictDataAddReq, userId uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysDictData.Ctx(ctx).Insert(do.SysDictData{
			DictSort:  req.DictSort,
			DictLabel: req.DictLabel,
			DictValue: req.DictValue,
			DictType:  req.DictType,
			CssClass:  req.CssClass,
			ListClass: req.ListClass,
			IsDefault: req.IsDefault,
			Status:    req.Status,
			CreateBy:  userId,
			Remark:    req.Remark,
		})
		liberr.ErrIsNil(ctx, err, "Failed to add dictionary data")
		//Clear cache
		service.Cache().RemoveByTag(ctx, consts.CacheSysDictTag)
	})
	return
}

// Get dictionary data
func (s *sSysDictData) Get(ctx context.Context, dictCode uint) (res *system.DictDataGetRes, err error) {
	res = new(system.DictDataGetRes)
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysDictData.Ctx(ctx).WherePri(dictCode).Scan(&res.Dict)
		liberr.ErrIsNil(ctx, err, "Failed to get dictionary data")
	})
	return
}

// Edit dictionary data
func (s *sSysDictData) Edit(ctx context.Context, req *system.DictDataEditReq, userId uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysDictData.Ctx(ctx).WherePri(req.DictCode).Update(do.SysDictData{
			DictSort:  req.DictSort,
			DictLabel: req.DictLabel,
			DictValue: req.DictValue,
			DictType:  req.DictType,
			CssClass:  req.CssClass,
			ListClass: req.ListClass,
			IsDefault: req.IsDefault,
			Status:    req.Status,
			UpdateBy:  userId,
			Remark:    req.Remark,
		})
		liberr.ErrIsNil(ctx, err, "Failed to modify dictionary data")
		//Clear cache
		service.Cache().RemoveByTag(ctx, consts.CacheSysDictTag)
	})
	return
}

// Delete dictionary data
func (s *sSysDictData) Delete(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysDictData.Ctx(ctx).Where(dao.SysDictData.Columns().DictCode+" in(?)", ids).Delete()
		liberr.ErrIsNil(ctx, err, "Failed to delete dictionary data")
		//Clear cache
		service.Cache().RemoveByTag(ctx, consts.CacheSysDictTag)
	})
	return
}
