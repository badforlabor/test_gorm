/**
 * Auth :   liubo
 * Date :   2021/9/13 17:15
 * Comment: 测试：增删改查，数目
 */

package main

import (
	"errors"
	"github.com/badforlabor/gocrazy/crazy3rd/glog"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strconv"
	"test_grom/model"
)

func addOneUser(user *model.SysUser) {

	// 比较username，如果没有找到，就创建一个
	var temp model.SysUser
	if errors.Is(globalDb.Limit(1).Where("username=?", user.Username).First(&temp).Error, gorm.ErrRecordNotFound) {
		user.UUID = uuid.NewV4()
		user.NickName = user.Username
		var e = globalDb.Create(user).Error
		if e != nil {
			glog.Info("db add user failed.", e.Error())
		} else {
			glog.Info("db add user succ.", user.Username)
		}
	}

}

// 创建1万个用户
func testCreateUser() {
	var baseid = 100000
	for i:=0; i<1000; i++ {
		var id = baseid + i
		var user = &model.SysUser{}
		user.Username = "test-" + strconv.Itoa(id)

		addOneUser(user)
	}
}

func testLoadUser() {
	var alluser []model.SysUser
	globalDb.Find(&alluser)
	glog.Info("user count:", len(alluser))
	if len(alluser) > 0 {
		glog.Infoln("load user[0]=", alluser[0].Username)
	}
}

func testDeleteUser() {

	var user model.SysUser

	var deletedUser = "test-100000"
	addOneUser(&model.SysUser{Username:deletedUser})

	var count1 int64
	globalDb.Model(&user).Count(&count1)

	globalDb.Where("username=?", "test-100000").Delete(&user)

	var count2 int64
	globalDb.Model(&user).Count(&count2)

	glog.Infoln("delete action, user count before=", count1, ", after=", count2)
}
func testUpdateUser() {
	var olduser = model.SysUser{Username:"test-100001"}
	var newNickName = "nickname-error-10000"

	var user model.SysUser

	// 更新一些列
	globalDb.Model(&user).Where("username=?", olduser.Username).Updates(&model.SysUser{NickName:newNickName})

	// 只更新某列
	globalDb.Model(&user).Where("username=?", olduser.Username).Update("header_img", "xxxx")

	globalDb.Where("username=?", olduser.Username).Find(&user)

	glog.Infoln("update user, new one=", user)


}