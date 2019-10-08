package member

import (
	"github.com/revel/revel"
)

// 帐户信息

type MemberUser struct {
	MemberBaseController
}

func (c MemberUser) Username() {
	c.SetUserInfo()
	c.SetLocale()
	c.ViewArgs["title"] = c.Message("Username")
	return c.RenderTemplate("member/user/username.html")
}

func (c MemberUser) Email() {
	c.SetUserInfo()
	c.SetLocale()
	c.ViewArgs["title"] = c.Message("Email")
	return c.RenderTemplate("member/user/email.html")
}

func (c MemberUser) Password() {
	c.SetUserInfo()
	c.SetLocale()
	c.ViewArgs["title"] = c.Message("Password")
	return c.RenderTemplate("member/user/password.html")
}

func (c MemberUser) Avatar() {
	c.SetUserInfo()
	c.SetLocale()
	c.ViewArgs["title"] = c.Message("Avatar")
	c.ViewArgs["globalConfigs"] = configService.GetGlobalConfigForUser()
	return c.RenderTemplate("member/user/avatar.html")
}
