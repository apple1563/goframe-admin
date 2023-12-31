package menuService

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-starter/api/vmenu"
	"goframe-starter/internal/consts"
	"goframe-starter/internal/dao"
	"goframe-starter/internal/model/entity"
	"goframe-starter/utility/xcasbin"
	"strings"
)

var menuCols = dao.Menu.Columns()

func AddMenu(ctx context.Context, req *vmenu.AddMenuReq) (res *vmenu.AddMenuRes, err error) {
	count, err := dao.Menu.Ctx(ctx).Where(g.Map{
		menuCols.Path: req.Path,
		menuCols.Pid:  req.Pid,
	}).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, consts.ErrMenuPathExists
	}
	_, err = dao.Menu.Ctx(ctx).Data(req).Insert()
	if err != nil {
		return nil, err
	}
	return
}
func DeleteMenu(ctx context.Context, req *vmenu.DeleteMenuReq) (res *vmenu.DeleteMenuRes, err error) {
	count, err := dao.Menu.Ctx(ctx).Where(menuCols.Pid, req.Id).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, consts.ErrMenuPathDeleteChildren
	}
	count, err = dao.Button.Ctx(ctx).Where(dao.Button.Columns().MenuId, req.Id).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, consts.ErrMenuPathDeleteChildren
	}
	_, err = dao.Menu.Ctx(ctx).Where(menuCols.Id, req.Id).Delete()
	if err != nil {
		return nil, err
	}
	// 删除关联的casbin，删除角色与菜单关联
	var obj = "menu " + gconv.String(req.Id)
	_, err = xcasbin.Enforcer.RemoveFilteredPolicy(1, obj)
	if err != nil {
		return nil, err
	}
	return
}
func AddMenuForRole(ctx context.Context, req *vmenu.MenuForRoleReq) (res *vmenu.MenuForRoleRes, err error) {
	var sub = consts.Role_Menu_Prefix + gconv.String(req.RoleId)
	_, err = xcasbin.Enforcer.RemoveFilteredPolicy(0, sub)
	if err != nil {
		return nil, err
	}
	for _, menuId := range req.MenuIds {
		var obj = "menu " + gconv.String(menuId)
		_, err := xcasbin.Enforcer.AddPolicy(sub, gconv.String(obj), "ALL")
		if err != nil {
			return nil, err
		}
	}
	return
}
func getRoleMenuIds(roleId uint) []int {
	var sub = consts.Role_Menu_Prefix + gconv.String(roleId)
	var rules = xcasbin.Enforcer.GetFilteredPolicy(0, sub)
	var res = make([]int, 0)
	for _, rule := range rules {
		parts := strings.Split(rule[1], " ")
		res = append(res, gconv.Int(parts[1]))
	}
	return res
}
func GetMenuByRole(ctx context.Context, req *vmenu.MenuByRoleReq) (res *vmenu.MenuByRoleRes, err error) {
	var resp = &vmenu.MenuByRoleRes{
		List: make([]int, 0),
	}
	var ids = getRoleMenuIds(req.RoleId)
	for _, id := range ids {
		resp.List = append(resp.List, id)
	}
	return resp, nil
}
func UpdateMenu(ctx context.Context, req *vmenu.UpdateMenuReq) (res *vmenu.UpdateMenuRes, err error) {
	_, err = dao.Menu.Ctx(ctx).Where(menuCols.Id, req.Id).Data(g.Map{
		menuCols.Pid:            req.Pid,
		menuCols.Title:          req.Title,
		menuCols.Name:           req.Name,
		menuCols.Path:           req.Path,
		menuCols.Icon:           req.Icon,
		menuCols.Type:           req.Type,
		menuCols.Redirect:       req.Redirect,
		menuCols.Permissions:    req.Permissions,
		menuCols.PermissionName: req.PermissionName,
		menuCols.Component:      req.Component,
		menuCols.AlwaysShow:     req.AlwaysShow,
		menuCols.ActiveMenu:     req.ActiveMenu,
		menuCols.IsRoot:         req.IsRoot,
		menuCols.IsFrame:        req.IsFrame,
		menuCols.FrameSrc:       req.FrameSrc,
		menuCols.KeepAlive:      req.KeepAlive,
		menuCols.Hidden:         req.Hidden,
		menuCols.Affix:          req.Affix,
		menuCols.Level:          req.Level,
		menuCols.Tree:           req.Tree,
		menuCols.Sort:           req.Sort,
		menuCols.Status:         req.Status,
	}).Update()
	if err != nil {
		return nil, err
	}
	return
}

