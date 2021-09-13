/**
 * Auth :   liubo
 * Date :   2021/9/13 17:04
 * Comment: 数据模型
 */

package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type GVA_MODEL struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Meta struct {
	KeepAlive   bool   `json:"keepAlive" gorm:"comment:是否缓存"`
	DefaultMenu bool   `json:"defaultMenu" gorm:"comment:是否是基础路由（开发中）"`
	Title       string `json:"title" gorm:"comment:菜单名"`
	Icon        string `json:"icon" gorm:"comment:菜单图标"`
	CloseTab    bool   `json:"closeTab" gorm:"comment:自动关闭tab"`
}


type SysBaseMenuParameter struct {
	GVA_MODEL
	SysBaseMenuID uint
	Type          string `json:"type" gorm:"comment:地址栏携带参数为params还是query"`
	Key           string `json:"key" gorm:"comment:地址栏携带参数的key"`
	Value         string `json:"value" gorm:"comment:地址栏携带参数的值"`
}

type SysBaseMenu struct {
	GVA_MODEL
	MenuLevel     uint   `json:"-"`
	ParentId      string `json:"parentId" gorm:"comment:父菜单ID"`
	Path          string `json:"path" gorm:"comment:路由path"`
	Name          string `json:"name" gorm:"comment:路由name"`
	Hidden        bool   `json:"hidden" gorm:"comment:是否在列表隐藏"`
	Component     string `json:"component" gorm:"comment:对应前端文件路径"`
	Sort          int    `json:"sort" gorm:"comment:排序标记"`
	Meta          `json:"meta" gorm:"comment:附加属性"`
	SysAuthoritys []SysAuthority         `json:"authoritys" gorm:"many2many:sys_authority_menus;"`
	Children      []SysBaseMenu          `json:"children" gorm:"-"`
	Parameters    []SysBaseMenuParameter `json:"parameters"`
}

type SysAuthority struct {
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time     `sql:"index"`
	AuthorityId     string         `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	AuthorityName   string         `json:"authorityName" gorm:"comment:角色名"`
	ParentId        string         `json:"parentId" gorm:"comment:父角色ID"`
	DataAuthorityId []SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id"`
	Children        []SysAuthority `json:"children" gorm:"-"`
	SysBaseMenus    []SysBaseMenu  `json:"menus" gorm:"many2many:sys_authority_menus;"`
	DefaultRouter   string         `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"`
}


type SysUser struct {
	GVA_MODEL
	UUID        uuid.UUID    `json:"uuid" gorm:"comment:用户UUID"`
	Username    string       `json:"userName" gorm:"comment:用户登录名" gorm:"index"`
	Password    string       `json:"-"  gorm:"comment:用户登录密码"`
	NickName    string       `json:"nickName" gorm:"default:系统用户;comment:用户昵称" `
	HeaderImg   string       `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
}



