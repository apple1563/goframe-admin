package menuService

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-starter/api/vmenu"
	"goframe-starter/internal/consts"
	"goframe-starter/internal/dao"
	"goframe-starter/internal/model/entity"
)

var menuCols = dao.Menu.Columns()

func AddMenu(ctx context.Context, req *vmenu.AddMenuReq) (res *vmenu.AddMenuRes, err error) {
	count, err := dao.Menu.Ctx(ctx).Where(menuCols.Path, req.Path).Count()
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
	_, err = dao.Menu.Ctx(ctx).Where(menuCols.Id, req.Id).Delete()
	if err != nil {
		return nil, err
	}
	// TODO 删除关联的casbin，删除角色与菜单关联
	return
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
func getAllMenus(ctx context.Context, excludeTypes ...int) (res []*entity.Menu, err error) {
	res = make([]*entity.Menu, 0)
	if len(excludeTypes) > 0 {
		err = dao.Menu.Ctx(ctx).WhereNotIn(menuCols.Type, excludeTypes).Scan(&res)
	} else {
		err = dao.Menu.Ctx(ctx).Scan(&res)
	}
	if err != nil {
		return nil, err
	}
	return
}

/*树形tree开始*/
func genTreeMenus(ctx context.Context, excludeTypes ...int) (res []*vmenu.TreeMenuItem, err error) {
	allMenus, err := getAllMenus(ctx, excludeTypes...)
	if err != nil {
		return nil, err
	}
	// 根级为0
	res = buildTreeMenus(allMenus, 0)
	return
}

func ListTreeMenus(ctx context.Context, req *vmenu.ListTReeMenuReq) (res *vmenu.ListTReeMenuRes, err error) {
	res = &vmenu.ListTReeMenuRes{}
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
			Title: men.Title,
			Icon:  men.Icon,
			/*KeepAlive:  men.KeepAlive == 1,
			Hidden:     men.Hidden == 1,
			Sort:       men.Sort,
			AlwaysShow: men.AlwaysShow == 1,
			ActiveMenu: men.ActiveMenu,
			IsRoot:     men.IsRoot == 1,
			FrameSrc:   men.FrameSrc,
			//Permissions: men.Permissions,
			Affix: men.Affix == 1,
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
	treeMenus, err := genTreeMenus(ctx, 3) // type 3 为按钮，不计入
	if err != nil {
		return nil, err
	}
	vueMenus := genVueMenus(treeMenus)

	res.List = append(res.List, vueMenus...)
	return
}

/*vueRouter结束*/
