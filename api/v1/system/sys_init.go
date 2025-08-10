package system

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

type DbInitIsInitReq struct {
	g.Meta `path:"/dbInit/isInit" tags:"System Initialization" method:"get" summary:"Check System Initialization Status"`
}

type DbInitIsInitRes bool

type DbInitGetEnvInfoReq struct {
	g.Meta `path:"/dbInit/getEnvInfo" tags:"System Initialization" method:"get" summary:"Get Environment Information"`
}

type DbInitGetEnvInfoRes g.Map

type DbInitCreateDbReq struct {
	g.Meta       `path:"/dbInit/createDb" tags:"System Initialization" method:"post" summary:"Create Configuration File"`
	DbHost       string `json:"dbHost" p:"dbHost" v:"required#Database host is required"`
	DbPort       int    `json:"dbPort" p:"dbPort" v:"required#Database port is required"`
	DbUser       string `json:"dbUser" p:"dbUser" v:"required#Database username is required"`
	DbPass       string `json:"dbPass"`
	DbName       string `json:"dbName" p:"dbName" v:"required#Database name is required"`
	DbCharset    string `json:"dbCharset" p:"dbCharset" v:"required#Database charset is required"`
	RedisAddress string `json:"redisAddress" p:"redisAddress" v:"required#RedisAddress"`
	RedisPort    int    `json:"redisPort" p:"redisPort" v:"required#RedisPort"`
	RedisDb      int    `json:"redisDb" p:"redisDb" v:"required#RedisDb"`
	RedisPass    string `json:"redisPass"`
}

type DbInitCreateDbRes bool

func (req *DbInitCreateDbReq) ToDbInitConfig() *model.DbInitConfig {
	return &model.DbInitConfig{
		Database: model.Database{
			Default: model.DbDefault{
				Host:        req.DbHost,
				Port:        req.DbPort,
				User:        req.DbUser,
				Pass:        req.DbPass,
				Name:        req.DbName,
				Type:        "mysql",
				Role:        "master",
				Debug:       true,
				Charset:     req.DbCharset,
				DryRun:      false,
				MaxIdle:     10,
				MaxOpen:     10,
				MaxLifetime: 10,
			},
		},
		Redis: model.Redis{
			Default: model.RedisDefault{
				Address:     fmt.Sprintf("%s:%d", req.RedisAddress, req.RedisPort),
				Db:          req.RedisDb,
				Pass:        req.RedisPass,
				IdleTimeout: 600,
				MaxActive:   100,
			},
		},
	}
}
