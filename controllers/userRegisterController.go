package controllers

import (
	"DataProject/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}
func (r *RegisterController) Post(){
	//1.解析请求数据
	var user models.User
	err := r.ParseForm(&user)
	if err !=nil{
		//返回错误信息给浏览器，提示用户
		r.TplName = "error.html"
		return
	}
	//2.保存用户信息到数据库
	_, err =user.SaveUser()
	//3.返回前端结果
	if err !=nil{
		r.TplName = "error.html"
		return
	}
	//用户注册成功
	r.TplName = "login.html"
}