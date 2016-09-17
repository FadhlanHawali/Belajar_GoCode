package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["paserta/controllers:PasertaController"] = append(beego.GlobalControllerRouter["paserta/controllers:PasertaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["paserta/controllers:PasertaController"] = append(beego.GlobalControllerRouter["paserta/controllers:PasertaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["paserta/controllers:PasertaController"] = append(beego.GlobalControllerRouter["paserta/controllers:PasertaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["paserta/controllers:PasertaController"] = append(beego.GlobalControllerRouter["paserta/controllers:PasertaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["paserta/controllers:PasertaController"] = append(beego.GlobalControllerRouter["paserta/controllers:PasertaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["paserta/controllers:PsertaController"] = append(beego.GlobalControllerRouter["paserta/controllers:PsertaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["paserta/controllers:PsertaController"] = append(beego.GlobalControllerRouter["paserta/controllers:PsertaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["paserta/controllers:PsertaController"] = append(beego.GlobalControllerRouter["paserta/controllers:PsertaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["paserta/controllers:PsertaController"] = append(beego.GlobalControllerRouter["paserta/controllers:PsertaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["paserta/controllers:PsertaController"] = append(beego.GlobalControllerRouter["paserta/controllers:PsertaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
