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

/*
	{
	    "path": "/list",
	    "name": "list",
	    "component": "LAYOUT",
	    "redirect": "/list/base",
	    "meta": {
	        "title": "列表页",
	        "icon": "view-list"
	    },
	    "children": [
	        {
	            "path": "base",
	            "name": "ListBase",
	            "component": "/list/base/index",
	            "meta": {
	                "title": "基础列表页"
	            }
	        },
	    ]
	}
*/
type TreeMenuItem struct {
	*entity.Menu
	Children []*TreeMenuItem
}
type ListTReeMenuReq struct {
	g.Meta `tags:"菜单" method:"get" path:"/menu/tree"  sm:"列表" dc:"菜单管理树形列表" `
}
type ListTReeMenuRes struct {
	List []*TreeMenuItem `json:"list"`
}
type VueMenuMeta struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
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
