// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"nxxlzx/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/admin",
			beego.NSInclude(
				&controllers.AdminController{},
			),
		),

		beego.NSNamespace("/banner",
			beego.NSInclude(
				&controllers.BannerController{},
			),
		),

		beego.NSNamespace("/community",
			beego.NSInclude(
				&controllers.CommunityController{},
			),
		),

		beego.NSNamespace("/community_class",
			beego.NSInclude(
				&controllers.CommunityClassController{},
			),
		),

		beego.NSNamespace("/community_class_sub",
			beego.NSInclude(
				&controllers.CommunityClassSubController{},
			),
		),

		beego.NSNamespace("/community_reply",
			beego.NSInclude(
				&controllers.CommunityReplyController{},
			),
		),

		beego.NSNamespace("/expert",
			beego.NSInclude(
				&controllers.ExpertController{},
			),
		),

		beego.NSNamespace("/expert_class",
			beego.NSInclude(
				&controllers.ExpertClassController{},
			),
		),

		beego.NSNamespace("/expert_auth",
			beego.NSInclude(
				&controllers.ExpertAuthController{},
			),
		),

		beego.NSNamespace("/info",
			beego.NSInclude(
				&controllers.InfoController{},
			),
		),

		beego.NSNamespace("/info_class",
			beego.NSInclude(
				&controllers.InfoClassController{},
			),
		),

		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),

		beego.NSNamespace("/upload",
			beego.NSInclude(
				&controllers.UploadController{},
			),
		),

		beego.NSNamespace("/menu",
			beego.NSInclude(
				&controllers.MenuController{},
			),
		),

		beego.NSNamespace("/opinion",
			beego.NSInclude(
				&controllers.OpinionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
