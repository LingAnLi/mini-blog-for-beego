package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myblog2/models"
	"myblog2/syserror"
	"regexp"
	"strings"
)

type UserControllers struct {
	BaseControllers
}
// @router /login [get]
func (this *IndexController) ShowLogin() {
	//this.Data["ttle"]="技术分享"
	//this.Layout="layout.html"
	this.TplName="admin/login.html"
}
//@router /login [post]
func(this*UserControllers)Lonin() {

	userName:=this.GetStrNotNil("userName","账号不能为空")
	passwd:=this.GetStrNotNil("passWord","密码不能为空")
	o:=orm.NewOrm()
	var user  models.User
	user.Name=userName
	err:=o.Read(&user,"Name")
	if err!=nil{
		this.Abort500(syserror.NewErr("账号错误",err))
		return
	}
	if user.PassWord!=passwd{
		this.Abort500(syserror.NewErr("密码错误",err))
	}

	if !user.IsActive{
		this.DelSession(SESSION_USER_KEY)
		this.Abort500(errors.New("账号未激活请前往邮箱激活"))
	}
	//设置免登入
	remrber:=this.GetString("remember")
	if remrber=="on"{
		beego.BConfig.WebConfig.Session.SessionCookieLifeTime=60*60*24*7
	}else {
		beego.BConfig.WebConfig.Session.SessionCookieLifeTime=1
	}

	this.SetSession(SESSION_USER_KEY,user)
	if user.Adminrivilege!=0{
		this.Redirect("admin/index.html",302)
	}else {
		this.Redirect("/",302)
	}

}
//@router /logout [get]
func(this*UserControllers)Lonout() {
	this.MustLogin()
	this.DelSession(SESSION_USER_KEY)
	this.Redirect("/",302)
}
//@router /register [get]
func(this*UserControllers)Register() {

	this.TplName="admin/register.html"
}
//@router /register [post]
func(this*UserControllers)Reg() {
	o:=orm.NewOrm()
	var user models.User
	//{"Email":,"userName":"passWord":,"passWord2":,"admin":};
	email:=this.GetStrNotNil("Email","Email不能为空")
	userName:=this.GetStrNotNil("userName","userName不能为空")
	passwd:=this.GetStrNotNil("passWord","passWord不能为空")
	passwd2:=this.GetStrNotNil("passWord2","确认密码不能为空")
	admin:=this.GetString("admin")
	beego.Info(strings.Split(email, "."))
	//校验邮箱
	if !IsEmail(email){
		this.Abort500(errors.New("邮箱格式不正确"))
	}
	user.Email=email
	//userName 唯一
	user.Name=userName
	err:=o.Read(&user,"Name")
	if err==nil {
		this.Abort500(errors.New("用户名已存在"))
	}
	//密码校验
	if passwd!=passwd2{
		this.Abort500(errors.New("两次密码不一致"))
	}
	user.PassWord=passwd
	//是否为管理员
	if admin=="on"{
		user.Adminrivilege=2
	}
	//开始事物
	o.Begin()
	_,err=o.Insert(&user)
	if err!=nil{
		this.Abort500(errors.New("注册失败，请重新尝试"))
	}
	err=SendEmail(user.Email)
	if err!=nil{
		o.Rollback()
		this.Abort500(err)
	}
	o.Commit()

	this.JsonOk("前往邮箱激活","/login")
}
//@router /active [get]
func(this*UserControllers)Active(){
	id,err:=this.GetInt("id")
	if err!=nil{
		this.Abort500(syserror.Error404{})
	}
	var user models.User
	user.Id=id
	o:=orm.NewOrm()
	err=o.Read(&user)
	if err!=nil{
		this.Abort500(errors.New("参数错误"))
	}
	if user.IsActive==true{
		this.Abort500(errors.New("已经激活"))
	}
	user.IsActive=true
	_,err=o.Update(&user)
	if err!=nil{
		this.Abort500(errors.New("激活失败"))
	}

	this.Redirect("/login",302)
}
func IsEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
func SendEmail(email string)error{
	//发送激活邮件--》》/active?id=user.Id

	return nil
}