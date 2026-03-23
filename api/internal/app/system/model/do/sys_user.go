// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure of table sys_user for DAO operations like Where/Data.
type SysUser struct {
	g.Meta        `orm:"table:sys_user, do:true"`
	Id            any         //
	UserName      any         // 用户名
	Mobile        any         // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	UserNickname  any         // 用户昵称
	Birthday      any         // 生日
	UserPassword  any         // 登录密码;cmf_password加密
	UserSalt      any         // 加密盐
	UserStatus    any         // 用户状态;0:禁用,1:正常,2:未验证
	UserEmail     any         // 用户登录邮箱
	Sex           any         // 性别;0:保密,1:男,2:女
	Avatar        any         // 用户头像
	Remark        any         // 备注
	IsAdmin       any         // 是否后台管理员 1 是  0   否
	Address       any         // 联系地址
	Describe      any         // 描述信息
	LastLoginIp   any         // 最后登录ip
	LastLoginTime *gtime.Time // 最后登录时间
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 删除时间
	RoleId        any         //
}