// getAllMenus 从数据库获取所有菜单，再把菜单生成树形结构
func getAllMenus(ctx context.Context) (res []*entity.Menu, err error) {
	res = make([]*entity.Menu, 0)
	// 超级管理员返回所有菜单，其他角色按权限设置的来
	var roleId = gconv.Uint(ctx.Value("roleId"))
	if roleId == 1024 { // 1024为超级管理员
		err = dao.Menu.Ctx(ctx).OrderAsc(menuCols.Sort).Scan(&res)
	} else {
		err = dao.Menu.Ctx(ctx).WhereIn(menuCols.Id, getRoleMenuIds(roleId)).OrderAsc(menuCols.Sort).Scan(&res)
	}
	if err != nil {
		return nil, err
	}
	return
}

/*树形tree开始*/
func genTreeMenus(ctx context.Context) (res []*vmenu.TreeMenuItem, err error) {
	allMenus, err := getAllMenus(ctx)
	if err != nil {
		return nil, err
	}
	// 根级为0
	res = buildTreeMenus(allMenus, 0)
	return
}

func ListTreeMenus(ctx context.Context, req *vmenu.ListTreeMenuReq) (res *vmenu.ListTreeMenuRes, err error) {
	res = &vmenu.ListTreeMenuRes{}
	res.List = make([]*vmenu.TreeMenuItem, 0)
	menus, err := genTreeMenus(ctx)
	if err != nil {
		return nil, err
	}
	res.List = append(res.List, menus...)
	return
}

func buildTreeMenus(menus []*entity.Menu, pid int64) (res []*vmenu.TreeMenuItem) {
	for _, menu := range menus {
		if menu.Pid == pid {
			var treeMenuItem = &vmenu.TreeMenuItem{
				Menu: menu,
			}
			treeMenuItem.Children = buildTreeMenus(menus, menu.Id)
			res = append(res, treeMenuItem)
		}
	}
	return res
}

/*树形tree结束*/

/*vueRouter开始*/
func genVueMenus(menus []*vmenu.TreeMenuItem) (sources []*vmenu.VueMenu) {
	for _, men := range menus {
		var source = new(vmenu.VueMenu)
		source.Name = men.Name
		source.Path = men.Path
		source.Redirect = men.Redirect
		source.Component = men.Component
		source.Meta = &vmenu.VueMenuMeta{
			Title:            men.Title,
			Icon:             men.Icon,
			OrderNo:          men.Sort,
			Hidden:           men.Hidden == 1,
			HiddenBreadcrumb: men.AlwaysShow == 1,
			KeepAlive:        men.KeepAlive == 1,
			FrameSrc:         men.FrameSrc,
			IsFrame:          men.IsFrame == 1,
			Affix:            men.Affix == 1,
			/*
				ActiveMenu: men.ActiveMenu,
				IsRoot:     men.IsRoot == 1,
				//Permissions: men.Permissions,
				Type:  men.Type,*/
		}
		if len(men.Children) > 0 {
			source.Children = append(source.Children, genVueMenus(men.Children)...)
		}
		sources = append(sources, source)
	}
	return
}

func ListVueMenus(ctx context.Context, req *vmenu.VueMenuReq) (res *vmenu.VueMenuRes, err error) {
	res = &vmenu.VueMenuRes{}
	res.List = make([]*vmenu.VueMenu, 0)
	treeMenus, err := genTreeMenus(ctx) // type 3 为按钮，不计入
	if err != nil {
		return nil, err
	}
	vueMenus := genVueMenus(treeMenus)

	res.List = append(res.List, vueMenus...)
	return
}

/*vueRouter结束*/
