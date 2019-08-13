package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["nxxlzx/controllers:AdminController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:AdminController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:AdminController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:AdminController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:AdminController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:BannerController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:BannerController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:BannerController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:BannerController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:BannerController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:BannerController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:BannerController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:BannerController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:CommunityClassController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:CommunityClassController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:CommunityClassController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:CommunityClassController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:CommunityClassController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:CommunityClassController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:CommunityClassController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:CommunityClassController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:CommunityClassSubController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:CommunityClassSubController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:CommunityClassSubController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:CommunityClassSubController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:CommunityClassSubController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:CommunityClassSubController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:CommunityController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:CommunityController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:CommunityController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:CommunityController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:CommunityController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:CommunityController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:CommunityController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:CommunityController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:CommunityController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:CommunityController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:CommunityReplyController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:CommunityReplyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:CommunityReplyController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:CommunityReplyController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:CommunityReplyController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:CommunityReplyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:ExpertAuthController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:ExpertAuthController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:ExpertAuthController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:ExpertAuthController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:ExpertAuthController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:ExpertAuthController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:ExpertAuthController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:ExpertAuthController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:ExpertClassController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:ExpertClassController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:ExpertClassController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:ExpertClassController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:ExpertClassController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:ExpertClassController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:ExpertClassController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:ExpertClassController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:ExpertController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:ExpertController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:ExpertController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:ExpertController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:ExpertController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:ExpertController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:ExpertController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:ExpertController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:ExpertController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:ExpertController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:ImController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:ImController"],
        beego.ControllerComments{
            Method: "Ws",
            Router: `/ws`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:InfoClassController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:InfoClassController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:InfoClassController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:InfoClassController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:InfoClassController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:InfoClassController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:InfoClassController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:InfoClassController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:InfoController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:InfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:InfoController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:InfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:InfoController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:InfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:InfoController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:InfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:InfoController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:InfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:MenuController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:MenuController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:MenuController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:MenuController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:MenuController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:MenuController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:MenuController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:MenuController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:OpinionController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:OpinionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:OpinionController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:OpinionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:OpinionController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:OpinionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:OpinionController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:OpinionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:UploadController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:UploadController"],
        beego.ControllerComments{
            Method: "Upload",
            Router: `/upload`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:UsersController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:UsersController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:UsersController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:UsersController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:UsersController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:UsersController"],
        beego.ControllerComments{
            Method: "SessionKey",
            Router: `/session_key`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nxxlzx/controllers:UsersController"] = append(beego.GlobalControllerRouter["nxxlzx/controllers:UsersController"],
        beego.ControllerComments{
            Method: "SignUp",
            Router: `/signup`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
