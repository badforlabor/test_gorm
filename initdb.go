/**
 * Auth :   liubo
 * Date :   2021/9/13 18:04
 * Comment: 连接数据库
			如果没有库，自动创建
			如果没有表，自动创建表
 */

package main

import (
	"database/sql"
	"fmt"
	"github.com/badforlabor/gocrazy/crazy3rd/glog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"test_grom/model"
)

type Mysql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	LogZap       string `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
}

func GormMysql() *gorm.DB {
	var m = Mysql{}

	m.Username = "root"
	m.Password = "123456"
	m.Path = "127.0.0.1:3306"
	m.Config = "charset=utf8mb4&parseTime=True&loc=Local"
	m.MaxIdleConns = 0
	m.MaxOpenConns = 0
	m.LogMode = true
	m.Dbname = "db123"

	if m.Dbname == "" {
		return nil
	}

	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	// 如果没有表，那么创建
	createDbIfNotExist(&m)

	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode)); err != nil {
		//global.GVA_LOG.Error("MySQL启动异常", zap.Any("err", err))
		//os.Exit(0)
		//return nil
		glog.Errorln("MYSQL启动失败...", err.Error())
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}

}

// 注册数据库表，如果没有，会自动创建对应的表
func RegisterAllSqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.SysUser{},
		model.SysAuthority{},
		model.SysBaseMenu{},
		model.SysBaseMenuParameter{},
	)

	if err != nil {
		glog.Errorln("register db table failed.", err.Error())
	}
}


func createDbIfNotExist(conf *Mysql) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", conf.Username, conf.Password, conf.Path)
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.Dbname)
	execSql(dsn, "mysql", createSql)
}

func execSql(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

func gormConfig(mod bool) *gorm.Config {
	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	return config
}
