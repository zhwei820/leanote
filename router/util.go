package router

import "github.com/gogf/gf/net/ghttp"

func BindGroup(s *ghttp.Server, path string, items []ghttp.GroupItem) {
	g := s.Group(path)
	g.Bind(items)
}
