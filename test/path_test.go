/*
* @desc:xxxxFunction Description
* @company:Yunnan Qixun Technology Co., Ltd
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/12/8 17:48
 */

package test

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"strings"
	"testing"
)

func TestContainer(t *testing.T) {
	//t.Run("testEncryptPassword", testEncryptPassword)
	//t.Run("testIpToAddr", testIpToAddr)
	//t.Run("testDb", testDb)
	t.Run("ScanPath", scanPath)
}

func scanPath(t *testing.T) {
	var (
		packages     []string
		err          error
		packageNames []string
	)
	//logicpath
	path := strings.Join([]string{"D:/goProject/p2022/gfast-v3/internal/app/demo/logic"}, "")
	//Get all packages to be loaded under this path
	packages, err = gfile.ScanDirFunc(path, "*", false, func(p string) string {
		if gfile.IsDir(p) {
			return p
		}
		return ""
	})
	if err != nil {
		panic(err)
	}
	packageNames = make([]string, len(packages))
	for k, v := range packages {
		fmt.Println("111111", v)
		v = gstr.Replace(v, "\\", "/")
		fmt.Println("2222222", v)
		packageNames[k] = gstr.SubStr(v, gstr.PosR(v, "/")+1)
	}
	fmt.Println("dddd", packageNames)
}
