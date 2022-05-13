package routers

import (
	"beegodemo01/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/student", &controllers.StudentController{})
	beego.Router("/student/addstu", &controllers.StudentController{}, "get:AddStudent")
	beego.Router("/student/doaddstu", &controllers.StudentController{}, "post:DoAddStudent")
	beego.Router("/student/formtostruct", &controllers.StudentController{}, "post:StructStudent")
	beego.Router("/student/structtojson", &controllers.StudentController{}, "get:JsonStudent")
}
