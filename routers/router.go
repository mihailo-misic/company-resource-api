// @APIVersion 1.0.0
// @Title Company Resource API
// @Description API for the company-resource project https://gitlab.com/mihailo-misic/company-resource-api
// @Url http://www.apache.org/licenses/LICENSE-2.0.html
// @Contact mihailo.misic@gmail.com
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"gitlab.com/mihailo-misic/company-resource-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
