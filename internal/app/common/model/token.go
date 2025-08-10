/*
* @desc:token options
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu
* @Date:   2022/3/8 16:02
 */

package model

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TokenOptions struct {
	//  server name
	ServerName string `json:"serverName"`
	// cachekey (theCacheKeymust be unique for each instance)
	CacheKey string `json:"cacheKey"`
	// timeout (default 10 days (seconds))
	Timeout int64 `json:"timeout"`
	// cache refresh time (default 5 days (seconds))
	// When processing a request with a token, if the current time is greater than the timeout and less than the cache refresh time, the token will automatically refresh, i.e., reset the token lifetime
	// If the MaxRefresh value is 0, the token will not automatically refresh
	MaxRefresh int64 `json:"maxRefresh"`
	// Allow multiple logins
	MultiLogin bool `json:"multiLogin"`
	// Token encryption key (32 bits)
	EncryptKey []byte `json:"encryptKey"`
	// Block excluded URLs
	ExcludePaths g.SliceStr `json:"excludePaths"`
	CacheModel   string     `json:"cacheModel"`
}
