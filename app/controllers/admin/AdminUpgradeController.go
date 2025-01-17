package admin

import (
	"github.com/revel/revel"
	//	"encoding/json"
	"github.com/leanote/leanote/app/info"
	//	"io/ioutil"
)

// Upgrade controller
type AdminUpgrade struct {
	AdminBaseController
}

func (c AdminUpgrade) UpgradeBlog() {
	upgradeService.UpgradeBlog()
	return nil
}

func (c AdminUpgrade) UpgradeBetaToBeta2() {
	re := info.NewRe()
	re.Ok, re.Msg = upgradeService.UpgradeBetaToBeta2(c.GetUserId())
	return c.RenderJSON(re)
}

func (c AdminUpgrade) UpgradeBeta3ToBeta4() {
	re := info.NewRe()
	re.Ok, re.Msg = upgradeService.Api(c.GetUserId())
	return c.RenderJSON(re)
}
