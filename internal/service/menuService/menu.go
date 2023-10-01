package menuService

import (
	"context"
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
