package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
	"strconv"
	"time"

	"utronshop.io/models"
	"utronshop.io/util"

	"github.com/gernest/utron/controller"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

//User is a controller for user
type User struct {
	controller.BaseController
	Routes []string
}

// Index 页面
func (u *User) Index() {
	cookie, err := u.Ctx.Request().Cookie("username")
	if err != nil {
		log.Println(">>INFO>> user not login")
	} else {
		u.Ctx.Data["User"] = models.User{Name: cookie.Value}
	}
	u.Ctx.Template = "index"
	u.HTML(http.StatusOK)
}

//Login 登录页面
func (u *User) Login() {
	if u.Ctx.Request().Method == http.MethodGet {
		u.Ctx.Template = "login"
		u.HTML(http.StatusOK)
		return
	}
	req := u.Ctx.Request()
	uname := req.FormValue("name")
	upwd := req.FormValue("pwd")
	user := &models.User{}
	if !u.Ctx.DB.Where("name = ?", uname).Find(&user).RecordNotFound() {
		if user.Pwd == upwd {
			u.Ctx.Data["User"] = user
		} else {
			u.Ctx.Data["Message"] = "密码错误!"
			u.Ctx.Template = "login"
			u.HTML(http.StatusOK)
			return
		}
	} else {
		u.Ctx.Data["Message"] = "用户不存在!"
		u.Ctx.Template = "login"

		u.HTML(http.StatusOK)
		return
	}
	// 设置cookie
	expiration := time.Now().Add(time.Minute * 30)
	cookie := &http.Cookie{Name: "username",
		Value:   uname,
		Expires: expiration}

	http.SetCookie(u.Ctx.Response(), cookie)

	log.Println("user [", uname, "] login success")
	u.Ctx.Redirect("/", http.StatusFound)
}

//Signup 注册页面
func (u *User) Signup() {
	if u.Ctx.Request().Method == http.MethodGet {
		u.Ctx.Template = "signup"
		u.HTML(http.StatusOK)
		return
	}
	user := &models.User{}
	req := u.Ctx.Request()
	_ = req.ParseForm()
	if err := decoder.Decode(user, req.PostForm); err != nil {
		u.Ctx.Data["Message"] = err.Error()
		u.Ctx.Template = "error"
		u.HTML(http.StatusInternalServerError)
		return
	}

	// 查看用户是否已经存在
	uname := req.FormValue("name")
	upwd := req.FormValue("pwd")
	email := req.FormValue("email")
	// 注意,需要加入Find() u.Ctx.DB.Where("name = ?", uname).RecordNotFound()始终返回false
	if u.Ctx.DB.Where("name = ?", uname).Find(&user).RecordNotFound() {
		user.Name = uname
		user.Pwd = upwd
		u.Ctx.DB.Create(user)
		// 注册成功
		if util.SendMail(email) {
			u.Ctx.Data["Message"] = "已经发送邮件到 " + email + ", 请登录邮箱验证"
		} else {
			u.Ctx.Data["Message"] = "发送邮件失败， 请重试！"
		}
		u.Ctx.Template = "signup"
		u.HTML(http.StatusOK)
	} else {
		// 用户已存在
		u.Ctx.Data["Message"] = "该用户名已存在, 请重新注册!"
		u.Ctx.Template = "signup"
		u.HTML(http.StatusOK)
	}
}

//Delete deletes a user
func (u *User) Delete() {
	UserID := u.Ctx.Params["id"]
	ID, err := strconv.Atoi(UserID)
	if err != nil {
		u.Ctx.Data["Message"] = err.Error()
		u.Ctx.Template = "error"
		u.HTML(http.StatusInternalServerError)
		return
	}
	u.Ctx.DB.Delete(&models.User{ID: ID})
	u.Ctx.Redirect("/", http.StatusFound)
}

//ActiveMail active email
func (u *User) ActiveMail() {
	emailToken := u.Ctx.Request().FormValue("emailToken")
	addr := u.Ctx.Request().FormValue("addr")
	// fmt.Println(emailToken)

	hasher := md5.New()
	hasher.Write([]byte(addr))
	if emailToken == hex.EncodeToString(hasher.Sum(nil)) {
		user := &models.User{}
		u.Ctx.DB.Find(&user).Update("valid", true)
		u.Ctx.Data["Message"] = "邮箱验证通过， 现在登录吧！"
	} else {
		u.Ctx.Data["Message"] = "邮箱验证未通过， 您可以先登录"
	}
	u.Ctx.Template = "login"
	u.HTML(http.StatusOK)
}

//Logout logout.
func (u *User) Logout() {
	// 清除cookie
	expiration := time.Now().Add(-time.Minute)
	cookie := &http.Cookie{Name: "username",
		Value:   "",
		Expires: expiration}

	http.SetCookie(u.Ctx.Response(), cookie)
	u.Ctx.Redirect("/", http.StatusFound)
}

//NewUser returns a new  user controller
func NewUser() controller.Controller {
	return &User{
		Routes: []string{
			"get,post;/;Index",
			"get,post;/login;Login",
			"get,post;/signup;Signup",
			// "get;/signup;Signup",
			"get;/activemail;ActiveMail",
			"get;/logout;Logout",
		},
	}
}
