// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-04 21:05:49
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        int         `json:"id"         ` //
	Avatar    string      `json:"avatar"     ` //
	Email     string      `json:"email"      ` //
	Name      string      `json:"name"       ` //
	Password  string      `json:"password"   ` //
	Type      int         `json:"type"       ` //
	CreatedAt *gtime.Time `json:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updated_at" ` //
	DeletedAt *gtime.Time `json:"deleted_at" ` //
	Introduce string      `json:"introduce"  ` //
}
