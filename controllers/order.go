package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gernest/utron/controller"
	"utronshop.io/models"
)

//Order is a controller for order
type Order struct {
	controller.BaseController
	Routes []string
}

//NewOrder returns a new  order controller
func NewOrder() controller.Controller {
	return &Order{
		Routes: []string{
			"get,post;/book;Book",
		},
	}
}

//Book 订单详情
func (order *Order) Book() {
	cookie, err := order.Ctx.Request().Cookie("username")
	if err != nil {
		log.Println(">>INFO>> product detail page, user not login")
	} else {
		order.Ctx.Data["User"] = models.User{Name: cookie.Value}

		if order.Ctx.Request().Method == "POST" {
			// 此时订单添加到数据库
			pid, _ := strconv.Atoi(order.Ctx.Request().FormValue("number"))
			username := order.Ctx.Request().FormValue("username")
			email := order.Ctx.Request().FormValue("email")
			tel := order.Ctx.Request().FormValue("tel")
			addr := order.Ctx.Request().FormValue("addr")

			norder := &models.Order{PID: pid, Username: username, Email: email, Tel: tel, Addr: addr}
			order.Ctx.DB.Create(norder)
		}

		// 查询当前用户的订单信息
		orders := &[]models.Order{}
		order.Ctx.DB.Where("username = ?", cookie.Value).Find(&orders)
		order.Ctx.Data["Orders"] = orders
	}
	order.Ctx.Template = "order"
	order.HTML(http.StatusOK)
}
