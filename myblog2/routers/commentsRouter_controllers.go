package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"] = append(beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"],
        beego.ControllerComments{
            Method: "ShowAddArticle",
            Router: `/admin/add.html`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"] = append(beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"],
        beego.ControllerComments{
            Method: "HanderAddArticle",
            Router: `/admin/addArticle/:key`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"] = append(beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"],
        beego.ControllerComments{
            Method: "AddLAble",
            Router: `/admin/addLable`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"] = append(beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"],
        beego.ControllerComments{
            Method: "ShowLAble",
            Router: `/admin/addLable.html`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"] = append(beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"],
        beego.ControllerComments{
            Method: "ShowAddType",
            Router: `/admin/addType.html`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"] = append(beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"],
        beego.ControllerComments{
            Method: "HanderAddType",
            Router: `/admin/addtype`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"] = append(beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"],
        beego.ControllerComments{
            Method: "ShowArticle",
            Router: `/admin/content`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"] = append(beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"],
        beego.ControllerComments{
            Method: "DeleteLAble",
            Router: `/admin/deleteLable`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"] = append(beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"],
        beego.ControllerComments{
            Method: "DeleteType",
            Router: `/admin/deleteType`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"] = append(beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"],
        beego.ControllerComments{
            Method: "ShowIndex",
            Router: `/admin/index.html`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"] = append(beego.GlobalControllerRouter["myblog2/controllers:ArticleControllers"],
        beego.ControllerComments{
            Method: "UploadImg",
            Router: `/admin/upload/:key`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:ContentControllers"] = append(beego.GlobalControllerRouter["myblog2/controllers:ContentControllers"],
        beego.ControllerComments{
            Method: "Praise",
            Router: `/Praise`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:ContentControllers"] = append(beego.GlobalControllerRouter["myblog2/controllers:ContentControllers"],
        beego.ControllerComments{
            Method: "ShowContent",
            Router: `/content`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:IndexController"] = append(beego.GlobalControllerRouter["myblog2/controllers:IndexController"],
        beego.ControllerComments{
            Method: "ShowIndex",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:IndexController"] = append(beego.GlobalControllerRouter["myblog2/controllers:IndexController"],
        beego.ControllerComments{
            Method: "ShowAbout",
            Router: `/about.html`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:IndexController"] = append(beego.GlobalControllerRouter["myblog2/controllers:IndexController"],
        beego.ControllerComments{
            Method: "ShowArticle",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:IndexController"] = append(beego.GlobalControllerRouter["myblog2/controllers:IndexController"],
        beego.ControllerComments{
            Method: "ShowDiary",
            Router: `/list.html`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:IndexController"] = append(beego.GlobalControllerRouter["myblog2/controllers:IndexController"],
        beego.ControllerComments{
            Method: "ShowLogin",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:SeekControllers"] = append(beego.GlobalControllerRouter["myblog2/controllers:SeekControllers"],
        beego.ControllerComments{
            Method: "Seek",
            Router: `/seek`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:UserControllers"] = append(beego.GlobalControllerRouter["myblog2/controllers:UserControllers"],
        beego.ControllerComments{
            Method: "Active",
            Router: `/active`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:UserControllers"] = append(beego.GlobalControllerRouter["myblog2/controllers:UserControllers"],
        beego.ControllerComments{
            Method: "Lonin",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:UserControllers"] = append(beego.GlobalControllerRouter["myblog2/controllers:UserControllers"],
        beego.ControllerComments{
            Method: "Lonout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:UserControllers"] = append(beego.GlobalControllerRouter["myblog2/controllers:UserControllers"],
        beego.ControllerComments{
            Method: "Reg",
            Router: `/regist`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog2/controllers:UserControllers"] = append(beego.GlobalControllerRouter["myblog2/controllers:UserControllers"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
