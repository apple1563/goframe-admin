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

type TreeMenuItem struct {
	*entity.Menu
	Children []*TreeMenuItem
}

// 从数据库获取所有菜单，再把菜单生成树形结构
func GetAllMenus(ctx context.Context) (res []*entity.Menu, err error) {
	res = make([]*entity.Menu, 0)
	err = dao.Menu.Ctx(ctx).Scan(&res)
	if err != nil {
		return nil, err
	}
	return
}

func GetTreeMenus(ctx context.Context) (res []*TreeMenuItem, err error) {
	allMenus, err := GetAllMenus(ctx)
	if err != nil {
		return nil, err
	}
	// 根级为0
	res = buildTreeMenus(allMenus, 0)
	return
}

func buildTreeMenus(menus []*entity.Menu, pid int64) (res []*TreeMenuItem) {
	for _, menu := range menus {
		if menu.Pid == pid {
			var treeMenuItem = &TreeMenuItem{
				Menu: menu,
			}
			treeMenuItem.Children = buildTreeMenus(menus, menu.Id)
			res = append(res, treeMenuItem)
		}
	}
	return res
}

//  这个方法要查多次数据库，放弃
/*func ListMenuByLevel(ctx context.Context, pid int64) (res []*TreeMenuItem, err error) {
	err = dao.Menu.Ctx(ctx).Where(g.Map{
		menuCols.Pid: pid,
	}).Scan(&res)
	if err != nil {
		return nil, err
	}
	return
}

// 递归真是玩不明白
func TreeMenu(ctx context.Context, pid int64) (res []*TreeMenuItem, err error) {
	menu, err := ListMenuByLevel(ctx, pid)
	if err != nil {
		return nil, err
	}
	if menu != nil && len(menu) > 0 {
		for _, item := range menu {
			item.Children, err = TreeMenu(ctx, item.Id)
			if err != nil {
				return nil, err
			}
		}
	}
	return menu, nil
}*/
