package vmenu

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-starter/internal/model/entity"
)

type AddMenuReq struct {
	g.Meta `tags:"菜单" method:"post" path:"/menu"  sm:"添加" dc:"添加菜单" `
	*entity.Menu
}
type AddMenuRes struct{}

type UpdateMenuReq struct {
	g.Meta `tags:"菜单" method:"put" path:"/menu"  sm:"修改" dc:"修改菜单" `
	*entity.Menu
}
type UpdateMenuRes struct{}

type ListMenuReq struct {
	g.Meta `tags:"菜单" method:"get" path:"/menu/list"  sm:"列表" dc:"菜单列表" `
	*entity.Menu
	//*vcommon.CommonPageReq
}

type ListMenuRes struct {
	List []*entity.Menu `json:"list"`
	//*vcommon.CommonPageRes
}

type DeleteMenuReq struct {
	g.Meta `tags:"菜单" method:"delete" path:"/menu"  sm:"删除" dc:"删除菜单" `
	Id     int64 `json:"id" v:"required" description:"菜单ID"`
}
type DeleteMenuRes struct{}

type TreeMenuItem struct {
	*entity.Menu
	Children []*TreeMenuItem `json:"children"`
}
type ListTreeMenuReq struct {
	g.Meta `tags:"菜单" method:"get" path:"/menu/tree"  sm:"列表" dc:"菜单管理树形列表" `
}
type ListTreeMenuRes struct {
	List []*TreeMenuItem `json:"list"`
}
type VueMenuMeta struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
	//expanded boolean
	OrderNo          int    `json:"orderNo"`
	Hidden           bool   `json:"hidden"`
	HiddenBreadcrumb bool   `json:"hiddenBreadcrumb"`
	Single           bool   `json:"single"`
	KeepAlive        bool   `json:"keepAlive"`
	FrameSrc         string `json:"frameSrc"`
	IsFrame          bool   `json:"isFrame"`
	Affix            bool   `json:"affix"`
}
type VueMenu struct {
	Path      string       `json:"path"`
	Name      string       `json:"name"`
	Component string       `json:"component"`
	Redirect  string       `json:"redirect"`
	Meta      *VueMenuMeta `json:"meta"`
	Children  []*VueMenu   `json:"children"`
}
type VueMenuReq struct {
	g.Meta `tags:"菜单" method:"get" path:"/menu/vue"  sm:"列表" dc:"返回给前端做动态路由" `
}
type VueMenuRes struct {
	List []*VueMenu `json:"list"`
}

type MenuForRoleReq struct {
	g.Meta  `tags:"菜单" method:"post" path:"/menu/role"  sm:"角色绑定菜单" dc:"为角色绑定菜单" `
	RoleId  uint  `json:"roleId"`
	MenuIds []int `json:"menuIds"`
}

type MenuForRoleRes struct{}

type MenuByRoleReq struct {
	g.Meta `tags:"菜单" method:"get" path:"/menu/role"  sm:"角色绑定菜单" dc:"获取角色绑定的菜单" `
	RoleId uint `json:"roleId"`
}

type MenuByRoleRes struct {
	List []int `json:"list"`
}
