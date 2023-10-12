// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Provinces is the golang structure of table provinces for DAO operations like Where/Data.
type Provinces struct {
	g.Meta    `orm:"table:provinces, do:true"`
	Id        interface{} // 省市区ID
	Title     interface{} // 栏目名称
	Pinyin    interface{} // 拼音
	Lng       interface{} // 经度
	Lat       interface{} // 纬度
	Pid       interface{} // 父栏目
	Level     interface{} // 关系树等级
	Tree      interface{} // 关系
	Sort      interface{} // 排序
	Status    interface{} // 状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
