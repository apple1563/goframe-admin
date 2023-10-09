// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Button is the golang structure for table button.
type Button struct {
	Id        uint   `json:"id"        description:""`
	MenuId    int64  `json:"menuId"    description:"按钮所在菜单id"`
	MenuTitle string `json:"menuTitle" description:"按钮所在菜单名称"`
	Name      string `json:"name"      description:"按钮标识符"`
	Title     string `json:"title"     description:"按钮名称"`
	Remark    string `json:"remark"    description:""`
}
