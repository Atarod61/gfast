package system

import (
	"github.com/gogf/gf/v2/frame/g"
)

type MonitorSearchReq struct {
	g.Meta `path:"/monitor/server" tags:"Service Monitoring" method:"get" summary:"Server Monitoring"`
}

type MonitorSearchRes g.Map
