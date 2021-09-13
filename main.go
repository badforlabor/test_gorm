/**
 * Auth :   liubo
 * Date :   2021/9/13 16:07
 * Comment:
 */

package main

import (
	"flag"
	"github.com/badforlabor/gocrazy/crazy3rd/glog"
	"github.com/badforlabor/gocrazy/crazyapp"
	"gorm.io/gorm"
)


var globalDb *gorm.DB

func main() {
	flag.Parse()

	crazyapp.CallNormalMain(func() {
		glog.Info("test grom")

		globalDb = GormMysql()
		if globalDb == nil {
			panic("sql is nil")
		}
		RegisterAllSqlTables(globalDb)

		glog.Infoln("初始化完毕...")

		testCreateUser()
		testLoadUser()
		testDeleteUser()
		testUpdateUser()

		// 执行查询操纵

	})
}
