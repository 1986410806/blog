// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-04 21:05:49
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Tag is the golang structure of table tag for DAO operations like Where/Data.
type Tag struct {
	g.Meta    `orm:"table:tag, do:true"`
	Id        interface{} //
	Name      interface{} //
	Desc      interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
