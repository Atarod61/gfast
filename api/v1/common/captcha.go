/*
* @desc:Verification code parameters
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu
* @Date:   2022/3/2 17:47
 */

package common

import "github.com/gogf/gf/v2/frame/g"

type CaptchaReq struct {
	g.Meta `path:"/get" tags:"Captcha" method:"get" summary:"Get captcha"`
}
type CaptchaRes struct {
	g.Meta `mime:"application/json"`
	Key    string `json:"key"`
	Img    string `json:"img"`
}
