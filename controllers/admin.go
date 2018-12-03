package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gernest/utron/controller"
	"utronshop.io/models"
)

//AdminCtl admin controller
type AdminCtl struct {
	controller.BaseController
	Routes []string
}

//Index 后台管理主页
func (ac *AdminCtl) Index() {
	// 从数据库获取所有的product数据
	products := &[]models.Product{}
	ac.Ctx.DB.Find(&products)
	ac.Ctx.Data["Products"] = products
	ac.Ctx.Data["IndexPage"] = true
	ac.Ctx.Template = "admin"
	ac.HTML(http.StatusOK)
}

//Create add a new record into DB
func (ac *AdminCtl) Create() {
	if ac.Ctx.Request().Method == "POST" {
		number, _ := strconv.Atoi(ac.Ctx.Request().FormValue("number"))
		name := ac.Ctx.Request().FormValue("name")
		mtype := ac.Ctx.Request().FormValue("type")
		count, _ := strconv.Atoi(strings.Trim(ac.Ctx.Request().FormValue("count"), " ")) // 去除空白字符
		price, _ := strconv.Atoi(ac.Ctx.Request().FormValue("price"))
		href := ac.Ctx.Request().FormValue("href")
		url := ac.Ctx.Request().FormValue("url")
		snumber := ac.Ctx.Request().FormValue("number")

		product := &models.Product{
			Number: number,
			Name:   name,
			Type:   mtype,
			Count:  count,
			Price:  price,
			Href:   href,
			URL:    url,
			Brief:  "/data/" + snumber + "/brief",
			Detail: "/data/" + snumber + "/detail",
		}
		ac.Ctx.DB.Create(&product)
		ac.Ctx.Redirect("/admin", http.StatusFound)
	} else {
		ac.Ctx.Data["AddPage"] = true
		ac.Ctx.Template = "admin-add"
		ac.HTML(http.StatusOK)
	}
}

//Update update record
func (ac *AdminCtl) Update() {
	if ac.Ctx.Request().Method == "POST" {
		number, _ := strconv.Atoi(ac.Ctx.Request().FormValue("number"))
		name := ac.Ctx.Request().FormValue("name")
		mtype := ac.Ctx.Request().FormValue("type")
		count, _ := strconv.Atoi(ac.Ctx.Request().FormValue("count"))
		price, _ := strconv.Atoi(ac.Ctx.Request().FormValue("price"))
		href := ac.Ctx.Request().FormValue("href")
		url := ac.Ctx.Request().FormValue("url")
		snumber := ac.Ctx.Request().FormValue("number")

		product := &models.Product{
			Number: number,
			Name:   name,
			Type:   mtype,
			Count:  count,
			Price:  price,
			Href:   href,
			URL:    url,
			Brief:  "/data/" + snumber + "/brief",
			Detail: "/data/" + snumber + "/detail",
		}
		ac.Ctx.DB.Where("number = ?", number).First(&models.Product{}).Updates(&product)
		ac.Ctx.Redirect("/admin", http.StatusFound)
	} else {
		number := ac.Ctx.Params["number"]
		product := &models.Product{}
		ac.Ctx.DB.Where("number = ?", number).First(&product)
		ac.Ctx.Data["Product"] = product
		ac.Ctx.Template = "admin-update"
		ac.HTML(http.StatusOK)
	}
}

//Delete delete record
func (ac *AdminCtl) Delete() {
	number := ac.Ctx.Request().FormValue("number")
	ac.Ctx.DB.Where("number = ?", number).Delete(&models.Product{})
	ac.Ctx.Redirect("/admin", http.StatusFound)
}

//NewAdminCtl returns a new  admin controller
func NewAdminCtl() controller.Controller {
	return &AdminCtl{
		Routes: []string{
			"get;/admin;Index",
			"get,post;/admin/add;Create",
			"get;/admin/update/{number:[0-9]+};Update",
			"post;/admin/update;Update",
			"post;/admin/delete;Delete",
		},
	}
}
