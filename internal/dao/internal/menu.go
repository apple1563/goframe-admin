// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MenuDao is the data access object for table menu.
type MenuDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns MenuColumns // columns contains all the column names of Table for convenient usage.
}

// MenuColumns defines and stores column names for table menu.
type MenuColumns struct {
	Id             string // 菜单ID
	Pid            string // 父菜单ID，0表示根级
	Title          string // 菜单名称
	Name           string // 名称编码
	Path           string // 路由地址
	Icon           string // 菜单图标
	Type           string // 菜单类型（1目录 2菜单）
	Redirect       string // 重定向地址
	Permissions    string // 菜单包含权限集合
	PermissionName string // 权限名称
	Component      string // 组件路径
	AlwaysShow     string // 取消自动计算根路由模式  1是2否
	ActiveMenu     string // 高亮菜单编码
	IsRoot         string // 是否跟路由
	IsFrame        string // 是否跳转外链
	FrameSrc       string // iframe地址
	KeepAlive      string // 缓存该路由，1是2否
	Hidden         string // 是否隐藏,1是2否
	Affix          string // 是否固定
	Level          string // 关系树等级 1根2子3孙
	Tree           string // 关系树
	Sort           string // 排序
	Remark         string // 备注
	Status         string // 菜单状态
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// menuColumns holds the columns for table menu.
var menuColumns = MenuColumns{
	Id:             "id",
	Pid:            "pid",
	Title:          "title",
	Name:           "name",
	Path:           "path",
	Icon:           "icon",
	Type:           "type",
	Redirect:       "redirect",
	Permissions:    "permissions",
	PermissionName: "permission_name",
	Component:      "component",
	AlwaysShow:     "always_show",
	ActiveMenu:     "active_menu",
	IsRoot:         "is_root",
	IsFrame:        "is_frame",
	FrameSrc:       "frame_src",
	KeepAlive:      "keep_alive",
	Hidden:         "hidden",
	Affix:          "affix",
	Level:          "level",
	Tree:           "tree",
	Sort:           "sort",
	Remark:         "remark",
	Status:         "status",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewMenuDao creates and returns a new DAO object for table data access.
func NewMenuDao() *MenuDao {
	return &MenuDao{
		group:   "default",
		table:   "menu",
		columns: menuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MenuDao) Columns() MenuColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
