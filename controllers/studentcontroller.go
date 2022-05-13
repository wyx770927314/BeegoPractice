package controllers

import (
	"beegodemo01/models"
	"fmt"
	"github.com/astaxie/beego"
)

type StudentController struct {
	beego.Controller
}

func (stu *StudentController) Get() {
	//如果有Ctx.WriteString的话，就不会进行模板渲染了
	//stu.Ctx.WriteString("这是直接返回的数据")
	stu.Data["title"] = "滴滴滴滴滴滴"
	stu.TplName = "home.html"

}
func (stu *StudentController) AddStudent() {
	//如果有Ctx.WriteString的话，就不会进行模板渲染了
	//is_ok := stu.GetString("id")
	//age, _ := stu.GetInt("age")
	//stu.Ctx.WriteString(is_ok + "  " + strconv.Itoa(age))
	//stu.Data["title"] = "滴滴滴滴滴滴"
	stu.TplName = "addstudent.html"

}
func (stu *StudentController) DoAddStudent() {
	//获取post请求--方法一
	username := stu.GetString("username")
	password := stu.GetString("password")
	hobby := stu.GetStrings("hobby")
	for i := 0; i < len(hobby); i++ {
		fmt.Println(hobby[i])
	}
	//获取post请求--方法二

	stu.Ctx.WriteString(username + password)

}

//form表单提交的内容转成struct
func (stu *StudentController) StructStudent() {
	s := models.Student{}
	err := stu.ParseForm(&s)
	if err != nil {
		stu.Ctx.WriteString("post解析失败")
	}
	fmt.Println(s)
	stu.Ctx.WriteString("post解析成功")

}

//struct转成json串
func (stu *StudentController) JsonStudent() {
	s := models.Student{"wyx", 27, []string{"追剧", "看电视"}}
	stu.Data["json"] = s
	stu.ServeJSON()

}
